// This file has automatically been generated on Wed Feb 26 15:50:30 +05 2020.
// DO NOT EDIT.
package pem

import (
	"encoding/pem"
	"io"
	_ "unsafe"
)

//go:linkname Encode encoding/pem.Encode
//go:noescape
func Encode(out io.Writer, b *pem.Block,) error

//go:linkname EncodeToMemory encoding/pem.EncodeToMemory
//go:noescape
func EncodeToMemory(b *pem.Block,) []byte

//go:linkname Decode encoding/pem.Decode
//go:noescape
func Decode(data []byte) (*pem.Block, []byte)
