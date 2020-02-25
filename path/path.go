// This file has automatically been generated on Wed Feb 26 02:10:10 +05 2020.
// DO NOT EDIT.
package path

import (
	_ "path"
	_ "unsafe"
)

//go:linkname Dir path.Dir
//go:noescape
func Dir(path string) string

//go:linkname Ext path.Ext
//go:noescape
func Ext(path string) string

//go:linkname IsAbs path.IsAbs
//go:noescape
func IsAbs(path string) bool

//go:linkname Join path.Join
//go:noescape
func Join(elem ...string) string

//go:linkname Match path.Match
//go:noescape
func Match(pattern, name string) (bool, error)

//go:linkname Split path.Split
//go:noescape
func Split(path string) string

//go:linkname Base path.Base
//go:noescape
func Base(path string) string

//go:linkname Clean path.Clean
//go:noescape
func Clean(path string) string
