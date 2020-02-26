// This file has automatically been generated on Wed Feb 26 15:50:22 +05 2020.
// DO NOT EDIT.
package rand

import (
	_ "crypto/rand"
	"io"
	"math/big"
	_ "unsafe"
)

//go:linkname Read crypto/rand.Read
//go:noescape
func Read(b []byte) (int, error)

//go:linkname Int crypto/rand.Int
//go:noescape
func Int(rand io.Reader, max *big.Int) (*big.Int, error)

//go:linkname Prime crypto/rand.Prime
//go:noescape
func Prime(rand io.Reader, bits int) (*big.Int, error)
