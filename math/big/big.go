// This file has automatically been generated on Wed Feb 26 02:10:03 +05 2020.
// DO NOT EDIT.
package big

import (
	"fmt"
	"math/big"
	"math/rand"
	_ "unsafe"
)

//go:linkname floatstring big.sub_floatstring
func floatstring(x *big.Float) string {
	return x.String()
}

//go:linkname FloatString big.sub_floatstring
//go:noescape
func FloatString(x *big.Float) string

//go:linkname ratsetfrac big.sub_ratsetfrac
func ratsetfrac(z *big.Rat, a, b *big.Int) *big.Rat {
	return z.SetFrac(a, b)
}

//go:linkname RatSetFrac big.sub_ratsetfrac
//go:noescape
func RatSetFrac(z *big.Rat, a, b *big.Int) *big.Rat

//go:linkname floatprec big.sub_floatprec
func floatprec(x *big.Float) uint {
	return x.Prec()
}

//go:linkname FloatPrec big.sub_floatprec
//go:noescape
func FloatPrec(x *big.Float) uint

//go:linkname floatadd big.sub_floatadd
func floatadd(z *big.Float, x, y *big.Float) *big.Float {
	return z.Add(x, y)
}

//go:linkname FloatAdd big.sub_floatadd
//go:noescape
func FloatAdd(z *big.Float, x, y *big.Float) *big.Float

//go:linkname floatmul big.sub_floatmul
func floatmul(z *big.Float, x, y *big.Float) *big.Float {
	return z.Mul(x, y)
}

//go:linkname FloatMul big.sub_floatmul
//go:noescape
func FloatMul(z *big.Float, x, y *big.Float) *big.Float

//go:linkname floatabs big.sub_floatabs
func floatabs(z *big.Float, x *big.Float) *big.Float {
	return z.Abs(x)
}

//go:linkname FloatAbs big.sub_floatabs
//go:noescape
func FloatAbs(z *big.Float, x *big.Float) *big.Float

//go:linkname floatcopy big.sub_floatcopy
func floatcopy(z *big.Float, x *big.Float) *big.Float {
	return z.Copy(x)
}

//go:linkname FloatCopy big.sub_floatcopy
//go:noescape
func FloatCopy(z *big.Float, x *big.Float) *big.Float

//go:linkname inttrailingzerobits big.sub_inttrailingzerobits
func inttrailingzerobits(x *big.Int) uint {
	return x.TrailingZeroBits()
}

//go:linkname IntTrailingZeroBits big.sub_inttrailingzerobits
//go:noescape
func IntTrailingZeroBits(x *big.Int) uint

//go:linkname ratmarshaltext big.sub_ratmarshaltext
func ratmarshaltext(x *big.Rat) ([]byte, error) {
	return x.MarshalText()
}

//go:linkname RatMarshalText big.sub_ratmarshaltext
//go:noescape
func RatMarshalText(x *big.Rat) ([]byte, error)

//go:linkname ratmul big.sub_ratmul
func ratmul(z *big.Rat, x, y *big.Rat) *big.Rat {
	return z.Mul(x, y)
}

//go:linkname RatMul big.sub_ratmul
//go:noescape
func RatMul(z *big.Rat, x, y *big.Rat) *big.Rat

//go:linkname floatacc big.sub_floatacc
func floatacc(x *big.Float) big.Accuracy {
	return x.Acc()
}

//go:linkname FloatAcc big.sub_floatacc
//go:noescape
func FloatAcc(x *big.Float) big.Accuracy

//go:linkname floatappend big.sub_floatappend
func floatappend(x *big.Float, buf []byte, fmt byte, prec int) []byte {
	return x.Append(buf, fmt, prec)
}

//go:linkname FloatAppend big.sub_floatappend
//go:noescape
func FloatAppend(x *big.Float, buf []byte, fmt byte, prec int) []byte

//go:linkname intor big.sub_intor
func intor(z *big.Int, x, y *big.Int) *big.Int {
	return z.Or(x, y)
}

//go:linkname IntOr big.sub_intor
//go:noescape
func IntOr(z *big.Int, x, y *big.Int) *big.Int

//go:linkname ratfloatstring big.sub_ratfloatstring
func ratfloatstring(x *big.Rat, prec int) string {
	return x.FloatString(prec)
}

