// This file has automatically been generated on Wed Feb 26 02:09:40 +05 2020.
// DO NOT EDIT.
package bytes

import (
	"bytes"
	"io"
	"unicode"
	_ "unsafe"
)

//go:linkname bufferwritestring bytes.sub_bufferwritestring
func bufferwritestring(b *bytes.Buffer, s string) (int, error) {
	return b.WriteString(s)
}

//go:linkname BufferWriteString bytes.sub_bufferwritestring
//go:noescape
func BufferWriteString(b *bytes.Buffer, s string) (int, error)

//go:linkname IndexRune bytes.IndexRune
//go:noescape
func IndexRune(s []byte, r rune) int

//go:linkname TrimRightFunc bytes.TrimRightFunc
//go:noescape
func TrimRightFunc(s []byte, f func(r rune) bool) []byte

//go:linkname bufferunreadbyte bytes.sub_bufferunreadbyte
func bufferunreadbyte(b *bytes.Buffer) error {
	return b.UnreadByte()
}

//go:linkname BufferUnreadByte bytes.sub_bufferunreadbyte
//go:noescape
func BufferUnreadByte(b *bytes.Buffer) error

//go:linkname bufferwriterune bytes.sub_bufferwriterune
func bufferwriterune(b *bytes.Buffer, r rune) (int, error) {
	return b.WriteRune(r)
}

//go:linkname BufferWriteRune bytes.sub_bufferwriterune
//go:noescape
func BufferWriteRune(b *bytes.Buffer, r rune) (int, error)

//go:linkname ContainsAny bytes.ContainsAny
//go:noescape
func ContainsAny(b []byte, chars string) bool

//go:linkname ToLower bytes.ToLower
//go:noescape
func ToLower(s []byte) []byte

//go:linkname TrimLeftFunc bytes.TrimLeftFunc
//go:noescape
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

//go:linkname NewBuffer bytes.NewBuffer
//go:noescape
func NewBuffer(buf []byte) *bytes.Buffer

//go:linkname readerunreadbyte bytes.sub_readerunreadbyte
func readerunreadbyte(r *bytes.Reader) error {
	return r.UnreadByte()
}

//go:linkname ReaderUnreadByte bytes.sub_readerunreadbyte
//go:noescape
func ReaderUnreadByte(r *bytes.Reader) error

//go:linkname bufferreadfrom bytes.sub_bufferreadfrom
func bufferreadfrom(b *bytes.Buffer, r io.Reader) (int64, error) {
	return b.ReadFrom(r)
}

//go:linkname BufferReadFrom bytes.sub_bufferreadfrom
//go:noescape
func BufferReadFrom(b *bytes.Buffer, r io.Reader) (int64, error)

//go:linkname TrimPrefix bytes.TrimPrefix
//go:noescape
func TrimPrefix(s, prefix []byte) []byte

//go:linkname NewBufferString bytes.NewBufferString
//go:noescape
func NewBufferString(s string) *bytes.Buffer

//go:linkname bufferreadbytes bytes.sub_bufferreadbytes
func bufferreadbytes(b *bytes.Buffer, delim byte) ([]byte, error) {
	return b.ReadBytes(delim)
}

//go:linkname BufferReadBytes bytes.sub_bufferreadbytes
//go:noescape
func BufferReadBytes(b *bytes.Buffer, delim byte) ([]byte, error)

//go:linkname FieldsFunc bytes.FieldsFunc
//go:noescape
func FieldsFunc(s []byte, f func(rune) bool) [][]byte

//go:linkname LastIndexFunc bytes.LastIndexFunc
//go:noescape
func LastIndexFunc(s []byte, f func(r rune) bool) int

//go:linkname readerwriteto bytes.sub_readerwriteto
func readerwriteto(r *bytes.Reader, w io.Writer) (int64, error) {
	return r.WriteTo(w)
}

//go:linkname ReaderWriteTo bytes.sub_readerwriteto
//go:noescape
func ReaderWriteTo(r *bytes.Reader, w io.Writer) (int64, error)

//go:linkname buffernext bytes.sub_buffernext
func buffernext(b *bytes.Buffer, n int) []byte {
	return b.Next(n)
}

//go:linkname BufferNext bytes.sub_buffernext
//go:noescape
func BufferNext(b *bytes.Buffer, n int) []byte

//go:linkname LastIndexByte bytes.LastIndexByte
//go:noescape
func LastIndexByte(s []byte, c byte) int

//go:linkname Title bytes.Title
//go:noescape
func Title(s []byte) []byte

//go:linkname ToUpperSpecial bytes.ToUpperSpecial
//go:noescape
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte

//go:linkname readerread bytes.sub_readerread
func readerread(r *bytes.Reader, b []byte) (int, error) {
	return r.Read(b)
}

//go:linkname ReaderRead bytes.sub_readerread
//go:noescape
func ReaderRead(r *bytes.Reader, b []byte) (int, error)

//go:linkname HasSuffix bytes.HasSuffix
//go:noescape
func HasSuffix(s, suffix []byte) bool

