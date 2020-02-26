// This file has automatically been generated on Wed Feb 26 15:50:52 +05 2020.
// DO NOT EDIT.
package strings

import (
	"io"
	"strings"
	"unicode"
	_ "unsafe"
)

//go:linkname readerreadbyte strings.sub_readerreadbyte
func readerreadbyte(r *strings.Reader,) (byte, error) {
	return r.ReadByte()
}

//go:linkname ReaderReadByte strings.sub_readerreadbyte
//go:noescape
func ReaderReadByte(r *strings.Reader,) (byte, error)

//go:linkname Count strings.Count
//go:noescape
func Count(s, substr string) int

//go:linkname LastIndex strings.LastIndex
//go:noescape
func LastIndex(s, substr string) int

//go:linkname Title strings.Title
//go:noescape
func Title(s string) string

//go:linkname TrimFunc strings.TrimFunc
//go:noescape
func TrimFunc(s string, f func(rune) bool) string

//go:linkname readerlen strings.sub_readerlen
func readerlen(r *strings.Reader,) int {
	return r.Len()
}

//go:linkname ReaderLen strings.sub_readerlen
//go:noescape
func ReaderLen(r *strings.Reader,) int

//go:linkname readerunreadrune strings.sub_readerunreadrune
func readerunreadrune(r *strings.Reader,) error {
	return r.UnreadRune()
}

//go:linkname ReaderUnreadRune strings.sub_readerunreadrune
//go:noescape
func ReaderUnreadRune(r *strings.Reader,) error

//go:linkname readerwriteto strings.sub_readerwriteto
func readerwriteto(r *strings.Reader, w io.Writer) (int64, error) {
	return r.WriteTo(w)
}

//go:linkname ReaderWriteTo strings.sub_readerwriteto
//go:noescape
func ReaderWriteTo(r *strings.Reader, w io.Writer) (int64, error)

//go:linkname TrimLeft strings.TrimLeft
//go:noescape
func TrimLeft(s string, cutset string) string

//go:linkname buildercap strings.sub_buildercap
func buildercap(b *strings.Builder,) int {
	return b.Cap()
}

//go:linkname BuilderCap strings.sub_buildercap
//go:noescape
func BuilderCap(b *strings.Builder,) int

//go:linkname readerread strings.sub_readerread
func readerread(r *strings.Reader, b []byte) (int, error) {
	return r.Read(b)
}

//go:linkname ReaderRead strings.sub_readerread
//go:noescape
func ReaderRead(r *strings.Reader, b []byte) (int, error)

//go:linkname IndexFunc strings.IndexFunc
//go:noescape
func IndexFunc(s string, f func(rune) bool) int

//go:linkname SplitAfterN strings.SplitAfterN
//go:noescape
func SplitAfterN(s, sep string, n int) []string

//go:linkname ToLowerSpecial strings.ToLowerSpecial
//go:noescape
func ToLowerSpecial(c unicode.SpecialCase, s string) string

//go:linkname ToTitleSpecial strings.ToTitleSpecial
//go:noescape
func ToTitleSpecial(c unicode.SpecialCase, s string) string

//go:linkname ToUpper strings.ToUpper
//go:noescape
func ToUpper(s string) string

//go:linkname builderlen strings.sub_builderlen
func builderlen(b *strings.Builder,) int {
	return b.Len()
}

//go:linkname BuilderLen strings.sub_builderlen
//go:noescape
func BuilderLen(b *strings.Builder,) int

//go:linkname readerreadat strings.sub_readerreadat
func readerreadat(r *strings.Reader, b []byte, off int64) (int, error) {
	return r.ReadAt(b, off)
}

//go:linkname ReaderReadAt strings.sub_readerreadat
//go:noescape
func ReaderReadAt(r *strings.Reader, b []byte, off int64) (int, error)

//go:linkname Fields strings.Fields
//go:noescape
func Fields(s string) []string

//go:linkname HasPrefix strings.HasPrefix
//go:noescape
func HasPrefix(s, prefix string) bool

//go:linkname readerunreadbyte strings.sub_readerunreadbyte
func readerunreadbyte(r *strings.Reader,) error {
	return r.UnreadByte()
}

//go:linkname ReaderUnreadByte strings.sub_readerunreadbyte
//go:noescape
func ReaderUnreadByte(r *strings.Reader,) error

//go:linkname EqualFold strings.EqualFold
//go:noescape
func EqualFold(s, t string) bool

//go:linkname ToTitle strings.ToTitle
//go:noescape
func ToTitle(s string) string

//go:linkname readerseek strings.sub_readerseek
func readerseek(r *strings.Reader, offset int64, whence int) (int64, error) {
	return r.Seek(offset, whence)
}

//go:linkname ReaderSeek strings.sub_readerseek
//go:noescape
func ReaderSeek(r *strings.Reader, offset int64, whence int) (int64, error)

//go:linkname replacerwritestring strings.sub_replacerwritestring
func replacerwritestring(r *strings.Replacer, w io.Writer, s string) (int, error) {
	return r.WriteString(w, s)
}

//go:linkname ReplacerWriteString strings.sub_replacerwritestring
//go:noescape
func ReplacerWriteString(r *strings.Replacer, w io.Writer, s string) (int, error)

//go:linkname LastIndexByte strings.LastIndexByte
//go:noescape
func LastIndexByte(s string, c byte) int

//go:linkname NewReader strings.NewReader
//go:noescape
func NewReader(s string) *strings.Reader

//go:linkname readersize strings.sub_readersize
func readersize(r *strings.Reader,) int64 {
	return r.Size()
}