//go:linkname RatFloatString big.sub_ratfloatstring
//go:noescape
func RatFloatString(x *big.Rat, prec int) string

//go:linkname floatsign big.sub_floatsign
func floatsign(x *big.Float) int {
	return x.Sign()
}

//go:linkname FloatSign big.sub_floatsign
//go:noescape
func FloatSign(x *big.Float) int

//go:linkname intlsh big.sub_intlsh
func intlsh(z *big.Int, x *big.Int, n uint) *big.Int {
	return z.Lsh(x, n)
}

//go:linkname IntLsh big.sub_intlsh
//go:noescape
func IntLsh(z *big.Int, x *big.Int, n uint) *big.Int

//go:linkname intunmarshaljson big.sub_intunmarshaljson
func intunmarshaljson(z *big.Int, text []byte) error {
	return z.UnmarshalJSON(text)
}

//go:linkname IntUnmarshalJSON big.sub_intunmarshaljson
//go:noescape
func IntUnmarshalJSON(z *big.Int, text []byte) error

//go:linkname ratadd big.sub_ratadd
func ratadd(z *big.Rat, x, y *big.Rat) *big.Rat {
	return z.Add(x, y)
}

//go:linkname RatAdd big.sub_ratadd
//go:noescape
func RatAdd(z *big.Rat, x, y *big.Rat) *big.Rat

//go:linkname ratsign big.sub_ratsign
func ratsign(x *big.Rat) int {
	return x.Sign()
}

//go:linkname RatSign big.sub_ratsign
//go:noescape
func RatSign(x *big.Rat) int

//go:linkname floatgobdecode big.sub_floatgobdecode
func floatgobdecode(z *big.Float, buf []byte) error {
	return z.GobDecode(buf)
}

//go:linkname FloatGobDecode big.sub_floatgobdecode
//go:noescape
func FloatGobDecode(z *big.Float, buf []byte) error

//go:linkname intstring big.sub_intstring
func intstring(x *big.Int) string {
	return x.String()
}

//go:linkname IntString big.sub_intstring
//go:noescape
func IntString(x *big.Int) string

//go:linkname floatcmp big.sub_floatcmp
func floatcmp(x *big.Float, y *big.Float) int {
	return x.Cmp(y)
}

//go:linkname FloatCmp big.sub_floatcmp
//go:noescape
func FloatCmp(x *big.Float, y *big.Float) int

//go:linkname intbit big.sub_intbit
func intbit(x *big.Int, i int) uint {
	return x.Bit(i)
}

//go:linkname IntBit big.sub_intbit
//go:noescape
func IntBit(x *big.Int, i int) uint

//go:linkname intgobencode big.sub_intgobencode
func intgobencode(x *big.Int) ([]byte, error) {
	return x.GobEncode()
}

//go:linkname IntGobEncode big.sub_intgobencode
//go:noescape
func IntGobEncode(x *big.Int) ([]byte, error)

//go:linkname intmulrange big.sub_intmulrange
func intmulrange(z *big.Int, a, b int64) *big.Int {
	return z.MulRange(a, b)
}

//go:linkname IntMulRange big.sub_intmulrange
//go:noescape
func IntMulRange(z *big.Int, a, b int64) *big.Int

//go:linkname intquorem big.sub_intquorem
func intquorem(z *big.Int, x, y, r *big.Int) (*big.Int, *big.Int) {
	return z.QuoRem(x, y, r)
}

//go:linkname IntQuoRem big.sub_intquorem
//go:noescape
func IntQuoRem(z *big.Int, x, y, r *big.Int) (*big.Int, *big.Int)

//go:linkname intxor big.sub_intxor
func intxor(z *big.Int, x, y *big.Int) *big.Int {
	return z.Xor(x, y)
}

//go:linkname IntXor big.sub_intxor
//go:noescape
func IntXor(z *big.Int, x, y *big.Int) *big.Int

//go:linkname ratset big.sub_ratset
func ratset(z *big.Rat, x *big.Rat) *big.Rat {
	return z.Set(x)
}

//go:linkname RatSet big.sub_ratset
//go:noescape
func RatSet(z *big.Rat, x *big.Rat) *big.Rat

//go:linkname ratunmarshaltext big.sub_ratunmarshaltext
func ratunmarshaltext(z *big.Rat, text []byte) error {
	return z.UnmarshalText(text)
}

