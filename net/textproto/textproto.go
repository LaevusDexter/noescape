// This file has automatically been generated on Wed Feb 26 15:50:46 +05 2020.
// DO NOT EDIT.
package textproto

import (
	"bufio"
	"io"
	"net/textproto"
	_ "unsafe"
)

//go:linkname readerreadcontinuedlinebytes textproto.sub_readerreadcontinuedlinebytes
func readerreadcontinuedlinebytes(r *textproto.Reader,) ([]byte, error) {
	return r.ReadContinuedLineBytes()
}

//go:linkname ReaderReadContinuedLineBytes textproto.sub_readerreadcontinuedlinebytes
//go:noescape
func ReaderReadContinuedLineBytes(r *textproto.Reader,) ([]byte, error)

//go:linkname readerreadline textproto.sub_readerreadline
func readerreadline(r *textproto.Reader,) (string, error) {
	return r.ReadLine()
}

//go:linkname ReaderReadLine textproto.sub_readerreadline
//go:noescape
func ReaderReadLine(r *textproto.Reader,) (string, error)

//go:linkname writerdotwriter textproto.sub_writerdotwriter
func writerdotwriter(w *textproto.Writer,) io.WriteCloser {
	return w.DotWriter()
}

//go:linkname WriterDotWriter textproto.sub_writerdotwriter
//go:noescape
func WriterDotWriter(w *textproto.Writer,) io.WriteCloser

//go:linkname CanonicalMIMEHeaderKey net/textproto.CanonicalMIMEHeaderKey
//go:noescape
func CanonicalMIMEHeaderKey(s string) string

//go:linkname TrimBytes net/textproto.TrimBytes
//go:noescape
func TrimBytes(b []byte) []byte

//go:linkname TrimString net/textproto.TrimString
//go:noescape
func TrimString(s string) string

//go:linkname mimeheaderget textproto.sub_mimeheaderget
func mimeheaderget(h textproto.MIMEHeader, key string) string {
	return h.Get(key)
}

//go:linkname MIMEHeaderGet textproto.sub_mimeheaderget
//go:noescape
func MIMEHeaderGet(h textproto.MIMEHeader, key string) string

//go:linkname readerreadcontinuedline textproto.sub_readerreadcontinuedline
func readerreadcontinuedline(r *textproto.Reader,) (string, error) {
	return r.ReadContinuedLine()
}

//go:linkname ReaderReadContinuedLine textproto.sub_readerreadcontinuedline
//go:noescape
func ReaderReadContinuedLine(r *textproto.Reader,) (string, error)

//go:linkname readerreaddotlines textproto.sub_readerreaddotlines
func readerreaddotlines(r *textproto.Reader,) ([]string, error) {
	return r.ReadDotLines()
}

//go:linkname ReaderReadDotLines textproto.sub_readerreaddotlines
//go:noescape
func ReaderReadDotLines(r *textproto.Reader,) ([]string, error)

//go:linkname readerreadlinebytes textproto.sub_readerreadlinebytes
func readerreadlinebytes(r *textproto.Reader,) ([]byte, error) {
	return r.ReadLineBytes()
}

//go:linkname ReaderReadLineBytes textproto.sub_readerreadlinebytes
//go:noescape
func ReaderReadLineBytes(r *textproto.Reader,) ([]byte, error)

//go:linkname NewConn net/textproto.NewConn
//go:noescape
func NewConn(conn io.ReadWriteCloser) *textproto.Conn

//go:linkname protocolerrorerror textproto.sub_protocolerrorerror
func protocolerrorerror(p textproto.ProtocolError,) string {
	return p.Error()
}

//go:linkname ProtocolErrorError textproto.sub_protocolerrorerror
//go:noescape
func ProtocolErrorError(p textproto.ProtocolError,) string

//go:linkname NewReader net/textproto.NewReader
//go:noescape
func NewReader(r *bufio.Reader) *textproto.Reader

//go:linkname readerdotreader textproto.sub_readerdotreader
func readerdotreader(r *textproto.Reader,) io.Reader {
	return r.DotReader()
}

