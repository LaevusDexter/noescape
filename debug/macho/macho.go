// This file has automatically been generated on Wed Feb 26 15:50:27 +05 2020.
// DO NOT EDIT.
package macho

import (
	"debug/dwarf"
	"debug/macho"
	"io"
	_ "unsafe"
)

//go:linkname cpustring macho.sub_cpustring
func cpustring(i macho.Cpu,) string {
	return i.String()
}

//go:linkname CpuString macho.sub_cpustring
//go:noescape
func CpuString(i macho.Cpu,) string

//go:linkname NewFatFile debug/macho.NewFatFile
//go:noescape
func NewFatFile(r io.ReaderAt) (*macho.FatFile, error)

//go:linkname fatfileclose macho.sub_fatfileclose
func fatfileclose(ff *macho.FatFile,) error {
	return ff.Close()
}

//go:linkname FatFileClose macho.sub_fatfileclose
//go:noescape
func FatFileClose(ff *macho.FatFile,) error

//go:linkname formaterrorerror macho.sub_formaterrorerror
func formaterrorerror(e *macho.FormatError,) string {
	return e.Error()
}

//go:linkname FormatErrorError macho.sub_formaterrorerror
//go:noescape
func FormatErrorError(e *macho.FormatError,) string

//go:linkname loadbytesraw macho.sub_loadbytesraw
func loadbytesraw(b macho.LoadBytes,) []byte {
	return b.Raw()
}

//go:linkname LoadBytesRaw macho.sub_loadbytesraw
//go:noescape
func LoadBytesRaw(b macho.LoadBytes,) []byte

//go:linkname loadcmdstring macho.sub_loadcmdstring
func loadcmdstring(i macho.LoadCmd,) string {
	return i.String()
}

//go:linkname LoadCmdString macho.sub_loadcmdstring
//go:noescape
func LoadCmdString(i macho.LoadCmd,) string

//go:linkname reloctypegenericstring macho.sub_reloctypegenericstring
func reloctypegenericstring(i macho.RelocTypeGeneric,) string {
	return i.String()
}

//go:linkname RelocTypeGenericString macho.sub_reloctypegenericstring
//go:noescape
func RelocTypeGenericString(i macho.RelocTypeGeneric,) string

//go:linkname loadcmdgostring macho.sub_loadcmdgostring
func loadcmdgostring(i macho.LoadCmd,) string {
	return i.GoString()
}

//go:linkname LoadCmdGoString macho.sub_loadcmdgostring
//go:noescape
func LoadCmdGoString(i macho.LoadCmd,) string

//go:linkname reloctypearmstring macho.sub_reloctypearmstring
func reloctypearmstring(i macho.RelocTypeARM,) string {
	return i.String()
}

//go:linkname RelocTypeARMString macho.sub_reloctypearmstring
//go:noescape
func RelocTypeARMString(i macho.RelocTypeARM,) string

//go:linkname Open debug/macho.Open
//go:noescape
func Open(name string) (*macho.File, error)

//go:linkname reloctypearm64gostring macho.sub_reloctypearm64gostring
func reloctypearm64gostring(r macho.RelocTypeARM64,) string {
	return r.GoString()
}

//go:linkname RelocTypeARM64GoString macho.sub_reloctypearm64gostring
//go:noescape
func RelocTypeARM64GoString(r macho.RelocTypeARM64,) string

//go:linkname fileclose macho.sub_fileclose
func fileclose(f *macho.File,) error {
	return f.Close()
}

//go:linkname FileClose macho.sub_fileclose
//go:noescape
func FileClose(f *macho.File,) error

//go:linkname fileimportedlibraries macho.sub_fileimportedlibraries
func fileimportedlibraries(f *macho.File,) ([]string, error) {
	return f.ImportedLibraries()
}

//go:linkname FileImportedLibraries macho.sub_fileimportedlibraries
//go:noescape
func FileImportedLibraries(f *macho.File,) ([]string, error)

//go:linkname reloctypex86_64gostring macho.sub_reloctypex86_64gostring
func reloctypex86_64gostring(r macho.RelocTypeX86_64,) string {
	return r.GoString()
}

//go:linkname RelocTypeX86_64GoString macho.sub_reloctypex86_64gostring
//go:noescape
func RelocTypeX86_64GoString(r macho.RelocTypeX86_64,) string

//go:linkname sectiondata macho.sub_sectiondata
func sectiondata(s *macho.Section,) ([]byte, error) {
	return s.Data()
}

//go:linkname SectionData macho.sub_sectiondata
//go:noescape
func SectionData(s *macho.Section,) ([]byte, error)