//go:linkname RatUnmarshalText big.sub_ratunmarshaltext
//go:noescape
func RatUnmarshalText(z *big.Rat, text []byte) error

//go:linkname ParseFloat math/big.ParseFloat
//go:noescape
func ParseFloat(s string, base int, prec uint, mode big.RoundingMode) (*big.Float, int, error)

//go:linkname floatscan big.sub_floatscan
func floatscan(z *big.Float, s fmt.ScanState, ch rune) error {
	return z.Scan(s, ch)
}

//go:linkname FloatScan big.sub_floatscan
//go:noescape
func FloatScan(z *big.Float, s fmt.ScanState, ch rune) error

//go:linkname intneg big.sub_intneg
func intneg(z *big.Int, x *big.Int) *big.Int {
	return z.Neg(x)
}

//go:linkname IntNeg big.sub_intneg
//go:noescape
func IntNeg(z *big.Int, x *big.Int) *big.Int

//go:linkname ratnum big.sub_ratnum
func ratnum(x *big.Rat) *big.Int {
	return x.Num()
}

//go:linkname RatNum big.sub_ratnum
//go:noescape
func RatNum(x *big.Rat) *big.Int

//go:linkname floatquo big.sub_floatquo
func floatquo(z *big.Float, x, y *big.Float) *big.Float {
	return z.Quo(x, y)
}

//go:linkname FloatQuo big.sub_floatquo
//go:noescape
func FloatQuo(z *big.Float, x, y *big.Float) *big.Float

//go:linkname floattext big.sub_floattext
func floattext(x *big.Float, format byte, prec int) string {
	return x.Text(format, prec)
}

//go:linkname FloatText big.sub_floattext
//go:noescape
func FloatText(x *big.Float, format byte, prec int) string

//go:linkname intsetbit big.sub_intsetbit
func intsetbit(z *big.Int, x *big.Int, i int, b uint) *big.Int {
	return z.SetBit(x, i, b)
}

//go:linkname IntSetBit big.sub_intsetbit
//go:noescape
func IntSetBit(z *big.Int, x *big.Int, i int, b uint) *big.Int

//go:linkname ratcmp big.sub_ratcmp
func ratcmp(x *big.Rat, y *big.Rat) int {
	return x.Cmp(y)
}

//go:linkname RatCmp big.sub_ratcmp
//go:noescape
func RatCmp(x *big.Rat, y *big.Rat) int

//go:linkname floatisint big.sub_floatisint
func floatisint(x *big.Float) bool {
	return x.IsInt()
}

//go:linkname FloatIsInt big.sub_floatisint
//go:noescape
func FloatIsInt(x *big.Float) bool

//go:linkname intgcd big.sub_intgcd
func intgcd(z *big.Int, x, y, a, b *big.Int) *big.Int {
	return z.GCD(x, y, a, b)
}

//go:linkname IntGCD big.sub_intgcd
//go:noescape
func IntGCD(z *big.Int, x, y, a, b *big.Int) *big.Int

//go:linkname intmodinverse big.sub_intmodinverse
func intmodinverse(z *big.Int, g, n *big.Int) *big.Int {
	return z.ModInverse(g, n)
}

//go:linkname IntModInverse big.sub_intmodinverse
//go:noescape
func IntModInverse(z *big.Int, g, n *big.Int) *big.Int

//go:linkname floatgobencode big.sub_floatgobencode
func floatgobencode(x *big.Float) ([]byte, error) {
	return x.GobEncode()
}

//go:linkname FloatGobEncode big.sub_floatgobencode
//go:noescape
func FloatGobEncode(x *big.Float) ([]byte, error)

//go:linkname intdiv big.sub_intdiv
func intdiv(z *big.Int, x, y *big.Int) *big.Int {
	return z.Div(x, y)
}

//go:linkname IntDiv big.sub_intdiv
//go:noescape
func IntDiv(z *big.Int, x, y *big.Int) *big.Int

//go:linkname intbits big.sub_intbits
func intbits(x *big.Int) []big.Word {
	return x.Bits()
}

//go:linkname IntBits big.sub_intbits
//go:noescape
func IntBits(x *big.Int) []big.Word

//go:linkname intexp big.sub_intexp
func intexp(z *big.Int, x, y, m *big.Int) *big.Int {
	return z.Exp(x, y, m)
}

//go:linkname IntExp big.sub_intexp
//go:noescape
func IntExp(z *big.Int, x, y, m *big.Int) *big.Int

