// This file has automatically been generated on Wed Feb 26 02:09:59 +05 2020.
// DO NOT EDIT.
package crc64

import (
	"hash"
	"hash/crc64"
	_ "unsafe"
)

//go:linkname Checksum hash/crc64.Checksum
//go:noescape
func Checksum(data []byte, tab *crc64.Table) uint64

//go:linkname New hash/crc64.New
//go:noescape
func New(tab *crc64.Table) hash.Hash64

//go:linkname Update hash/crc64.Update
//go:noescape
func Update(crc uint64, tab *crc64.Table, p []byte) uint64

//go:linkname MakeTable hash/crc64.MakeTable
//go:noescape
func MakeTable(poly uint64) *crc64.Table
