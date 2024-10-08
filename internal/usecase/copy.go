package usecase

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/lavinas/vooo-etl/internal/domain"
	"github.com/lavinas/vooo-etl/internal/port"
)

// Copy is a struct that represents the use case copy a object from database to another
type Copy struct {
	Base
}

// Run runs the use case
func (c *Copy) Run(job port.Domain, back bool, txTarget interface{}) (string, int64, error) {
	now := time.Now()
	j := job.(*domain.Job)
	if j.Type != "table" {
		return "", 0, errors.New(port.ErrJobTypeNotImplemented)
	}
	tmiss := float64(0)
	processed, missing, copied, sub, err := c.runFactory(j, back, txTarget)
	if err != nil {
		return "", 0, err
	}
	if processed != 0 {
		tmiss = (time.Since(now).Seconds() * float64(missing)) / float64(3600*processed)
	}
	return fmt.Sprintf(port.CopyReturnMessage, processed, copied, sub, tmiss), missing, nil
}

// runFactory runs the factory use case
func (c *Copy) runFactory(j *domain.Job, back bool, txTarget interface{}) (int64, int64, int64, int64, error) {
	switch j.Action {
	case "all":
		return c.runCopyAll(j, txTarget)
	case "copy":
		return c.runCopy(j, back, txTarget)
	default:
		return 0, 0, 0, 0, fmt.Errorf(port.ErrActionNotFound, j.Action)
	}
}

// runCopyAll runs the copy all use case
func (c *Copy) runCopyAll(j *domain.Job, txTarget interface{}) (int64, int64, int64, int64, error) {
	cols, rows, err := c.getSourceAll(j)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	if err := c.deleteTargetAll(j, txTarget); err != nil {
		return 0, 0, 0, 0, err
	}
	colsMap := make(map[string]int)
	for i, col := range cols {
		colsMap[col] = i
	}
	rows, err = c.selectRefs(j, colsMap, rows, txTarget)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	err = c.putSource(j, cols, rows, txTarget)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return int64(len(rows)), 0, int64(len(rows)), 0, nil
}

// runCopy runs the copy use case
func (c *Copy) runCopy(j *domain.Job, back bool, txTarget interface{}) (int64, int64, int64, int64, error) {
	keys, missing, processed, err := c.getLimits(j, back)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	cols, rows, err := c.getSource(j, back, keys)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	rows, keys, sub, err := c.filterRefs(j, cols, rows, keys, txTarget)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	if err = c.move(j, keys, rows, txTarget); err != nil {
		return 0, 0, 0, 0, err
	}
	if err := j.SetKeysLast(keys, c.RepoTarget, txTarget); err != nil {
		return 0, 0, 0, 0, err
	}
	return processed, missing, int64(len(rows)), sub, nil
}

// filterRefs filters the references
func (c *Copy) filterRefs(j *domain.Job, cols []string, rows [][]*string, keys []int64, tx interface{}) ([][]*string, []int64, int64, error) {
	if len(rows) == 0 || len(j.Refs) == 0 {
		return rows, keys, 0, nil
	}
	colsMap := make(map[string]int)
	for i, col := range cols {
		colsMap[col] = i
	}
	rows, keys, sub, err := c.limitRefs(j, colsMap, rows, keys)
	if err != nil {
		return nil, nil, 0, err
	}
	rows, err = c.selectRefs(j, colsMap, rows, tx)
	if err != nil {
		return nil, nil, 0, err
	}
	return rows, keys, sub, nil
}

// limitRefs limits rows to the limits of the references
func (c *Copy) limitRefs(j *domain.Job, cols map[string]int, rows [][]*string, keys []int64) ([][]*string, []int64, int64, error) {
	message, err := c.checkRefs(j, cols, rows)
	if err != nil {
		return nil, nil, 0, err
	}
	if message == "" {
		return rows, keys, 0, nil
	}
	rows, reduced, err := c.limitRefsReduces(j, cols, rows, message)
	if err != nil {
		return nil, nil, 0, err
	}
	newKeys, err := c.limitNewKeys(j, cols, rows)
	if err != nil {
		return nil, nil, 0, err
	}
	return rows, newKeys, reduced, nil
}