//go:linkname Runes bytes.Runes
//go:noescape
func Runes(s []byte) []rune

//go:linkname bufferreadstring bytes.sub_bufferreadstring
func bufferreadstring(b *bytes.Buffer, delim byte) (string, error) {
	return b.ReadString(delim)
}

//go:linkname BufferReadString bytes.sub_bufferreadstring
//go:noescape
func BufferReadString(b *bytes.Buffer, delim byte) (string, error)

//go:linkname bufferwritebyte bytes.sub_bufferwritebyte
func bufferwritebyte(b *bytes.Buffer, c byte) error {
	return b.WriteByte(c)
}

//go:linkname BufferWriteByte bytes.sub_bufferwritebyte
//go:noescape
func BufferWriteByte(b *bytes.Buffer, c byte) error

//go:linkname NewReader bytes.NewReader
//go:noescape
func NewReader(b []byte) *bytes.Reader

//go:linkname EqualFold bytes.EqualFold
//go:noescape
func EqualFold(s, t []byte) bool

//go:linkname SplitAfterN bytes.SplitAfterN
//go:noescape
func SplitAfterN(s, sep []byte, n int) [][]byte

//go:linkname ToLowerSpecial bytes.ToLowerSpecial
//go:noescape
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte

//go:linkname bufferreadbyte bytes.sub_bufferreadbyte
func bufferreadbyte(b *bytes.Buffer) (byte, error) {
	return b.ReadByte()
}

//go:linkname BufferReadByte bytes.sub_bufferreadbyte
//go:noescape
func BufferReadByte(b *bytes.Buffer) (byte, error)

//go:linkname bufferstring bytes.sub_bufferstring
func bufferstring(b *bytes.Buffer) string {
	return b.String()
}

//go:linkname BufferString bytes.sub_bufferstring
//go:noescape
func BufferString(b *bytes.Buffer) string

//go:linkname bufferwrite bytes.sub_bufferwrite
func bufferwrite(b *bytes.Buffer, p []byte) (int, error) {
	return b.Write(p)
}

//go:linkname BufferWrite bytes.sub_bufferwrite
//go:noescape
func BufferWrite(b *bytes.Buffer, p []byte) (int, error)

//go:linkname IndexFunc bytes.IndexFunc
//go:noescape
func IndexFunc(s []byte, f func(r rune) bool) int

//go:linkname ToUpper bytes.ToUpper
//go:noescape
func ToUpper(s []byte) []byte

//go:linkname TrimSpace bytes.TrimSpace
//go:noescape
func TrimSpace(s []byte) []byte

//go:linkname ToTitleSpecial bytes.ToTitleSpecial
//go:noescape
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte

//go:linkname TrimSuffix bytes.TrimSuffix
//go:noescape
func TrimSuffix(s, suffix []byte) []byte

//go:linkname bufferread bytes.sub_bufferread
func bufferread(b *bytes.Buffer, p []byte) (int, error) {
	return b.Read(p)
}

//go:linkname BufferRead bytes.sub_bufferread
//go:noescape
func BufferRead(b *bytes.Buffer, p []byte) (int, error)

//go:linkname readerunreadrune bytes.sub_readerunreadrune
func readerunreadrune(r *bytes.Reader) error {
	return r.UnreadRune()
}

//go:linkname ReaderUnreadRune bytes.sub_readerunreadrune
//go:noescape
func ReaderUnreadRune(r *bytes.Reader) error

//go:linkname Equal bytes.Equal
//go:noescape
func Equal(a, b []byte) bool

//go:linkname HasPrefix bytes.HasPrefix
//go:noescape
func HasPrefix(s, prefix []byte) bool

//go:linkname IndexByte bytes.IndexByte
//go:noescape
func IndexByte(b []byte, c byte) int

//go:linkname Fields bytes.Fields
//go:noescape
func Fields(s []byte) [][]byte

//go:linkname Join bytes.Join
//go:noescape
func Join(s [][]byte, sep []byte) []byte

//go:linkname bufferunreadrune bytes.sub_bufferunreadrune
func bufferunreadrune(b *bytes.Buffer) error {
	return b.UnreadRune()
}

//go:linkname BufferUnreadRune bytes.sub_bufferunreadrune
//go:noescape
func BufferUnreadRune(b *bytes.Buffer) error

//go:linkname LastIndex bytes.LastIndex
//go:noescape
func LastIndex(s, sep []byte) int

//go:linkname Map bytes.Map
//go:noescape
func Map(mapping func(r rune) rune, s []byte) []byte

//go:linkname buffercap bytes.sub_buffercap
func buffercap(b *bytes.Buffer) int {
	return b.Cap()
}

//go:linkname BufferCap bytes.sub_buffercap
//go:noescape
func BufferCap(b *bytes.Buffer) int

//go:linkname bufferbytes bytes.sub_bufferbytes
func bufferbytes(b *bytes.Buffer) []byte {
	return b.Bytes()
}

