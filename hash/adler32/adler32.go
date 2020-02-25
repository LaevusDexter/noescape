// This file has automatically been generated on Wed Feb 26 02:09:58 +05 2020.
// DO NOT EDIT.
package adler32

import (
	"hash"
	_ "hash/adler32"
	_ "unsafe"
)

//go:linkname Checksum hash/adler32.Checksum
//go:noescape
func Checksum(data []byte) uint32

//go:linkname New hash/adler32.New
//go:noescape
func New() hash.Hash32
