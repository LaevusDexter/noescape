// This file has automatically been generated on Wed Feb 26 02:10:17 +05 2020.
// DO NOT EDIT.
package utf16

import (
	_ "unicode/utf16"
	_ "unsafe"
)

//go:linkname Decode unicode/utf16.Decode
//go:noescape
func Decode(s []uint16) []rune

//go:linkname DecodeRune unicode/utf16.DecodeRune
//go:noescape
func DecodeRune(r1, r2 rune) rune

//go:linkname Encode unicode/utf16.Encode
//go:noescape
func Encode(s []rune) []uint16

//go:linkname EncodeRune unicode/utf16.EncodeRune
//go:noescape
func EncodeRune(r rune) rune

//go:linkname IsSurrogate unicode/utf16.IsSurrogate
//go:noescape
func IsSurrogate(r rune) bool
