// This file has automatically been generated on Wed Feb 26 02:09:44 +05 2020.
// DO NOT EDIT.
package dsa

import (
	"crypto/dsa"
	"io"
	"math/big"
	_ "unsafe"
)

//go:linkname GenerateParameters crypto/dsa.GenerateParameters
//go:noescape
func GenerateParameters(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error

//go:linkname Sign crypto/dsa.Sign
//go:noescape
func Sign(rand io.Reader, priv *dsa.PrivateKey, hash []byte) (*big.Int, error)

//go:linkname Verify crypto/dsa.Verify
//go:noescape
func Verify(pub *dsa.PublicKey, hash []byte, r, s *big.Int) bool

//go:linkname GenerateKey crypto/dsa.GenerateKey
//go:noescape
func GenerateKey(priv *dsa.PrivateKey, rand io.Reader) error
