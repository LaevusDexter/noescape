// This file has automatically been generated on Wed Feb 26 15:50:49 +05 2020.
// DO NOT EDIT.
package regexp

import (
	"io"
	"regexp"
	_ "unsafe"
)

//go:linkname regexpfindallstringsubmatchindex regexp.sub_regexpfindallstringsubmatchindex
func regexpfindallstringsubmatchindex(re *regexp.Regexp, s string, n int) [][]int {
	return re.FindAllStringSubmatchIndex(s, n)
}

//go:linkname RegexpFindAllStringSubmatchIndex regexp.sub_regexpfindallstringsubmatchindex
//go:noescape
func RegexpFindAllStringSubmatchIndex(re *regexp.Regexp, s string, n int) [][]int

//go:linkname MatchReader regexp.MatchReader
//go:noescape
func MatchReader(pattern string, r io.RuneReader) (bool, error)

//go:linkname Compile regexp.Compile
//go:noescape
func Compile(expr string) (*regexp.Regexp, error)

//go:linkname regexpfindall regexp.sub_regexpfindall
func regexpfindall(re *regexp.Regexp, b []byte, n int) [][]byte {
	return re.FindAll(b, n)
}

//go:linkname RegexpFindAll regexp.sub_regexpfindall
//go:noescape
func RegexpFindAll(re *regexp.Regexp, b []byte, n int) [][]byte

//go:linkname regexpfindallstringindex regexp.sub_regexpfindallstringindex
func regexpfindallstringindex(re *regexp.Regexp, s string, n int) [][]int {
	return re.FindAllStringIndex(s, n)
}

//go:linkname RegexpFindAllStringIndex regexp.sub_regexpfindallstringindex
//go:noescape
func RegexpFindAllStringIndex(re *regexp.Regexp, s string, n int) [][]int

//go:linkname regexpfindallsubmatchindex regexp.sub_regexpfindallsubmatchindex
func regexpfindallsubmatchindex(re *regexp.Regexp, b []byte, n int) [][]int {
	return re.FindAllSubmatchIndex(b, n)
}

//go:linkname RegexpFindAllSubmatchIndex regexp.sub_regexpfindallsubmatchindex
//go:noescape
func RegexpFindAllSubmatchIndex(re *regexp.Regexp, b []byte, n int) [][]int

//go:linkname regexpfindreaderindex regexp.sub_regexpfindreaderindex
func regexpfindreaderindex(re *regexp.Regexp, r io.RuneReader) []int {
	return re.FindReaderIndex(r)
}

//go:linkname RegexpFindReaderIndex regexp.sub_regexpfindreaderindex
//go:noescape
func RegexpFindReaderIndex(re *regexp.Regexp, r io.RuneReader) []int

//go:linkname regexpfindstring regexp.sub_regexpfindstring
func regexpfindstring(re *regexp.Regexp, s string) string {
	return re.FindString(s)
}

//go:linkname RegexpFindString regexp.sub_regexpfindstring
//go:noescape
func RegexpFindString(re *regexp.Regexp, s string) string

//go:linkname regexpfindsubmatch regexp.sub_regexpfindsubmatch
func regexpfindsubmatch(re *regexp.Regexp, b []byte) [][]byte {
	return re.FindSubmatch(b)
}

//go:linkname RegexpFindSubmatch regexp.sub_regexpfindsubmatch
//go:noescape
func RegexpFindSubmatch(re *regexp.Regexp, b []byte) [][]byte

//go:linkname MatchString regexp.MatchString
//go:noescape
func MatchString(pattern string, s string) (bool, error)

//go:linkname QuoteMeta regexp.QuoteMeta
//go:noescape
func QuoteMeta(s string) string

//go:linkname regexpfind regexp.sub_regexpfind
func regexpfind(re *regexp.Regexp, b []byte) []byte {
	return re.Find(b)
}

//go:linkname RegexpFind regexp.sub_regexpfind
//go:noescape
func RegexpFind(re *regexp.Regexp, b []byte) []byte

//go:linkname regexpfindallsubmatch regexp.sub_regexpfindallsubmatch
func regexpfindallsubmatch(re *regexp.Regexp, b []byte, n int) [][][]byte {
	return re.FindAllSubmatch(b, n)
}

//go:linkname RegexpFindAllSubmatch regexp.sub_regexpfindallsubmatch
//go:noescape
func RegexpFindAllSubmatch(re *regexp.Regexp, b []byte, n int) [][][]byte

//go:linkname regexpsplit regexp.sub_regexpsplit
func regexpsplit(re *regexp.Regexp, s string, n int) []string {
	return re.Split(s, n)
}

//go:linkname RegexpSplit regexp.sub_regexpsplit
//go:noescape
func RegexpSplit(re *regexp.Regexp, s string, n int) []string

//go:linkname regexpmatchstring regexp.sub_regexpmatchstring
func regexpmatchstring(re *regexp.Regexp, s string) bool {
	return re.MatchString(s)
}

//go:linkname RegexpMatchString regexp.sub_regexpmatchstring
//go:noescape
func RegexpMatchString(re *regexp.Regexp, s string) bool

