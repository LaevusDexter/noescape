// This file has automatically been generated on Wed Feb 26 02:09:43 +05 2020.
// DO NOT EDIT.
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	_ "unsafe"
)

//go:linkname NewCipher crypto/aes.NewCipher
//go:noescape
func NewCipher(key []byte) (cipher.Block, error)

//go:linkname keysizeerrorerror aes.sub_keysizeerrorerror
func keysizeerrorerror(k aes.KeySizeError) string {
	return k.Error()
}

//go:linkname KeySizeErrorError aes.sub_keysizeerrorerror
//go:noescape
func KeySizeErrorError(k aes.KeySizeError) string
