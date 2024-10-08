package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/lavinas/vooo-etl/internal/domain"
	"github.com/lavinas/vooo-etl/internal/port"
)

// Run is a struct that represents the use case
type Run struct {
	Base
}

// NewRun creates a new use case
func NewRun(repoSource port.Repository, repoTarget port.Repository, signal chan os.Signal) *Run {
	return &Run{
		Base: Base{
			RepoSource: repoSource,
			RepoTarget: repoTarget,
			Signal:     signal,
		},
	}
}

// Run runs app with given parameters
func (r *Run) Run(in *port.RunIn, out chan *port.RunOut) {
	if in == nil || out == nil {
		panic("Run: Invalid input parameters")
	}
	start := time.Now()
	defer r.finish(r.RepoTarget, out, start)
	count := int64(0)
	back := false
	for {
		r.RepoSource.Reload()
		r.RepoTarget.Reload()
		if in.Back > 0 && (count+1)%in.Back == 0 && in.Shifts == 0 {
			back = true
		}
		if r.runCycle(in, out, start, back) {
			break
		}
		count++
		if in.Repeat != 0 && count >= in.Repeat {
			break
		}
		if r.sleep(in) {
			r.sendOut(out, -1, -1, -1, port.InterrupedStatus, "interrupted", start, back)
			break
		}
	}
}

// runCycle runs a cycle of jobs
func (r *Run) runCycle(in *port.RunIn, out chan *port.RunOut, start time.Time, back bool) bool {
	jobs, err := r.getJobsId(in, r.RepoTarget)
	if err != nil {
		r.sendOut(out, -1, -1, -1, port.ErrorStatus, err.Error(), start, back)
		return true
	}
	for _, j := range jobs {
		ret := r.runJob(j, in, out, back)
		if r.getOut(ret, in.ErrorSkip) {
			return true
		}
	}
	return false
}

// runUntil runs all jobs until finish all registers
func (r *Run) runJob(job *domain.Job, in *port.RunIn, out chan *port.RunOut, back bool) *port.RunOut {
	shift := int64(1)
	for {
		start := time.Now()
		ret := r.runJobCycle(job.Id, out, start, shift, back)
		if ret.Status == port.ErrorStatus || ret.Status == port.InterrupedStatus {
			return ret
		}
		if ret.More == 0 {
			return ret
		}
		if in.Shifts != 0 && shift >= in.Shifts {
			return ret
		}
		back = false
		shift++
	}
}

// runJobCycle runs a cycle of job
func (r *Run) runJobCycle(jobId int64, out chan *port.RunOut, start time.Time, shift int64, back bool) *port.RunOut {
	exec := &domain.Log{}
	if err := exec.Init(r.RepoTarget, jobId, start, shift); err != nil {
		return r.sendOut(out, jobId, shift, -1, port.ErrorStatus, err.Error(), start, back)
	}
	tx := r.RepoTarget.Begin("")
	defer r.RepoTarget.Rollback(tx)
	var close context.CancelFunc
	r.Ctx, close = context.WithTimeout(context.Background(), port.RunTimeout)
	defer close()
	ret := r.runWait(jobId, out, start, shift, back, tx)
	if err := exec.SetStatus(r.RepoTarget, ret); err != nil {
		return r.sendOut(out, jobId, shift, -1, port.ErrorStatus, err.Error(), start, back)
	}
	return ret
}

// runWait waits for a given time
func (r *Run) runWait(jobId int64, out chan *port.RunOut, start time.Time, shift int64, back bool, tx interface{}) *port.RunOut {
	run := make(chan *port.RunOut)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(port.DbRelief)
		wg.Done()
	}()
	go r.runAtom(jobId, run, tx, start, shift, &wg, back)
	select {
	case <-r.Ctx.Done():
		return r.sendOut(out, jobId, shift, -1, port.ErrorStatus, "timeout", start, back)
	case <-r.Signal:
		return r.sendOut(out, jobId, shift, -1, port.InterrupedStatus, "interrupted", start, back)
	case r := <-run:
		wg.Wait()
		r.Duration = time.Since(start).Seconds()
		out <- r
		return r
	}
}

