// This file has automatically been generated on Wed Feb 26 02:09:44 +05 2020.
// DO NOT EDIT.
package ed25519

import (
	"crypto"
	"crypto/ed25519"
	"io"
	_ "unsafe"
)

//go:linkname privatekeypublic ed25519.sub_privatekeypublic
func privatekeypublic(priv ed25519.PrivateKey) crypto.PublicKey {
	return priv.Public()
}

//go:linkname PrivateKeyPublic ed25519.sub_privatekeypublic
//go:noescape
func PrivateKeyPublic(priv ed25519.PrivateKey) crypto.PublicKey

//go:linkname privatekeyseed ed25519.sub_privatekeyseed
func privatekeyseed(priv ed25519.PrivateKey) []byte {
	return priv.Seed()
}

//go:linkname PrivateKeySeed ed25519.sub_privatekeyseed
//go:noescape
func PrivateKeySeed(priv ed25519.PrivateKey) []byte

//go:linkname privatekeysign ed25519.sub_privatekeysign
func privatekeysign(priv ed25519.PrivateKey, rand io.Reader, message []byte, opts crypto.SignerOpts) ([]byte, error) {
	return priv.Sign(rand, message, opts)
}

//go:linkname PrivateKeySign ed25519.sub_privatekeysign
//go:noescape
func PrivateKeySign(priv ed25519.PrivateKey, rand io.Reader, message []byte, opts crypto.SignerOpts) ([]byte, error)

//go:linkname GenerateKey crypto/ed25519.GenerateKey
//go:noescape
func GenerateKey(rand io.Reader) (ed25519.PublicKey, ed25519.PrivateKey, error)

//go:linkname Sign crypto/ed25519.Sign
//go:noescape
func Sign(privateKey ed25519.PrivateKey, message []byte) []byte

//go:linkname Verify crypto/ed25519.Verify
//go:noescape
func Verify(publicKey ed25519.PublicKey, message, sig []byte) bool

//go:linkname NewKeyFromSeed crypto/ed25519.NewKeyFromSeed
//go:noescape
func NewKeyFromSeed(seed []byte) ed25519.PrivateKey
