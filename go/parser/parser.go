// This file has automatically been generated on Wed Feb 26 02:09:56 +05 2020.
// DO NOT EDIT.
package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	_ "unsafe"
)

//go:linkname ParseDir go/parser.ParseDir
//go:noescape
func ParseDir(fset *token.FileSet, path string, filter func(os.FileInfo) bool, mode parser.Mode) (map[string]*ast.Package, error)

//go:linkname ParseExpr go/parser.ParseExpr
//go:noescape
func ParseExpr(x string) (ast.Expr, error)

//go:linkname ParseExprFrom go/parser.ParseExprFrom
//go:noescape
func ParseExprFrom(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (ast.Expr, error)

//go:linkname ParseFile go/parser.ParseFile
//go:noescape
func ParseFile(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (*ast.File, error)
