// This file has automatically been generated on Wed Feb 26 02:09:48 +05 2020.
// DO NOT EDIT.
package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	_ "unsafe"
)

//go:linkname nulltimescan sql.sub_nulltimescan
func nulltimescan(n *sql.NullTime, value interface{}) error {
	return n.Scan(value)
}

//go:linkname NullTimeScan sql.sub_nulltimescan
//go:noescape
func NullTimeScan(n *sql.NullTime, value interface{}) error

//go:linkname stmtqueryrowcontext sql.sub_stmtqueryrowcontext
func stmtqueryrowcontext(s *sql.Stmt, ctx context.Context, args ...interface{}) *sql.Row {
	return s.QueryRowContext(ctx, args)
}

//go:linkname StmtQueryRowContext sql.sub_stmtqueryrowcontext
//go:noescape
func StmtQueryRowContext(s *sql.Stmt, ctx context.Context, args ...interface{}) *sql.Row

//go:linkname columntypedecimalsize sql.sub_columntypedecimalsize
func columntypedecimalsize(ci *sql.ColumnType) (int64, bool) {
	return ci.DecimalSize()
}

//go:linkname ColumnTypeDecimalSize sql.sub_columntypedecimalsize
//go:noescape
func ColumnTypeDecimalSize(ci *sql.ColumnType) (int64, bool)

//go:linkname dbclose sql.sub_dbclose
func dbclose(db *sql.DB) error {
	return db.Close()
}

//go:linkname DBClose sql.sub_dbclose
//go:noescape
func DBClose(db *sql.DB) error

//go:linkname dbconn sql.sub_dbconn
func dbconn(db *sql.DB, ctx context.Context) (*sql.Conn, error) {
	return db.Conn(ctx)
}

//go:linkname DBConn sql.sub_dbconn
//go:noescape
func DBConn(db *sql.DB, ctx context.Context) (*sql.Conn, error)

//go:linkname nullboolvalue sql.sub_nullboolvalue
func nullboolvalue(n sql.NullBool) (driver.Value, error) {
	return n.Value()
}

//go:linkname NullBoolValue sql.sub_nullboolvalue
//go:noescape
func NullBoolValue(n sql.NullBool) (driver.Value, error)

//go:linkname nullint64value sql.sub_nullint64value
func nullint64value(n sql.NullInt64) (driver.Value, error) {
	return n.Value()
}

//go:linkname NullInt64Value sql.sub_nullint64value
//go:noescape
func NullInt64Value(n sql.NullInt64) (driver.Value, error)

//go:linkname columntypename sql.sub_columntypename
func columntypename(ci *sql.ColumnType) string {
	return ci.Name()
}

//go:linkname ColumnTypeName sql.sub_columntypename
//go:noescape
func ColumnTypeName(ci *sql.ColumnType) string

//go:linkname stmtclose sql.sub_stmtclose
func stmtclose(s *sql.Stmt) error {
	return s.Close()
}

//go:linkname StmtClose sql.sub_stmtclose
//go:noescape
func StmtClose(s *sql.Stmt) error

//go:linkname columntypedatabasetypename sql.sub_columntypedatabasetypename
func columntypedatabasetypename(ci *sql.ColumnType) string {
	return ci.DatabaseTypeName()
}

//go:linkname ColumnTypeDatabaseTypeName sql.sub_columntypedatabasetypename
//go:noescape
func ColumnTypeDatabaseTypeName(ci *sql.ColumnType) string

//go:linkname connpingcontext sql.sub_connpingcontext
func connpingcontext(c *sql.Conn, ctx context.Context) error {
	return c.PingContext(ctx)
}

//go:linkname ConnPingContext sql.sub_connpingcontext
//go:noescape
func ConnPingContext(c *sql.Conn, ctx context.Context) error

//go:linkname nullfloat64value sql.sub_nullfloat64value
func nullfloat64value(n sql.NullFloat64) (driver.Value, error) {
	return n.Value()
}

//go:linkname NullFloat64Value sql.sub_nullfloat64value
//go:noescape
func NullFloat64Value(n sql.NullFloat64) (driver.Value, error)

//go:linkname nullint32value sql.sub_nullint32value
func nullint32value(n sql.NullInt32) (driver.Value, error) {
	return n.Value()
}

//go:linkname NullInt32Value sql.sub_nullint32value
//go:noescape
func NullInt32Value(n sql.NullInt32) (driver.Value, error)

//go:linkname rowserr sql.sub_rowserr
func rowserr(rs *sql.Rows) error {
	return rs.Err()
}

