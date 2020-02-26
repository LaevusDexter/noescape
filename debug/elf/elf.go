// This file has automatically been generated on Wed Feb 26 15:50:27 +05 2020.
// DO NOT EDIT.
package elf

import (
	"debug/dwarf"
	"debug/elf"
	"io"
	_ "unsafe"
)

//go:linkname ST_INFO debug/elf.ST_INFO
//go:noescape
func ST_INFO(bind elf.SymBind, typ elf.SymType,) uint8

//go:linkname dyntaggostring elf.sub_dyntaggostring
func dyntaggostring(i elf.DynTag,) string {
	return i.GoString()
}

//go:linkname DynTagGoString elf.sub_dyntaggostring
//go:noescape
func DynTagGoString(i elf.DynTag,) string

//go:linkname filesection elf.sub_filesection
func filesection(f *elf.File, name string) *elf.Section {
	return f.Section(name)
}

//go:linkname FileSection elf.sub_filesection
//go:noescape
func FileSection(f *elf.File, name string) *elf.Section

//go:linkname r_armgostring elf.sub_r_armgostring
func r_armgostring(i elf.R_ARM,) string {
	return i.GoString()
}

//go:linkname R_ARMGoString elf.sub_r_armgostring
//go:noescape
func R_ARMGoString(i elf.R_ARM,) string

//go:linkname R_INFO debug/elf.R_INFO
//go:noescape
func R_INFO(sym, typ uint32) uint64

//go:linkname compressiontypestring elf.sub_compressiontypestring
func compressiontypestring(i elf.CompressionType,) string {
	return i.String()
}

//go:linkname CompressionTypeString elf.sub_compressiontypestring
//go:noescape
func CompressionTypeString(i elf.CompressionType,) string

//go:linkname NewFile debug/elf.NewFile
//go:noescape
func NewFile(r io.ReaderAt) (*elf.File, error)

//go:linkname progflagstring elf.sub_progflagstring
func progflagstring(i elf.ProgFlag,) string {
	return i.String()
}

//go:linkname ProgFlagString elf.sub_progflagstring
//go:noescape
func ProgFlagString(i elf.ProgFlag,) string

//go:linkname r_390string elf.sub_r_390string
func r_390string(i elf.R_390,) string {
	return i.String()
}

//go:linkname R_390String elf.sub_r_390string
//go:noescape
func R_390String(i elf.R_390,) string

//go:linkname r_riscvgostring elf.sub_r_riscvgostring
func r_riscvgostring(i elf.R_RISCV,) string {
	return i.GoString()
}

//go:linkname R_RISCVGoString elf.sub_r_riscvgostring
//go:noescape
func R_RISCVGoString(i elf.R_RISCV,) string

//go:linkname sectionopen elf.sub_sectionopen
func sectionopen(s *elf.Section,) io.ReadSeeker {
	return s.Open()
}

//go:linkname SectionOpen elf.sub_sectionopen
//go:noescape
func SectionOpen(s *elf.Section,) io.ReadSeeker

//go:linkname sectionflagstring elf.sub_sectionflagstring
func sectionflagstring(i elf.SectionFlag,) string {
	return i.String()
}

//go:linkname SectionFlagString elf.sub_sectionflagstring
//go:noescape
func SectionFlagString(i elf.SectionFlag,) string

//go:linkname ST_VISIBILITY debug/elf.ST_VISIBILITY
//go:noescape
func ST_VISIBILITY(other uint8) elf.SymVis

//go:linkname typegostring elf.sub_typegostring
func typegostring(i elf.Type,) string {
	return i.GoString()
}

//go:linkname TypeGoString elf.sub_typegostring
//go:noescape
func TypeGoString(i elf.Type,) string

//go:linkname datastring elf.sub_datastring
func datastring(i elf.Data,) string {
	return i.String()
}

//go:linkname DataString elf.sub_datastring
//go:noescape
func DataString(i elf.Data,) string

//go:linkname dynflaggostring elf.sub_dynflaggostring
func dynflaggostring(i elf.DynFlag,) string {
	return i.GoString()
}

//go:linkname DynFlagGoString elf.sub_dynflaggostring
//go:noescape
func DynFlagGoString(i elf.DynFlag,) string

//go:linkname Open debug/elf.Open
//go:noescape
func Open(name string) (*elf.File, error)

