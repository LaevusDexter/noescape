// This file has automatically been generated on Wed Feb 26 15:50:17 +05 2020.
// DO NOT EDIT.
package zip

import (
	"archive/zip"
	"io"
	"os"
	"time"
	_ "unsafe"
)

//go:linkname fileheadermodtime zip.sub_fileheadermodtime
func fileheadermodtime(h *zip.FileHeader,) time.Time {
	return h.ModTime()
}

//go:linkname FileHeaderModTime zip.sub_fileheadermodtime
//go:noescape
func FileHeaderModTime(h *zip.FileHeader,) time.Time

//go:linkname writerflush zip.sub_writerflush
func writerflush(w *zip.Writer,) error {
	return w.Flush()
}

//go:linkname WriterFlush zip.sub_writerflush
//go:noescape
func WriterFlush(w *zip.Writer,) error

//go:linkname writersetcomment zip.sub_writersetcomment
func writersetcomment(w *zip.Writer, comment string) error {
	return w.SetComment(comment)
}

//go:linkname WriterSetComment zip.sub_writersetcomment
//go:noescape
func WriterSetComment(w *zip.Writer, comment string) error

//go:linkname NewWriter archive/zip.NewWriter
//go:noescape
func NewWriter(w io.Writer) *zip.Writer

//go:linkname fileopen zip.sub_fileopen
func fileopen(f *zip.File,) (io.ReadCloser, error) {
	return f.Open()
}

//go:linkname FileOpen zip.sub_fileopen
//go:noescape
func FileOpen(f *zip.File,) (io.ReadCloser, error)

//go:linkname fileheaderfileinfo zip.sub_fileheaderfileinfo
func fileheaderfileinfo(h *zip.FileHeader,) os.FileInfo {
	return h.FileInfo()
}

//go:linkname FileHeaderFileInfo zip.sub_fileheaderfileinfo
//go:noescape
func FileHeaderFileInfo(h *zip.FileHeader,) os.FileInfo

//go:linkname fileheadermode zip.sub_fileheadermode
func fileheadermode(h *zip.FileHeader,) os.FileMode {
	return h.Mode()
}

//go:linkname FileHeaderMode zip.sub_fileheadermode
//go:noescape
func FileHeaderMode(h *zip.FileHeader,) os.FileMode

//go:linkname writercreate zip.sub_writercreate
func writercreate(w *zip.Writer, name string) (io.Writer, error) {
	return w.Create(name)
}

//go:linkname WriterCreate zip.sub_writercreate
//go:noescape
func WriterCreate(w *zip.Writer, name string) (io.Writer, error)

//go:linkname writercreateheader zip.sub_writercreateheader
func writercreateheader(w *zip.Writer, fh *zip.FileHeader,) (io.Writer, error) {
	return w.CreateHeader(fh)
}

//go:linkname WriterCreateHeader zip.sub_writercreateheader
//go:noescape
func WriterCreateHeader(w *zip.Writer, fh *zip.FileHeader,) (io.Writer, error)

//go:linkname FileInfoHeader archive/zip.FileInfoHeader
//go:noescape
func FileInfoHeader(fi os.FileInfo) (*zip.FileHeader, error)

//go:linkname NewReader archive/zip.NewReader
//go:noescape
func NewReader(r io.ReaderAt, size int64) (*zip.Reader, error)

//go:linkname writerclose zip.sub_writerclose
func writerclose(w *zip.Writer,) error {
	return w.Close()
}

//go:linkname WriterClose zip.sub_writerclose
//go:noescape
func WriterClose(w *zip.Writer,) error

//go:linkname filedataoffset zip.sub_filedataoffset
func filedataoffset(f *zip.File,) (int64, error) {
	return f.DataOffset()
}

//go:linkname FileDataOffset zip.sub_filedataoffset
//go:noescape
func FileDataOffset(f *zip.File,) (int64, error)

//go:linkname OpenReader archive/zip.OpenReader
//go:noescape
func OpenReader(name string) (*zip.ReadCloser, error)

//go:linkname readcloserclose zip.sub_readcloserclose
func readcloserclose(rc *zip.ReadCloser,) error {
	return rc.Close()
}

//go:linkname ReadCloserClose zip.sub_readcloserclose
//go:noescape
func ReadCloserClose(rc *zip.ReadCloser,) error
