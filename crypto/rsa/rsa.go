// This file has automatically been generated on Wed Feb 26 15:50:23 +05 2020.
// DO NOT EDIT.
package rsa

import (
	"crypto"
	"crypto/rsa"
	"hash"
	"io"
	_ "unsafe"
)

//go:linkname pssoptionshashfunc rsa.sub_pssoptionshashfunc
func pssoptionshashfunc(pssOpts *rsa.PSSOptions,) crypto.Hash {
	return pssOpts.HashFunc()
}

//go:linkname PSSOptionsHashFunc rsa.sub_pssoptionshashfunc
//go:noescape
func PSSOptionsHashFunc(pssOpts *rsa.PSSOptions,) crypto.Hash

//go:linkname GenerateKey crypto/rsa.GenerateKey
//go:noescape
func GenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error)

//go:linkname privatekeydecrypt rsa.sub_privatekeydecrypt
func privatekeydecrypt(priv *rsa.PrivateKey, rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) ([]byte, error) {
	return priv.Decrypt(rand, ciphertext, opts)
}

//go:linkname PrivateKeyDecrypt rsa.sub_privatekeydecrypt
//go:noescape
func PrivateKeyDecrypt(priv *rsa.PrivateKey, rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) ([]byte, error)

//go:linkname privatekeypublic rsa.sub_privatekeypublic
func privatekeypublic(priv *rsa.PrivateKey,) crypto.PublicKey {
	return priv.Public()
}

//go:linkname PrivateKeyPublic rsa.sub_privatekeypublic
//go:noescape
func PrivateKeyPublic(priv *rsa.PrivateKey,) crypto.PublicKey

//go:linkname privatekeyvalidate rsa.sub_privatekeyvalidate
func privatekeyvalidate(priv *rsa.PrivateKey,) error {
	return priv.Validate()
}

//go:linkname PrivateKeyValidate rsa.sub_privatekeyvalidate
//go:noescape
func PrivateKeyValidate(priv *rsa.PrivateKey,) error

//go:linkname publickeysize rsa.sub_publickeysize
func publickeysize(pub *rsa.PublicKey,) int {
	return pub.Size()
}

//go:linkname PublicKeySize rsa.sub_publickeysize
//go:noescape
func PublicKeySize(pub *rsa.PublicKey,) int

//go:linkname DecryptOAEP crypto/rsa.DecryptOAEP
//go:noescape
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error)

//go:linkname EncryptOAEP crypto/rsa.EncryptOAEP
//go:noescape
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *rsa.PublicKey, msg []byte, label []byte) ([]byte, error)

//go:linkname SignPSS crypto/rsa.SignPSS
//go:noescape
func SignPSS(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions,) ([]byte, error)

//go:linkname VerifyPSS crypto/rsa.VerifyPSS
//go:noescape
func VerifyPSS(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *rsa.PSSOptions,) error

//go:linkname GenerateMultiPrimeKey crypto/rsa.GenerateMultiPrimeKey
//go:noescape
func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*rsa.PrivateKey, error)

//go:linkname privatekeysign rsa.sub_privatekeysign
func privatekeysign(priv *rsa.PrivateKey, rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	return priv.Sign(rand, digest, opts)
}

//go:linkname PrivateKeySign rsa.sub_privatekeysign
//go:noescape
func PrivateKeySign(priv *rsa.PrivateKey, rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
