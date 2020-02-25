// This file has automatically been generated on Wed Feb 26 02:09:56 +05 2020.
// DO NOT EDIT.
package doc

import (
	"go/ast"
	"go/doc"
	"go/token"
	_ "unsafe"
)

//go:linkname IsPredeclared go/doc.IsPredeclared
//go:noescape
func IsPredeclared(s string) bool

//go:linkname Synopsis go/doc.Synopsis
//go:noescape
func Synopsis(s string) string

//go:linkname Examples go/doc.Examples
//go:noescape
func Examples(testFiles ...*ast.File) []*doc.Example

//go:linkname New go/doc.New
//go:noescape
func New(pkg *ast.Package, importPath string, mode doc.Mode) *doc.Package

//go:linkname NewFromFiles go/doc.NewFromFiles
//go:noescape
func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...interface{}) (*doc.Package, error)
