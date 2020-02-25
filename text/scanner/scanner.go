// This file has automatically been generated on Wed Feb 26 02:10:16 +05 2020.
// DO NOT EDIT.
package scanner

import (
	"io"
	"text/scanner"
	_ "unsafe"
)

//go:linkname TokenString text/scanner.TokenString
//go:noescape
func TokenString(tok rune) string

//go:linkname positionstring scanner.sub_positionstring
func positionstring(pos scanner.Position) string {
	return pos.String()
}

//go:linkname PositionString scanner.sub_positionstring
//go:noescape
func PositionString(pos scanner.Position) string

//go:linkname scannerinit scanner.sub_scannerinit
func scannerinit(s *scanner.Scanner, src io.Reader) *scanner.Scanner {
	return s.Init(src)
}

//go:linkname ScannerInit scanner.sub_scannerinit
//go:noescape
func ScannerInit(s *scanner.Scanner, src io.Reader) *scanner.Scanner

//go:linkname scannerpeek scanner.sub_scannerpeek
func scannerpeek(s *scanner.Scanner) rune {
	return s.Peek()
}

//go:linkname ScannerPeek scanner.sub_scannerpeek
//go:noescape
func ScannerPeek(s *scanner.Scanner) rune

//go:linkname scannerpos scanner.sub_scannerpos
func scannerpos(s *scanner.Scanner) scanner.Position {
	return s.Pos()
}

//go:linkname ScannerPos scanner.sub_scannerpos
//go:noescape
func ScannerPos(s *scanner.Scanner) scanner.Position

//go:linkname scannerscan scanner.sub_scannerscan
func scannerscan(s *scanner.Scanner) rune {
	return s.Scan()
}

//go:linkname ScannerScan scanner.sub_scannerscan
//go:noescape
func ScannerScan(s *scanner.Scanner) rune

//go:linkname scannertokentext scanner.sub_scannertokentext
func scannertokentext(s *scanner.Scanner) string {
	return s.TokenText()
}

//go:linkname ScannerTokenText scanner.sub_scannertokentext
//go:noescape
func ScannerTokenText(s *scanner.Scanner) string

//go:linkname positionisvalid scanner.sub_positionisvalid
func positionisvalid(pos *scanner.Position) bool {
	return pos.IsValid()
}

//go:linkname PositionIsValid scanner.sub_positionisvalid
//go:noescape
func PositionIsValid(pos *scanner.Position) bool

//go:linkname scannernext scanner.sub_scannernext
func scannernext(s *scanner.Scanner) rune {
	return s.Next()
}

//go:linkname ScannerNext scanner.sub_scannernext
//go:noescape
func ScannerNext(s *scanner.Scanner) rune
