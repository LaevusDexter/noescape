// This file has automatically been generated on Wed Feb 26 02:09:45 +05 2020.
// DO NOT EDIT.
package elliptic

import (
	"crypto/elliptic"
	"io"
	"math/big"
	_ "unsafe"
)

//go:linkname curveparamsisoncurve elliptic.sub_curveparamsisoncurve
func curveparamsisoncurve(curve *elliptic.CurveParams, x, y *big.Int) bool {
	return curve.IsOnCurve(x, y)
}

//go:linkname CurveParamsIsOnCurve elliptic.sub_curveparamsisoncurve
//go:noescape
func CurveParamsIsOnCurve(curve *elliptic.CurveParams, x, y *big.Int) bool

//go:linkname curveparamsscalarmult elliptic.sub_curveparamsscalarmult
func curveparamsscalarmult(curve *elliptic.CurveParams, Bx, By *big.Int, k []byte) (*big.Int, *big.Int) {
	return curve.ScalarMult(Bx, By, k)
}

//go:linkname CurveParamsScalarMult elliptic.sub_curveparamsscalarmult
//go:noescape
func CurveParamsScalarMult(curve *elliptic.CurveParams, Bx, By *big.Int, k []byte) (*big.Int, *big.Int)

//go:linkname GenerateKey crypto/elliptic.GenerateKey
//go:noescape
func GenerateKey(curve elliptic.Curve, rand io.Reader) ([]byte, *big.Int, error)

//go:linkname curveparamsdouble elliptic.sub_curveparamsdouble
func curveparamsdouble(curve *elliptic.CurveParams, x1, y1 *big.Int) (*big.Int, *big.Int) {
	return curve.Double(x1, y1)
}

//go:linkname CurveParamsDouble elliptic.sub_curveparamsdouble
//go:noescape
func CurveParamsDouble(curve *elliptic.CurveParams, x1, y1 *big.Int) (*big.Int, *big.Int)

//go:linkname curveparamsadd elliptic.sub_curveparamsadd
func curveparamsadd(curve *elliptic.CurveParams, x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int) {
	return curve.Add(x1, y1, x2, y2)
}

//go:linkname CurveParamsAdd elliptic.sub_curveparamsadd
//go:noescape
func CurveParamsAdd(curve *elliptic.CurveParams, x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)

//go:linkname curveparamsparams elliptic.sub_curveparamsparams
func curveparamsparams(curve *elliptic.CurveParams) *elliptic.CurveParams {
	return curve.Params()
}

//go:linkname CurveParamsParams elliptic.sub_curveparamsparams
//go:noescape
func CurveParamsParams(curve *elliptic.CurveParams) *elliptic.CurveParams

//go:linkname curveparamsscalarbasemult elliptic.sub_curveparamsscalarbasemult
func curveparamsscalarbasemult(curve *elliptic.CurveParams, k []byte) (*big.Int, *big.Int) {
	return curve.ScalarBaseMult(k)
}

//go:linkname CurveParamsScalarBaseMult elliptic.sub_curveparamsscalarbasemult
//go:noescape
func CurveParamsScalarBaseMult(curve *elliptic.CurveParams, k []byte) (*big.Int, *big.Int)

//go:linkname Marshal crypto/elliptic.Marshal
//go:noescape
func Marshal(curve elliptic.Curve, x, y *big.Int) []byte

//go:linkname Unmarshal crypto/elliptic.Unmarshal
//go:noescape
func Unmarshal(curve elliptic.Curve, data []byte) *big.Int