//go:linkname RowsErr sql.sub_rowserr
//go:noescape
func RowsErr(rs *sql.Rows) error

//go:linkname connpreparecontext sql.sub_connpreparecontext
func connpreparecontext(c *sql.Conn, ctx context.Context, query string) (*sql.Stmt, error) {
	return c.PrepareContext(ctx, query)
}

//go:linkname ConnPrepareContext sql.sub_connpreparecontext
//go:noescape
func ConnPrepareContext(c *sql.Conn, ctx context.Context, query string) (*sql.Stmt, error)

//go:linkname nulltimevalue sql.sub_nulltimevalue
func nulltimevalue(n sql.NullTime) (driver.Value, error) {
	return n.Value()
}

//go:linkname NullTimeValue sql.sub_nulltimevalue
//go:noescape
func NullTimeValue(n sql.NullTime) (driver.Value, error)

//go:linkname txstmt sql.sub_txstmt
func txstmt(tx *sql.Tx, stmt *sql.Stmt) *sql.Stmt {
	return tx.Stmt(stmt)
}

//go:linkname TxStmt sql.sub_txstmt
//go:noescape
func TxStmt(tx *sql.Tx, stmt *sql.Stmt) *sql.Stmt

//go:linkname nullstringscan sql.sub_nullstringscan
func nullstringscan(ns *sql.NullString, value interface{}) error {
	return ns.Scan(value)
}

//go:linkname NullStringScan sql.sub_nullstringscan
//go:noescape
func NullStringScan(ns *sql.NullString, value interface{}) error

//go:linkname rowsclose sql.sub_rowsclose
func rowsclose(rs *sql.Rows) error {
	return rs.Close()
}

//go:linkname RowsClose sql.sub_rowsclose
//go:noescape
func RowsClose(rs *sql.Rows) error

//go:linkname rowsnextresultset sql.sub_rowsnextresultset
func rowsnextresultset(rs *sql.Rows) bool {
	return rs.NextResultSet()
}

//go:linkname RowsNextResultSet sql.sub_rowsnextresultset
//go:noescape
func RowsNextResultSet(rs *sql.Rows) bool

//go:linkname txquery sql.sub_txquery
func txquery(tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.Query(query, args)
}

//go:linkname TxQuery sql.sub_txquery
//go:noescape
func TxQuery(tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, error)

//go:linkname txqueryrowcontext sql.sub_txqueryrowcontext
func txqueryrowcontext(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) *sql.Row {
	return tx.QueryRowContext(ctx, query, args)
}

//go:linkname TxQueryRowContext sql.sub_txqueryrowcontext
//go:noescape
func TxQueryRowContext(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) *sql.Row

//go:linkname txexeccontext sql.sub_txexeccontext
func txexeccontext(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.ExecContext(ctx, query, args)
}

//go:linkname TxExecContext sql.sub_txexeccontext
//go:noescape
func TxExecContext(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) (sql.Result, error)

//go:linkname dbping sql.sub_dbping
func dbping(db *sql.DB) error {
	return db.Ping()
}

//go:linkname DBPing sql.sub_dbping
//go:noescape
func DBPing(db *sql.DB) error

//go:linkname dbpingcontext sql.sub_dbpingcontext
func dbpingcontext(db *sql.DB, ctx context.Context) error {
	return db.PingContext(ctx)
}

//go:linkname DBPingContext sql.sub_dbpingcontext
//go:noescape
func DBPingContext(db *sql.DB, ctx context.Context) error

//go:linkname dbqueryrow sql.sub_dbqueryrow
func dbqueryrow(db *sql.DB, query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args)
}

//go:linkname DBQueryRow sql.sub_dbqueryrow
//go:noescape
func DBQueryRow(db *sql.DB, query string, args ...interface{}) *sql.Row

//go:linkname Named database/sql.Named
//go:noescape
func Named(name string, value interface{}) sql.NamedArg

//go:linkname rowsnext sql.sub_rowsnext
func rowsnext(rs *sql.Rows) bool {
	return rs.Next()
}

//go:linkname RowsNext sql.sub_rowsnext
//go:noescape
func RowsNext(rs *sql.Rows) bool

//go:linkname Drivers database/sql.Drivers
//go:noescape
func Drivers() []string

//go:linkname dbbegin sql.sub_dbbegin
func dbbegin(db *sql.DB) (*sql.Tx, error) {
	return db.Begin()
}

//go:linkname DBBegin sql.sub_dbbegin
//go:noescape
func DBBegin(db *sql.DB) (*sql.Tx, error)