//go:linkname filedwarf elf.sub_filedwarf
func filedwarf(f *elf.File,) (*dwarf.Data, error) {
	return f.DWARF()
}

//go:linkname FileDWARF elf.sub_filedwarf
//go:noescape
func FileDWARF(f *elf.File,) (*dwarf.Data, error)

//go:linkname r_386string elf.sub_r_386string
func r_386string(i elf.R_386,) string {
	return i.String()
}

//go:linkname R_386String elf.sub_r_386string
//go:noescape
func R_386String(i elf.R_386,) string

//go:linkname versionstring elf.sub_versionstring
func versionstring(i elf.Version,) string {
	return i.String()
}

//go:linkname VersionString elf.sub_versionstring
//go:noescape
func VersionString(i elf.Version,) string

//go:linkname fileclose elf.sub_fileclose
func fileclose(f *elf.File,) error {
	return f.Close()
}

//go:linkname FileClose elf.sub_fileclose
//go:noescape
func FileClose(f *elf.File,) error

//go:linkname ntypegostring elf.sub_ntypegostring
func ntypegostring(i elf.NType,) string {
	return i.GoString()
}

//go:linkname NTypeGoString elf.sub_ntypegostring
//go:noescape
func NTypeGoString(i elf.NType,) string

//go:linkname r_armstring elf.sub_r_armstring
func r_armstring(i elf.R_ARM,) string {
	return i.String()
}

//go:linkname R_ARMString elf.sub_r_armstring
//go:noescape
func R_ARMString(i elf.R_ARM,) string

//go:linkname r_ppcgostring elf.sub_r_ppcgostring
func r_ppcgostring(i elf.R_PPC,) string {
	return i.GoString()
}

//go:linkname R_PPCGoString elf.sub_r_ppcgostring
//go:noescape
func R_PPCGoString(i elf.R_PPC,) string

//go:linkname r_alphagostring elf.sub_r_alphagostring
func r_alphagostring(i elf.R_ALPHA,) string {
	return i.GoString()
}

//go:linkname R_ALPHAGoString elf.sub_r_alphagostring
//go:noescape
func R_ALPHAGoString(i elf.R_ALPHA,) string

//go:linkname r_alphastring elf.sub_r_alphastring
func r_alphastring(i elf.R_ALPHA,) string {
	return i.String()
}

//go:linkname R_ALPHAString elf.sub_r_alphastring
//go:noescape
func R_ALPHAString(i elf.R_ALPHA,) string

//go:linkname sectiontypegostring elf.sub_sectiontypegostring
func sectiontypegostring(i elf.SectionType,) string {
	return i.GoString()
}

//go:linkname SectionTypeGoString elf.sub_sectiontypegostring
//go:noescape
func SectionTypeGoString(i elf.SectionType,) string

//go:linkname filedynamicsymbols elf.sub_filedynamicsymbols
func filedynamicsymbols(f *elf.File,) ([]elf.Symbol, error) {
	return f.DynamicSymbols()
}

//go:linkname FileDynamicSymbols elf.sub_filedynamicsymbols
//go:noescape
func FileDynamicSymbols(f *elf.File,) ([]elf.Symbol, error)

//go:linkname fileimportedlibraries elf.sub_fileimportedlibraries
func fileimportedlibraries(f *elf.File,) ([]string, error) {
	return f.ImportedLibraries()
}

//go:linkname FileImportedLibraries elf.sub_fileimportedlibraries
//go:noescape
func FileImportedLibraries(f *elf.File,) ([]string, error)

//go:linkname osabigostring elf.sub_osabigostring
func osabigostring(i elf.OSABI,) string {
	return i.GoString()
}

//go:linkname OSABIGoString elf.sub_osabigostring
//go:noescape
func OSABIGoString(i elf.OSABI,) string

//go:linkname r_390gostring elf.sub_r_390gostring
func r_390gostring(i elf.R_390,) string {
	return i.GoString()
}

//go:linkname R_390GoString elf.sub_r_390gostring
//go:noescape
func R_390GoString(i elf.R_390,) string

//go:linkname r_aarch64string elf.sub_r_aarch64string
func r_aarch64string(i elf.R_AARCH64,) string {
	return i.String()
}

//go:linkname R_AARCH64String elf.sub_r_aarch64string
//go:noescape
func R_AARCH64String(i elf.R_AARCH64,) string