//go:linkname ReaderDotReader textproto.sub_readerdotreader
//go:noescape
func ReaderDotReader(r *textproto.Reader,) io.Reader

//go:linkname readerreaddotbytes textproto.sub_readerreaddotbytes
func readerreaddotbytes(r *textproto.Reader,) ([]byte, error) {
	return r.ReadDotBytes()
}

//go:linkname ReaderReadDotBytes textproto.sub_readerreaddotbytes
//go:noescape
func ReaderReadDotBytes(r *textproto.Reader,) ([]byte, error)

//go:linkname mimeheadervalues textproto.sub_mimeheadervalues
func mimeheadervalues(h textproto.MIMEHeader, key string) []string {
	return h.Values(key)
}

//go:linkname MIMEHeaderValues textproto.sub_mimeheadervalues
//go:noescape
func MIMEHeaderValues(h textproto.MIMEHeader, key string) []string

//go:linkname readerreadmimeheader textproto.sub_readerreadmimeheader
func readerreadmimeheader(r *textproto.Reader,) (textproto.MIMEHeader, error) {
	return r.ReadMIMEHeader()
}

//go:linkname ReaderReadMIMEHeader textproto.sub_readerreadmimeheader
//go:noescape
func ReaderReadMIMEHeader(r *textproto.Reader,) (textproto.MIMEHeader, error)

//go:linkname readerreadresponse textproto.sub_readerreadresponse
func readerreadresponse(r *textproto.Reader, expectCode int) (int, string, error) {
	return r.ReadResponse(expectCode)
}

//go:linkname ReaderReadResponse textproto.sub_readerreadresponse
//go:noescape
func ReaderReadResponse(r *textproto.Reader, expectCode int) (int, string, error)

//go:linkname NewWriter net/textproto.NewWriter
//go:noescape
func NewWriter(w *bufio.Writer) *textproto.Writer

//go:linkname readerreadcodeline textproto.sub_readerreadcodeline
func readerreadcodeline(r *textproto.Reader, expectCode int) (int, string, error) {
	return r.ReadCodeLine(expectCode)
}

//go:linkname ReaderReadCodeLine textproto.sub_readerreadcodeline
//go:noescape
func ReaderReadCodeLine(r *textproto.Reader, expectCode int) (int, string, error)

//go:linkname writerprintfline textproto.sub_writerprintfline
func writerprintfline(w *textproto.Writer, format string, args ...interface{}) error {
	return w.PrintfLine(format, args...)
}

//go:linkname WriterPrintfLine textproto.sub_writerprintfline
//go:noescape
func WriterPrintfLine(w *textproto.Writer, format string, args ...interface{}) error

//go:linkname Dial net/textproto.Dial
//go:noescape
func Dial(network, addr string) (*textproto.Conn, error)

//go:linkname connclose textproto.sub_connclose
func connclose(c *textproto.Conn,) error {
	return c.Close()
}

//go:linkname ConnClose textproto.sub_connclose
//go:noescape
func ConnClose(c *textproto.Conn,) error

//go:linkname conncmd textproto.sub_conncmd
func conncmd(c *textproto.Conn, format string, args ...interface{}) (uint, error) {
	return c.Cmd(format, args...)
}

//go:linkname ConnCmd textproto.sub_conncmd
//go:noescape
func ConnCmd(c *textproto.Conn, format string, args ...interface{}) (uint, error)

//go:linkname errorerror textproto.sub_errorerror
func errorerror(e *textproto.Error,) string {
	return e.Error()
}

//go:linkname ErrorError textproto.sub_errorerror
//go:noescape
func ErrorError(e *textproto.Error,) string

//go:linkname pipelinenext textproto.sub_pipelinenext
func pipelinenext(p *textproto.Pipeline,) uint {
	return p.Next()
}

//go:linkname PipelineNext textproto.sub_pipelinenext
//go:noescape
func PipelineNext(p *textproto.Pipeline,) uint