//go:linkname intmod big.sub_intmod
func intmod(z *big.Int, x, y *big.Int) *big.Int {
	return z.Mod(x, y)
}

//go:linkname IntMod big.sub_intmod
//go:noescape
func IntMod(z *big.Int, x, y *big.Int) *big.Int

//go:linkname intrand big.sub_intrand
func intrand(z *big.Int, rnd *rand.Rand, n *big.Int) *big.Int {
	return z.Rand(rnd, n)
}

//go:linkname IntRand big.sub_intrand
//go:noescape
func IntRand(z *big.Int, rnd *rand.Rand, n *big.Int) *big.Int

//go:linkname floatunmarshaltext big.sub_floatunmarshaltext
func floatunmarshaltext(z *big.Float, text []byte) error {
	return z.UnmarshalText(text)
}

//go:linkname FloatUnmarshalText big.sub_floatunmarshaltext
//go:noescape
func FloatUnmarshalText(z *big.Float, text []byte) error

//go:linkname intadd big.sub_intadd
func intadd(z *big.Int, x, y *big.Int) *big.Int {
	return z.Add(x, y)
}

//go:linkname IntAdd big.sub_intadd
//go:noescape
func IntAdd(z *big.Int, x, y *big.Int) *big.Int

//go:linkname intsub big.sub_intsub
func intsub(z *big.Int, x, y *big.Int) *big.Int {
	return z.Sub(x, y)
}

//go:linkname IntSub big.sub_intsub
//go:noescape
func IntSub(z *big.Int, x, y *big.Int) *big.Int

//go:linkname ratdenom big.sub_ratdenom
func ratdenom(x *big.Rat) *big.Int {
	return x.Denom()
}

//go:linkname RatDenom big.sub_ratdenom
//go:noescape
func RatDenom(x *big.Rat) *big.Int

//go:linkname floatsetinf big.sub_floatsetinf
func floatsetinf(z *big.Float, signbit bool) *big.Float {
	return z.SetInf(signbit)
}

//go:linkname FloatSetInf big.sub_floatsetinf
//go:noescape
func FloatSetInf(z *big.Float, signbit bool) *big.Float

//go:linkname intcmp big.sub_intcmp
func intcmp(x *big.Int, y *big.Int) int {
	return x.Cmp(y)
}

//go:linkname IntCmp big.sub_intcmp
//go:noescape
func IntCmp(x *big.Int, y *big.Int) int

//go:linkname intdivmod big.sub_intdivmod
func intdivmod(z *big.Int, x, y, m *big.Int) (*big.Int, *big.Int) {
	return z.DivMod(x, y, m)
}

//go:linkname IntDivMod big.sub_intdivmod
//go:noescape
func IntDivMod(z *big.Int, x, y, m *big.Int) (*big.Int, *big.Int)

//go:linkname intmarshaltext big.sub_intmarshaltext
func intmarshaltext(x *big.Int) ([]byte, error) {
	return x.MarshalText()
}

//go:linkname IntMarshalText big.sub_intmarshaltext
//go:noescape
func IntMarshalText(x *big.Int) ([]byte, error)

//go:linkname floatsetstring big.sub_floatsetstring
func floatsetstring(z *big.Float, s string) (*big.Float, bool) {
	return z.SetString(s)
}

//go:linkname FloatSetString big.sub_floatsetstring
//go:noescape
func FloatSetString(z *big.Float, s string) (*big.Float, bool)

//go:linkname intcmpabs big.sub_intcmpabs
func intcmpabs(x *big.Int, y *big.Int) int {
	return x.CmpAbs(y)
}

//go:linkname IntCmpAbs big.sub_intcmpabs
//go:noescape
func IntCmpAbs(x *big.Int, y *big.Int) int

//go:linkname intand big.sub_intand
func intand(z *big.Int, x, y *big.Int) *big.Int {
	return z.And(x, y)
}

//go:linkname IntAnd big.sub_intand
//go:noescape
func IntAnd(z *big.Int, x, y *big.Int) *big.Int

//go:linkname intquo big.sub_intquo
func intquo(z *big.Int, x, y *big.Int) *big.Int {
	return z.Quo(x, y)
}

//go:linkname IntQuo big.sub_intquo
//go:noescape
func IntQuo(z *big.Int, x, y *big.Int) *big.Int

