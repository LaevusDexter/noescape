// This file has automatically been generated on Wed Feb 26 15:50:35 +05 2020.
// DO NOT EDIT.
package token

import (
	"go/token"
	_ "unsafe"
)

//go:linkname NewFileSet go/token.NewFileSet
//go:noescape
func NewFileSet() *token.FileSet

//go:linkname posisvalid token.sub_posisvalid
func posisvalid(p token.Pos,) bool {
	return p.IsValid()
}

//go:linkname PosIsValid token.sub_posisvalid
//go:noescape
func PosIsValid(p token.Pos,) bool

//go:linkname tokenisoperator token.sub_tokenisoperator
func tokenisoperator(tok token.Token,) bool {
	return tok.IsOperator()
}

//go:linkname TokenIsOperator token.sub_tokenisoperator
//go:noescape
func TokenIsOperator(tok token.Token,) bool

//go:linkname filepos token.sub_filepos
func filepos(f *token.File, offset int) token.Pos {
	return f.Pos(offset)
}

//go:linkname FilePos token.sub_filepos
//go:noescape
func FilePos(f *token.File, offset int) token.Pos

//go:linkname filesetbase token.sub_filesetbase
func filesetbase(s *token.FileSet,) int {
	return s.Base()
}

//go:linkname FileSetBase token.sub_filesetbase
//go:noescape
func FileSetBase(s *token.FileSet,) int

//go:linkname filesetfile token.sub_filesetfile
func filesetfile(s *token.FileSet, p token.Pos,) *token.File {
	return s.File(p)
}

//go:linkname FileSetFile token.sub_filesetfile
//go:noescape
func FileSetFile(s *token.FileSet, p token.Pos,) *token.File

//go:linkname filesetiterate token.sub_filesetiterate
func filesetiterate(s *token.FileSet, f func(*token.File,) bool) {
	s.Iterate(f)
}

//go:linkname FileSetIterate token.sub_filesetiterate
//go:noescape
func FileSetIterate(s *token.FileSet, f func(*token.File,) bool)

//go:linkname filesetposition token.sub_filesetposition
func filesetposition(s *token.FileSet, p token.Pos,) token.Position {
	return s.Position(p)
}

//go:linkname FileSetPosition token.sub_filesetposition
//go:noescape
func FileSetPosition(s *token.FileSet, p token.Pos,) token.Position

//go:linkname filesetpositionfor token.sub_filesetpositionfor
func filesetpositionfor(s *token.FileSet, p token.Pos, adjusted bool) token.Position {
	return s.PositionFor(p, adjusted)
}

//go:linkname FileSetPositionFor token.sub_filesetpositionfor
//go:noescape
func FileSetPositionFor(s *token.FileSet, p token.Pos, adjusted bool) token.Position

//go:linkname tokenprecedence token.sub_tokenprecedence
func tokenprecedence(op token.Token,) int {
	return op.Precedence()
}

//go:linkname TokenPrecedence token.sub_tokenprecedence
//go:noescape
func TokenPrecedence(op token.Token,) int

//go:linkname IsExported go/token.IsExported
//go:noescape
func IsExported(name string) bool

//go:linkname IsIdentifier go/token.IsIdentifier
//go:noescape
func IsIdentifier(name string) bool

//go:linkname filesetlines token.sub_filesetlines
func filesetlines(f *token.File, lines []int) bool {
	return f.SetLines(lines)
}

//go:linkname FileSetLines token.sub_filesetlines
//go:noescape
func FileSetLines(f *token.File, lines []int) bool

//go:linkname positionisvalid token.sub_positionisvalid
func positionisvalid(pos *token.Position,) bool {
	return pos.IsValid()
}

//go:linkname PositionIsValid token.sub_positionisvalid
//go:noescape
func PositionIsValid(pos *token.Position,) bool

//go:linkname tokenstring token.sub_tokenstring
func tokenstring(tok token.Token,) string {
	return tok.String()
}

//go:linkname TokenString token.sub_tokenstring
//go:noescape
func TokenString(tok token.Token,) string

//go:linkname filebase token.sub_filebase
func filebase(f *token.File,) int {
	return f.Base()
}

//go:linkname FileBase token.sub_filebase
//go:noescape
func FileBase(f *token.File,) int

