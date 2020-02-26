// This file has automatically been generated on Wed Feb 26 15:50:17 +05 2020.
// DO NOT EDIT.
package bufio

import (
	"bufio"
	"io"
	_ "unsafe"
)

//go:linkname readerreadline bufio.sub_readerreadline
func readerreadline(b *bufio.Reader,) ([]byte, bool, error) {
	return b.ReadLine()
}

//go:linkname ReaderReadLine bufio.sub_readerreadline
//go:noescape
func ReaderReadLine(b *bufio.Reader,) ([]byte, bool, error)

//go:linkname readerreadstring bufio.sub_readerreadstring
func readerreadstring(b *bufio.Reader, delim byte) (string, error) {
	return b.ReadString(delim)
}

//go:linkname ReaderReadString bufio.sub_readerreadstring
//go:noescape
func ReaderReadString(b *bufio.Reader, delim byte) (string, error)

//go:linkname readersize bufio.sub_readersize
func readersize(b *bufio.Reader,) int {
	return b.Size()
}

//go:linkname ReaderSize bufio.sub_readersize
//go:noescape
func ReaderSize(b *bufio.Reader,) int

//go:linkname scannerbytes bufio.sub_scannerbytes
func scannerbytes(s *bufio.Scanner,) []byte {
	return s.Bytes()
}

//go:linkname ScannerBytes bufio.sub_scannerbytes
//go:noescape
func ScannerBytes(s *bufio.Scanner,) []byte

//go:linkname writersize bufio.sub_writersize
func writersize(b *bufio.Writer,) int {
	return b.Size()
}

//go:linkname WriterSize bufio.sub_writersize
//go:noescape
func WriterSize(b *bufio.Writer,) int

//go:linkname writerwrite bufio.sub_writerwrite
func writerwrite(b *bufio.Writer, p []byte) (int, error) {
	return b.Write(p)
}

//go:linkname WriterWrite bufio.sub_writerwrite
//go:noescape
func WriterWrite(b *bufio.Writer, p []byte) (int, error)

//go:linkname NewReaderSize bufio.NewReaderSize
//go:noescape
func NewReaderSize(rd io.Reader, size int) *bufio.Reader

//go:linkname readerreadbytes bufio.sub_readerreadbytes
func readerreadbytes(b *bufio.Reader, delim byte) ([]byte, error) {
	return b.ReadBytes(delim)
}

//go:linkname ReaderReadBytes bufio.sub_readerreadbytes
//go:noescape
func ReaderReadBytes(b *bufio.Reader, delim byte) ([]byte, error)

//go:linkname scannertext bufio.sub_scannertext
func scannertext(s *bufio.Scanner,) string {
	return s.Text()
}

//go:linkname ScannerText bufio.sub_scannertext
//go:noescape
func ScannerText(s *bufio.Scanner,) string

//go:linkname NewWriter bufio.NewWriter
//go:noescape
func NewWriter(w io.Writer) *bufio.Writer

//go:linkname writerflush bufio.sub_writerflush
func writerflush(b *bufio.Writer,) error {
	return b.Flush()
}

//go:linkname WriterFlush bufio.sub_writerflush
//go:noescape
func WriterFlush(b *bufio.Writer,) error

//go:linkname writerwritestring bufio.sub_writerwritestring
func writerwritestring(b *bufio.Writer, s string) (int, error) {
	return b.WriteString(s)
}

//go:linkname WriterWriteString bufio.sub_writerwritestring
//go:noescape
func WriterWriteString(b *bufio.Writer, s string) (int, error)

//go:linkname ScanBytes bufio.ScanBytes
//go:noescape
func ScanBytes(data []byte, atEOF bool) (int, []byte, error)

//go:linkname ScanRunes bufio.ScanRunes
//go:noescape
func ScanRunes(data []byte, atEOF bool) (int, []byte, error)

//go:linkname ScanWords bufio.ScanWords
//go:noescape
func ScanWords(data []byte, atEOF bool) (int, []byte, error)

//go:linkname readerreadslice bufio.sub_readerreadslice
func readerreadslice(b *bufio.Reader, delim byte) ([]byte, error) {
	return b.ReadSlice(delim)
}

//go:linkname ReaderReadSlice bufio.sub_readerreadslice
//go:noescape
func ReaderReadSlice(b *bufio.Reader, delim byte) ([]byte, error)

//go:linkname NewScanner bufio.NewScanner
//go:noescape
func NewScanner(r io.Reader) *bufio.Scanner

//go:linkname scannererr bufio.sub_scannererr
func scannererr(s *bufio.Scanner,) error {
	return s.Err()
}

//go:linkname ScannerErr bufio.sub_scannererr
//go:noescape
func ScannerErr(s *bufio.Scanner,) error

//go:linkname NewWriterSize bufio.NewWriterSize
//go:noescape
func NewWriterSize(w io.Writer, size int) *bufio.Writer

//go:linkname writeravailable bufio.sub_writeravailable
func writeravailable(b *bufio.Writer,) int {
	return b.Available()
}

