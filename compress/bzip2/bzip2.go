// This file has automatically been generated on Wed Feb 26 15:50:18 +05 2020.
// DO NOT EDIT.
package bzip2

import (
	"compress/bzip2"
	"io"
	_ "unsafe"
)

//go:linkname NewReader compress/bzip2.NewReader
//go:noescape
func NewReader(r io.Reader) io.Reader

//go:linkname structuralerrorerror bzip2.sub_structuralerrorerror
func structuralerrorerror(s bzip2.StructuralError,) string {
	return s.Error()
}

//go:linkname StructuralErrorError bzip2.sub_structuralerrorerror
//go:noescape
func StructuralErrorError(s bzip2.StructuralError,) string