//go:linkname intappend big.sub_intappend
func intappend(x *big.Int, buf []byte, base int) []byte {
	return x.Append(buf, base)
}

//go:linkname IntAppend big.sub_intappend
//go:noescape
func IntAppend(x *big.Int, buf []byte, base int) []byte

//go:linkname NewFloat math/big.NewFloat
//go:noescape
func NewFloat(x float64) *big.Float

//go:linkname floatmarshaltext big.sub_floatmarshaltext
func floatmarshaltext(x *big.Float) ([]byte, error) {
	return x.MarshalText()
}

//go:linkname FloatMarshalText big.sub_floatmarshaltext
//go:noescape
func FloatMarshalText(x *big.Float) ([]byte, error)

//go:linkname intrem big.sub_intrem
func intrem(z *big.Int, x, y *big.Int) *big.Int {
	return z.Rem(x, y)
}

//go:linkname IntRem big.sub_intrem
//go:noescape
func IntRem(z *big.Int, x, y *big.Int) *big.Int

//go:linkname intsqrt big.sub_intsqrt
func intsqrt(z *big.Int, x *big.Int) *big.Int {
	return z.Sqrt(x)
}

//go:linkname IntSqrt big.sub_intsqrt
//go:noescape
func IntSqrt(z *big.Int, x *big.Int) *big.Int

//go:linkname ratscan big.sub_ratscan
func ratscan(z *big.Rat, s fmt.ScanState, ch rune) error {
	return z.Scan(s, ch)
}

//go:linkname RatScan big.sub_ratscan
//go:noescape
func RatScan(z *big.Rat, s fmt.ScanState, ch rune) error

//go:linkname errnanerror big.sub_errnanerror
func errnanerror(err big.ErrNaN) string {
	return err.Error()
}

//go:linkname ErrNaNError big.sub_errnanerror
//go:noescape
func ErrNaNError(err big.ErrNaN) string

//go:linkname intbytes big.sub_intbytes
func intbytes(x *big.Int) []byte {
	return x.Bytes()
}

//go:linkname IntBytes big.sub_intbytes
//go:noescape
func IntBytes(x *big.Int) []byte

//go:linkname intmarshaljson big.sub_intmarshaljson
func intmarshaljson(x *big.Int) ([]byte, error) {
	return x.MarshalJSON()
}

//go:linkname IntMarshalJSON big.sub_intmarshaljson
//go:noescape
func IntMarshalJSON(x *big.Int) ([]byte, error)

//go:linkname intnot big.sub_intnot
func intnot(z *big.Int, x *big.Int) *big.Int {
	return z.Not(x)
}

//go:linkname IntNot big.sub_intnot
//go:noescape
func IntNot(z *big.Int, x *big.Int) *big.Int

//go:linkname intrsh big.sub_intrsh
func intrsh(z *big.Int, x *big.Int, n uint) *big.Int {
	return z.Rsh(x, n)
}

//go:linkname IntRsh big.sub_intrsh
//go:noescape
func IntRsh(z *big.Int, x *big.Int, n uint) *big.Int

//go:linkname intscan big.sub_intscan
func intscan(z *big.Int, s fmt.ScanState, ch rune) error {
	return z.Scan(s, ch)
}

//go:linkname IntScan big.sub_intscan
//go:noescape
func IntScan(z *big.Int, s fmt.ScanState, ch rune) error

//go:linkname inttext big.sub_inttext
func inttext(x *big.Int, base int) string {
	return x.Text(base)
}

//go:linkname IntText big.sub_inttext
//go:noescape
func IntText(x *big.Int, base int) string

//go:linkname intunmarshaltext big.sub_intunmarshaltext
func intunmarshaltext(z *big.Int, text []byte) error {
	return z.UnmarshalText(text)
}

//go:linkname IntUnmarshalText big.sub_intunmarshaltext
//go:noescape
func IntUnmarshalText(z *big.Int, text []byte) error

//go:linkname floatmode big.sub_floatmode
func floatmode(x *big.Float) big.RoundingMode {
	return x.Mode()
}

//go:linkname FloatMode big.sub_floatmode
//go:noescape
func FloatMode(x *big.Float) big.RoundingMode

//go:linkname floatsetint big.sub_floatsetint
func floatsetint(z *big.Float, x *big.Int) *big.Float {
	return z.SetInt(x)
}

