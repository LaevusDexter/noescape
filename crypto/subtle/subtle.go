// This file has automatically been generated on Wed Feb 26 15:50:24 +05 2020.
// DO NOT EDIT.
package subtle

import (
	_ "crypto/subtle"
	_ "unsafe"
)

//go:linkname ConstantTimeByteEq crypto/subtle.ConstantTimeByteEq
//go:noescape
func ConstantTimeByteEq(x, y uint8) int

//go:linkname ConstantTimeCompare crypto/subtle.ConstantTimeCompare
//go:noescape
func ConstantTimeCompare(x, y []byte) int

//go:linkname ConstantTimeEq crypto/subtle.ConstantTimeEq
//go:noescape
func ConstantTimeEq(x, y int32) int

//go:linkname ConstantTimeLessOrEq crypto/subtle.ConstantTimeLessOrEq
//go:noescape
func ConstantTimeLessOrEq(x, y int) int

//go:linkname ConstantTimeSelect crypto/subtle.ConstantTimeSelect
//go:noescape
func ConstantTimeSelect(v, x, y int) int
