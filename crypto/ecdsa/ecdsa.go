// This file has automatically been generated on Wed Feb 26 15:50:21 +05 2020.
// DO NOT EDIT.
package ecdsa

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"io"
	"math/big"
	_ "unsafe"
)

//go:linkname GenerateKey crypto/ecdsa.GenerateKey
//go:noescape
func GenerateKey(c elliptic.Curve, rand io.Reader) (*ecdsa.PrivateKey, error)

//go:linkname privatekeypublic ecdsa.sub_privatekeypublic
func privatekeypublic(priv *ecdsa.PrivateKey,) crypto.PublicKey {
	return priv.Public()
}

//go:linkname PrivateKeyPublic ecdsa.sub_privatekeypublic
//go:noescape
func PrivateKeyPublic(priv *ecdsa.PrivateKey,) crypto.PublicKey

//go:linkname privatekeysign ecdsa.sub_privatekeysign
func privatekeysign(priv *ecdsa.PrivateKey, rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	return priv.Sign(rand, digest, opts)
}

//go:linkname PrivateKeySign ecdsa.sub_privatekeysign
//go:noescape
func PrivateKeySign(priv *ecdsa.PrivateKey, rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)

//go:linkname Sign crypto/ecdsa.Sign
//go:noescape
func Sign(rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) (*big.Int, *big.Int, error)

//go:linkname Verify crypto/ecdsa.Verify
//go:noescape
func Verify(pub *ecdsa.PublicKey, hash []byte, r, s *big.Int) bool