// limitRefsReduces limits rows to the limits of the references
func (c *Copy) limitRefsReduces(j *domain.Job, cols map[string]int, rows [][]*string, message string) ([][]*string, int64, error) {
	rows = c.sortRows(j, cols, rows)
	red := port.CopyReduce
	total := int64(len(rows))
	for ; ; red += port.CopyReduce {
		if red >= port.CopyReduceLimit || red >= total {
			return nil, 0, errors.New(message)
		}
		rows = rows[:int(total-red)]
		message, err := c.checkRefs(j, cols, rows)
		if err != nil {
			return nil, 0, err
		}
		if message == "" {
			break
		}
	}
	return rows, red, nil
}

// limitNewKeys limits the new keys
func (c *Copy) limitNewKeys(j *domain.Job, cols map[string]int, rows [][]*string) ([]int64, error) {
	limits := make([]int64, len(j.Keys))
	var err error
	for i, key := range j.Keys {
		limits[i], err = strconv.ParseInt(*rows[len(rows)-1][cols[key.Name]], 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return limits, nil
}

// sortRows sorts the rows by the keys
func (c *Copy) sortRows(j *domain.Job, cols map[string]int, rows [][]*string) [][]*string {
	iKeys := make([]int, len(j.Keys))
	for i, key := range j.Keys {
		iKeys[i] = cols[key.Name]
	}
	sort.Slice(rows, func(i, j int) bool {
		for _, key := range iKeys {
			ik, _ := strconv.ParseInt(*rows[i][key], 10, 64)
			jk, _ := strconv.ParseInt(*rows[j][key], 10, 64)
			if ik > jk {
				return false
			}
		}
		return true
	})
	return rows
}

// limitRefs limits rows to the limits of the references
func (c *Copy) checkRefs(j *domain.Job, cols map[string]int, rows [][]*string) (string, error) {
	for r := range j.Refs {
		for i := range j.Refs[r].Keys {
			_, max, err := c.getRefRange(j.Refs[r].Keys[i].Referrer, cols, rows)
			if err != nil {
				return "", err
			}
			if message, err := c.limitRefKey(&j.Refs[r].Job, j.Refs[r].Keys[i].Referred, max); err != nil || message != "" {
				return message, err
			}
		}
	}
	return "", nil
}

// filterRefLimits filters the references by limits
func (c *Copy) limitRefKey(job *domain.Job, name string, max int64) (string, error) {
	keys := job.Keys
	for _, key := range keys {
		if key.Name == name {
			if max > key.Last {
				message := fmt.Sprintf(port.ErrReferenceNotDone, job.Name, name, max, key.Last)
				return message, nil
			}
			return "", nil
		}
	}
	return c.limitRefNotKey(job, name, max)
}

// filterRefLimitsNotKey filters the references by limits if reference key is not a job key
func (c *Copy) limitRefNotKey(job *domain.Job, name string, max int64) (string, error) {
	rows, err := c.getlimitRefNotKey(job, name, max)
	if err != nil {
		return "", err
	}
	for i, key := range job.Keys {
		val, err := strconv.ParseInt(*rows[0][i], 10, 64)
		if err != nil {
			return "", err
		}
		if val > key.Last {
			message := fmt.Sprintf(port.ErrReferenceNotDone, job.Name, name, max, key.Last)
			return message, nil
		}
	}
	return "", nil
}

// getFilterRefLimits gets the filter reference limits
func (c *Copy) getlimitRefNotKey(job *domain.Job, name string, max int64) ([][]*string, error) {
	kCols := ""
	for _, key := range job.Keys {
		kCols += key.Name + ", "
	}
	kCols = kCols[:len(kCols)-2]
	q := fmt.Sprintf(port.CopyMaxExists, kCols, job.Base, job.Object, name, max)
	tx := c.RepoSource.Begin(job.Base)
	defer c.RepoSource.Rollback(tx)
	_, rows, err := c.RepoSource.Query(tx, q)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 || len(rows[0]) == 0 || rows[0][0] == nil {
		return nil, fmt.Errorf(port.ErrReferenceNotFound, job.Base, job.Object, name, max)
	}
	return rows, nil
}

// filterRefs filters the references
func (c *Copy) selectRefs(j *domain.Job, cols map[string]int, rows [][]*string, tx interface{}) ([][]*string, error) {
	for r := range j.Refs {
		mRows := [][]*string{}
		for i := range j.Refs[r].Keys {
			r, err := c.filterRefbyKey(j, r, i, cols, rows, tx)
			if err != nil {
				return nil, err
			}
			mRows = append(mRows, r...)
		}
		rows = mRows
	}
	return rows, nil
}

// filterRefbyKey filters the references by key
func (c *Copy) filterRefbyKey(j *domain.Job, r int, i int, cols map[string]int, rows [][]*string, tx interface{}) ([][]*string, error) {
	min, max, err := c.getRefRange(j.Refs[r].Keys[i].Referrer, cols, rows)
	if err != nil {
		return nil, err
	}
	possibles, err := c.getRefPossibles(&j.Refs[r], i, min, max, tx)
	if err != nil {
		return nil, err
	}
	rows, err = c.filterRef(j.Refs[r].Keys[i].Referrer, possibles, cols, rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// getRefPossibles gets the possible references
func (c *Copy) getRefPossibles(ref *domain.Ref, i int, min int64, max int64, txTarget interface{}) (map[int64]bool, error) {
	field := ref.Keys[i].Referred
	sql := fmt.Sprintf(port.CopySelectRef, field, ref.Job.Base, ref.Job.Object, field, min-1, field, max)
	_, rows, err := c.RepoTarget.Query(txTarget, sql)
	if err != nil {
		return nil, err
	}
	possibles := make(map[int64]bool)
	for _, row := range rows {
		val, err := strconv.ParseInt(*row[0], 10, 64)
		if err != nil {
			return nil, err
		}
		possibles[val] = true
	}
	return possibles, nil
}

// filterRef filters the references
func (c *Copy) filterRef(field string, possibles map[int64]bool, cols map[string]int, rows [][]*string) ([][]*string, error) {
	iField, ok := cols[field]
	if !ok {
		return nil, errors.New(port.ErrFieldReferrerNotFound)
	}
	var filtered [][]*string
	for _, row := range rows {
		if row[iField] == nil {
			continue
		}
		val, err := strconv.ParseInt(*row[iField], 10, 64)
		if err != nil {
			return nil, err
		}
		if _, ok := possibles[val]; ok {
			filtered = append(filtered, row)
		}
	}
	return filtered, nil
}

// getRefRange gets the reference range of ids based on the source data
func (c *Copy) getRefRange(field string, cols map[string]int, rows [][]*string) (int64, int64, error) {
	iField, ok := cols[field]
	if !ok {
		return 0, 0, errors.New(port.ErrFieldReferrerNotFound)
	}
	var min, max int64
	for _, row := range rows {
		if row[iField] == nil {
			continue
		}
		val, err := strconv.ParseInt(*row[iField], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		if min == 0 || val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max, nil
}

// deleteTargetAll deletes the target data
func (c *Copy) deleteTargetAll(j *domain.Job, txTarget interface{}) error {
	if _, err := c.RepoTarget.Exec(txTarget, port.CopyDisableFK); err != nil {
		return err
	}
	q := fmt.Sprintf(port.CopyDeleteAll, j.Base, j.Object)
	if _, err := c.RepoTarget.Exec(txTarget, q); err != nil {
		return err
	}
	if _, err := c.RepoTarget.Exec(txTarget, port.CopyEnableFK); err != nil {
		return err
	}
	return nil
}

// getSource gets the source data for the insert
func (c *Copy) getSource(j *domain.Job, back bool, limit []int64) ([]string, [][]*string, error) {
	fields, err := c.mountSourceFields(j)
	if err != nil {
		return nil, nil, err
	}
	txSource := c.RepoSource.Begin(j.Base)
	defer c.RepoSource.Rollback(txSource)
	allFields, err := c.getAllFields(j, txSource)
	if err != nil {
		return nil, nil, err
	}
	cols := []string{}
	rows := [][]*string{}
	for i := range j.Keys {
		co, ro, err := c.getSourceByKey(j, i, fields, allFields, limit[i], back, txSource)
		if err != nil {
			return nil, nil, err
		}
		if len(ro) != 0 {
			cols = co
			rows = append(rows, ro...)
		}
	}
	return cols, rows, nil
}

// mountSourceFields mounts the source fields name for the select
func (c *Copy) mountSourceFields(j *domain.Job) (string, error) {
	ret := ""
	for _, key := range j.Keys {
		ret += key.Name + ", "
	}
	for _, ref := range j.Refs {
		for _, key := range ref.Keys {
			ret += key.Referrer + ", "
		}
	}
	if ret == "" {
		return "", errors.New(port.ErrFieldNotFound)
	}
	return ret[:len(ret)-2], nil
}

// getFields gets the fields from target table
func (c *Copy) getAllFields(j *domain.Job, tx interface{}) (string, error) {
	q := fmt.Sprintf(port.CopyGetFields, j.Base, j.Object)
	_, rows, err := c.RepoSource.Query(tx, q)
	if err != nil {
		return "", err
	}
	mountedFields, err := c.mountAllFields(rows)
	if err != nil {
		return "", err
	}
	return mountedFields, nil
}

// mountAllFields mounts the fields to a string
func (c *Copy) mountAllFields(fields [][]*string) (string, error) {
	if len(fields) == 0 {
		return "", errors.New(port.ErrInvalidUpdateFields)
	}
	ret := ""
	pat := "ifnull(`%s`, ''), "
	for _, field := range fields {
		ret += fmt.Sprintf(pat, *field[0])
	}
	return ret[:len(ret)-2], nil
}

// getSourceByKey gets the source data for the insert
func (c *Copy) getSourceByKey(j *domain.Job, i int, fields, allFields string, limit int64, back bool, txSource interface{}) ([]string, [][]*string, error) {
	last := j.Keys[i].Last
	if back {
		last -= j.Keys[i].Back
		if last < 0 {
			last = -1
		}
	}
	sql := fmt.Sprintf(port.CopySelectF, fields, allFields, j.Base, j.Object, j.Keys[i].Name, last, j.Keys[i].Name, limit)
	cols, rows, err := c.RepoSource.Query(txSource, sql)
	if err != nil {
		return nil, nil, err
	}
	return cols, rows, nil
}

// getSourceAll gets all the source data for the insert
func (c *Copy) getSourceAll(j *domain.Job) ([]string, [][]*string, error) {
	tx := c.RepoSource.Begin(j.Base)
	defer c.RepoSource.Rollback(tx)
	sql := fmt.Sprintf(port.CopySelectAllCount, j.Base, j.Object)
	_, rows, err := c.RepoSource.Query(tx, sql)
	if err != nil {
		return nil, nil, err
	}
	count, _ := strconv.ParseInt(*rows[0][0], 10, 64)
	if count > port.AllLimit {
		return nil, nil, fmt.Errorf(port.ErrTooManyRows, count, port.InLimit)
	}
	sql = fmt.Sprintf(port.CopySelectAll, j.Base, j.Object)
	cols, rows, err := c.RepoSource.Query(tx, sql)
	if err != nil {
		return nil, nil, err
	}
	return cols, rows, nil
}

// getAllSource gets all the source data for the insert
func (c *Copy) move(j *domain.Job, keys []int64, rows [][]*string, txTarget interface{}) error {
	if len(rows) == 0 {
		return nil
	}
	rows, err := c.filterMoved(j, keys, rows, txTarget)
	if err != nil {
		return err
	}
	txSource := c.RepoSource.Begin(j.Base)
	defer c.RepoSource.Rollback(txSource)
	last := int64(len(rows))
	for i := int64(0); i < last; i += port.InLimit {
		if err := c.moveAtomic(j, rows[i:min(i+port.InLimit, last)], txTarget, txSource); err != nil {
			return err
		}
	}
	return nil
}

// filterMoved filters the moved data
func (c *Copy) filterMoved(j *domain.Job, limits []int64, rows [][]*string, txTarget interface{}) ([][]*string, error) {
	keys, err := c.mountKeys(j)
	if err != nil {
		return nil, err
	}
	target, err := c.getMD5Target(j, keys, limits, txTarget)
	if err != nil {
		return nil, err
	}
	ret := [][]*string{}
	for _, row := range rows {
		f, err := c.mountIds(j, [][]*string{row})
		if err != nil {
			return nil, err
		}
		if _, ok := target[f]; ok && target[f] == *row[len(row)-1] {
			delete (target, f)
			continue
		}
		ret = append(ret, row)
	}
	if err := c.deleteMoved(j, keys, target, txTarget); err != nil {
		return nil, err
	}
	return ret, nil
}

// deleteMoved deletes the moved data
func (c *Copy) deleteMoved(j *domain.Job, keys string, target map[string]string, txTarget interface{}) error {
	if len(target) == 0 {
		return nil
	}
	ids := ""
	for i := range target {
		ids += i + ", "
	}
	ids = ids[:len(ids)-2]
	if _, err := c.RepoTarget.Exec(txTarget, port.CopyDisableFK); err != nil {
		return err
	}
	q := fmt.Sprintf(port.CopyDeleteIn, j.Base, j.Object, keys, ids)
	_, err := c.RepoTarget.Exec(txTarget, q)
	if err != nil {
		return err
	}
	if _, err := c.RepoTarget.Exec(txTarget, port.CopyEnableFK); err != nil {
		return err
	}
	return nil
}

// getMD5Target gets the md5 of the target data
func (c *Copy) getMD5Target(j *domain.Job, keys string, limits []int64, txTarget interface{}) (map[string]string, error) {
	ret := make(map[string]string)
	allFields, err := c.getAllFields(j, txTarget)
	if err != nil {
		return nil, err
	}
	for i := range j.Keys {
		r, err := c.getMD5TargetbyKey(j, keys, allFields, i, limits, txTarget)
		if err != nil {
			return nil, err
		}
		for k, v := range r {
			ret[k] = v
		}
	}
	return ret, nil
}

// getMD5TargetbyKey gets the md5 of the target data by key
func (c *Copy) getMD5TargetbyKey(j *domain.Job, keys string, allFields string, i int, limits []int64, tx interface{}) (map[string]string, error) {
	ret := make(map[string]string)
	q := fmt.Sprintf(port.CopySelectF, keys, allFields, j.Base, j.Object, j.Keys[i].Name, j.Keys[i].Last, j.Keys[i].Name, limits[i])
	_, rows, err := c.RepoTarget.Query(tx, q)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		kVal := ""
		for i := 0; i < len(row)-1; i++ {
			kVal += *row[i] + ", "
		}
		kVal = "(" + kVal[:len(kVal)-2] + ")"
		ret[kVal] = *row[len(row)-1]
	}
	return ret, nil
}


// runWait waits for the run to finish or db relief time
func (c *Copy) moveAtomic(j *domain.Job, rows [][]*string, txTarget interface{}, txSource interface{}) error {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(2)
	chn := make(chan bool, 1)
	go func(chn chan bool) {
		err = c.moveAtom(j, rows, txTarget, txSource)
		wg.Done()
		chn <- true
	}(chn)
	go func() {
		time.Sleep(port.DbRelief)
		wg.Done()
	}()
	select {
	case <-chn:
		wg.Wait()
	case <-c.Ctx.Done():
		return errors.New(port.ErrTimeout)
	case <-c.Signal:
		return errors.New(port.ErrInterrupted)
	}
	return err
}

// getAllSource steps the source data for the insert
func (c *Copy) moveAtom(j *domain.Job, rows [][]*string, txTarget interface{}, txSource interface{}) error {
	fields, err := c.mountKeys(j)
	if err != nil {
		return err
	}
	ids, err := c.mountIds(j, rows)
	if err != nil {
		return err
	}
	sql := fmt.Sprintf(port.CopySelectIn, j.Base, j.Object, fields, ids, fields)
	cols, rows, err := c.RepoSource.Query(txSource, sql)
	if err != nil {
		return err
	}
	if err = c.putSourceAtomic(j, c.mountInsertCols(cols), rows, txTarget); err != nil {
		return err
	}
	return nil
}

// mountAllSourceFields mounts the source fields name for the select
func (c *Copy) mountKeys(j *domain.Job) (string, error) {
	ret := ""
	for _, key := range j.Keys {
		ret += key.Name + ", "
	}
	if ret == "" {
		return "", errors.New(port.ErrFieldNotFound)
	}
	return ret[:len(ret)-2], nil
}

// mountAllSourceIds mounts the source ids for the select
func (c *Copy) mountIds(j *domain.Job, rows [][]*string) (string, error) {
	ret := ""
	for _, row := range rows {
		r := ""
		for i := 0; i < len(j.Keys); i++ {
			r += *row[i] + ", "
		}
		ret += "(" + r[:len(r)-2] + "), "
	}
	if ret == "" {
		return "", errors.New(port.ErrFieldNotFound)
	}
	return ret[:len(ret)-2], nil
}

// putSource puts the source data into the target
func (c *Copy) putSource(j *domain.Job, cols []string, rows [][]*string, txTarget interface{}) error {
	if len(rows) == 0 {
		return nil
	}
	scols := c.mountInsertCols(cols)
	last := int64(len(rows))
	for i := int64(0); i < last; i += port.OutLimit {
		if err := c.putSourceAtomic(j, scols, rows[i:min(i+port.OutLimit, last)], txTarget); err != nil {
			return err
		}
	}
	return nil
}

// putSourceAtomic puts the source data in installments way
func (c *Copy) putSourceAtomic(j *domain.Job, cols string, rows [][]*string, txTarget interface{}) error {
	if _, err := c.RepoTarget.Exec(txTarget, port.CopyDisableFK); err != nil {
		return err
	}
	cmd := c.mountInsert(j.Base, j.Object, cols, rows)
	if cmd == "" {
		return nil
	}
	if _, err := c.RepoTarget.Exec(txTarget, cmd); err != nil {
		return err
	}
	if _, err := c.RepoTarget.Exec(txTarget, port.CopyEnableFK); err != nil {
		return err
	}
	return nil
}

// mountInsertCols mounts coluns for insert
func (c *Copy) mountInsertCols(cols []string) string {
	strCols := "("
	for _, col := range cols {
		strCols += fmt.Sprintf("`%s`, ", col)
	}
	return strCols[:len(strCols)-2] + ")"
}

// mountInsert mounts the insert sql
func (c *Copy) mountInsert(base string, tablename string, cols string, rows [][]*string) string {
	values := ""
	for _, row := range rows {
		value := ""
		for _, col := range row {
			value += c.formatValue(col) + ", "
		}
		values += fmt.Sprintf("(%s), ", value[:len(value)-2])
	}
	return fmt.Sprintf(port.CopyInsert, base, tablename, cols, values[:len(values)-2])
}

// formatValue formats the value to insert
func (c *Copy) formatValue(col *string) string {
	if col == nil {
		return "NULL"
	}
	ret := *col
	ret = strings.Replace(ret, "'", "''", -1)
	ret = strings.Replace(ret, "\\", "\\\\", -1)
	ret = strings.Replace(ret, "\n", "", -1)
	ret = strings.Replace(ret, "\r", "", -1)
	ret = strings.Replace(ret, "\t", "", -1)
	ret = strings.Replace(ret, "0000-00-00 00:00:00", "2001-01-01 00:00:00", -1)
	return fmt.Sprintf("'%s'", ret)
}

// getMaxClient gets the max id from the client table
func (c *Copy) getLimits(j *domain.Job, back bool) ([]int64, int64, int64, error) {
	tx := c.RepoSource.Begin(j.Base)
	defer c.RepoSource.Rollback(tx)
	limits := make([]int64, len(j.Keys))
	var missing, processed int64
	for i := range j.Keys {
		l, m, p, err := c.getKeyLimit(j, i, back, tx)
		if err != nil {
			return nil, 0, 0, err
		}
		limits[i] = l
		missing += m
		processed += p
	}
	return limits, missing, processed, nil
}

// getKeyLimit gets the limit of the job key
func (c *Copy) getKeyLimit(j *domain.Job, i int, back bool, tx interface{}) (int64, int64, int64, error) {
	q := fmt.Sprintf(port.CopyMaxClient, j.Keys[i].Name, j.Object)
	_, rows, err := c.RepoSource.Query(tx, q)
	if err != nil {
		return 0, 0, 0, err
	}
	if len(rows[0]) == 0 || rows[0][0] == nil {
		return 0, 0, 0, nil
	}
	max, err := strconv.ParseInt(*rows[0][0], 10, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	last := j.Keys[i].Last
	if back {
		last -= j.Keys[i].Back
		if last < 0 {
			last = -1
		}
	}
	l := last + j.Keys[i].Step
	if l > max {
		l = max
	}
	return l, max - l, l - last, nil
}
