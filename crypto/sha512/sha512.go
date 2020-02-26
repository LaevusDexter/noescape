// This file has automatically been generated on Wed Feb 26 15:50:23 +05 2020.
// DO NOT EDIT.
package sha512

import (
	_ "crypto/sha512"
	"hash"
	_ "unsafe"
)

//go:linkname New crypto/sha512.New
//go:noescape
func New() hash.Hash