//go:linkname FloatSetInt big.sub_floatsetint
//go:noescape
func FloatSetInt(z *big.Float, x *big.Int) *big.Float

//go:linkname ratgobdecode big.sub_ratgobdecode
func ratgobdecode(z *big.Rat, buf []byte) error {
	return z.GobDecode(buf)
}

//go:linkname RatGobDecode big.sub_ratgobdecode
//go:noescape
func RatGobDecode(z *big.Rat, buf []byte) error

//go:linkname ratsetstring big.sub_ratsetstring
func ratsetstring(z *big.Rat, s string) (*big.Rat, bool) {
	return z.SetString(s)
}

//go:linkname RatSetString big.sub_ratsetstring
//go:noescape
func RatSetString(z *big.Rat, s string) (*big.Rat, bool)

//go:linkname ratinv big.sub_ratinv
func ratinv(z *big.Rat, x *big.Rat) *big.Rat {
	return z.Inv(x)
}

//go:linkname RatInv big.sub_ratinv
//go:noescape
func RatInv(z *big.Rat, x *big.Rat) *big.Rat

//go:linkname roundingmodestring big.sub_roundingmodestring
func roundingmodestring(i big.RoundingMode) string {
	return i.String()
}

//go:linkname RoundingModeString big.sub_roundingmodestring
//go:noescape
func RoundingModeString(i big.RoundingMode) string

//go:linkname floatparse big.sub_floatparse
func floatparse(z *big.Float, s string, base int) (*big.Float, int, error) {
	return z.Parse(s, base)
}

//go:linkname FloatParse big.sub_floatparse
//go:noescape
func FloatParse(z *big.Float, s string, base int) (*big.Float, int, error)

//go:linkname floatsub big.sub_floatsub
func floatsub(z *big.Float, x, y *big.Float) *big.Float {
	return z.Sub(x, y)
}

//go:linkname FloatSub big.sub_floatsub
//go:noescape
func FloatSub(z *big.Float, x, y *big.Float) *big.Float

//go:linkname intabs big.sub_intabs
func intabs(z *big.Int, x *big.Int) *big.Int {
	return z.Abs(x)
}

//go:linkname IntAbs big.sub_intabs
//go:noescape
func IntAbs(z *big.Int, x *big.Int) *big.Int

//go:linkname intsetstring big.sub_intsetstring
func intsetstring(z *big.Int, s string, base int) (*big.Int, bool) {
	return z.SetString(s, base)
}

//go:linkname IntSetString big.sub_intsetstring
//go:noescape
func IntSetString(z *big.Int, s string, base int) (*big.Int, bool)

//go:linkname floatneg big.sub_floatneg
func floatneg(z *big.Float, x *big.Float) *big.Float {
	return z.Neg(x)
}

//go:linkname FloatNeg big.sub_floatneg
//go:noescape
func FloatNeg(z *big.Float, x *big.Float) *big.Float

//go:linkname floatsetprec big.sub_floatsetprec
func floatsetprec(z *big.Float, prec uint) *big.Float {
	return z.SetPrec(prec)
}

//go:linkname FloatSetPrec big.sub_floatsetprec
//go:noescape
func FloatSetPrec(z *big.Float, prec uint) *big.Float

//go:linkname floatisinf big.sub_floatisinf
func floatisinf(x *big.Float) bool {
	return x.IsInf()
}

//go:linkname FloatIsInf big.sub_floatisinf
//go:noescape
func FloatIsInf(x *big.Float) bool

//go:linkname floatmantexp big.sub_floatmantexp
func floatmantexp(x *big.Float, mant *big.Float) int {
	return x.MantExp(mant)
}

//go:linkname FloatMantExp big.sub_floatmantexp
//go:noescape
func FloatMantExp(x *big.Float, mant *big.Float) int

//go:linkname ratabs big.sub_ratabs
func ratabs(z *big.Rat, x *big.Rat) *big.Rat {
	return z.Abs(x)
}

//go:linkname RatAbs big.sub_ratabs
//go:noescape
func RatAbs(z *big.Rat, x *big.Rat) *big.Rat

//go:linkname floatrat big.sub_floatrat
func floatrat(x *big.Float, z *big.Rat) (*big.Rat, big.Accuracy) {
	return x.Rat(z)
}

