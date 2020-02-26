// This file has automatically been generated on Wed Feb 26 15:50:41 +05 2020.
// DO NOT EDIT.
package rand

import (
	"math/rand"
	_ "unsafe"
)

//go:linkname NewSource math/rand.NewSource
//go:noescape
func NewSource(seed int64) rand.Source

//go:linkname NewZipf math/rand.NewZipf
//go:noescape
func NewZipf(r *rand.Rand, s float64, v float64, imax uint64) *rand.Zipf

//go:linkname Int math/rand.Int
//go:noescape
func Int() int

//go:linkname Read math/rand.Read
//go:noescape
func Read(p []byte) (int, error)

//go:linkname randread rand.sub_randread
func randread(r *rand.Rand, p []byte) (int, error) {
	return r.Read(p)
}

//go:linkname RandRead rand.sub_randread
//go:noescape
func RandRead(r *rand.Rand, p []byte) (int, error)

//go:linkname randint rand.sub_randint
func randint(r *rand.Rand,) int {
	return r.Int()
}

//go:linkname RandInt rand.sub_randint
//go:noescape
func RandInt(r *rand.Rand,) int

//go:linkname randintn rand.sub_randintn
func randintn(r *rand.Rand, n int) int {
	return r.Intn(n)
}

//go:linkname RandIntn rand.sub_randintn
//go:noescape
func RandIntn(r *rand.Rand, n int) int

//go:linkname randperm rand.sub_randperm
func randperm(r *rand.Rand, n int) []int {
	return r.Perm(n)
}

//go:linkname RandPerm rand.sub_randperm
//go:noescape
func RandPerm(r *rand.Rand, n int) []int

//go:linkname Intn math/rand.Intn
//go:noescape
func Intn(n int) int

//go:linkname Perm math/rand.Perm
//go:noescape
func Perm(n int) []int

//go:linkname New math/rand.New
//go:noescape
func New(src rand.Source,) *rand.Rand
