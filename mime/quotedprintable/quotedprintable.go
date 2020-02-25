// This file has automatically been generated on Wed Feb 26 02:10:04 +05 2020.
// DO NOT EDIT.
package quotedprintable

import (
	"io"
	"mime/quotedprintable"
	_ "unsafe"
)

//go:linkname NewReader mime/quotedprintable.NewReader
//go:noescape
func NewReader(r io.Reader) *quotedprintable.Reader

//go:linkname readerread quotedprintable.sub_readerread
func readerread(r *quotedprintable.Reader, p []byte) (int, error) {
	return r.Read(p)
}

//go:linkname ReaderRead quotedprintable.sub_readerread
//go:noescape
func ReaderRead(r *quotedprintable.Reader, p []byte) (int, error)

//go:linkname NewWriter mime/quotedprintable.NewWriter
//go:noescape
func NewWriter(w io.Writer) *quotedprintable.Writer

//go:linkname writerclose quotedprintable.sub_writerclose
func writerclose(w *quotedprintable.Writer) error {
	return w.Close()
}

//go:linkname WriterClose quotedprintable.sub_writerclose
//go:noescape
func WriterClose(w *quotedprintable.Writer) error

//go:linkname writerwrite quotedprintable.sub_writerwrite
func writerwrite(w *quotedprintable.Writer, p []byte) (int, error) {
	return w.Write(p)
}

//go:linkname WriterWrite quotedprintable.sub_writerwrite
//go:noescape
func WriterWrite(w *quotedprintable.Writer, p []byte) (int, error)
