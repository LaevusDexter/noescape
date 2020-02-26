// This file has automatically been generated on Wed Feb 26 15:50:22 +05 2020.
// DO NOT EDIT.
package md5

import (
	"crypto/md5"
	"hash"
	_ "unsafe"
)

//go:linkname New crypto/md5.New
//go:noescape
func New() hash.Hash

//go:linkname Sum crypto/md5.Sum
//go:noescape
func Sum(data []byte) [md5.Size]byte
