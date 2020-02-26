// This file has automatically been generated on Wed Feb 26 15:50:50 +05 2020.
// DO NOT EDIT.
package runtime

import (
	"runtime"
	_ "unsafe"
)

//go:linkname ReadTrace runtime.ReadTrace
//go:noescape
func ReadTrace() []byte

//go:linkname SetMutexProfileFraction runtime.SetMutexProfileFraction
//go:noescape
func SetMutexProfileFraction(rate int) int

//go:linkname StartTrace runtime.StartTrace
//go:noescape
func StartTrace() error

//go:linkname Version runtime.Version
//go:noescape
func Version() string

//go:linkname BlockProfile runtime.BlockProfile
//go:noescape
func BlockProfile(p []runtime.BlockProfileRecord,) (int, bool)

//go:linkname MemProfile runtime.MemProfile
//go:noescape
func MemProfile(p []runtime.MemProfileRecord, inuseZero bool) (int, bool)

//go:linkname NumGoroutine runtime.NumGoroutine
//go:noescape
func NumGoroutine() int

//go:linkname Stack runtime.Stack
//go:noescape
func Stack(buf []byte, all bool) int

//go:linkname memprofilerecordinuseobjects runtime.sub_memprofilerecordinuseobjects
func memprofilerecordinuseobjects(r *runtime.MemProfileRecord,) int64 {
	return r.InUseObjects()
}

//go:linkname MemProfileRecordInUseObjects runtime.sub_memprofilerecordinuseobjects
//go:noescape
func MemProfileRecordInUseObjects(r *runtime.MemProfileRecord,) int64

//go:linkname GOMAXPROCS runtime.GOMAXPROCS
//go:noescape
func GOMAXPROCS(n int) int

//go:linkname FuncForPC runtime.FuncForPC
//go:noescape
func FuncForPC(pc uintptr) *runtime.Func

//go:linkname funcname runtime.sub_funcname
func funcname(f *runtime.Func,) string {
	return f.Name()
}

//go:linkname FuncName runtime.sub_funcname
//go:noescape
func FuncName(f *runtime.Func,) string

//go:linkname GoroutineProfile runtime.GoroutineProfile
//go:noescape
func GoroutineProfile(p []runtime.StackRecord,) (int, bool)

//go:linkname Callers runtime.Callers
//go:noescape
func Callers(skip int, pc []uintptr) int

//go:linkname NumCPU runtime.NumCPU
//go:noescape
func NumCPU() int

//go:linkname typeassertionerrorerror runtime.sub_typeassertionerrorerror
func typeassertionerrorerror(e *runtime.TypeAssertionError,) string {
	return e.Error()
}

//go:linkname TypeAssertionErrorError runtime.sub_typeassertionerrorerror
//go:noescape
func TypeAssertionErrorError(e *runtime.TypeAssertionError,) string

//go:linkname Caller runtime.Caller
//go:noescape
func Caller(skip int) (uintptr, string, int, bool)

//go:linkname ThreadCreateProfile runtime.ThreadCreateProfile
//go:noescape
func ThreadCreateProfile(p []runtime.StackRecord,) (int, bool)

//go:linkname CallersFrames runtime.CallersFrames
//go:noescape
func CallersFrames(callers []uintptr) *runtime.Frames

//go:linkname funcentry runtime.sub_funcentry
func funcentry(f *runtime.Func,) uintptr {
	return f.Entry()
}

//go:linkname FuncEntry runtime.sub_funcentry
//go:noescape
func FuncEntry(f *runtime.Func,) uintptr

//go:linkname memprofilerecordinusebytes runtime.sub_memprofilerecordinusebytes
func memprofilerecordinusebytes(r *runtime.MemProfileRecord,) int64 {
	return r.InUseBytes()
}

//go:linkname MemProfileRecordInUseBytes runtime.sub_memprofilerecordinusebytes
//go:noescape
func MemProfileRecordInUseBytes(r *runtime.MemProfileRecord,) int64

//go:linkname stackrecordstack runtime.sub_stackrecordstack
func stackrecordstack(r *runtime.StackRecord,) []uintptr {
	return r.Stack()
}

//go:linkname StackRecordStack runtime.sub_stackrecordstack
//go:noescape
func StackRecordStack(r *runtime.StackRecord,) []uintptr

//go:linkname CPUProfile runtime.CPUProfile
//go:noescape
func CPUProfile() []byte

//go:linkname framesnext runtime.sub_framesnext
func framesnext(ci *runtime.Frames,) (runtime.Frame, bool) {
	return ci.Next()
}

//go:linkname FramesNext runtime.sub_framesnext
//go:noescape
func FramesNext(ci *runtime.Frames,) (runtime.Frame, bool)

//go:linkname GOROOT runtime.GOROOT
//go:noescape
func GOROOT() string

//go:linkname memprofilerecordstack runtime.sub_memprofilerecordstack
func memprofilerecordstack(r *runtime.MemProfileRecord,) []uintptr {
	return r.Stack()
}

//go:linkname MemProfileRecordStack runtime.sub_memprofilerecordstack
//go:noescape
func MemProfileRecordStack(r *runtime.MemProfileRecord,) []uintptr

//go:linkname MutexProfile runtime.MutexProfile
//go:noescape
func MutexProfile(p []runtime.BlockProfileRecord,) (int, bool)

//go:linkname funcfileline runtime.sub_funcfileline
func funcfileline(f *runtime.Func, pc uintptr) (string, int) {
	return f.FileLine(pc)
}

//go:linkname FuncFileLine runtime.sub_funcfileline
//go:noescape
func FuncFileLine(f *runtime.Func, pc uintptr) (string, int)

//go:linkname NumCgoCall runtime.NumCgoCall
//go:noescape
func NumCgoCall() int64
