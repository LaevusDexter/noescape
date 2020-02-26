// This file has automatically been generated on Wed Feb 26 15:50:20 +05 2020.
// DO NOT EDIT.
package crypto

import (
	"crypto"
	"hash"
	_ "unsafe"
)

//go:linkname hashnew crypto.sub_hashnew
func hashnew(h crypto.Hash,) hash.Hash {
	return h.New()
}

//go:linkname HashNew crypto.sub_hashnew
//go:noescape
func HashNew(h crypto.Hash,) hash.Hash

//go:linkname hashsize crypto.sub_hashsize
func hashsize(h crypto.Hash,) int {
	return h.Size()
}

//go:linkname HashSize crypto.sub_hashsize
//go:noescape
func HashSize(h crypto.Hash,) int

//go:linkname RegisterHash crypto.RegisterHash
//go:noescape
func RegisterHash(h crypto.Hash, f func() hash.Hash)

//go:linkname hashavailable crypto.sub_hashavailable
func hashavailable(h crypto.Hash,) bool {
	return h.Available()
}

//go:linkname HashAvailable crypto.sub_hashavailable
//go:noescape
func HashAvailable(h crypto.Hash,) bool

//go:linkname hashhashfunc crypto.sub_hashhashfunc
func hashhashfunc(h crypto.Hash,) crypto.Hash {
	return h.HashFunc()
}

//go:linkname HashHashFunc crypto.sub_hashhashfunc
//go:noescape
func HashHashFunc(h crypto.Hash,) crypto.Hash