//go:linkname BufferBytes bytes.sub_bufferbytes
//go:noescape
func BufferBytes(b *bytes.Buffer) []byte

//go:linkname readerlen bytes.sub_readerlen
func readerlen(r *bytes.Reader) int {
	return r.Len()
}

//go:linkname ReaderLen bytes.sub_readerlen
//go:noescape
func ReaderLen(r *bytes.Reader) int

//go:linkname Compare bytes.Compare
//go:noescape
func Compare(a, b []byte) int

//go:linkname Repeat bytes.Repeat
//go:noescape
func Repeat(b []byte, count int) []byte

//go:linkname ReplaceAll bytes.ReplaceAll
//go:noescape
func ReplaceAll(s, old, new []byte) []byte

//go:linkname SplitN bytes.SplitN
//go:noescape
func SplitN(s, sep []byte, n int) [][]byte

//go:linkname ToTitle bytes.ToTitle
//go:noescape
func ToTitle(s []byte) []byte

//go:linkname TrimLeft bytes.TrimLeft
//go:noescape
func TrimLeft(s []byte, cutset string) []byte

//go:linkname readerreadrune bytes.sub_readerreadrune
func readerreadrune(r *bytes.Reader) (rune, int, error) {
	return r.ReadRune()
}

//go:linkname ReaderReadRune bytes.sub_readerreadrune
//go:noescape
func ReaderReadRune(r *bytes.Reader) (rune, int, error)

//go:linkname readersize bytes.sub_readersize
func readersize(r *bytes.Reader) int64 {
	return r.Size()
}

//go:linkname ReaderSize bytes.sub_readersize
//go:noescape
func ReaderSize(r *bytes.Reader) int64

//go:linkname Replace bytes.Replace
//go:noescape
func Replace(s, old, new []byte, n int) []byte

//go:linkname Split bytes.Split
//go:noescape
func Split(s, sep []byte) [][]byte

//go:linkname SplitAfter bytes.SplitAfter
//go:noescape
func SplitAfter(s, sep []byte) [][]byte

//go:linkname readerreadat bytes.sub_readerreadat
func readerreadat(r *bytes.Reader, b []byte, off int64) (int, error) {
	return r.ReadAt(b, off)
}

//go:linkname ReaderReadAt bytes.sub_readerreadat
//go:noescape
func ReaderReadAt(r *bytes.Reader, b []byte, off int64) (int, error)

//go:linkname readerreadbyte bytes.sub_readerreadbyte
func readerreadbyte(r *bytes.Reader) (byte, error) {
	return r.ReadByte()
}

//go:linkname ReaderReadByte bytes.sub_readerreadbyte
//go:noescape
func ReaderReadByte(r *bytes.Reader) (byte, error)

//go:linkname readerseek bytes.sub_readerseek
func readerseek(r *bytes.Reader, offset int64, whence int) (int64, error) {
	return r.Seek(offset, whence)
}

//go:linkname ReaderSeek bytes.sub_readerseek
//go:noescape
func ReaderSeek(r *bytes.Reader, offset int64, whence int) (int64, error)

//go:linkname ContainsRune bytes.ContainsRune
//go:noescape
func ContainsRune(b []byte, r rune) bool

//go:linkname LastIndexAny bytes.LastIndexAny
//go:noescape
func LastIndexAny(s []byte, chars string) int

//go:linkname Trim bytes.Trim
//go:noescape
func Trim(s []byte, cutset string) []byte

//go:linkname IndexAny bytes.IndexAny
//go:noescape
func IndexAny(s []byte, chars string) int

//go:linkname TrimFunc bytes.TrimFunc
//go:noescape
func TrimFunc(s []byte, f func(r rune) bool) []byte

//go:linkname TrimRight bytes.TrimRight
//go:noescape
func TrimRight(s []byte, cutset string) []byte

//go:linkname bufferlen bytes.sub_bufferlen
func bufferlen(b *bytes.Buffer) int {
	return b.Len()
}

//go:linkname BufferLen bytes.sub_bufferlen
//go:noescape
func BufferLen(b *bytes.Buffer) int

//go:linkname bufferreadrune bytes.sub_bufferreadrune
func bufferreadrune(b *bytes.Buffer) (rune, int, error) {
	return b.ReadRune()
}

//go:linkname BufferReadRune bytes.sub_bufferreadrune
//go:noescape
func BufferReadRune(b *bytes.Buffer) (rune, int, error)

//go:linkname Contains bytes.Contains
//go:noescape
func Contains(b, subslice []byte) bool

//go:linkname Count bytes.Count
//go:noescape
func Count(s, sep []byte) int

//go:linkname Index bytes.Index
//go:noescape
func Index(s, sep []byte) int

//go:linkname bufferwriteto bytes.sub_bufferwriteto
func bufferwriteto(b *bytes.Buffer, w io.Writer) (int64, error) {
	return b.WriteTo(w)
}

//go:linkname BufferWriteTo bytes.sub_bufferwriteto
//go:noescape
func BufferWriteTo(b *bytes.Buffer, w io.Writer) (int64, error)