//go:linkname FloatRat big.sub_floatrat
//go:noescape
func FloatRat(x *big.Float, z *big.Rat) (*big.Rat, big.Accuracy)

//go:linkname intandnot big.sub_intandnot
func intandnot(z *big.Int, x, y *big.Int) *big.Int {
	return z.AndNot(x, y)
}

//go:linkname IntAndNot big.sub_intandnot
//go:noescape
func IntAndNot(z *big.Int, x, y *big.Int) *big.Int

//go:linkname ratsub big.sub_ratsub
func ratsub(z *big.Rat, x, y *big.Rat) *big.Rat {
	return z.Sub(x, y)
}

//go:linkname RatSub big.sub_ratsub
//go:noescape
func RatSub(z *big.Rat, x, y *big.Rat) *big.Rat

//go:linkname floatset big.sub_floatset
func floatset(z *big.Float, x *big.Float) *big.Float {
	return z.Set(x)
}

//go:linkname FloatSet big.sub_floatset
//go:noescape
func FloatSet(z *big.Float, x *big.Float) *big.Float

//go:linkname ratsetint big.sub_ratsetint
func ratsetint(z *big.Rat, x *big.Int) *big.Rat {
	return z.SetInt(x)
}

//go:linkname RatSetInt big.sub_ratsetint
//go:noescape
func RatSetInt(z *big.Rat, x *big.Int) *big.Rat

//go:linkname intmodsqrt big.sub_intmodsqrt
func intmodsqrt(z *big.Int, x, p *big.Int) *big.Int {
	return z.ModSqrt(x, p)
}

//go:linkname IntModSqrt big.sub_intmodsqrt
//go:noescape
func IntModSqrt(z *big.Int, x, p *big.Int) *big.Int

//go:linkname intsetbits big.sub_intsetbits
func intsetbits(z *big.Int, abs []big.Word) *big.Int {
	return z.SetBits(abs)
}

//go:linkname IntSetBits big.sub_intsetbits
//go:noescape
func IntSetBits(z *big.Int, abs []big.Word) *big.Int

//go:linkname ratquo big.sub_ratquo
func ratquo(z *big.Rat, x, y *big.Rat) *big.Rat {
	return z.Quo(x, y)
}

//go:linkname RatQuo big.sub_ratquo
//go:noescape
func RatQuo(z *big.Rat, x, y *big.Rat) *big.Rat

//go:linkname Jacobi math/big.Jacobi
//go:noescape
func Jacobi(x, y *big.Int) int

//go:linkname intbitlen big.sub_intbitlen
func intbitlen(x *big.Int) int {
	return x.BitLen()
}

//go:linkname IntBitLen big.sub_intbitlen
//go:noescape
func IntBitLen(x *big.Int) int

//go:linkname intsetbytes big.sub_intsetbytes
func intsetbytes(z *big.Int, buf []byte) *big.Int {
	return z.SetBytes(buf)
}

//go:linkname IntSetBytes big.sub_intsetbytes
//go:noescape
func IntSetBytes(z *big.Int, buf []byte) *big.Int

//go:linkname accuracystring big.sub_accuracystring
func accuracystring(i big.Accuracy) string {
	return i.String()
}

//go:linkname AccuracyString big.sub_accuracystring
//go:noescape
func AccuracyString(i big.Accuracy) string

//go:linkname intmul big.sub_intmul
func intmul(z *big.Int, x, y *big.Int) *big.Int {
	return z.Mul(x, y)
}

//go:linkname IntMul big.sub_intmul
//go:noescape
func IntMul(z *big.Int, x, y *big.Int) *big.Int

//go:linkname ratisint big.sub_ratisint
func ratisint(x *big.Rat) bool {
	return x.IsInt()
}

//go:linkname RatIsInt big.sub_ratisint
//go:noescape
func RatIsInt(x *big.Rat) bool

//go:linkname ratgobencode big.sub_ratgobencode
func ratgobencode(x *big.Rat) ([]byte, error) {
	return x.GobEncode()
}

//go:linkname RatGobEncode big.sub_ratgobencode
//go:noescape
func RatGobEncode(x *big.Rat) ([]byte, error)

//go:linkname floatsetrat big.sub_floatsetrat
func floatsetrat(z *big.Float, x *big.Rat) *big.Float {
	return z.SetRat(x)
}

//go:linkname FloatSetRat big.sub_floatsetrat
//go:noescape
func FloatSetRat(z *big.Float, x *big.Rat) *big.Float

