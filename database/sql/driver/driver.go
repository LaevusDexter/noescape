// This file has automatically been generated on Wed Feb 26 15:50:26 +05 2020.
// DO NOT EDIT.
package driver

import (
	"database/sql/driver"
	_ "unsafe"
)

//go:linkname IsScanValue database/sql/driver.IsScanValue
//go:noescape
func IsScanValue(v interface{}) bool

//go:linkname IsValue database/sql/driver.IsValue
//go:noescape
func IsValue(v interface{}) bool

//go:linkname notnullconvertvalue driver.sub_notnullconvertvalue
func notnullconvertvalue(n driver.NotNull, v interface{}) (driver.Value, error) {
	return n.ConvertValue(v)
}

//go:linkname NotNullConvertValue driver.sub_notnullconvertvalue
//go:noescape
func NotNullConvertValue(n driver.NotNull, v interface{}) (driver.Value, error)

//go:linkname nullconvertvalue driver.sub_nullconvertvalue
func nullconvertvalue(n driver.Null, v interface{}) (driver.Value, error) {
	return n.ConvertValue(v)
}

//go:linkname NullConvertValue driver.sub_nullconvertvalue
//go:noescape
func NullConvertValue(n driver.Null, v interface{}) (driver.Value, error)

//go:linkname rowsaffectedlastinsertid driver.sub_rowsaffectedlastinsertid
func rowsaffectedlastinsertid(receiverV driver.RowsAffected,) (int64, error) {
	return receiverV.LastInsertId()
}

//go:linkname RowsAffectedLastInsertId driver.sub_rowsaffectedlastinsertid
//go:noescape
func RowsAffectedLastInsertId(receiverV driver.RowsAffected,) (int64, error)

//go:linkname rowsaffectedrowsaffected driver.sub_rowsaffectedrowsaffected
func rowsaffectedrowsaffected(v driver.RowsAffected,) (int64, error) {
	return v.RowsAffected()
}

//go:linkname RowsAffectedRowsAffected driver.sub_rowsaffectedrowsaffected
//go:noescape
func RowsAffectedRowsAffected(v driver.RowsAffected,) (int64, error)
