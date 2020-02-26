// This file has automatically been generated on Wed Feb 26 15:50:19 +05 2020.
// DO NOT EDIT.
package lzw

import (
	"compress/lzw"
	"io"
	_ "unsafe"
)

//go:linkname NewReader compress/lzw.NewReader
//go:noescape
func NewReader(r io.Reader, order lzw.Order, litWidth int) io.ReadCloser

//go:linkname NewWriter compress/lzw.NewWriter
//go:noescape
func NewWriter(w io.Writer, order lzw.Order, litWidth int) io.WriteCloser
