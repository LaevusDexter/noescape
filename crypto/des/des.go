// This file has automatically been generated on Wed Feb 26 02:09:44 +05 2020.
// DO NOT EDIT.
package des

import (
	"crypto/cipher"
	"crypto/des"
	_ "unsafe"
)

//go:linkname NewCipher crypto/des.NewCipher
//go:noescape
func NewCipher(key []byte) (cipher.Block, error)

//go:linkname NewTripleDESCipher crypto/des.NewTripleDESCipher
//go:noescape
func NewTripleDESCipher(key []byte) (cipher.Block, error)

//go:linkname keysizeerrorerror des.sub_keysizeerrorerror
func keysizeerrorerror(k des.KeySizeError) string {
	return k.Error()
}

//go:linkname KeySizeErrorError des.sub_keysizeerrorerror
//go:noescape
func KeySizeErrorError(k des.KeySizeError) string
