// This file has automatically been generated on Wed Feb 26 15:50:53 +05 2020.
// DO NOT EDIT.
package iotest

import (
	"io"
	_ "testing/iotest"
	_ "unsafe"
)

//go:linkname TimeoutReader testing/iotest.TimeoutReader
//go:noescape
func TimeoutReader(r io.Reader) io.Reader

//go:linkname TruncateWriter testing/iotest.TruncateWriter
//go:noescape
func TruncateWriter(w io.Writer, n int64) io.Writer

//go:linkname DataErrReader testing/iotest.DataErrReader
//go:noescape
func DataErrReader(r io.Reader) io.Reader

//go:linkname HalfReader testing/iotest.HalfReader
//go:noescape
func HalfReader(r io.Reader) io.Reader

//go:linkname NewReadLogger testing/iotest.NewReadLogger
//go:noescape
func NewReadLogger(prefix string, r io.Reader) io.Reader

//go:linkname NewWriteLogger testing/iotest.NewWriteLogger
//go:noescape
func NewWriteLogger(prefix string, w io.Writer) io.Writer

//go:linkname OneByteReader testing/iotest.OneByteReader
//go:noescape
func OneByteReader(r io.Reader) io.Reader
