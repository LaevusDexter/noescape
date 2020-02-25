// This file has automatically been generated on Wed Feb 26 02:10:12 +05 2020.
// DO NOT EDIT.
package debug

import (
	"runtime/debug"
	_ "unsafe"
)

//go:linkname SetGCPercent runtime/debug.SetGCPercent
//go:noescape
func SetGCPercent(percent int) int

//go:linkname SetMaxStack runtime/debug.SetMaxStack
//go:noescape
func SetMaxStack(bytes int) int

//go:linkname SetMaxThreads runtime/debug.SetMaxThreads
//go:noescape
func SetMaxThreads(threads int) int

//go:linkname SetPanicOnFault runtime/debug.SetPanicOnFault
//go:noescape
func SetPanicOnFault(enabled bool) bool

//go:linkname Stack runtime/debug.Stack
//go:noescape
func Stack() []byte

//go:linkname ReadBuildInfo runtime/debug.ReadBuildInfo
//go:noescape
func ReadBuildInfo() (*debug.BuildInfo, bool)
