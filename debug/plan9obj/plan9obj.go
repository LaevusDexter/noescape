// This file has automatically been generated on Wed Feb 26 15:50:28 +05 2020.
// DO NOT EDIT.
package plan9obj

import (
	"debug/plan9obj"
	"io"
	_ "unsafe"
)

//go:linkname Open debug/plan9obj.Open
//go:noescape
func Open(name string) (*plan9obj.File, error)

//go:linkname fileclose plan9obj.sub_fileclose
func fileclose(f *plan9obj.File,) error {
	return f.Close()
}

//go:linkname FileClose plan9obj.sub_fileclose
//go:noescape
func FileClose(f *plan9obj.File,) error

//go:linkname filesection plan9obj.sub_filesection
func filesection(f *plan9obj.File, name string) *plan9obj.Section {
	return f.Section(name)
}

//go:linkname FileSection plan9obj.sub_filesection
//go:noescape
func FileSection(f *plan9obj.File, name string) *plan9obj.Section

//go:linkname filesymbols plan9obj.sub_filesymbols
func filesymbols(f *plan9obj.File,) ([]plan9obj.Sym, error) {
	return f.Symbols()
}

//go:linkname FileSymbols plan9obj.sub_filesymbols
//go:noescape
func FileSymbols(f *plan9obj.File,) ([]plan9obj.Sym, error)

//go:linkname sectiondata plan9obj.sub_sectiondata
func sectiondata(s *plan9obj.Section,) ([]byte, error) {
	return s.Data()
}

//go:linkname SectionData plan9obj.sub_sectiondata
//go:noescape
func SectionData(s *plan9obj.Section,) ([]byte, error)

//go:linkname sectionopen plan9obj.sub_sectionopen
func sectionopen(s *plan9obj.Section,) io.ReadSeeker {
	return s.Open()
}

//go:linkname SectionOpen plan9obj.sub_sectionopen
//go:noescape
func SectionOpen(s *plan9obj.Section,) io.ReadSeeker

//go:linkname NewFile debug/plan9obj.NewFile
//go:noescape
func NewFile(r io.ReaderAt) (*plan9obj.File, error)