//go:linkname intgobdecode big.sub_intgobdecode
func intgobdecode(z *big.Int, buf []byte) error {
	return z.GobDecode(buf)
}

//go:linkname IntGobDecode big.sub_intgobdecode
//go:noescape
func IntGobDecode(z *big.Int, buf []byte) error

//go:linkname intbinomial big.sub_intbinomial
func intbinomial(z *big.Int, n, k int64) *big.Int {
	return z.Binomial(n, k)
}

//go:linkname IntBinomial big.sub_intbinomial
//go:noescape
func IntBinomial(z *big.Int, n, k int64) *big.Int

//go:linkname floatminprec big.sub_floatminprec
func floatminprec(x *big.Float) uint {
	return x.MinPrec()
}

//go:linkname FloatMinPrec big.sub_floatminprec
//go:noescape
func FloatMinPrec(x *big.Float) uint

//go:linkname floatsignbit big.sub_floatsignbit
func floatsignbit(x *big.Float) bool {
	return x.Signbit()
}

//go:linkname FloatSignbit big.sub_floatsignbit
//go:noescape
func FloatSignbit(x *big.Float) bool

//go:linkname intprobablyprime big.sub_intprobablyprime
func intprobablyprime(x *big.Int, n int) bool {
	return x.ProbablyPrime(n)
}

//go:linkname IntProbablyPrime big.sub_intprobablyprime
//go:noescape
func IntProbablyPrime(x *big.Int, n int) bool

//go:linkname ratstring big.sub_ratstring
func ratstring(x *big.Rat) string {
	return x.String()
}

//go:linkname RatString big.sub_ratstring
//go:noescape
func RatString(x *big.Rat) string

//go:linkname floatint big.sub_floatint
func floatint(x *big.Float, z *big.Int) (*big.Int, big.Accuracy) {
	return x.Int(z)
}

//go:linkname FloatInt big.sub_floatint
//go:noescape
func FloatInt(x *big.Float, z *big.Int) (*big.Int, big.Accuracy)

//go:linkname NewInt math/big.NewInt
//go:noescape
func NewInt(x int64) *big.Int

//go:linkname intset big.sub_intset
func intset(z *big.Int, x *big.Int) *big.Int {
	return z.Set(x)
}

//go:linkname IntSet big.sub_intset
//go:noescape
func IntSet(z *big.Int, x *big.Int) *big.Int

//go:linkname intsign big.sub_intsign
func intsign(x *big.Int) int {
	return x.Sign()
}

//go:linkname IntSign big.sub_intsign
//go:noescape
func IntSign(x *big.Int) int

//go:linkname ratneg big.sub_ratneg
func ratneg(z *big.Rat, x *big.Rat) *big.Rat {
	return z.Neg(x)
}

//go:linkname RatNeg big.sub_ratneg
//go:noescape
func RatNeg(z *big.Rat, x *big.Rat) *big.Rat

//go:linkname floatsetmantexp big.sub_floatsetmantexp
func floatsetmantexp(z *big.Float, mant *big.Float, exp int) *big.Float {
	return z.SetMantExp(mant, exp)
}

//go:linkname FloatSetMantExp big.sub_floatsetmantexp
//go:noescape
func FloatSetMantExp(z *big.Float, mant *big.Float, exp int) *big.Float

//go:linkname floatsqrt big.sub_floatsqrt
func floatsqrt(z *big.Float, x *big.Float) *big.Float {
	return z.Sqrt(x)
}

//go:linkname FloatSqrt big.sub_floatsqrt
//go:noescape
func FloatSqrt(z *big.Float, x *big.Float) *big.Float

//go:linkname ratratstring big.sub_ratratstring
func ratratstring(x *big.Rat) string {
	return x.RatString()
}

//go:linkname RatRatString big.sub_ratratstring
//go:noescape
func RatRatString(x *big.Rat) string

//go:linkname floatsetmode big.sub_floatsetmode
func floatsetmode(z *big.Float, mode big.RoundingMode) *big.Float {
	return z.SetMode(mode)
}

//go:linkname FloatSetMode big.sub_floatsetmode
//go:noescape
func FloatSetMode(z *big.Float, mode big.RoundingMode) *big.Float

//go:linkname NewRat math/big.NewRat
//go:noescape
func NewRat(a, b int64) *big.Rat