//go:linkname dbstats sql.sub_dbstats
func dbstats(db *sql.DB) sql.DBStats {
	return db.Stats()
}

//go:linkname DBStats sql.sub_dbstats
//go:noescape
func DBStats(db *sql.DB) sql.DBStats

//go:linkname stmtquery sql.sub_stmtquery
func stmtquery(s *sql.Stmt, args ...interface{}) (*sql.Rows, error) {
	return s.Query(args)
}

//go:linkname StmtQuery sql.sub_stmtquery
//go:noescape
func StmtQuery(s *sql.Stmt, args ...interface{}) (*sql.Rows, error)

//go:linkname stmtquerycontext sql.sub_stmtquerycontext
func stmtquerycontext(s *sql.Stmt, ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	return s.QueryContext(ctx, args)
}

//go:linkname StmtQueryContext sql.sub_stmtquerycontext
//go:noescape
func StmtQueryContext(s *sql.Stmt, ctx context.Context, args ...interface{}) (*sql.Rows, error)

//go:linkname stmtqueryrow sql.sub_stmtqueryrow
func stmtqueryrow(s *sql.Stmt, args ...interface{}) *sql.Row {
	return s.QueryRow(args)
}

//go:linkname StmtQueryRow sql.sub_stmtqueryrow
//go:noescape
func StmtQueryRow(s *sql.Stmt, args ...interface{}) *sql.Row

//go:linkname txrollback sql.sub_txrollback
func txrollback(tx *sql.Tx) error {
	return tx.Rollback()
}

//go:linkname TxRollback sql.sub_txrollback
//go:noescape
func TxRollback(tx *sql.Tx) error

//go:linkname connclose sql.sub_connclose
func connclose(c *sql.Conn) error {
	return c.Close()
}

//go:linkname ConnClose sql.sub_connclose
//go:noescape
func ConnClose(c *sql.Conn) error

//go:linkname connexeccontext sql.sub_connexeccontext
func connexeccontext(c *sql.Conn, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.ExecContext(ctx, query, args)
}

//go:linkname ConnExecContext sql.sub_connexeccontext
//go:noescape
func ConnExecContext(c *sql.Conn, ctx context.Context, query string, args ...interface{}) (sql.Result, error)

//go:linkname dbbegintx sql.sub_dbbegintx
func dbbegintx(db *sql.DB, ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.BeginTx(ctx, opts)
}

//go:linkname DBBeginTx sql.sub_dbbegintx
//go:noescape
func DBBeginTx(db *sql.DB, ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

//go:linkname dbquerycontext sql.sub_dbquerycontext
func dbquerycontext(db *sql.DB, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, args)
}

