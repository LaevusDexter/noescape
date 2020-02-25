// This file has automatically been generated on Wed Feb 26 02:10:04 +05 2020.
// DO NOT EDIT.
package multipart

import (
	"io"
	"mime/multipart"
	"net/textproto"
	_ "unsafe"
)

//go:linkname partfilename multipart.sub_partfilename
func partfilename(p *multipart.Part) string {
	return p.FileName()
}

//go:linkname PartFileName multipart.sub_partfilename
//go:noescape
func PartFileName(p *multipart.Part) string

//go:linkname writercreateformfield multipart.sub_writercreateformfield
func writercreateformfield(w *multipart.Writer, fieldname string) (io.Writer, error) {
	return w.CreateFormField(fieldname)
}

//go:linkname WriterCreateFormField multipart.sub_writercreateformfield
//go:noescape
func WriterCreateFormField(w *multipart.Writer, fieldname string) (io.Writer, error)

//go:linkname writercreateformfile multipart.sub_writercreateformfile
func writercreateformfile(w *multipart.Writer, fieldname, filename string) (io.Writer, error) {
	return w.CreateFormFile(fieldname, filename)
}

//go:linkname WriterCreateFormFile multipart.sub_writercreateformfile
//go:noescape
func WriterCreateFormFile(w *multipart.Writer, fieldname, filename string) (io.Writer, error)

//go:linkname writercreatepart multipart.sub_writercreatepart
func writercreatepart(w *multipart.Writer, header textproto.MIMEHeader) (io.Writer, error) {
	return w.CreatePart(header)
}

//go:linkname WriterCreatePart multipart.sub_writercreatepart
//go:noescape
func WriterCreatePart(w *multipart.Writer, header textproto.MIMEHeader) (io.Writer, error)

//go:linkname writerformdatacontenttype multipart.sub_writerformdatacontenttype
func writerformdatacontenttype(w *multipart.Writer) string {
	return w.FormDataContentType()
}

//go:linkname WriterFormDataContentType multipart.sub_writerformdatacontenttype
//go:noescape
func WriterFormDataContentType(w *multipart.Writer) string

//go:linkname writerwritefield multipart.sub_writerwritefield
func writerwritefield(w *multipart.Writer, fieldname, value string) error {
	return w.WriteField(fieldname, value)
}

//go:linkname WriterWriteField multipart.sub_writerwritefield
//go:noescape
func WriterWriteField(w *multipart.Writer, fieldname, value string) error

//go:linkname partclose multipart.sub_partclose
func partclose(p *multipart.Part) error {
	return p.Close()
}

//go:linkname PartClose multipart.sub_partclose
//go:noescape
func PartClose(p *multipart.Part) error

//go:linkname partformname multipart.sub_partformname
func partformname(p *multipart.Part) string {
	return p.FormName()
}

//go:linkname PartFormName multipart.sub_partformname
//go:noescape
func PartFormName(p *multipart.Part) string

//go:linkname NewReader mime/multipart.NewReader
//go:noescape
func NewReader(r io.Reader, boundary string) *multipart.Reader

//go:linkname NewWriter mime/multipart.NewWriter
//go:noescape
func NewWriter(w io.Writer) *multipart.Writer

//go:linkname writerboundary multipart.sub_writerboundary
func writerboundary(w *multipart.Writer) string {
	return w.Boundary()
}

//go:linkname WriterBoundary multipart.sub_writerboundary
//go:noescape
func WriterBoundary(w *multipart.Writer) string

//go:linkname writersetboundary multipart.sub_writersetboundary
func writersetboundary(w *multipart.Writer, boundary string) error {
	return w.SetBoundary(boundary)
}

//go:linkname WriterSetBoundary multipart.sub_writersetboundary
//go:noescape
func WriterSetBoundary(w *multipart.Writer, boundary string) error

//go:linkname fileheaderopen multipart.sub_fileheaderopen
func fileheaderopen(fh *multipart.FileHeader) (multipart.File, error) {
	return fh.Open()
}

//go:linkname FileHeaderOpen multipart.sub_fileheaderopen
//go:noescape
func FileHeaderOpen(fh *multipart.FileHeader) (multipart.File, error)

//go:linkname formremoveall multipart.sub_formremoveall
func formremoveall(f *multipart.Form) error {
	return f.RemoveAll()
}

//go:linkname FormRemoveAll multipart.sub_formremoveall
//go:noescape
func FormRemoveAll(f *multipart.Form) error

//go:linkname readernextpart multipart.sub_readernextpart
func readernextpart(r *multipart.Reader) (*multipart.Part, error) {
	return r.NextPart()
}

//go:linkname ReaderNextPart multipart.sub_readernextpart
//go:noescape
func ReaderNextPart(r *multipart.Reader) (*multipart.Part, error)

//go:linkname readernextrawpart multipart.sub_readernextrawpart
func readernextrawpart(r *multipart.Reader) (*multipart.Part, error) {
	return r.NextRawPart()
}

//go:linkname ReaderNextRawPart multipart.sub_readernextrawpart
//go:noescape
func ReaderNextRawPart(r *multipart.Reader) (*multipart.Part, error)

//go:linkname writerclose multipart.sub_writerclose
func writerclose(w *multipart.Writer) error {
	return w.Close()
}

//go:linkname WriterClose multipart.sub_writerclose
//go:noescape
func WriterClose(w *multipart.Writer) error

//go:linkname partread multipart.sub_partread
func partread(p *multipart.Part, d []byte) (int, error) {
	return p.Read(d)
}

//go:linkname PartRead multipart.sub_partread
//go:noescape
func PartRead(p *multipart.Part, d []byte) (int, error)

//go:linkname readerreadform multipart.sub_readerreadform
func readerreadform(r *multipart.Reader, maxMemory int64) (*multipart.Form, error) {
	return r.ReadForm(maxMemory)
}

//go:linkname ReaderReadForm multipart.sub_readerreadform
//go:noescape
func ReaderReadForm(r *multipart.Reader, maxMemory int64) (*multipart.Form, error)