//go:linkname symvisgostring elf.sub_symvisgostring
func symvisgostring(i elf.SymVis,) string {
	return i.GoString()
}

//go:linkname SymVisGoString elf.sub_symvisgostring
//go:noescape
func SymVisGoString(i elf.SymVis,) string

//go:linkname classgostring elf.sub_classgostring
func classgostring(i elf.Class,) string {
	return i.GoString()
}

//go:linkname ClassGoString elf.sub_classgostring
//go:noescape
func ClassGoString(i elf.Class,) string

//go:linkname datagostring elf.sub_datagostring
func datagostring(i elf.Data,) string {
	return i.GoString()
}

//go:linkname DataGoString elf.sub_datagostring
//go:noescape
func DataGoString(i elf.Data,) string

//go:linkname sectionindexgostring elf.sub_sectionindexgostring
func sectionindexgostring(i elf.SectionIndex,) string {
	return i.GoString()
}

//go:linkname SectionIndexGoString elf.sub_sectionindexgostring
//go:noescape
func SectionIndexGoString(i elf.SectionIndex,) string

//go:linkname sectionindexstring elf.sub_sectionindexstring
func sectionindexstring(i elf.SectionIndex,) string {
	return i.String()
}

//go:linkname SectionIndexString elf.sub_sectionindexstring
//go:noescape
func SectionIndexString(i elf.SectionIndex,) string

//go:linkname symbindgostring elf.sub_symbindgostring
func symbindgostring(i elf.SymBind,) string {
	return i.GoString()
}

//go:linkname SymBindGoString elf.sub_symbindgostring
//go:noescape
func SymBindGoString(i elf.SymBind,) string

//go:linkname machinegostring elf.sub_machinegostring
func machinegostring(i elf.Machine,) string {
	return i.GoString()
}

//go:linkname MachineGoString elf.sub_machinegostring
//go:noescape
func MachineGoString(i elf.Machine,) string

//go:linkname sectionflaggostring elf.sub_sectionflaggostring
func sectionflaggostring(i elf.SectionFlag,) string {
	return i.GoString()
}

//go:linkname SectionFlagGoString elf.sub_sectionflaggostring
//go:noescape
func SectionFlagGoString(i elf.SectionFlag,) string

//go:linkname symbindstring elf.sub_symbindstring
func symbindstring(i elf.SymBind,) string {
	return i.String()
}

//go:linkname SymBindString elf.sub_symbindstring
//go:noescape
func SymBindString(i elf.SymBind,) string

//go:linkname symtypestring elf.sub_symtypestring
func symtypestring(i elf.SymType,) string {
	return i.String()
}

//go:linkname SymTypeString elf.sub_symtypestring
//go:noescape
func SymTypeString(i elf.SymType,) string

//go:linkname fileimportedsymbols elf.sub_fileimportedsymbols
func fileimportedsymbols(f *elf.File,) ([]elf.ImportedSymbol, error) {
	return f.ImportedSymbols()
}

//go:linkname FileImportedSymbols elf.sub_fileimportedsymbols
//go:noescape
func FileImportedSymbols(f *elf.File,) ([]elf.ImportedSymbol, error)

//go:linkname r_386gostring elf.sub_r_386gostring
func r_386gostring(i elf.R_386,) string {
	return i.GoString()
}

//go:linkname R_386GoString elf.sub_r_386gostring
//go:noescape
func R_386GoString(i elf.R_386,) string

//go:linkname r_x86_64gostring elf.sub_r_x86_64gostring
func r_x86_64gostring(i elf.R_X86_64,) string {
	return i.GoString()
}

//go:linkname R_X86_64GoString elf.sub_r_x86_64gostring
//go:noescape
func R_X86_64GoString(i elf.R_X86_64,) string

//go:linkname sectiondata elf.sub_sectiondata
func sectiondata(s *elf.Section,) ([]byte, error) {
	return s.Data()
}

//go:linkname SectionData elf.sub_sectiondata
//go:noescape
func SectionData(s *elf.Section,) ([]byte, error)

//go:linkname r_riscvstring elf.sub_r_riscvstring
func r_riscvstring(i elf.R_RISCV,) string {
	return i.String()
}

//go:linkname R_RISCVString elf.sub_r_riscvstring
//go:noescape
func R_RISCVString(i elf.R_RISCV,) string

