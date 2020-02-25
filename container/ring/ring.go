// This file has automatically been generated on Wed Feb 26 02:09:42 +05 2020.
// DO NOT EDIT.
package ring

import (
	"container/ring"
	_ "unsafe"
)

//go:linkname ringunlink ring.sub_ringunlink
func ringunlink(r *ring.Ring, n int) *ring.Ring {
	return r.Unlink(n)
}

//go:linkname RingUnlink ring.sub_ringunlink
//go:noescape
func RingUnlink(r *ring.Ring, n int) *ring.Ring

//go:linkname New container/ring.New
//go:noescape
func New(n int) *ring.Ring

//go:linkname ringlen ring.sub_ringlen
func ringlen(r *ring.Ring) int {
	return r.Len()
}

//go:linkname RingLen ring.sub_ringlen
//go:noescape
func RingLen(r *ring.Ring) int

//go:linkname ringlink ring.sub_ringlink
func ringlink(r *ring.Ring, s *ring.Ring) *ring.Ring {
	return r.Link(s)
}

//go:linkname RingLink ring.sub_ringlink
//go:noescape
func RingLink(r *ring.Ring, s *ring.Ring) *ring.Ring

//go:linkname ringmove ring.sub_ringmove
func ringmove(r *ring.Ring, n int) *ring.Ring {
	return r.Move(n)
}

//go:linkname RingMove ring.sub_ringmove
//go:noescape
func RingMove(r *ring.Ring, n int) *ring.Ring

//go:linkname ringnext ring.sub_ringnext
func ringnext(r *ring.Ring) *ring.Ring {
	return r.Next()
}

//go:linkname RingNext ring.sub_ringnext
//go:noescape
func RingNext(r *ring.Ring) *ring.Ring

//go:linkname ringprev ring.sub_ringprev
func ringprev(r *ring.Ring) *ring.Ring {
	return r.Prev()
}

//go:linkname RingPrev ring.sub_ringprev
//go:noescape
func RingPrev(r *ring.Ring) *ring.Ring
