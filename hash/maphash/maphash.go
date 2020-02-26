// This file has automatically been generated on Wed Feb 26 15:50:36 +05 2020.
// DO NOT EDIT.
package maphash

import (
	"hash/maphash"
	_ "unsafe"
)

//go:linkname hashwritebyte maphash.sub_hashwritebyte
func hashwritebyte(h *maphash.Hash, b byte) error {
	return h.WriteByte(b)
}

//go:linkname HashWriteByte maphash.sub_hashwritebyte
//go:noescape
func HashWriteByte(h *maphash.Hash, b byte) error

//go:linkname hashwritestring maphash.sub_hashwritestring
func hashwritestring(h *maphash.Hash, s string) (int, error) {
	return h.WriteString(s)
}

//go:linkname HashWriteString maphash.sub_hashwritestring
//go:noescape
func HashWriteString(h *maphash.Hash, s string) (int, error)

//go:linkname MakeSeed hash/maphash.MakeSeed
//go:noescape
func MakeSeed() maphash.Seed

//go:linkname hashblocksize maphash.sub_hashblocksize
func hashblocksize(h *maphash.Hash,) int {
	return h.BlockSize()
}

//go:linkname HashBlockSize maphash.sub_hashblocksize
//go:noescape
func HashBlockSize(h *maphash.Hash,) int

//go:linkname hashseed maphash.sub_hashseed
func hashseed(h *maphash.Hash,) maphash.Seed {
	return h.Seed()
}

//go:linkname HashSeed maphash.sub_hashseed
//go:noescape
func HashSeed(h *maphash.Hash,) maphash.Seed

//go:linkname hashsize maphash.sub_hashsize
func hashsize(h *maphash.Hash,) int {
	return h.Size()
}

//go:linkname HashSize maphash.sub_hashsize
//go:noescape
func HashSize(h *maphash.Hash,) int

//go:linkname hashsum maphash.sub_hashsum
func hashsum(h *maphash.Hash, b []byte) []byte {
	return h.Sum(b)
}

//go:linkname HashSum maphash.sub_hashsum
//go:noescape
func HashSum(h *maphash.Hash, b []byte) []byte

//go:linkname hashwrite maphash.sub_hashwrite
func hashwrite(h *maphash.Hash, b []byte) (int, error) {
	return h.Write(b)
}

//go:linkname HashWrite maphash.sub_hashwrite
//go:noescape
func HashWrite(h *maphash.Hash, b []byte) (int, error)