//go:linkname DBQueryContext sql.sub_dbquerycontext
//go:noescape
func DBQueryContext(db *sql.DB, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

//go:linkname nullboolscan sql.sub_nullboolscan
func nullboolscan(n *sql.NullBool, value interface{}) error {
	return n.Scan(value)
}

//go:linkname NullBoolScan sql.sub_nullboolscan
//go:noescape
func NullBoolScan(n *sql.NullBool, value interface{}) error

//go:linkname txqueryrow sql.sub_txqueryrow
func txqueryrow(tx *sql.Tx, query string, args ...interface{}) *sql.Row {
	return tx.QueryRow(query, args)
}

//go:linkname TxQueryRow sql.sub_txqueryrow
//go:noescape
func TxQueryRow(tx *sql.Tx, query string, args ...interface{}) *sql.Row

//go:linkname columntypescantype sql.sub_columntypescantype
func columntypescantype(ci *sql.ColumnType) reflect.Type {
	return ci.ScanType()
}

//go:linkname ColumnTypeScanType sql.sub_columntypescantype
//go:noescape
func ColumnTypeScanType(ci *sql.ColumnType) reflect.Type

//go:linkname OpenDB database/sql.OpenDB
//go:noescape
func OpenDB(c driver.Connector) *sql.DB

//go:linkname dbprepare sql.sub_dbprepare
func dbprepare(db *sql.DB, query string) (*sql.Stmt, error) {
	return db.Prepare(query)
}

//go:linkname DBPrepare sql.sub_dbprepare
//go:noescape
func DBPrepare(db *sql.DB, query string) (*sql.Stmt, error)

//go:linkname isolationlevelstring sql.sub_isolationlevelstring
func isolationlevelstring(i sql.IsolationLevel) string {
	return i.String()
}

//go:linkname IsolationLevelString sql.sub_isolationlevelstring
//go:noescape
func IsolationLevelString(i sql.IsolationLevel) string

//go:linkname txpreparecontext sql.sub_txpreparecontext
func txpreparecontext(tx *sql.Tx, ctx context.Context, query string) (*sql.Stmt, error) {
	return tx.PrepareContext(ctx, query)
}

//go:linkname TxPrepareContext sql.sub_txpreparecontext
//go:noescape
func TxPrepareContext(tx *sql.Tx, ctx context.Context, query string) (*sql.Stmt, error)

//go:linkname dbexeccontext sql.sub_dbexeccontext
func dbexeccontext(db *sql.DB, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.ExecContext(ctx, query, args)
}

//go:linkname DBExecContext sql.sub_dbexeccontext
//go:noescape
func DBExecContext(db *sql.DB, ctx context.Context, query string, args ...interface{}) (sql.Result, error)

//go:linkname nullint64scan sql.sub_nullint64scan
func nullint64scan(n *sql.NullInt64, value interface{}) error {
	return n.Scan(value)
}

//go:linkname NullInt64Scan sql.sub_nullint64scan
//go:noescape
func NullInt64Scan(n *sql.NullInt64, value interface{}) error

//go:linkname txcommit sql.sub_txcommit
func txcommit(tx *sql.Tx) error {
	return tx.Commit()
}

//go:linkname TxCommit sql.sub_txcommit
//go:noescape
func TxCommit(tx *sql.Tx) error

//go:linkname connbegintx sql.sub_connbegintx
func connbegintx(c *sql.Conn, ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return c.BeginTx(ctx, opts)
}

//go:linkname ConnBeginTx sql.sub_connbegintx
//go:noescape
func ConnBeginTx(c *sql.Conn, ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

//go:linkname connqueryrowcontext sql.sub_connqueryrowcontext
func connqueryrowcontext(c *sql.Conn, ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.QueryRowContext(ctx, query, args)
}

//go:linkname ConnQueryRowContext sql.sub_connqueryrowcontext
//go:noescape
func ConnQueryRowContext(c *sql.Conn, ctx context.Context, query string, args ...interface{}) *sql.Row

//go:linkname rowscolumntypes sql.sub_rowscolumntypes
func rowscolumntypes(rs *sql.Rows) ([]*sql.ColumnType, error) {
	return rs.ColumnTypes()
}

//go:linkname RowsColumnTypes sql.sub_rowscolumntypes
//go:noescape
func RowsColumnTypes(rs *sql.Rows) ([]*sql.ColumnType, error)

//go:linkname stmtexeccontext sql.sub_stmtexeccontext
func stmtexeccontext(s *sql.Stmt, ctx context.Context, args ...interface{}) (sql.Result, error) {
	return s.ExecContext(ctx, args)
}

//go:linkname StmtExecContext sql.sub_stmtexeccontext
//go:noescape
func StmtExecContext(s *sql.Stmt, ctx context.Context, args ...interface{}) (sql.Result, error)

//go:linkname rowsscan sql.sub_rowsscan
func rowsscan(rs *sql.Rows, dest ...interface{}) error {
	return rs.Scan(dest)
}

//go:linkname RowsScan sql.sub_rowsscan
//go:noescape
func RowsScan(rs *sql.Rows, dest ...interface{}) error

//go:linkname txquerycontext sql.sub_txquerycontext
func txquerycontext(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.QueryContext(ctx, query, args)
}

//go:linkname TxQueryContext sql.sub_txquerycontext
//go:noescape
func TxQueryContext(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

//go:linkname txstmtcontext sql.sub_txstmtcontext
func txstmtcontext(tx *sql.Tx, ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	return tx.StmtContext(ctx, stmt)
}

//go:linkname TxStmtContext sql.sub_txstmtcontext
//go:noescape
func TxStmtContext(tx *sql.Tx, ctx context.Context, stmt *sql.Stmt) *sql.Stmt

//go:linkname connquerycontext sql.sub_connquerycontext
func connquerycontext(c *sql.Conn, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.QueryContext(ctx, query, args)
}

//go:linkname ConnQueryContext sql.sub_connquerycontext
//go:noescape
func ConnQueryContext(c *sql.Conn, ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

//go:linkname connraw sql.sub_connraw
func connraw(c *sql.Conn, f func(driverConn interface{}) error) error {
	return c.Raw(f)
}

//go:linkname ConnRaw sql.sub_connraw
//go:noescape
func ConnRaw(c *sql.Conn, f func(driverConn interface{}) error) error

//go:linkname dbexec sql.sub_dbexec
func dbexec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args)
}

//go:linkname DBExec sql.sub_dbexec
//go:noescape
func DBExec(db *sql.DB, query string, args ...interface{}) (sql.Result, error)

//go:linkname rowscan sql.sub_rowscan
func rowscan(r *sql.Row, dest ...interface{}) error {
	return r.Scan(dest)
}

//go:linkname RowScan sql.sub_rowscan
//go:noescape
func RowScan(r *sql.Row, dest ...interface{}) error

//go:linkname rowscolumns sql.sub_rowscolumns
func rowscolumns(rs *sql.Rows) ([]string, error) {
	return rs.Columns()
}

//go:linkname RowsColumns sql.sub_rowscolumns
//go:noescape
func RowsColumns(rs *sql.Rows) ([]string, error)

//go:linkname columntypelength sql.sub_columntypelength
func columntypelength(ci *sql.ColumnType) (int64, bool) {
	return ci.Length()
}

//go:linkname ColumnTypeLength sql.sub_columntypelength
//go:noescape
func ColumnTypeLength(ci *sql.ColumnType) (int64, bool)

//go:linkname columntypenullable sql.sub_columntypenullable
func columntypenullable(ci *sql.ColumnType) bool {
	return ci.Nullable()
}

//go:linkname ColumnTypeNullable sql.sub_columntypenullable
//go:noescape
func ColumnTypeNullable(ci *sql.ColumnType) bool

//go:linkname nullfloat64scan sql.sub_nullfloat64scan
func nullfloat64scan(n *sql.NullFloat64, value interface{}) error {
	return n.Scan(value)
}

//go:linkname NullFloat64Scan sql.sub_nullfloat64scan
//go:noescape
func NullFloat64Scan(n *sql.NullFloat64, value interface{}) error

//go:linkname nullstringvalue sql.sub_nullstringvalue
func nullstringvalue(ns sql.NullString) (driver.Value, error) {
	return ns.Value()
}

//go:linkname NullStringValue sql.sub_nullstringvalue
//go:noescape
func NullStringValue(ns sql.NullString) (driver.Value, error)

//go:linkname stmtexec sql.sub_stmtexec
func stmtexec(s *sql.Stmt, args ...interface{}) (sql.Result, error) {
	return s.Exec(args)
}

//go:linkname StmtExec sql.sub_stmtexec
//go:noescape
func StmtExec(s *sql.Stmt, args ...interface{}) (sql.Result, error)

//go:linkname txexec sql.sub_txexec
func txexec(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return tx.Exec(query, args)
}

//go:linkname TxExec sql.sub_txexec
//go:noescape
func TxExec(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error)

//go:linkname txprepare sql.sub_txprepare
func txprepare(tx *sql.Tx, query string) (*sql.Stmt, error) {
	return tx.Prepare(query)
}

//go:linkname TxPrepare sql.sub_txprepare
//go:noescape
func TxPrepare(tx *sql.Tx, query string) (*sql.Stmt, error)

//go:linkname Open database/sql.Open
//go:noescape
func Open(driverName, dataSourceName string) (*sql.DB, error)

//go:linkname dbdriver sql.sub_dbdriver
func dbdriver(db *sql.DB) driver.Driver {
	return db.Driver()
}

//go:linkname DBDriver sql.sub_dbdriver
//go:noescape
func DBDriver(db *sql.DB) driver.Driver

//go:linkname dbquery sql.sub_dbquery
func dbquery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args)
}

//go:linkname DBQuery sql.sub_dbquery
//go:noescape
func DBQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error)

//go:linkname dbqueryrowcontext sql.sub_dbqueryrowcontext
func dbqueryrowcontext(db *sql.DB, ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.QueryRowContext(ctx, query, args)
}

//go:linkname DBQueryRowContext sql.sub_dbqueryrowcontext
//go:noescape
func DBQueryRowContext(db *sql.DB, ctx context.Context, query string, args ...interface{}) *sql.Row

//go:linkname nullint32scan sql.sub_nullint32scan
func nullint32scan(n *sql.NullInt32, value interface{}) error {
	return n.Scan(value)
}

//go:linkname NullInt32Scan sql.sub_nullint32scan
//go:noescape
func NullInt32Scan(n *sql.NullInt32, value interface{}) error

//go:linkname dbpreparecontext sql.sub_dbpreparecontext
func dbpreparecontext(db *sql.DB, ctx context.Context, query string) (*sql.Stmt, error) {
	return db.PrepareContext(ctx, query)
}

//go:linkname DBPrepareContext sql.sub_dbpreparecontext
//go:noescape
func DBPrepareContext(db *sql.DB, ctx context.Context, query string) (*sql.Stmt, error)
