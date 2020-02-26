// This file has automatically been generated on Wed Feb 26 15:50:27 +05 2020.
// DO NOT EDIT.
package pe

import (
	"debug/dwarf"
	"debug/pe"
	"io"
	_ "unsafe"
)

//go:linkname NewFile debug/pe.NewFile
//go:noescape
func NewFile(r io.ReaderAt) (*pe.File, error)

//go:linkname Open debug/pe.Open
//go:noescape
func Open(name string) (*pe.File, error)

//go:linkname filedwarf pe.sub_filedwarf
func filedwarf(f *pe.File,) (*dwarf.Data, error) {
	return f.DWARF()
}

//go:linkname FileDWARF pe.sub_filedwarf
//go:noescape
func FileDWARF(f *pe.File,) (*dwarf.Data, error)

//go:linkname fileimportedsymbols pe.sub_fileimportedsymbols
func fileimportedsymbols(f *pe.File,) ([]string, error) {
	return f.ImportedSymbols()
}

//go:linkname FileImportedSymbols pe.sub_fileimportedsymbols
//go:noescape
func FileImportedSymbols(f *pe.File,) ([]string, error)

//go:linkname sectiondata pe.sub_sectiondata
func sectiondata(s *pe.Section,) ([]byte, error) {
	return s.Data()
}

//go:linkname SectionData pe.sub_sectiondata
//go:noescape
func SectionData(s *pe.Section,) ([]byte, error)

//go:linkname stringtablestring pe.sub_stringtablestring
func stringtablestring(st pe.StringTable, start uint32) (string, error) {
	return st.String(start)
}

//go:linkname StringTableString pe.sub_stringtablestring
//go:noescape
func StringTableString(st pe.StringTable, start uint32) (string, error)

//go:linkname coffsymbolfullname pe.sub_coffsymbolfullname
func coffsymbolfullname(sym *pe.COFFSymbol, st pe.StringTable,) (string, error) {
	return sym.FullName(st)
}

//go:linkname COFFSymbolFullName pe.sub_coffsymbolfullname
//go:noescape
func COFFSymbolFullName(sym *pe.COFFSymbol, st pe.StringTable,) (string, error)

//go:linkname fileclose pe.sub_fileclose
func fileclose(f *pe.File,) error {
	return f.Close()
}

//go:linkname FileClose pe.sub_fileclose
//go:noescape
func FileClose(f *pe.File,) error

//go:linkname fileimportedlibraries pe.sub_fileimportedlibraries
func fileimportedlibraries(f *pe.File,) ([]string, error) {
	return f.ImportedLibraries()
}

//go:linkname FileImportedLibraries pe.sub_fileimportedlibraries
//go:noescape
func FileImportedLibraries(f *pe.File,) ([]string, error)

//go:linkname filesection pe.sub_filesection
func filesection(f *pe.File, name string) *pe.Section {
	return f.Section(name)
}

//go:linkname FileSection pe.sub_filesection
//go:noescape
func FileSection(f *pe.File, name string) *pe.Section

//go:linkname formaterrorerror pe.sub_formaterrorerror
func formaterrorerror(e *pe.FormatError,) string {
	return e.Error()
}

//go:linkname FormatErrorError pe.sub_formaterrorerror
//go:noescape
func FormatErrorError(e *pe.FormatError,) string

//go:linkname sectionopen pe.sub_sectionopen
func sectionopen(s *pe.Section,) io.ReadSeeker {
	return s.Open()
}

//go:linkname SectionOpen pe.sub_sectionopen
//go:noescape
func SectionOpen(s *pe.Section,) io.ReadSeeker
