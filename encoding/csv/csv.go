// This file has automatically been generated on Wed Feb 26 02:09:52 +05 2020.
// DO NOT EDIT.
package csv

import (
	"encoding/csv"
	"io"
	_ "unsafe"
)

//go:linkname parseerrorerror csv.sub_parseerrorerror
func parseerrorerror(e *csv.ParseError) string {
	return e.Error()
}

//go:linkname ParseErrorError csv.sub_parseerrorerror
//go:noescape
func ParseErrorError(e *csv.ParseError) string

//go:linkname readerread csv.sub_readerread
func readerread(r *csv.Reader) ([]string, error) {
	return r.Read()
}

//go:linkname ReaderRead csv.sub_readerread
//go:noescape
func ReaderRead(r *csv.Reader) ([]string, error)

//go:linkname readerreadall csv.sub_readerreadall
func readerreadall(r *csv.Reader) ([][]string, error) {
	return r.ReadAll()
}

//go:linkname ReaderReadAll csv.sub_readerreadall
//go:noescape
func ReaderReadAll(r *csv.Reader) ([][]string, error)

//go:linkname NewWriter encoding/csv.NewWriter
//go:noescape
func NewWriter(w io.Writer) *csv.Writer

//go:linkname writerwrite csv.sub_writerwrite
func writerwrite(w *csv.Writer, record []string) error {
	return w.Write(record)
}

//go:linkname WriterWrite csv.sub_writerwrite
//go:noescape
func WriterWrite(w *csv.Writer, record []string) error

//go:linkname parseerrorunwrap csv.sub_parseerrorunwrap
func parseerrorunwrap(e *csv.ParseError) error {
	return e.Unwrap()
}

//go:linkname ParseErrorUnwrap csv.sub_parseerrorunwrap
//go:noescape
func ParseErrorUnwrap(e *csv.ParseError) error

//go:linkname NewReader encoding/csv.NewReader
//go:noescape
func NewReader(r io.Reader) *csv.Reader

//go:linkname writererror csv.sub_writererror
func writererror(w *csv.Writer) error {
	return w.Error()
}

//go:linkname WriterError csv.sub_writererror
//go:noescape
func WriterError(w *csv.Writer) error

//go:linkname writerwriteall csv.sub_writerwriteall
func writerwriteall(w *csv.Writer, records [][]string) error {
	return w.WriteAll(records)
}

//go:linkname WriterWriteAll csv.sub_writerwriteall
//go:noescape
func WriterWriteAll(w *csv.Writer, records [][]string) error
