// This file has automatically been generated on Wed Feb 26 15:50:39 +05 2020.
// DO NOT EDIT.
package io

import (
	"io"
	_ "unsafe"
)

//go:linkname limitedreaderread io.sub_limitedreaderread
func limitedreaderread(l *io.LimitedReader, p []byte) (int, error) {
	return l.Read(p)
}

//go:linkname LimitedReaderRead io.sub_limitedreaderread
//go:noescape
func LimitedReaderRead(l *io.LimitedReader, p []byte) (int, error)

//go:linkname pipereaderclose io.sub_pipereaderclose
func pipereaderclose(r *io.PipeReader,) error {
	return r.Close()
}

//go:linkname PipeReaderClose io.sub_pipereaderclose
//go:noescape
func PipeReaderClose(r *io.PipeReader,) error

//go:linkname pipereaderread io.sub_pipereaderread
func pipereaderread(r *io.PipeReader, data []byte) (int, error) {
	return r.Read(data)
}

//go:linkname PipeReaderRead io.sub_pipereaderread
//go:noescape
func PipeReaderRead(r *io.PipeReader, data []byte) (int, error)

//go:linkname MultiReader io.MultiReader
//go:noescape
func MultiReader(readers ...io.Reader,) io.Reader

//go:linkname sectionreaderread io.sub_sectionreaderread
func sectionreaderread(s *io.SectionReader, p []byte) (int, error) {
	return s.Read(p)
}

//go:linkname SectionReaderRead io.sub_sectionreaderread
//go:noescape
func SectionReaderRead(s *io.SectionReader, p []byte) (int, error)

//go:linkname CopyBuffer io.CopyBuffer
//go:noescape
func CopyBuffer(dst io.Writer, src io.Reader, buf []byte) (int64, error)

//go:linkname CopyN io.CopyN
//go:noescape
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)

//go:linkname ReadAtLeast io.ReadAtLeast
//go:noescape
func ReadAtLeast(r io.Reader, buf []byte, min int) (int, error)

//go:linkname sectionreadersize io.sub_sectionreadersize
func sectionreadersize(s *io.SectionReader,) int64 {
	return s.Size()
}

//go:linkname SectionReaderSize io.sub_sectionreadersize
//go:noescape
func SectionReaderSize(s *io.SectionReader,) int64

//go:linkname Pipe io.Pipe
//go:noescape
func Pipe() (*io.PipeReader, *io.PipeWriter,)

//go:linkname pipewriterwrite io.sub_pipewriterwrite
func pipewriterwrite(w *io.PipeWriter, data []byte) (int, error) {
	return w.Write(data)
}

//go:linkname PipeWriterWrite io.sub_pipewriterwrite
//go:noescape
func PipeWriterWrite(w *io.PipeWriter, data []byte) (int, error)

//go:linkname sectionreaderseek io.sub_sectionreaderseek
func sectionreaderseek(s *io.SectionReader, offset int64, whence int) (int64, error) {
	return s.Seek(offset, whence)
}

//go:linkname SectionReaderSeek io.sub_sectionreaderseek
//go:noescape
func SectionReaderSeek(s *io.SectionReader, offset int64, whence int) (int64, error)

//go:linkname NewSectionReader io.NewSectionReader
//go:noescape
func NewSectionReader(r io.ReaderAt, off int64, n int64) *io.SectionReader

//go:linkname MultiWriter io.MultiWriter
//go:noescape
func MultiWriter(writers ...io.Writer,) io.Writer

//go:linkname Copy io.Copy
//go:noescape
func Copy(dst io.Writer, src io.Reader,) (int64, error)

//go:linkname pipewriterclosewitherror io.sub_pipewriterclosewitherror
func pipewriterclosewitherror(w *io.PipeWriter, err error) error {
	return w.CloseWithError(err)
}

//go:linkname PipeWriterCloseWithError io.sub_pipewriterclosewitherror
//go:noescape
func PipeWriterCloseWithError(w *io.PipeWriter, err error) error

//go:linkname TeeReader io.TeeReader
//go:noescape
func TeeReader(r io.Reader, w io.Writer,) io.Reader

//go:linkname pipewriterclose io.sub_pipewriterclose
func pipewriterclose(w *io.PipeWriter,) error {
	return w.Close()
}

//go:linkname PipeWriterClose io.sub_pipewriterclose
//go:noescape
func PipeWriterClose(w *io.PipeWriter,) error

//go:linkname LimitReader io.LimitReader
//go:noescape
func LimitReader(r io.Reader, n int64) io.Reader

//go:linkname sectionreaderreadat io.sub_sectionreaderreadat
func sectionreaderreadat(s *io.SectionReader, p []byte, off int64) (int, error) {
	return s.ReadAt(p, off)
}

//go:linkname SectionReaderReadAt io.sub_sectionreaderreadat
//go:noescape
func SectionReaderReadAt(s *io.SectionReader, p []byte, off int64) (int, error)

//go:linkname ReadFull io.ReadFull
//go:noescape
func ReadFull(r io.Reader, buf []byte) (int, error)

//go:linkname WriteString io.WriteString
//go:noescape
func WriteString(w io.Writer, s string) (int, error)

//go:linkname pipereaderclosewitherror io.sub_pipereaderclosewitherror
func pipereaderclosewitherror(r *io.PipeReader, err error) error {
	return r.CloseWithError(err)
}

//go:linkname PipeReaderCloseWithError io.sub_pipereaderclosewitherror
//go:noescape
func PipeReaderCloseWithError(r *io.PipeReader, err error) error
