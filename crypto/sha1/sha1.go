// This file has automatically been generated on Wed Feb 26 15:50:23 +05 2020.
// DO NOT EDIT.
package sha1

import (
	"crypto/sha1"
	"hash"
	_ "unsafe"
)

//go:linkname New crypto/sha1.New
//go:noescape
func New() hash.Hash

//go:linkname Sum crypto/sha1.Sum
//go:noescape
func Sum(data []byte) [sha1.Size]byte
