// This file has automatically been generated on Wed Feb 26 02:09:57 +05 2020.
// DO NOT EDIT.
package scanner

import (
	"go/scanner"
	"go/token"
	_ "unsafe"
)

//go:linkname errorerror scanner.sub_errorerror
func errorerror(e scanner.Error) string {
	return e.Error()
}

//go:linkname ErrorError scanner.sub_errorerror
//go:noescape
func ErrorError(e scanner.Error) string

//go:linkname errorlisterr scanner.sub_errorlisterr
func errorlisterr(p scanner.ErrorList) error {
	return p.Err()
}

//go:linkname ErrorListErr scanner.sub_errorlisterr
//go:noescape
func ErrorListErr(p scanner.ErrorList) error

//go:linkname errorlisterror scanner.sub_errorlisterror
func errorlisterror(p scanner.ErrorList) string {
	return p.Error()
}

//go:linkname ErrorListError scanner.sub_errorlisterror
//go:noescape
func ErrorListError(p scanner.ErrorList) string

//go:linkname errorlistlen scanner.sub_errorlistlen
func errorlistlen(p scanner.ErrorList) int {
	return p.Len()
}

//go:linkname ErrorListLen scanner.sub_errorlistlen
//go:noescape
func ErrorListLen(p scanner.ErrorList) int

//go:linkname errorlistless scanner.sub_errorlistless
func errorlistless(p scanner.ErrorList, i, j int) bool {
	return p.Less(i, j)
}

//go:linkname ErrorListLess scanner.sub_errorlistless
//go:noescape
func ErrorListLess(p scanner.ErrorList, i, j int) bool

//go:linkname scannerscan scanner.sub_scannerscan
func scannerscan(s *scanner.Scanner) (token.Pos, token.Token, string) {
	return s.Scan()
}

//go:linkname ScannerScan scanner.sub_scannerscan
//go:noescape
func ScannerScan(s *scanner.Scanner) (token.Pos, token.Token, string)
