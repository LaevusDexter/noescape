// This file has automatically been generated on Wed Feb 26 02:09:52 +05 2020.
// DO NOT EDIT.
package hex

import (
	"encoding/hex"
	"io"
	_ "unsafe"
)

//go:linkname Dump encoding/hex.Dump
//go:noescape
func Dump(data []byte) string

//go:linkname Dumper encoding/hex.Dumper
//go:noescape
func Dumper(w io.Writer) io.WriteCloser

//go:linkname Encode encoding/hex.Encode
//go:noescape
func Encode(dst, src []byte) int

//go:linkname NewDecoder encoding/hex.NewDecoder
//go:noescape
func NewDecoder(r io.Reader) io.Reader

//go:linkname NewEncoder encoding/hex.NewEncoder
//go:noescape
func NewEncoder(w io.Writer) io.Writer

//go:linkname invalidbyteerrorerror hex.sub_invalidbyteerrorerror
func invalidbyteerrorerror(e hex.InvalidByteError) string {
	return e.Error()
}

//go:linkname InvalidByteErrorError hex.sub_invalidbyteerrorerror
//go:noescape
func InvalidByteErrorError(e hex.InvalidByteError) string

//go:linkname Decode encoding/hex.Decode
//go:noescape
func Decode(dst, src []byte) (int, error)

//go:linkname DecodeString encoding/hex.DecodeString
//go:noescape
func DecodeString(s string) ([]byte, error)

//go:linkname DecodedLen encoding/hex.DecodedLen
//go:noescape
func DecodedLen(x int) int

//go:linkname EncodeToString encoding/hex.EncodeToString
//go:noescape
func EncodeToString(src []byte) string

//go:linkname EncodedLen encoding/hex.EncodedLen
//go:noescape
func EncodedLen(n int) int
