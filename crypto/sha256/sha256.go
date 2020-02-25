// This file has automatically been generated on Wed Feb 26 02:09:46 +05 2020.
// DO NOT EDIT.
package sha256

import (
	_ "crypto/sha256"
	"hash"
	_ "unsafe"
)

//go:linkname New crypto/sha256.New
//go:noescape
func New() hash.Hash