//go:linkname WriterAvailable bufio.sub_writeravailable
//go:noescape
func WriterAvailable(b *bufio.Writer,) int

//go:linkname ScanLines bufio.ScanLines
//go:noescape
func ScanLines(data []byte, atEOF bool) (int, []byte, error)

//go:linkname NewReadWriter bufio.NewReadWriter
//go:noescape
func NewReadWriter(r *bufio.Reader, w *bufio.Writer,) *bufio.ReadWriter

//go:linkname readerreadbyte bufio.sub_readerreadbyte
func readerreadbyte(b *bufio.Reader,) (byte, error) {
	return b.ReadByte()
}

//go:linkname ReaderReadByte bufio.sub_readerreadbyte
//go:noescape
func ReaderReadByte(b *bufio.Reader,) (byte, error)

//go:linkname readerunreadbyte bufio.sub_readerunreadbyte
func readerunreadbyte(b *bufio.Reader,) error {
	return b.UnreadByte()
}

//go:linkname ReaderUnreadByte bufio.sub_readerunreadbyte
//go:noescape
func ReaderUnreadByte(b *bufio.Reader,) error

//go:linkname readerwriteto bufio.sub_readerwriteto
func readerwriteto(b *bufio.Reader, w io.Writer) (int64, error) {
	return b.WriteTo(w)
}

//go:linkname ReaderWriteTo bufio.sub_readerwriteto
//go:noescape
func ReaderWriteTo(b *bufio.Reader, w io.Writer) (int64, error)

//go:linkname NewReader bufio.NewReader
//go:noescape
func NewReader(rd io.Reader) *bufio.Reader

//go:linkname readerread bufio.sub_readerread
func readerread(b *bufio.Reader, p []byte) (int, error) {
	return b.Read(p)
}

//go:linkname ReaderRead bufio.sub_readerread
//go:noescape
func ReaderRead(b *bufio.Reader, p []byte) (int, error)

//go:linkname writerreadfrom bufio.sub_writerreadfrom
func writerreadfrom(b *bufio.Writer, r io.Reader) (int64, error) {
	return b.ReadFrom(r)
}

//go:linkname WriterReadFrom bufio.sub_writerreadfrom
//go:noescape
func WriterReadFrom(b *bufio.Writer, r io.Reader) (int64, error)

//go:linkname readerpeek bufio.sub_readerpeek
func readerpeek(b *bufio.Reader, n int) ([]byte, error) {
	return b.Peek(n)
}

//go:linkname ReaderPeek bufio.sub_readerpeek
//go:noescape
func ReaderPeek(b *bufio.Reader, n int) ([]byte, error)

//go:linkname scannerscan bufio.sub_scannerscan
func scannerscan(s *bufio.Scanner,) bool {
	return s.Scan()
}

//go:linkname ScannerScan bufio.sub_scannerscan
//go:noescape
func ScannerScan(s *bufio.Scanner,) bool

//go:linkname readerbuffered bufio.sub_readerbuffered
func readerbuffered(b *bufio.Reader,) int {
	return b.Buffered()
}

//go:linkname ReaderBuffered bufio.sub_readerbuffered
//go:noescape
func ReaderBuffered(b *bufio.Reader,) int

//go:linkname readerunreadrune bufio.sub_readerunreadrune
func readerunreadrune(b *bufio.Reader,) error {
	return b.UnreadRune()
}

//go:linkname ReaderUnreadRune bufio.sub_readerunreadrune
//go:noescape
func ReaderUnreadRune(b *bufio.Reader,) error

//go:linkname writerwriterune bufio.sub_writerwriterune
func writerwriterune(b *bufio.Writer, r rune) (int, error) {
	return b.WriteRune(r)
}

//go:linkname WriterWriteRune bufio.sub_writerwriterune
//go:noescape
func WriterWriteRune(b *bufio.Writer, r rune) (int, error)

//go:linkname readerdiscard bufio.sub_readerdiscard
func readerdiscard(b *bufio.Reader, n int) (int, error) {
	return b.Discard(n)
}

//go:linkname ReaderDiscard bufio.sub_readerdiscard
//go:noescape
func ReaderDiscard(b *bufio.Reader, n int) (int, error)

//go:linkname readerreadrune bufio.sub_readerreadrune
func readerreadrune(b *bufio.Reader,) (rune, int, error) {
	return b.ReadRune()
}

//go:linkname ReaderReadRune bufio.sub_readerreadrune
//go:noescape
func ReaderReadRune(b *bufio.Reader,) (rune, int, error)

//go:linkname writerbuffered bufio.sub_writerbuffered
func writerbuffered(b *bufio.Writer,) int {
	return b.Buffered()
}

//go:linkname WriterBuffered bufio.sub_writerbuffered
//go:noescape
func WriterBuffered(b *bufio.Writer,) int

//go:linkname writerwritebyte bufio.sub_writerwritebyte
func writerwritebyte(b *bufio.Writer, c byte) error {
	return b.WriteByte(c)
}

//go:linkname WriterWriteByte bufio.sub_writerwritebyte
//go:noescape
func WriterWriteByte(b *bufio.Writer, c byte) error
