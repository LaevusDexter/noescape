// This file has automatically been generated on Wed Feb 26 15:50:48 +05 2020.
// DO NOT EDIT.
package filepath

import (
	"path/filepath"
	_ "unsafe"
)

//go:linkname IsAbs path/filepath.IsAbs
//go:noescape
func IsAbs(path string) bool

//go:linkname Join path/filepath.Join
//go:noescape
func Join(elem ...string) string

//go:linkname Split path/filepath.Split
//go:noescape
func Split(path string) (string, string)

//go:linkname Glob path/filepath.Glob
//go:noescape
func Glob(pattern string) ([]string, error)

//go:linkname Match path/filepath.Match
//go:noescape
func Match(pattern, name string) (bool, error)

//go:linkname Walk path/filepath.Walk
//go:noescape
func Walk(root string, walkFn filepath.WalkFunc,) error

//go:linkname Base path/filepath.Base
//go:noescape
func Base(path string) string

//go:linkname EvalSymlinks path/filepath.EvalSymlinks
//go:noescape
func EvalSymlinks(path string) (string, error)

//go:linkname FromSlash path/filepath.FromSlash
//go:noescape
func FromSlash(path string) string

//go:linkname VolumeName path/filepath.VolumeName
//go:noescape
func VolumeName(path string) string

//go:linkname Ext path/filepath.Ext
//go:noescape
func Ext(path string) string

//go:linkname Rel path/filepath.Rel
//go:noescape
func Rel(basepath, targpath string) (string, error)

//go:linkname SplitList path/filepath.SplitList
//go:noescape
func SplitList(path string) []string

//go:linkname HasPrefix path/filepath.HasPrefix
//go:noescape
func HasPrefix(p, prefix string) bool

//go:linkname ToSlash path/filepath.ToSlash
//go:noescape
func ToSlash(path string) string

//go:linkname Abs path/filepath.Abs
//go:noescape
func Abs(path string) (string, error)

//go:linkname Clean path/filepath.Clean
//go:noescape
func Clean(path string) string

//go:linkname Dir path/filepath.Dir
//go:noescape
func Dir(path string) string
