package port

import (
	"gorm.io/gorm/logger"
	"math"
	"time"
)

// messages
const (
	ErrRepoNilTx                = "tx informed is nil"
	ErrRepoInvalidTX            = "tx informed is invalid"
	ErrRepoNilObject            = "object informed is nil"
	ErrRepoInvalidObject        = "object informed is invalid"
	ErrRepoSshInvalid           = "ssh dns is invalid"
	ErrJobNotFound              = "job not found"
	ErrJobNotReady              = "job is not ready"
	ErrJobNotRunning            = "job is not running"
	ErrRepoPassNotImplemented   = "just file implemented"
	ErrRepoProtoNotImplemented  = "just tcp implemented"
	ErrActionNotFound           = "action not found: %s"
	ErrJobTypeNotImplemented    = "job type not implemented"
	ErrFieldNotFound            = "field not found"
	ErrReferenceNotFound        = "reference %s.%s.%s = %d not found"
	ErrReferenceNotDone         = "reference %s not done. Field %s %d > %d"
	ErrAggregatorNotFound       = "aggregator not found"
	ErrJobsNotFound             = "no jobs is found"
	ErrRepoSshTimeout           = "ssh timeout connecting to database"
	ErrFieldReferrerNotFound    = "referrer field not found in reference table"
	ErrFieldReferredNotFound    = "referred field not found in reference table"
	ErrInvalidUpadateReferences = "invalid update references. Just one reference is allowed"
	ErrInvalidUpdateSource      = "source and ids have different length"
	ErrInvelidCopySource        = "source consulted has shorter length than expected %d < %d"
	ErrInvalidUpdateFields      = "no fields field found on update"
	ErrTimeout                  = "timeout"
	ErrInterrupted              = "interrupted"
	ErrJobKeyNotFound           = "job key not found"
	ErrKeysLength               = "keys length is different"
	ErrSetupSchemaEmpty         = "schema is empty"
	ErrNoTablesFound            = "no tables found"
	ErrTableReferceNotFound     = "table reference %s not found"
	ErrCircularReference        = "circular reference in table %s"
	ErrNoSchemasFound           = "no schemas found"
	ErrTooManyRows              = "too many rows for all action: done %d, max %d"
	ErrJobKeyMismatch           = "reference key %s not found in job %s"
)

// queries
const (
	CopyGetFields              = "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s' ORDER BY COLUMN_NAME;"
	CopyDisableFK              = "SET FOREIGN_KEY_CHECKS = 0;"
	CopyEnableFK               = "SET FOREIGN_KEY_CHECKS = 1;"
	CopyMaxClient              = "SELECT max(%s) FROM `%s`;"
	CopyMaxExists              = "SELECT %s FROM %s.`%s` where `%s` = %d;"
	CopySelectIn               = "SELECT * FROM %s.`%s` WHERE (%s) in (%s) order by %s;"
	CopySelectAllCount         = "SELECT count(1) FROM %s.`%s`;"
	CopySelectAll              = "SELECT * FROM %s.`%s`;"
	CopyDeleteAll              = "DELETE FROM %s.`%s`;"
	CopySelectF                = "SELECT %s, md5(concat(%s)) md5 FROM %s.`%s` WHERE %s > %d and %s <= %d;"
	CopySelectRef			   = "SELECT %s FROM %s.`%s` WHERE %s > %d and %s <= %d;"
	CopyDeleteF                = "DELETE FROM %s.`%s` WHERE %s > %d and %s <= %d;"
	CopyDeleteIn               = "DELETE FROM %s.`%s` WHERE (%s) in (%s);"
	CopyInsert                 = "REPLACE INTO %s.`%s` %s VALUES %s;"
	LoadClientAggregator       = "SELECT id FROM aggregator_ref;"
	LoadClientSelect           = "SELECT id FROM client WHERE id_aggregator in (%s) and id > %d and id <= %d order by id;"
	LoadClientInsert           = "INSERT IGNORE INTO client_ref VALUES %s;"
	LoadClientMax              = "SELECT max(id) FROM client;"
	UpdateDisableFK            = "SET FOREIGN_KEY_CHECKS = 0;"
	UpdateEnableFK             = "SET FOREIGN_KEY_CHECKS = 1;"
	UpdateGetFields            = "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s';"
	UpdateSelectID1            = "SELECT %s, md5(concat(%s)) FROM %s.%s WHERE %s > %d and %s <= %d order by 1;"
	UpdateSelectID2            = "SELECT %s, md5(concat(%s)) FROM %s.%s WHERE (%s) in (%s) order by 1;"
	UpdateSelectAll            = "SELECT * FROM %s.%s WHERE (%s) in (%s);"
	UpdateInsert               = "REPLACE INTO %s.%s(%s) VALUES %s;"
	SetupDisableFK             = "SET FOREIGN_KEY_CHECKS = 0;"
	SetupEnableFK              = "SET FOREIGN_KEY_CHECKS = 1;"
	SetUpSelectTables          = "SELECT table_schema, table_name from information_schema.tables where table_schema in (%s) and not (%s) order by 1,2;"
	SetUpSelectPrime           = "SELECT table_name, column_name, table_schema FROM information_schema.columns WHERE table_schema in (%s) AND table_name in (%s) AND column_key = 'PRI';"
	SetUpSelectForeignInternal = "SELECT table_name, column_name, referenced_table_schema, referenced_table_name, referenced_column_name FROM information_schema.key_column_usage WHERE table_schema in (%s) and table_name IN (%s) and referenced_table_name is not null;"
	SetUpSelectForeignExternal = "SELECT referrer_table, referrer_field, referenced_schema, referenced_table, referenced_field FROM _ref WHERE referrer_schema in (%s) AND referrer_table in (%s);"
	SetUpExcepts               = "SELECT stat FROM _expt;"
	SetUpKey                   = "SELECT object, field, init_key from _key;"
	SetUpTruncate              = "TRUNCATE TABLE %s;"
	SetUpSchemas               = "SELECT name, ref_type from _schma where name = '%s' or %s;"
	SetUpAllSchemas            = "SELECT name, init_key from _schma;"
	TruncateDisableFK          = "SET FOREIGN_KEY_CHECKS = 0;"
	TruncateEnableFK           = "SET FOREIGN_KEY_CHECKS = 1;"
	TruncateTruncate           = "TRUNCATE TABLE %s.%s;"
)

// domain variables
var (
	Int64Nil = int64(math.MinInt64)
)

// usecases variables
var (
	RunTimeout            = 30 * time.Second
	DbRelief              = 0 * time.Millisecond
	LoadClientSourceBase  = "vooo_prod_backend"
	CopySourceBase        = "vooo_prod_backend"
	CopyReduce            = int64(1)
	CopyReduceLimit       = int64(2000)
	UpdateReturnMessage   = "%d processed, %d updated"
	CopyReturnMessage     = "%d rows proc, %d cop, %d sub, %.2f hrs miss"
	TruncateReturnMessage = "%s.%s truncated"
	InLimit               = int64(1000)
	OutLimit              = int64(1000)
	AllLimit  		      = int64(100000)
)

// MySQL variables
var (
	LoggerType        = logger.Default.LogMode(logger.Silent)
	ConnectionTimeout = time.Second * 5
)
