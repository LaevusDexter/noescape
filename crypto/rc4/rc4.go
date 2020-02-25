// This file has automatically been generated on Wed Feb 26 02:09:45 +05 2020.
// DO NOT EDIT.
package rc4

import (
	"crypto/rc4"
	_ "unsafe"
)

//go:linkname NewCipher crypto/rc4.NewCipher
//go:noescape
func NewCipher(key []byte) (*rc4.Cipher, error)

//go:linkname keysizeerrorerror rc4.sub_keysizeerrorerror
func keysizeerrorerror(k rc4.KeySizeError) string {
	return k.Error()
}

//go:linkname KeySizeErrorError rc4.sub_keysizeerrorerror
//go:noescape
func KeySizeErrorError(k rc4.KeySizeError) string
