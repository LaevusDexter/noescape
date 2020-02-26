// This file has automatically been generated on Wed Feb 26 15:50:56 +05 2020.
// DO NOT EDIT.
package utf8

import (
	_ "unicode/utf8"
	_ "unsafe"
)

//go:linkname RuneCount unicode/utf8.RuneCount
//go:noescape
func RuneCount(p []byte) int

//go:linkname RuneCountInString unicode/utf8.RuneCountInString
//go:noescape
func RuneCountInString(s string) int

//go:linkname FullRuneInString unicode/utf8.FullRuneInString
//go:noescape
func FullRuneInString(s string) bool

//go:linkname ValidRune unicode/utf8.ValidRune
//go:noescape
func ValidRune(r rune) bool

//go:linkname DecodeLastRuneInString unicode/utf8.DecodeLastRuneInString
//go:noescape
func DecodeLastRuneInString(s string) (rune, int)

//go:linkname FullRune unicode/utf8.FullRune
//go:noescape
func FullRune(p []byte) bool

//go:linkname DecodeRuneInString unicode/utf8.DecodeRuneInString
//go:noescape
func DecodeRuneInString(s string) (rune, int)

//go:linkname RuneLen unicode/utf8.RuneLen
//go:noescape
func RuneLen(r rune) int

//go:linkname Valid unicode/utf8.Valid
//go:noescape
func Valid(p []byte) bool

//go:linkname DecodeLastRune unicode/utf8.DecodeLastRune
//go:noescape
func DecodeLastRune(p []byte) (rune, int)

//go:linkname DecodeRune unicode/utf8.DecodeRune
//go:noescape
func DecodeRune(p []byte) (rune, int)

//go:linkname ValidString unicode/utf8.ValidString
//go:noescape
func ValidString(s string) bool

//go:linkname EncodeRune unicode/utf8.EncodeRune
//go:noescape
func EncodeRune(p []byte, r rune) int

//go:linkname RuneStart unicode/utf8.RuneStart
//go:noescape
func RuneStart(b byte) bool
