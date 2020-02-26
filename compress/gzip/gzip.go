// This file has automatically been generated on Wed Feb 26 15:50:18 +05 2020.
// DO NOT EDIT.
package gzip

import (
	"compress/gzip"
	"io"
	_ "unsafe"
)

//go:linkname readerclose gzip.sub_readerclose
func readerclose(z *gzip.Reader,) error {
	return z.Close()
}

//go:linkname ReaderClose gzip.sub_readerclose
//go:noescape
func ReaderClose(z *gzip.Reader,) error

//go:linkname readerreset gzip.sub_readerreset
func readerreset(z *gzip.Reader, r io.Reader) error {
	return z.Reset(r)
}

//go:linkname ReaderReset gzip.sub_readerreset
//go:noescape
func ReaderReset(z *gzip.Reader, r io.Reader) error

//go:linkname NewWriterLevel compress/gzip.NewWriterLevel
//go:noescape
func NewWriterLevel(w io.Writer, level int) (*gzip.Writer, error)

//go:linkname writerclose gzip.sub_writerclose
func writerclose(z *gzip.Writer,) error {
	return z.Close()
}

//go:linkname WriterClose gzip.sub_writerclose
//go:noescape
func WriterClose(z *gzip.Writer,) error

//go:linkname writerflush gzip.sub_writerflush
func writerflush(z *gzip.Writer,) error {
	return z.Flush()
}

//go:linkname WriterFlush gzip.sub_writerflush
//go:noescape
func WriterFlush(z *gzip.Writer,) error

//go:linkname writerwrite gzip.sub_writerwrite
func writerwrite(z *gzip.Writer, p []byte) (int, error) {
	return z.Write(p)
}

//go:linkname WriterWrite gzip.sub_writerwrite
//go:noescape
func WriterWrite(z *gzip.Writer, p []byte) (int, error)

//go:linkname NewReader compress/gzip.NewReader
//go:noescape
func NewReader(r io.Reader) (*gzip.Reader, error)

//go:linkname readerread gzip.sub_readerread
func readerread(z *gzip.Reader, p []byte) (int, error) {
	return z.Read(p)
}

//go:linkname ReaderRead gzip.sub_readerread
//go:noescape
func ReaderRead(z *gzip.Reader, p []byte) (int, error)

//go:linkname NewWriter compress/gzip.NewWriter
//go:noescape
func NewWriter(w io.Writer) *gzip.Writer
