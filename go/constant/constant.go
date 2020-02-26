// This file has automatically been generated on Wed Feb 26 15:50:33 +05 2020.
// DO NOT EDIT.
package constant

import (
	"go/constant"
	"go/token"
	_ "unsafe"
)

//go:linkname StringVal go/constant.StringVal
//go:noescape
func StringVal(x constant.Value,) string

//go:linkname BinaryOp go/constant.BinaryOp
//go:noescape
func BinaryOp(x_ constant.Value, op token.Token, y_ constant.Value,) constant.Value

//go:linkname Denom go/constant.Denom
//go:noescape
func Denom(x constant.Value,) constant.Value

//go:linkname Make go/constant.Make
//go:noescape
func Make(x interface{}) constant.Value

//go:linkname MakeFromLiteral go/constant.MakeFromLiteral
//go:noescape
func MakeFromLiteral(lit string, tok token.Token, zero uint) constant.Value

//go:linkname MakeUnknown go/constant.MakeUnknown
//go:noescape
func MakeUnknown() constant.Value

//go:linkname Compare go/constant.Compare
//go:noescape
func Compare(x_ constant.Value, op token.Token, y_ constant.Value,) bool

//go:linkname Val go/constant.Val
//go:noescape
func Val(x constant.Value,) interface{}

//go:linkname Num go/constant.Num
//go:noescape
func Num(x constant.Value,) constant.Value

//go:linkname Real go/constant.Real
//go:noescape
func Real(x constant.Value,) constant.Value

//go:linkname Shift go/constant.Shift
//go:noescape
func Shift(x constant.Value, op token.Token, s uint) constant.Value

//go:linkname ToFloat go/constant.ToFloat
//go:noescape
func ToFloat(x constant.Value,) constant.Value

//go:linkname BitLen go/constant.BitLen
//go:noescape
func BitLen(x constant.Value,) int

//go:linkname Imag go/constant.Imag
//go:noescape
func Imag(x constant.Value,) constant.Value

//go:linkname MakeBool go/constant.MakeBool
//go:noescape
func MakeBool(b bool) constant.Value

//go:linkname MakeFromBytes go/constant.MakeFromBytes
//go:noescape
func MakeFromBytes(bytes []byte) constant.Value

//go:linkname ToInt go/constant.ToInt
//go:noescape
func ToInt(x constant.Value,) constant.Value

//go:linkname UnaryOp go/constant.UnaryOp
//go:noescape
func UnaryOp(op token.Token, y constant.Value, prec uint) constant.Value

//go:linkname BoolVal go/constant.BoolVal
//go:noescape
func BoolVal(x constant.Value,) bool

//go:linkname Sign go/constant.Sign
//go:noescape
func Sign(x constant.Value,) int

//go:linkname MakeImag go/constant.MakeImag
//go:noescape
func MakeImag(x constant.Value,) constant.Value

//go:linkname MakeString go/constant.MakeString
//go:noescape
func MakeString(s string) constant.Value

//go:linkname ToComplex go/constant.ToComplex
//go:noescape
func ToComplex(x constant.Value,) constant.Value

//go:linkname Bytes go/constant.Bytes
//go:noescape
func Bytes(x constant.Value,) []byte
