// This file has automatically been generated on Wed Feb 26 15:50:28 +05 2020.
// DO NOT EDIT.
package ascii85

import (
	"encoding/ascii85"
	"io"
	_ "unsafe"
)

//go:linkname NewDecoder encoding/ascii85.NewDecoder
//go:noescape
func NewDecoder(r io.Reader) io.Reader

//go:linkname NewEncoder encoding/ascii85.NewEncoder
//go:noescape
func NewEncoder(w io.Writer) io.WriteCloser

//go:linkname corruptinputerrorerror ascii85.sub_corruptinputerrorerror
func corruptinputerrorerror(e ascii85.CorruptInputError,) string {
	return e.Error()
}

//go:linkname CorruptInputErrorError ascii85.sub_corruptinputerrorerror
//go:noescape
func CorruptInputErrorError(e ascii85.CorruptInputError,) string

//go:linkname Decode encoding/ascii85.Decode
//go:noescape
func Decode(dst, src []byte, flush bool) (int, int, error)

//go:linkname Encode encoding/ascii85.Encode
//go:noescape
func Encode(dst, src []byte) int

//go:linkname MaxEncodedLen encoding/ascii85.MaxEncodedLen
//go:noescape
func MaxEncodedLen(n int) int
