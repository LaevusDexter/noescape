// This file has automatically been generated on Wed Feb 26 15:50:18 +05 2020.
// DO NOT EDIT.
package flate

import (
	"compress/flate"
	"io"
	_ "unsafe"
)

//go:linkname NewReader compress/flate.NewReader
//go:noescape
func NewReader(r io.Reader) io.ReadCloser

//go:linkname corruptinputerrorerror flate.sub_corruptinputerrorerror
func corruptinputerrorerror(e flate.CorruptInputError,) string {
	return e.Error()
}

//go:linkname CorruptInputErrorError flate.sub_corruptinputerrorerror
//go:noescape
func CorruptInputErrorError(e flate.CorruptInputError,) string

//go:linkname readerrorerror flate.sub_readerrorerror
func readerrorerror(e *flate.ReadError,) string {
	return e.Error()
}

//go:linkname ReadErrorError flate.sub_readerrorerror
//go:noescape
func ReadErrorError(e *flate.ReadError,) string

//go:linkname NewWriter compress/flate.NewWriter
//go:noescape
func NewWriter(w io.Writer, level int) (*flate.Writer, error)

//go:linkname writerclose flate.sub_writerclose
func writerclose(w *flate.Writer,) error {
	return w.Close()
}

//go:linkname WriterClose flate.sub_writerclose
//go:noescape
func WriterClose(w *flate.Writer,) error

//go:linkname writerflush flate.sub_writerflush
func writerflush(w *flate.Writer,) error {
	return w.Flush()
}

//go:linkname WriterFlush flate.sub_writerflush
//go:noescape
func WriterFlush(w *flate.Writer,) error

//go:linkname NewReaderDict compress/flate.NewReaderDict
//go:noescape
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser

//go:linkname internalerrorerror flate.sub_internalerrorerror
func internalerrorerror(e flate.InternalError,) string {
	return e.Error()
}

//go:linkname InternalErrorError flate.sub_internalerrorerror
//go:noescape
func InternalErrorError(e flate.InternalError,) string

//go:linkname writeerrorerror flate.sub_writeerrorerror
func writeerrorerror(e *flate.WriteError,) string {
	return e.Error()
}

//go:linkname WriteErrorError flate.sub_writeerrorerror
//go:noescape
func WriteErrorError(e *flate.WriteError,) string

//go:linkname NewWriterDict compress/flate.NewWriterDict
//go:noescape
func NewWriterDict(w io.Writer, level int, dict []byte) (*flate.Writer, error)

//go:linkname writerwrite flate.sub_writerwrite
func writerwrite(w *flate.Writer, data []byte) (int, error) {
	return w.Write(data)
}

//go:linkname WriterWrite flate.sub_writerwrite
//go:noescape
func WriterWrite(w *flate.Writer, data []byte) (int, error)
