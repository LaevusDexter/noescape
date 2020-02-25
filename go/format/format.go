// This file has automatically been generated on Wed Feb 26 02:09:56 +05 2020.
// DO NOT EDIT.
package format

import (
	_ "go/format"
	"go/token"
	"io"
	_ "unsafe"
)

//go:linkname Node go/format.Node
//go:noescape
func Node(dst io.Writer, fset *token.FileSet, node interface{}) error

//go:linkname Source go/format.Source
//go:noescape
func Source(src []byte) ([]byte, error)