//go:linkname ReaderSize strings.sub_readersize
//go:noescape
func ReaderSize(r *strings.Reader,) int64

//go:linkname HasSuffix strings.HasSuffix
//go:noescape
func HasSuffix(s, suffix string) bool

//go:linkname Map strings.Map
//go:noescape
func Map(mapping func(rune) rune, s string) string

//go:linkname Repeat strings.Repeat
//go:noescape
func Repeat(s string, count int) string

//go:linkname Contains strings.Contains
//go:noescape
func Contains(s, substr string) bool

//go:linkname TrimLeftFunc strings.TrimLeftFunc
//go:noescape
func TrimLeftFunc(s string, f func(rune) bool) string

//go:linkname TrimRight strings.TrimRight
//go:noescape
func TrimRight(s string, cutset string) string

//go:linkname builderstring strings.sub_builderstring
func builderstring(b *strings.Builder,) string {
	return b.String()
}

//go:linkname BuilderString strings.sub_builderstring
//go:noescape
func BuilderString(b *strings.Builder,) string

//go:linkname ContainsRune strings.ContainsRune
//go:noescape
func ContainsRune(s string, r rune) bool

//go:linkname LastIndexFunc strings.LastIndexFunc
//go:noescape
func LastIndexFunc(s string, f func(rune) bool) int

//go:linkname TrimPrefix strings.TrimPrefix
//go:noescape
func TrimPrefix(s, prefix string) string

//go:linkname TrimSpace strings.TrimSpace
//go:noescape
func TrimSpace(s string) string

//go:linkname NewReplacer strings.NewReplacer
//go:noescape
func NewReplacer(oldnew ...string) *strings.Replacer

//go:linkname Join strings.Join
//go:noescape
func Join(elems []string, sep string) string

//go:linkname SplitAfter strings.SplitAfter
//go:noescape
func SplitAfter(s, sep string) []string

//go:linkname builderwrite strings.sub_builderwrite
func builderwrite(b *strings.Builder, p []byte) (int, error) {
	return b.Write(p)
}

//go:linkname BuilderWrite strings.sub_builderwrite
//go:noescape
func BuilderWrite(b *strings.Builder, p []byte) (int, error)

//go:linkname builderwritestring strings.sub_builderwritestring
func builderwritestring(b *strings.Builder, s string) (int, error) {
	return b.WriteString(s)
}

//go:linkname BuilderWriteString strings.sub_builderwritestring
//go:noescape
func BuilderWriteString(b *strings.Builder, s string) (int, error)

//go:linkname replacerreplace strings.sub_replacerreplace
func replacerreplace(r *strings.Replacer, s string) string {
	return r.Replace(s)
}

//go:linkname ReplacerReplace strings.sub_replacerreplace
//go:noescape
func ReplacerReplace(r *strings.Replacer, s string) string

//go:linkname ContainsAny strings.ContainsAny
//go:noescape
func ContainsAny(s, chars string) bool

//go:linkname Index strings.Index
//go:noescape
func Index(s, substr string) int

//go:linkname IndexRune strings.IndexRune
//go:noescape
func IndexRune(s string, r rune) int

//go:linkname TrimSuffix strings.TrimSuffix
//go:noescape
func TrimSuffix(s, suffix string) string

//go:linkname builderwritebyte strings.sub_builderwritebyte
func builderwritebyte(b *strings.Builder, c byte) error {
	return b.WriteByte(c)
}

//go:linkname BuilderWriteByte strings.sub_builderwritebyte
//go:noescape
func BuilderWriteByte(b *strings.Builder, c byte) error

//go:linkname readerreadrune strings.sub_readerreadrune
func readerreadrune(r *strings.Reader,) (rune, int, error) {
	return r.ReadRune()
}

//go:linkname ReaderReadRune strings.sub_readerreadrune
//go:noescape
func ReaderReadRune(r *strings.Reader,) (rune, int, error)

//go:linkname Compare strings.Compare
//go:noescape
func Compare(a, b string) int

//go:linkname Split strings.Split
//go:noescape
func Split(s, sep string) []string

//go:linkname SplitN strings.SplitN
//go:noescape
func SplitN(s, sep string, n int) []string

//go:linkname builderwriterune strings.sub_builderwriterune
func builderwriterune(b *strings.Builder, r rune) (int, error) {
	return b.WriteRune(r)
}

//go:linkname BuilderWriteRune strings.sub_builderwriterune
//go:noescape
func BuilderWriteRune(b *strings.Builder, r rune) (int, error)

//go:linkname FieldsFunc strings.FieldsFunc
//go:noescape
func FieldsFunc(s string, f func(rune) bool) []string

//go:linkname IndexByte strings.IndexByte
//go:noescape
func IndexByte(s string, c byte) int

//go:linkname LastIndexAny strings.LastIndexAny
//go:noescape
func LastIndexAny(s, chars string) int

//go:linkname ReplaceAll strings.ReplaceAll
//go:noescape
func ReplaceAll(s, old, new string) string

//go:linkname Trim strings.Trim
//go:noescape
func Trim(s string, cutset string) string

//go:linkname TrimRightFunc strings.TrimRightFunc
//go:noescape
func TrimRightFunc(s string, f func(rune) bool) string

//go:linkname IndexAny strings.IndexAny
//go:noescape
func IndexAny(s, chars string) int

//go:linkname Replace strings.Replace
//go:noescape
func Replace(s, old, new string, n int) string

//go:linkname ToLower strings.ToLower
//go:noescape
func ToLower(s string) string

//go:linkname ToUpperSpecial strings.ToUpperSpecial
//go:noescape
func ToUpperSpecial(c unicode.SpecialCase, s string) string