//go:linkname sectiontypestring elf.sub_sectiontypestring
func sectiontypestring(i elf.SectionType,) string {
	return i.String()
}

//go:linkname SectionTypeString elf.sub_sectiontypestring
//go:noescape
func SectionTypeString(i elf.SectionType,) string

//go:linkname versiongostring elf.sub_versiongostring
func versiongostring(i elf.Version,) string {
	return i.GoString()
}

//go:linkname VersionGoString elf.sub_versiongostring
//go:noescape
func VersionGoString(i elf.Version,) string

//go:linkname classstring elf.sub_classstring
func classstring(i elf.Class,) string {
	return i.String()
}

//go:linkname ClassString elf.sub_classstring
//go:noescape
func ClassString(i elf.Class,) string

//go:linkname filesymbols elf.sub_filesymbols
func filesymbols(f *elf.File,) ([]elf.Symbol, error) {
	return f.Symbols()
}

//go:linkname FileSymbols elf.sub_filesymbols
//go:noescape
func FileSymbols(f *elf.File,) ([]elf.Symbol, error)

//go:linkname r_ppcstring elf.sub_r_ppcstring
func r_ppcstring(i elf.R_PPC,) string {
	return i.String()
}

//go:linkname R_PPCString elf.sub_r_ppcstring
//go:noescape
func R_PPCString(i elf.R_PPC,) string

//go:linkname r_ppc64gostring elf.sub_r_ppc64gostring
func r_ppc64gostring(i elf.R_PPC64,) string {
	return i.GoString()
}

//go:linkname R_PPC64GoString elf.sub_r_ppc64gostring
//go:noescape
func R_PPC64GoString(i elf.R_PPC64,) string

//go:linkname r_ppc64string elf.sub_r_ppc64string
func r_ppc64string(i elf.R_PPC64,) string {
	return i.String()
}

//go:linkname R_PPC64String elf.sub_r_ppc64string
//go:noescape
func R_PPC64String(i elf.R_PPC64,) string

//go:linkname progopen elf.sub_progopen
func progopen(p *elf.Prog,) io.ReadSeeker {
	return p.Open()
}

//go:linkname ProgOpen elf.sub_progopen
//go:noescape
func ProgOpen(p *elf.Prog,) io.ReadSeeker

//go:linkname progtypegostring elf.sub_progtypegostring
func progtypegostring(i elf.ProgType,) string {
	return i.GoString()
}

//go:linkname ProgTypeGoString elf.sub_progtypegostring
//go:noescape
func ProgTypeGoString(i elf.ProgType,) string

//go:linkname r_x86_64string elf.sub_r_x86_64string
func r_x86_64string(i elf.R_X86_64,) string {
	return i.String()
}

//go:linkname R_X86_64String elf.sub_r_x86_64string
//go:noescape
func R_X86_64String(i elf.R_X86_64,) string

//go:linkname ST_BIND debug/elf.ST_BIND
//go:noescape
func ST_BIND(info uint8) elf.SymBind

//go:linkname dynflagstring elf.sub_dynflagstring
func dynflagstring(i elf.DynFlag,) string {
	return i.String()
}

//go:linkname DynFlagString elf.sub_dynflagstring
//go:noescape
func DynFlagString(i elf.DynFlag,) string

//go:linkname filedynstring elf.sub_filedynstring
func filedynstring(f *elf.File, tag elf.DynTag,) ([]string, error) {
	return f.DynString(tag)
}

//go:linkname FileDynString elf.sub_filedynstring
//go:noescape
func FileDynString(f *elf.File, tag elf.DynTag,) ([]string, error)

//go:linkname osabistring elf.sub_osabistring
func osabistring(i elf.OSABI,) string {
	return i.String()
}

//go:linkname OSABIString elf.sub_osabistring
//go:noescape
func OSABIString(i elf.OSABI,) string

//go:linkname r_aarch64gostring elf.sub_r_aarch64gostring
func r_aarch64gostring(i elf.R_AARCH64,) string {
	return i.GoString()
}

//go:linkname R_AARCH64GoString elf.sub_r_aarch64gostring
//go:noescape
func R_AARCH64GoString(i elf.R_AARCH64,) string

//go:linkname r_sparcgostring elf.sub_r_sparcgostring
func r_sparcgostring(i elf.R_SPARC,) string {
	return i.GoString()
}

