// This file has automatically been generated on Wed Feb 26 02:09:50 +05 2020.
// DO NOT EDIT.
package gosym

import (
	"debug/gosym"
	_ "unsafe"
)

//go:linkname tablepctoline gosym.sub_tablepctoline
func tablepctoline(t *gosym.Table, pc uint64) (string, int, *gosym.Func) {
	return t.PCToLine(pc)
}

//go:linkname TablePCToLine gosym.sub_tablepctoline
//go:noescape
func TablePCToLine(t *gosym.Table, pc uint64) (string, int, *gosym.Func)

//go:linkname unknownlineerrorerror gosym.sub_unknownlineerrorerror
func unknownlineerrorerror(e *gosym.UnknownLineError) string {
	return e.Error()
}

//go:linkname UnknownLineErrorError gosym.sub_unknownlineerrorerror
//go:noescape
func UnknownLineErrorError(e *gosym.UnknownLineError) string

//go:linkname NewLineTable debug/gosym.NewLineTable
//go:noescape
func NewLineTable(data []byte, text uint64) *gosym.LineTable

//go:linkname linetablepctoline gosym.sub_linetablepctoline
func linetablepctoline(t *gosym.LineTable, pc uint64) int {
	return t.PCToLine(pc)
}

//go:linkname LineTablePCToLine gosym.sub_linetablepctoline
//go:noescape
func LineTablePCToLine(t *gosym.LineTable, pc uint64) int

//go:linkname symbasename gosym.sub_symbasename
func symbasename(s *gosym.Sym) string {
	return s.BaseName()
}

//go:linkname SymBaseName gosym.sub_symbasename
//go:noescape
func SymBaseName(s *gosym.Sym) string

//go:linkname tablepctofunc gosym.sub_tablepctofunc
func tablepctofunc(t *gosym.Table, pc uint64) *gosym.Func {
	return t.PCToFunc(pc)
}

//go:linkname TablePCToFunc gosym.sub_tablepctofunc
//go:noescape
func TablePCToFunc(t *gosym.Table, pc uint64) *gosym.Func

//go:linkname tablesymbyaddr gosym.sub_tablesymbyaddr
func tablesymbyaddr(t *gosym.Table, addr uint64) *gosym.Sym {
	return t.SymByAddr(addr)
}

//go:linkname TableSymByAddr gosym.sub_tablesymbyaddr
//go:noescape
func TableSymByAddr(t *gosym.Table, addr uint64) *gosym.Sym

//go:linkname unknownfileerrorerror gosym.sub_unknownfileerrorerror
func unknownfileerrorerror(e gosym.UnknownFileError) string {
	return e.Error()
}

//go:linkname UnknownFileErrorError gosym.sub_unknownfileerrorerror
//go:noescape
func UnknownFileErrorError(e gosym.UnknownFileError) string

//go:linkname decodingerrorerror gosym.sub_decodingerrorerror
func decodingerrorerror(e *gosym.DecodingError) string {
	return e.Error()
}

//go:linkname DecodingErrorError gosym.sub_decodingerrorerror
//go:noescape
func DecodingErrorError(e *gosym.DecodingError) string

//go:linkname tablelinetopc gosym.sub_tablelinetopc
func tablelinetopc(t *gosym.Table, file string, line int) (uint64, *gosym.Func, error) {
	return t.LineToPC(file, line)
}

//go:linkname TableLineToPC gosym.sub_tablelinetopc
//go:noescape
func TableLineToPC(t *gosym.Table, file string, line int) (uint64, *gosym.Func, error)

//go:linkname tablelookupfunc gosym.sub_tablelookupfunc
func tablelookupfunc(t *gosym.Table, name string) *gosym.Func {
	return t.LookupFunc(name)
}

//go:linkname TableLookupFunc gosym.sub_tablelookupfunc
//go:noescape
func TableLookupFunc(t *gosym.Table, name string) *gosym.Func

//go:linkname tablelookupsym gosym.sub_tablelookupsym
func tablelookupsym(t *gosym.Table, name string) *gosym.Sym {
	return t.LookupSym(name)
}

//go:linkname TableLookupSym gosym.sub_tablelookupsym
//go:noescape
func TableLookupSym(t *gosym.Table, name string) *gosym.Sym

//go:linkname sympackagename gosym.sub_sympackagename
func sympackagename(s *gosym.Sym) string {
	return s.PackageName()
}

//go:linkname SymPackageName gosym.sub_sympackagename
//go:noescape
func SymPackageName(s *gosym.Sym) string

//go:linkname symreceivername gosym.sub_symreceivername
func symreceivername(s *gosym.Sym) string {
	return s.ReceiverName()
}

//go:linkname SymReceiverName gosym.sub_symreceivername
//go:noescape
func SymReceiverName(s *gosym.Sym) string

//go:linkname symstatic gosym.sub_symstatic
func symstatic(s *gosym.Sym) bool {
	return s.Static()
}

//go:linkname SymStatic gosym.sub_symstatic
//go:noescape
func SymStatic(s *gosym.Sym) bool

//go:linkname NewTable debug/gosym.NewTable
//go:noescape
func NewTable(symtab []byte, pcln *gosym.LineTable) (*gosym.Table, error)

//go:linkname linetablelinetopc gosym.sub_linetablelinetopc
func linetablelinetopc(t *gosym.LineTable, line int, maxpc uint64) uint64 {
	return t.LineToPC(line, maxpc)
}

//go:linkname LineTableLineToPC gosym.sub_linetablelinetopc
//go:noescape
func LineTableLineToPC(t *gosym.LineTable, line int, maxpc uint64) uint64
