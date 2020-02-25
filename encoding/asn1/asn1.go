// This file has automatically been generated on Wed Feb 26 02:09:51 +05 2020.
// DO NOT EDIT.
package asn1

import (
	"encoding/asn1"
	_ "unsafe"
)

//go:linkname MarshalWithParams encoding/asn1.MarshalWithParams
//go:noescape
func MarshalWithParams(val interface{}, params string) ([]byte, error)

//go:linkname bitstringat asn1.sub_bitstringat
func bitstringat(b asn1.BitString, i int) int {
	return b.At(i)
}

//go:linkname BitStringAt asn1.sub_bitstringat
//go:noescape
func BitStringAt(b asn1.BitString, i int) int

//go:linkname objectidentifierequal asn1.sub_objectidentifierequal
func objectidentifierequal(oi asn1.ObjectIdentifier, other asn1.ObjectIdentifier) bool {
	return oi.Equal(other)
}

//go:linkname ObjectIdentifierEqual asn1.sub_objectidentifierequal
//go:noescape
func ObjectIdentifierEqual(oi asn1.ObjectIdentifier, other asn1.ObjectIdentifier) bool

//go:linkname Marshal encoding/asn1.Marshal
//go:noescape
func Marshal(val interface{}) ([]byte, error)

//go:linkname Unmarshal encoding/asn1.Unmarshal
//go:noescape
func Unmarshal(b []byte, val interface{}) ([]byte, error)

//go:linkname UnmarshalWithParams encoding/asn1.UnmarshalWithParams
//go:noescape
func UnmarshalWithParams(b []byte, val interface{}, params string) ([]byte, error)

//go:linkname bitstringrightalign asn1.sub_bitstringrightalign
func bitstringrightalign(b asn1.BitString) []byte {
	return b.RightAlign()
}

//go:linkname BitStringRightAlign asn1.sub_bitstringrightalign
//go:noescape
func BitStringRightAlign(b asn1.BitString) []byte

//go:linkname objectidentifierstring asn1.sub_objectidentifierstring
func objectidentifierstring(oi asn1.ObjectIdentifier) string {
	return oi.String()
}

//go:linkname ObjectIdentifierString asn1.sub_objectidentifierstring
//go:noescape
func ObjectIdentifierString(oi asn1.ObjectIdentifier) string

//go:linkname structuralerrorerror asn1.sub_structuralerrorerror
func structuralerrorerror(e asn1.StructuralError) string {
	return e.Error()
}

//go:linkname StructuralErrorError asn1.sub_structuralerrorerror
//go:noescape
func StructuralErrorError(e asn1.StructuralError) string

//go:linkname syntaxerrorerror asn1.sub_syntaxerrorerror
func syntaxerrorerror(e asn1.SyntaxError) string {
	return e.Error()
}

//go:linkname SyntaxErrorError asn1.sub_syntaxerrorerror
//go:noescape
func SyntaxErrorError(e asn1.SyntaxError) string