//go:linkname filesetaddfile token.sub_filesetaddfile
func filesetaddfile(s *token.FileSet, filename string, base, size int) *token.File {
	return s.AddFile(filename, base, size)
}

//go:linkname FileSetAddFile token.sub_filesetaddfile
//go:noescape
func FileSetAddFile(s *token.FileSet, filename string, base, size int) *token.File

//go:linkname IsKeyword go/token.IsKeyword
//go:noescape
func IsKeyword(name string) bool

//go:linkname fileposition token.sub_fileposition
func fileposition(f *token.File, p token.Pos,) token.Position {
	return f.Position(p)
}

//go:linkname FilePosition token.sub_fileposition
//go:noescape
func FilePosition(f *token.File, p token.Pos,) token.Position

//go:linkname fileline token.sub_fileline
func fileline(f *token.File, p token.Pos,) int {
	return f.Line(p)
}

//go:linkname FileLine token.sub_fileline
//go:noescape
func FileLine(f *token.File, p token.Pos,) int

//go:linkname filelinecount token.sub_filelinecount
func filelinecount(f *token.File,) int {
	return f.LineCount()
}

//go:linkname FileLineCount token.sub_filelinecount
//go:noescape
func FileLineCount(f *token.File,) int

//go:linkname filelinestart token.sub_filelinestart
func filelinestart(f *token.File, line int) token.Pos {
	return f.LineStart(line)
}

//go:linkname FileLineStart token.sub_filelinestart
//go:noescape
func FileLineStart(f *token.File, line int) token.Pos

//go:linkname fileoffset token.sub_fileoffset
func fileoffset(f *token.File, p token.Pos,) int {
	return f.Offset(p)
}

//go:linkname FileOffset token.sub_fileoffset
//go:noescape
func FileOffset(f *token.File, p token.Pos,) int

//go:linkname filesetread token.sub_filesetread
func filesetread(s *token.FileSet, decode func(interface{}) error) error {
	return s.Read(decode)
}

//go:linkname FileSetRead token.sub_filesetread
//go:noescape
func FileSetRead(s *token.FileSet, decode func(interface{}) error) error

//go:linkname filesetwrite token.sub_filesetwrite
func filesetwrite(s *token.FileSet, encode func(interface{}) error) error {
	return s.Write(encode)
}

//go:linkname FileSetWrite token.sub_filesetwrite
//go:noescape
func FileSetWrite(s *token.FileSet, encode func(interface{}) error) error

//go:linkname positionstring token.sub_positionstring
func positionstring(pos token.Position,) string {
	return pos.String()
}

//go:linkname PositionString token.sub_positionstring
//go:noescape
func PositionString(pos token.Position,) string

//go:linkname filename token.sub_filename
func filename(f *token.File,) string {
	return f.Name()
}

//go:linkname FileName token.sub_filename
//go:noescape
func FileName(f *token.File,) string

//go:linkname filepositionfor token.sub_filepositionfor
func filepositionfor(f *token.File, p token.Pos, adjusted bool) token.Position {
	return f.PositionFor(p, adjusted)
}

//go:linkname FilePositionFor token.sub_filepositionfor
//go:noescape
func FilePositionFor(f *token.File, p token.Pos, adjusted bool) token.Position

//go:linkname filesize token.sub_filesize
func filesize(f *token.File,) int {
	return f.Size()
}

//go:linkname FileSize token.sub_filesize
//go:noescape
func FileSize(f *token.File,) int

//go:linkname Lookup go/token.Lookup
//go:noescape
func Lookup(ident string) token.Token

//go:linkname tokeniskeyword token.sub_tokeniskeyword
func tokeniskeyword(tok token.Token,) bool {
	return tok.IsKeyword()
}

//go:linkname TokenIsKeyword token.sub_tokeniskeyword
//go:noescape
func TokenIsKeyword(tok token.Token,) bool

//go:linkname tokenisliteral token.sub_tokenisliteral
func tokenisliteral(tok token.Token,) bool {
	return tok.IsLiteral()
}

//go:linkname TokenIsLiteral token.sub_tokenisliteral
//go:noescape
func TokenIsLiteral(tok token.Token,) bool
