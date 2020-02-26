// This file has automatically been generated on Wed Feb 26 15:50:17 +05 2020.
// DO NOT EDIT.
package tar

import (
	"archive/tar"
	"io"
	"os"
	_ "unsafe"
)

//go:linkname FileInfoHeader archive/tar.FileInfoHeader
//go:noescape
func FileInfoHeader(fi os.FileInfo, link string) (*tar.Header, error)

//go:linkname headerfileinfo tar.sub_headerfileinfo
func headerfileinfo(h *tar.Header,) os.FileInfo {
	return h.FileInfo()
}

//go:linkname HeaderFileInfo tar.sub_headerfileinfo
//go:noescape
func HeaderFileInfo(h *tar.Header,) os.FileInfo

//go:linkname NewReader archive/tar.NewReader
//go:noescape
func NewReader(r io.Reader) *tar.Reader

//go:linkname NewWriter archive/tar.NewWriter
//go:noescape
func NewWriter(w io.Writer) *tar.Writer

//go:linkname writerclose tar.sub_writerclose
func writerclose(tw *tar.Writer,) error {
	return tw.Close()
}

//go:linkname WriterClose tar.sub_writerclose
//go:noescape
func WriterClose(tw *tar.Writer,) error

//go:linkname writerflush tar.sub_writerflush
func writerflush(tw *tar.Writer,) error {
	return tw.Flush()
}

//go:linkname WriterFlush tar.sub_writerflush
//go:noescape
func WriterFlush(tw *tar.Writer,) error

//go:linkname formatstring tar.sub_formatstring
func formatstring(f tar.Format,) string {
	return f.String()
}

//go:linkname FormatString tar.sub_formatstring
//go:noescape
func FormatString(f tar.Format,) string

//go:linkname readernext tar.sub_readernext
func readernext(tr *tar.Reader,) (*tar.Header, error) {
	return tr.Next()
}

//go:linkname ReaderNext tar.sub_readernext
//go:noescape
func ReaderNext(tr *tar.Reader,) (*tar.Header, error)

//go:linkname readerread tar.sub_readerread
func readerread(tr *tar.Reader, b []byte) (int, error) {
	return tr.Read(b)
}

//go:linkname ReaderRead tar.sub_readerread
//go:noescape
func ReaderRead(tr *tar.Reader, b []byte) (int, error)

//go:linkname writerwrite tar.sub_writerwrite
func writerwrite(tw *tar.Writer, b []byte) (int, error) {
	return tw.Write(b)
}

//go:linkname WriterWrite tar.sub_writerwrite
//go:noescape
func WriterWrite(tw *tar.Writer, b []byte) (int, error)

//go:linkname writerwriteheader tar.sub_writerwriteheader
func writerwriteheader(tw *tar.Writer, hdr *tar.Header,) error {
	return tw.WriteHeader(hdr)
}

//go:linkname WriterWriteHeader tar.sub_writerwriteheader
//go:noescape
func WriterWriteHeader(tw *tar.Writer, hdr *tar.Header,) error
