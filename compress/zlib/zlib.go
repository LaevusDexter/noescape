// This file has automatically been generated on Wed Feb 26 02:09:42 +05 2020.
// DO NOT EDIT.
package zlib

import (
	"compress/zlib"
	"io"
	_ "unsafe"
)

//go:linkname NewReader compress/zlib.NewReader
//go:noescape
func NewReader(r io.Reader) (io.ReadCloser, error)

//go:linkname NewReaderDict compress/zlib.NewReaderDict
//go:noescape
func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)

//go:linkname NewWriter compress/zlib.NewWriter
//go:noescape
func NewWriter(w io.Writer) *zlib.Writer

//go:linkname NewWriterLevel compress/zlib.NewWriterLevel
//go:noescape
func NewWriterLevel(w io.Writer, level int) (*zlib.Writer, error)

//go:linkname NewWriterLevelDict compress/zlib.NewWriterLevelDict
//go:noescape
func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*zlib.Writer, error)

//go:linkname writerclose zlib.sub_writerclose
func writerclose(z *zlib.Writer) error {
	return z.Close()
}

//go:linkname WriterClose zlib.sub_writerclose
//go:noescape
func WriterClose(z *zlib.Writer) error

//go:linkname writerflush zlib.sub_writerflush
func writerflush(z *zlib.Writer) error {
	return z.Flush()
}

//go:linkname WriterFlush zlib.sub_writerflush
//go:noescape
func WriterFlush(z *zlib.Writer) error

//go:linkname writerwrite zlib.sub_writerwrite
func writerwrite(z *zlib.Writer, p []byte) (int, error) {
	return z.Write(p)
}

//go:linkname WriterWrite zlib.sub_writerwrite
//go:noescape
func WriterWrite(z *zlib.Writer, p []byte) (int, error)
