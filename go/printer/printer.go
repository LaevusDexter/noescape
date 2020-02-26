// This file has automatically been generated on Wed Feb 26 15:50:34 +05 2020.
// DO NOT EDIT.
package printer

import (
	"go/printer"
	"go/token"
	"io"
	_ "unsafe"
)

//go:linkname Fprint go/printer.Fprint
//go:noescape
func Fprint(output io.Writer, fset *token.FileSet, node interface{}) error

//go:linkname configfprint printer.sub_configfprint
func configfprint(cfg *printer.Config, output io.Writer, fset *token.FileSet, node interface{}) error {
	return cfg.Fprint(output, fset, node)
}

//go:linkname ConfigFprint printer.sub_configfprint
//go:noescape
func ConfigFprint(cfg *printer.Config, output io.Writer, fset *token.FileSet, node interface{}) error
