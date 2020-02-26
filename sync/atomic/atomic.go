// This file has automatically been generated on Wed Feb 26 15:50:53 +05 2020.
// DO NOT EDIT.
package atomic

import (
	"sync/atomic"
	"unsafe"
)

//go:linkname LoadPointer sync/atomic.LoadPointer
//go:noescape
func LoadPointer(addr *unsafe.Pointer) unsafe.Pointer

//go:linkname LoadUintptr sync/atomic.LoadUintptr
//go:noescape
func LoadUintptr(addr *uintptr) uintptr

//go:linkname SwapPointer sync/atomic.SwapPointer
//go:noescape
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) unsafe.Pointer

//go:linkname SwapUintptr sync/atomic.SwapUintptr
//go:noescape
func SwapUintptr(addr *uintptr, new uintptr) uintptr

//go:linkname valueload atomic.sub_valueload
func valueload(v *atomic.Value,) interface{} {
	return v.Load()
}

//go:linkname ValueLoad atomic.sub_valueload
//go:noescape
func ValueLoad(v *atomic.Value,) interface{}

//go:linkname AddUintptr sync/atomic.AddUintptr
//go:noescape
func AddUintptr(addr *uintptr, delta uintptr) uintptr

//go:linkname CompareAndSwapPointer sync/atomic.CompareAndSwapPointer
//go:noescape
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) bool

//go:linkname CompareAndSwapUintptr sync/atomic.CompareAndSwapUintptr
//go:noescape
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) bool
