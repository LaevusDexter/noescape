// This file has automatically been generated on Wed Feb 26 02:09:56 +05 2020.
// DO NOT EDIT.
package importer

import (
	"go/importer"
	"go/token"
	"go/types"
	_ "unsafe"
)

//go:linkname Default go/importer.Default
//go:noescape
func Default() types.Importer

//go:linkname For go/importer.For
//go:noescape
func For(compiler string, lookup importer.Lookup) types.Importer

//go:linkname ForCompiler go/importer.ForCompiler
//go:noescape
func ForCompiler(fset *token.FileSet, compiler string, lookup importer.Lookup) types.Importer
