// This file has automatically been generated on Wed Feb 26 02:10:03 +05 2020.
// DO NOT EDIT.
package bits

import (
	_ "math/bits"
	_ "unsafe"
)

//go:linkname Rem math/bits.Rem
//go:noescape
func Rem(hi, lo, y uint) uint

//go:linkname ReverseBytes math/bits.ReverseBytes
//go:noescape
func ReverseBytes(x uint) uint

//go:linkname Add math/bits.Add
//go:noescape
func Add(x, y, carry uint) uint

//go:linkname Div math/bits.Div
//go:noescape
func Div(hi, lo, y uint) uint

//go:linkname LeadingZeros math/bits.LeadingZeros
//go:noescape
func LeadingZeros(x uint) int

//go:linkname Len math/bits.Len
//go:noescape
func Len(x uint) int

//go:linkname Sub math/bits.Sub
//go:noescape
func Sub(x, y, borrow uint) uint

//go:linkname TrailingZeros math/bits.TrailingZeros
//go:noescape
func TrailingZeros(x uint) int

//go:linkname Mul math/bits.Mul
//go:noescape
func Mul(x, y uint) uint

//go:linkname OnesCount math/bits.OnesCount
//go:noescape
func OnesCount(x uint) int

//go:linkname Reverse math/bits.Reverse
//go:noescape
func Reverse(x uint) uint

//go:linkname RotateLeft math/bits.RotateLeft
//go:noescape
func RotateLeft(x uint, k int) uint