//go:linkname regexpnumsubexp regexp.sub_regexpnumsubexp
func regexpnumsubexp(re *regexp.Regexp,) int {
	return re.NumSubexp()
}

//go:linkname RegexpNumSubexp regexp.sub_regexpnumsubexp
//go:noescape
func RegexpNumSubexp(re *regexp.Regexp,) int

//go:linkname regexpreplaceallliteral regexp.sub_regexpreplaceallliteral
func regexpreplaceallliteral(re *regexp.Regexp, src, repl []byte) []byte {
	return re.ReplaceAllLiteral(src, repl)
}

//go:linkname RegexpReplaceAllLiteral regexp.sub_regexpreplaceallliteral
//go:noescape
func RegexpReplaceAllLiteral(re *regexp.Regexp, src, repl []byte) []byte

//go:linkname regexpreplaceallliteralstring regexp.sub_regexpreplaceallliteralstring
func regexpreplaceallliteralstring(re *regexp.Regexp, src, repl string) string {
	return re.ReplaceAllLiteralString(src, repl)
}

//go:linkname RegexpReplaceAllLiteralString regexp.sub_regexpreplaceallliteralstring
//go:noescape
func RegexpReplaceAllLiteralString(re *regexp.Regexp, src, repl string) string

//go:linkname regexpreplaceallstring regexp.sub_regexpreplaceallstring
func regexpreplaceallstring(re *regexp.Regexp, src, repl string) string {
	return re.ReplaceAllString(src, repl)
}

//go:linkname RegexpReplaceAllString regexp.sub_regexpreplaceallstring
//go:noescape
func RegexpReplaceAllString(re *regexp.Regexp, src, repl string) string

//go:linkname MustCompilePOSIX regexp.MustCompilePOSIX
//go:noescape
func MustCompilePOSIX(str string) *regexp.Regexp

//go:linkname regexpcopy regexp.sub_regexpcopy
func regexpcopy(re *regexp.Regexp,) *regexp.Regexp {
	return re.Copy()
}

//go:linkname RegexpCopy regexp.sub_regexpcopy
//go:noescape
func RegexpCopy(re *regexp.Regexp,) *regexp.Regexp

//go:linkname regexpfindstringsubmatchindex regexp.sub_regexpfindstringsubmatchindex
func regexpfindstringsubmatchindex(re *regexp.Regexp, s string) []int {
	return re.FindStringSubmatchIndex(s)
}

//go:linkname RegexpFindStringSubmatchIndex regexp.sub_regexpfindstringsubmatchindex
//go:noescape
func RegexpFindStringSubmatchIndex(re *regexp.Regexp, s string) []int

//go:linkname regexpliteralprefix regexp.sub_regexpliteralprefix
func regexpliteralprefix(re *regexp.Regexp,) (string, bool) {
	return re.LiteralPrefix()
}

//go:linkname RegexpLiteralPrefix regexp.sub_regexpliteralprefix
//go:noescape
func RegexpLiteralPrefix(re *regexp.Regexp,) (string, bool)

//go:linkname regexpsubexpnames regexp.sub_regexpsubexpnames
func regexpsubexpnames(re *regexp.Regexp,) []string {
	return re.SubexpNames()
}

//go:linkname RegexpSubexpNames regexp.sub_regexpsubexpnames
//go:noescape
func RegexpSubexpNames(re *regexp.Regexp,) []string

//go:linkname regexpmatch regexp.sub_regexpmatch
func regexpmatch(re *regexp.Regexp, b []byte) bool {
	return re.Match(b)
}

//go:linkname RegexpMatch regexp.sub_regexpmatch
//go:noescape
func RegexpMatch(re *regexp.Regexp, b []byte) bool

//go:linkname regexpmatchreader regexp.sub_regexpmatchreader
func regexpmatchreader(re *regexp.Regexp, r io.RuneReader) bool {
	return re.MatchReader(r)
}

//go:linkname RegexpMatchReader regexp.sub_regexpmatchreader
//go:noescape
func RegexpMatchReader(re *regexp.Regexp, r io.RuneReader) bool

//go:linkname regexpreplaceallfunc regexp.sub_regexpreplaceallfunc
func regexpreplaceallfunc(re *regexp.Regexp, src []byte, repl func([]byte) []byte) []byte {
	return re.ReplaceAllFunc(src, repl)
}

//go:linkname RegexpReplaceAllFunc regexp.sub_regexpreplaceallfunc
//go:noescape
func RegexpReplaceAllFunc(re *regexp.Regexp, src []byte, repl func([]byte) []byte) []byte

//go:linkname regexpreplaceallstringfunc regexp.sub_regexpreplaceallstringfunc
func regexpreplaceallstringfunc(re *regexp.Regexp, src string, repl func(string) string) string {
	return re.ReplaceAllStringFunc(src, repl)
}

//go:linkname RegexpReplaceAllStringFunc regexp.sub_regexpreplaceallstringfunc
//go:noescape
func RegexpReplaceAllStringFunc(re *regexp.Regexp, src string, repl func(string) string) string

//go:linkname CompilePOSIX regexp.CompilePOSIX
//go:noescape
func CompilePOSIX(expr string) (*regexp.Regexp, error)

