// This file has automatically been generated on Wed Feb 26 15:50:22 +05 2020.
// DO NOT EDIT.
package hmac

import (
	_ "crypto/hmac"
	"hash"
	_ "unsafe"
)

//go:linkname Equal crypto/hmac.Equal
//go:noescape
func Equal(mac1, mac2 []byte) bool

//go:linkname New crypto/hmac.New
//go:noescape
func New(h func() hash.Hash, key []byte) hash.Hash