//go:linkname cpugostring macho.sub_cpugostring
func cpugostring(i macho.Cpu,) string {
	return i.GoString()
}

//go:linkname CpuGoString macho.sub_cpugostring
//go:noescape
func CpuGoString(i macho.Cpu,) string

//go:linkname reloctypearmgostring macho.sub_reloctypearmgostring
func reloctypearmgostring(r macho.RelocTypeARM,) string {
	return r.GoString()
}

//go:linkname RelocTypeARMGoString macho.sub_reloctypearmgostring
//go:noescape
func RelocTypeARMGoString(r macho.RelocTypeARM,) string

//go:linkname reloctypearm64string macho.sub_reloctypearm64string
func reloctypearm64string(i macho.RelocTypeARM64,) string {
	return i.String()
}

//go:linkname RelocTypeARM64String macho.sub_reloctypearm64string
//go:noescape
func RelocTypeARM64String(i macho.RelocTypeARM64,) string

//go:linkname reloctypegenericgostring macho.sub_reloctypegenericgostring
func reloctypegenericgostring(r macho.RelocTypeGeneric,) string {
	return r.GoString()
}

//go:linkname RelocTypeGenericGoString macho.sub_reloctypegenericgostring
//go:noescape
func RelocTypeGenericGoString(r macho.RelocTypeGeneric,) string

//go:linkname sectionopen macho.sub_sectionopen
func sectionopen(s *macho.Section,) io.ReadSeeker {
	return s.Open()
}

//go:linkname SectionOpen macho.sub_sectionopen
//go:noescape
func SectionOpen(s *macho.Section,) io.ReadSeeker

//go:linkname typegostring macho.sub_typegostring
func typegostring(t macho.Type,) string {
	return t.GoString()
}

//go:linkname TypeGoString macho.sub_typegostring
//go:noescape
func TypeGoString(t macho.Type,) string

//go:linkname OpenFat debug/macho.OpenFat
//go:noescape
func OpenFat(name string) (*macho.FatFile, error)

//go:linkname segmentopen macho.sub_segmentopen
func segmentopen(s *macho.Segment,) io.ReadSeeker {
	return s.Open()
}

//go:linkname SegmentOpen macho.sub_segmentopen
//go:noescape
func SegmentOpen(s *macho.Segment,) io.ReadSeeker

//go:linkname NewFile debug/macho.NewFile
//go:noescape
func NewFile(r io.ReaderAt) (*macho.File, error)

//go:linkname typestring macho.sub_typestring
func typestring(t macho.Type,) string {
	return t.String()
}

//go:linkname TypeString macho.sub_typestring
//go:noescape
func TypeString(t macho.Type,) string

//go:linkname filedwarf macho.sub_filedwarf
func filedwarf(f *macho.File,) (*dwarf.Data, error) {
	return f.DWARF()
}

//go:linkname FileDWARF macho.sub_filedwarf
//go:noescape
func FileDWARF(f *macho.File,) (*dwarf.Data, error)

//go:linkname fileimportedsymbols macho.sub_fileimportedsymbols
func fileimportedsymbols(f *macho.File,) ([]string, error) {
	return f.ImportedSymbols()
}

//go:linkname FileImportedSymbols macho.sub_fileimportedsymbols
//go:noescape
func FileImportedSymbols(f *macho.File,) ([]string, error)

//go:linkname filesection macho.sub_filesection
func filesection(f *macho.File, name string) *macho.Section {
	return f.Section(name)
}

//go:linkname FileSection macho.sub_filesection
//go:noescape
func FileSection(f *macho.File, name string) *macho.Section

//go:linkname filesegment macho.sub_filesegment
func filesegment(f *macho.File, name string) *macho.Segment {
	return f.Segment(name)
}

//go:linkname FileSegment macho.sub_filesegment
//go:noescape
func FileSegment(f *macho.File, name string) *macho.Segment

//go:linkname reloctypex86_64string macho.sub_reloctypex86_64string
func reloctypex86_64string(i macho.RelocTypeX86_64,) string {
	return i.String()
}

//go:linkname RelocTypeX86_64String macho.sub_reloctypex86_64string
//go:noescape
func RelocTypeX86_64String(i macho.RelocTypeX86_64,) string

//go:linkname segmentdata macho.sub_segmentdata
func segmentdata(s *macho.Segment,) ([]byte, error) {
	return s.Data()
}

//go:linkname SegmentData macho.sub_segmentdata
//go:noescape
func SegmentData(s *macho.Segment,) ([]byte, error)
