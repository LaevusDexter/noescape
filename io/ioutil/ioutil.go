// This file has automatically been generated on Wed Feb 26 02:10:02 +05 2020.
// DO NOT EDIT.
package ioutil

import (
	"io"
	_ "io/ioutil"
	"os"
	_ "unsafe"
)

//go:linkname WriteFile io/ioutil.WriteFile
//go:noescape
func WriteFile(filename string, data []byte, perm os.FileMode) error

//go:linkname NopCloser io/ioutil.NopCloser
//go:noescape
func NopCloser(r io.Reader) io.ReadCloser

//go:linkname ReadAll io/ioutil.ReadAll
//go:noescape
func ReadAll(r io.Reader) ([]byte, error)

//go:linkname ReadDir io/ioutil.ReadDir
//go:noescape
func ReadDir(dirname string) ([]os.FileInfo, error)

//go:linkname ReadFile io/ioutil.ReadFile
//go:noescape
func ReadFile(filename string) ([]byte, error)

//go:linkname TempDir io/ioutil.TempDir
//go:noescape
func TempDir(dir, pattern string) (string, error)

//go:linkname TempFile io/ioutil.TempFile
//go:noescape
func TempFile(dir, pattern string) (*os.File, error)