// run runs the use case
func (r *Run) runAtom(jobId int64, out chan *port.RunOut, tx interface{},
	start time.Time, shift int64, wg *sync.WaitGroup, back bool) *port.RunOut {
	defer wg.Done()
	job := &domain.Job{Id: jobId}
	if err := job.Load(r.RepoTarget, tx, true); err != nil {
		return r.sendOut(out, jobId, shift, -1, port.ErrorStatus, err.Error(), start, back)
	}
	action, err := r.factory(job.Action)
	if err != nil {
		return r.sendOut(out, jobId, shift, -1, port.ErrorStatus, err.Error(), start, back)
	}
	if action == nil {
		return r.sendOut(out, job.Id, 0, 0, port.SuccessStatus, "none action", time.Now(), back)
	}
	message, more, err := action.Run(job, back, tx)
	if err != nil {
		return r.sendOut(out, jobId, shift, more, port.ErrorStatus, err.Error(), start, back)
	}
	if err := r.RepoTarget.Commit(tx); err != nil {
		return r.sendOut(out, jobId, shift, more, port.ErrorStatus, err.Error(), start, back)
	}
	return r.sendOut(out, jobId, shift, more, port.SuccessStatus, message, start, back)
}

// sendMessage sends a message to the channel
func (r *Run) sendOut(out chan *port.RunOut, id int64, shift, more int64, status, detail string, start time.Time, back bool) *port.RunOut {
	dur := time.Since(start).Seconds()
	ret := &port.RunOut{JobID: id, Shift: shift, Status: status, Detail: detail, Duration: dur, More: more, Backed: back}
	out <- ret
	return ret
}

// sleep sleeps for a given time or until a signal is received
func (r *Run) sleep(in *port.RunIn) bool {
	select {
	case <-r.Signal:
		return true
	case <-time.After(time.Duration(in.Delay) * time.Second):
		return false
	}
}

// finish sends a finish message to the channel
func (r *Run) finish(repo port.Repository, out chan *port.RunOut, start time.Time) {
	ret := port.RunOut{
		JobID:    -1,
		Shift:    -1,
		Status:   port.FinishedStatus,
		Detail:   "finished signal",
		Start:    start,
		Duration: time.Since(start).Seconds(),
		More:     -1,
	}
	log := &domain.Log{}
	if err := log.Init(repo, -1, start, -1); err != nil {
		r.sendOut(out, -1, -1, -1, port.ErrorStatus, err.Error(), start, false)
		r.sendOut(out, -1, -1, -1, port.FinishedStatus, "", start, false)
		return
	}
	if err := log.SetStatus(repo, &ret); err != nil {
		r.sendOut(out, -1, -1, -1, port.ErrorStatus, err.Error(), start, false)
		r.sendOut(out, -1, -1, -1, port.FinishedStatus, "", start, false)
		return
	}
	out <- &ret
}

// getOut returns true if the RunOut has an error or is interrupted
func (r *Run) getOut(ret *port.RunOut, skipError bool) bool {
	return (ret.Status == port.ErrorStatus && skipError) || ret.Status == port.InterrupedStatus
}

// getJobsId
func (r *Run) getJobsId(in *port.RunIn, repo port.Repository) ([]*domain.Job, error) {
	job := domain.Job{}
	jobs, err := job.GetAll(repo)
	if err != nil {
		return nil, err
	}
	if jobs == nil || len(*jobs) == 0 {
		return nil, errors.New(port.ErrJobsNotFound)
	}
	ret := make([]*domain.Job, 0)
	for _, j := range *jobs {
		if j.Id >= in.JobID && j.Id <= in.Until {
			ret = append(ret, &j)
		}
	}
	return ret, err
}

// factoryAction creates a new action use case
func (r *Run) factory(action string) (port.RunAction, error) {
	base := Base{RepoSource: r.RepoSource, RepoTarget: r.RepoTarget, Ctx: r.Ctx, Signal: r.Signal}
	switch action {
	case "copy":
		return &Copy{Base: base}, nil
	case "all":
		return &Copy{Base: base}, nil
	case "update":
		return &Update{Base: base}, nil
	case "none":
		return nil, nil
	}
	return nil, fmt.Errorf(port.ErrActionNotFound, action)
}
