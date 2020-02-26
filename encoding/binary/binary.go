// This file has automatically been generated on Wed Feb 26 15:50:29 +05 2020.
// DO NOT EDIT.
package binary

import (
	"encoding/binary"
	"io"
	_ "unsafe"
)

//go:linkname Uvarint encoding/binary.Uvarint
//go:noescape
func Uvarint(buf []byte) (uint64, int)

//go:linkname Varint encoding/binary.Varint
//go:noescape
func Varint(buf []byte) (int64, int)

//go:linkname PutUvarint encoding/binary.PutUvarint
//go:noescape
func PutUvarint(buf []byte, x uint64) int

//go:linkname ReadUvarint encoding/binary.ReadUvarint
//go:noescape
func ReadUvarint(r io.ByteReader) (uint64, error)

//go:linkname ReadVarint encoding/binary.ReadVarint
//go:noescape
func ReadVarint(r io.ByteReader) (int64, error)

//go:linkname Size encoding/binary.Size
//go:noescape
func Size(v interface{}) int

//go:linkname Write encoding/binary.Write
//go:noescape
func Write(w io.Writer, order binary.ByteOrder, data interface{}) error

//go:linkname PutVarint encoding/binary.PutVarint
//go:noescape
func PutVarint(buf []byte, x int64) int

//go:linkname Read encoding/binary.Read
//go:noescape
func Read(r io.Reader, order binary.ByteOrder, data interface{}) error
