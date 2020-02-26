// This file has automatically been generated on Wed Feb 26 15:50:36 +05 2020.
// DO NOT EDIT.
package crc32

import (
	"hash"
	"hash/crc32"
	_ "unsafe"
)

//go:linkname MakeTable hash/crc32.MakeTable
//go:noescape
func MakeTable(poly uint32) *crc32.Table

//go:linkname Checksum hash/crc32.Checksum
//go:noescape
func Checksum(data []byte, tab *crc32.Table,) uint32

//go:linkname ChecksumIEEE hash/crc32.ChecksumIEEE
//go:noescape
func ChecksumIEEE(data []byte) uint32

//go:linkname New hash/crc32.New
//go:noescape
func New(tab *crc32.Table,) hash.Hash32

//go:linkname NewIEEE hash/crc32.NewIEEE
//go:noescape
func NewIEEE() hash.Hash32

//go:linkname Update hash/crc32.Update
//go:noescape
func Update(crc uint32, tab *crc32.Table, p []byte) uint32