//go:linkname regexpfindallindex regexp.sub_regexpfindallindex
func regexpfindallindex(re *regexp.Regexp, b []byte, n int) [][]int {
	return re.FindAllIndex(b, n)
}

//go:linkname RegexpFindAllIndex regexp.sub_regexpfindallindex
//go:noescape
func RegexpFindAllIndex(re *regexp.Regexp, b []byte, n int) [][]int

//go:linkname regexpreplaceall regexp.sub_regexpreplaceall
func regexpreplaceall(re *regexp.Regexp, src, repl []byte) []byte {
	return re.ReplaceAll(src, repl)
}

//go:linkname RegexpReplaceAll regexp.sub_regexpreplaceall
//go:noescape
func RegexpReplaceAll(re *regexp.Regexp, src, repl []byte) []byte

//go:linkname regexpfindstringsubmatch regexp.sub_regexpfindstringsubmatch
func regexpfindstringsubmatch(re *regexp.Regexp, s string) []string {
	return re.FindStringSubmatch(s)
}

//go:linkname RegexpFindStringSubmatch regexp.sub_regexpfindstringsubmatch
//go:noescape
func RegexpFindStringSubmatch(re *regexp.Regexp, s string) []string

//go:linkname MustCompile regexp.MustCompile
//go:noescape
func MustCompile(str string) *regexp.Regexp

//go:linkname regexpexpand regexp.sub_regexpexpand
func regexpexpand(re *regexp.Regexp, dst []byte, template []byte, src []byte, match []int) []byte {
	return re.Expand(dst, template, src, match)
}

//go:linkname RegexpExpand regexp.sub_regexpexpand
//go:noescape
func RegexpExpand(re *regexp.Regexp, dst []byte, template []byte, src []byte, match []int) []byte

//go:linkname regexpexpandstring regexp.sub_regexpexpandstring
func regexpexpandstring(re *regexp.Regexp, dst []byte, template string, src string, match []int) []byte {
	return re.ExpandString(dst, template, src, match)
}

//go:linkname RegexpExpandString regexp.sub_regexpexpandstring
//go:noescape
func RegexpExpandString(re *regexp.Regexp, dst []byte, template string, src string, match []int) []byte

//go:linkname regexpfindallstring regexp.sub_regexpfindallstring
func regexpfindallstring(re *regexp.Regexp, s string, n int) []string {
	return re.FindAllString(s, n)
}

//go:linkname RegexpFindAllString regexp.sub_regexpfindallstring
//go:noescape
func RegexpFindAllString(re *regexp.Regexp, s string, n int) []string

//go:linkname Match regexp.Match
//go:noescape
func Match(pattern string, b []byte) (bool, error)

//go:linkname regexpfindindex regexp.sub_regexpfindindex
func regexpfindindex(re *regexp.Regexp, b []byte) []int {
	return re.FindIndex(b)
}

//go:linkname RegexpFindIndex regexp.sub_regexpfindindex
//go:noescape
func RegexpFindIndex(re *regexp.Regexp, b []byte) []int

//go:linkname regexpfindreadersubmatchindex regexp.sub_regexpfindreadersubmatchindex
func regexpfindreadersubmatchindex(re *regexp.Regexp, r io.RuneReader) []int {
	return re.FindReaderSubmatchIndex(r)
}

//go:linkname RegexpFindReaderSubmatchIndex regexp.sub_regexpfindreadersubmatchindex
//go:noescape
func RegexpFindReaderSubmatchIndex(re *regexp.Regexp, r io.RuneReader) []int

//go:linkname regexpfindstringindex regexp.sub_regexpfindstringindex
func regexpfindstringindex(re *regexp.Regexp, s string) []int {
	return re.FindStringIndex(s)
}

//go:linkname RegexpFindStringIndex regexp.sub_regexpfindstringindex
//go:noescape
func RegexpFindStringIndex(re *regexp.Regexp, s string) []int

//go:linkname regexpfindallstringsubmatch regexp.sub_regexpfindallstringsubmatch
func regexpfindallstringsubmatch(re *regexp.Regexp, s string, n int) [][]string {
	return re.FindAllStringSubmatch(s, n)
}

//go:linkname RegexpFindAllStringSubmatch regexp.sub_regexpfindallstringsubmatch
//go:noescape
func RegexpFindAllStringSubmatch(re *regexp.Regexp, s string, n int) [][]string

//go:linkname regexpfindsubmatchindex regexp.sub_regexpfindsubmatchindex
func regexpfindsubmatchindex(re *regexp.Regexp, b []byte) []int {
	return re.FindSubmatchIndex(b)
}

//go:linkname RegexpFindSubmatchIndex regexp.sub_regexpfindsubmatchindex
//go:noescape
func RegexpFindSubmatchIndex(re *regexp.Regexp, b []byte) []int

//go:linkname regexpstring regexp.sub_regexpstring
func regexpstring(re *regexp.Regexp,) string {
	return re.String()
}

//go:linkname RegexpString regexp.sub_regexpstring
//go:noescape
func RegexpString(re *regexp.Regexp,) string
