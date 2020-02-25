// This file has automatically been generated on Wed Feb 26 02:09:42 +05 2020.
// DO NOT EDIT.
package heap

import (
	"container/heap"
	_ "unsafe"
)

//go:linkname Pop container/heap.Pop
//go:noescape
func Pop(h heap.Interface) interface{}

//go:linkname Remove container/heap.Remove
//go:noescape
func Remove(h heap.Interface, i int) interface{}