//go:linkname R_SPARCGoString elf.sub_r_sparcgostring
//go:noescape
func R_SPARCGoString(i elf.R_SPARC,) string

//go:linkname typestring elf.sub_typestring
func typestring(i elf.Type,) string {
	return i.String()
}

//go:linkname TypeString elf.sub_typestring
//go:noescape
func TypeString(i elf.Type,) string

//go:linkname compressiontypegostring elf.sub_compressiontypegostring
func compressiontypegostring(i elf.CompressionType,) string {
	return i.GoString()
}

//go:linkname CompressionTypeGoString elf.sub_compressiontypegostring
//go:noescape
func CompressionTypeGoString(i elf.CompressionType,) string

//go:linkname dyntagstring elf.sub_dyntagstring
func dyntagstring(i elf.DynTag,) string {
	return i.String()
}

//go:linkname DynTagString elf.sub_dyntagstring
//go:noescape
func DynTagString(i elf.DynTag,) string

//go:linkname formaterrorerror elf.sub_formaterrorerror
func formaterrorerror(e *elf.FormatError,) string {
	return e.Error()
}

//go:linkname FormatErrorError elf.sub_formaterrorerror
//go:noescape
func FormatErrorError(e *elf.FormatError,) string

//go:linkname r_mipsgostring elf.sub_r_mipsgostring
func r_mipsgostring(i elf.R_MIPS,) string {
	return i.GoString()
}

//go:linkname R_MIPSGoString elf.sub_r_mipsgostring
//go:noescape
func R_MIPSGoString(i elf.R_MIPS,) string

//go:linkname symtypegostring elf.sub_symtypegostring
func symtypegostring(i elf.SymType,) string {
	return i.GoString()
}

//go:linkname SymTypeGoString elf.sub_symtypegostring
//go:noescape
func SymTypeGoString(i elf.SymType,) string

//go:linkname ST_TYPE debug/elf.ST_TYPE
//go:noescape
func ST_TYPE(info uint8) elf.SymType

//go:linkname filesectionbytype elf.sub_filesectionbytype
func filesectionbytype(f *elf.File, typ elf.SectionType,) *elf.Section {
	return f.SectionByType(typ)
}

//go:linkname FileSectionByType elf.sub_filesectionbytype
//go:noescape
func FileSectionByType(f *elf.File, typ elf.SectionType,) *elf.Section

//go:linkname machinestring elf.sub_machinestring
func machinestring(i elf.Machine,) string {
	return i.String()
}

//go:linkname MachineString elf.sub_machinestring
//go:noescape
func MachineString(i elf.Machine,) string

//go:linkname progflaggostring elf.sub_progflaggostring
func progflaggostring(i elf.ProgFlag,) string {
	return i.GoString()
}

//go:linkname ProgFlagGoString elf.sub_progflaggostring
//go:noescape
func ProgFlagGoString(i elf.ProgFlag,) string

//go:linkname progtypestring elf.sub_progtypestring
func progtypestring(i elf.ProgType,) string {
	return i.String()
}

//go:linkname ProgTypeString elf.sub_progtypestring
//go:noescape
func ProgTypeString(i elf.ProgType,) string

//go:linkname r_sparcstring elf.sub_r_sparcstring
func r_sparcstring(i elf.R_SPARC,) string {
	return i.String()
}

//go:linkname R_SPARCString elf.sub_r_sparcstring
//go:noescape
func R_SPARCString(i elf.R_SPARC,) string

//go:linkname ntypestring elf.sub_ntypestring
func ntypestring(i elf.NType,) string {
	return i.String()
}

//go:linkname NTypeString elf.sub_ntypestring
//go:noescape
func NTypeString(i elf.NType,) string

//go:linkname r_mipsstring elf.sub_r_mipsstring
func r_mipsstring(i elf.R_MIPS,) string {
	return i.String()
}

//go:linkname R_MIPSString elf.sub_r_mipsstring
//go:noescape
func R_MIPSString(i elf.R_MIPS,) string

//go:linkname symvisstring elf.sub_symvisstring
func symvisstring(i elf.SymVis,) string {
	return i.String()
}

//go:linkname SymVisString elf.sub_symvisstring
//go:noescape
func SymVisString(i elf.SymVis,) string
