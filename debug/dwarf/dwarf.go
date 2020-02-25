// This file has automatically been generated on Wed Feb 26 02:09:49 +05 2020.
// DO NOT EDIT.
package dwarf

import (
	"debug/dwarf"
	"encoding/binary"
	_ "unsafe"
)

//go:linkname functypestring dwarf.sub_functypestring
func functypestring(t *dwarf.FuncType) string {
	return t.String()
}

//go:linkname FuncTypeString dwarf.sub_functypestring
//go:noescape
func FuncTypeString(t *dwarf.FuncType) string

//go:linkname qualtypestring dwarf.sub_qualtypestring
func qualtypestring(t *dwarf.QualType) string {
	return t.String()
}

//go:linkname QualTypeString dwarf.sub_qualtypestring
//go:noescape
func QualTypeString(t *dwarf.QualType) string

//go:linkname readernext dwarf.sub_readernext
func readernext(r *dwarf.Reader) (*dwarf.Entry, error) {
	return r.Next()
}

//go:linkname ReaderNext dwarf.sub_readernext
//go:noescape
func ReaderNext(r *dwarf.Reader) (*dwarf.Entry, error)

//go:linkname structtypestring dwarf.sub_structtypestring
func structtypestring(t *dwarf.StructType) string {
	return t.String()
}

//go:linkname StructTypeString dwarf.sub_structtypestring
//go:noescape
func StructTypeString(t *dwarf.StructType) string

//go:linkname datareader dwarf.sub_datareader
func datareader(d *dwarf.Data) *dwarf.Reader {
	return d.Reader()
}

//go:linkname DataReader dwarf.sub_datareader
//go:noescape
func DataReader(d *dwarf.Data) *dwarf.Reader

//go:linkname typedeftypestring dwarf.sub_typedeftypestring
func typedeftypestring(t *dwarf.TypedefType) string {
	return t.String()
}

//go:linkname TypedefTypeString dwarf.sub_typedeftypestring
//go:noescape
func TypedefTypeString(t *dwarf.TypedefType) string

//go:linkname linereaderfiles dwarf.sub_linereaderfiles
func linereaderfiles(r *dwarf.LineReader) []*dwarf.LineFile {
	return r.Files()
}

//go:linkname LineReaderFiles dwarf.sub_linereaderfiles
//go:noescape
func LineReaderFiles(r *dwarf.LineReader) []*dwarf.LineFile

//go:linkname classstring dwarf.sub_classstring
func classstring(i dwarf.Class) string {
	return i.String()
}

//go:linkname ClassString dwarf.sub_classstring
//go:noescape
func ClassString(i dwarf.Class) string

//go:linkname dataaddtypes dwarf.sub_dataaddtypes
func dataaddtypes(d *dwarf.Data, name string, types []byte) error {
	return d.AddTypes(name, types)
}

//go:linkname DataAddTypes dwarf.sub_dataaddtypes
//go:noescape
func DataAddTypes(d *dwarf.Data, name string, types []byte) error

//go:linkname qualtypesize dwarf.sub_qualtypesize
func qualtypesize(t *dwarf.QualType) int64 {
	return t.Size()
}

//go:linkname QualTypeSize dwarf.sub_qualtypesize
//go:noescape
func QualTypeSize(t *dwarf.QualType) int64

//go:linkname arraytypestring dwarf.sub_arraytypestring
func arraytypestring(t *dwarf.ArrayType) string {
	return t.String()
}

//go:linkname ArrayTypeString dwarf.sub_arraytypestring
//go:noescape
func ArrayTypeString(t *dwarf.ArrayType) string

//go:linkname dataaddsection dwarf.sub_dataaddsection
func dataaddsection(d *dwarf.Data, name string, contents []byte) error {
	return d.AddSection(name, contents)
}

//go:linkname DataAddSection dwarf.sub_dataaddsection
//go:noescape
func DataAddSection(d *dwarf.Data, name string, contents []byte) error

//go:linkname datatype dwarf.sub_datatype
func datatype(d *dwarf.Data, off dwarf.Offset) (dwarf.Type, error) {
	return d.Type(off)
}

//go:linkname DataType dwarf.sub_datatype
//go:noescape
func DataType(d *dwarf.Data, off dwarf.Offset) (dwarf.Type, error)

//go:linkname decodeerrorerror dwarf.sub_decodeerrorerror
func decodeerrorerror(e dwarf.DecodeError) string {
	return e.Error()
}

//go:linkname DecodeErrorError dwarf.sub_decodeerrorerror
//go:noescape
func DecodeErrorError(e dwarf.DecodeError) string

//go:linkname entryattrfield dwarf.sub_entryattrfield
func entryattrfield(e *dwarf.Entry, a dwarf.Attr) *dwarf.Field {
	return e.AttrField(a)
}

//go:linkname EntryAttrField dwarf.sub_entryattrfield
//go:noescape
func EntryAttrField(e *dwarf.Entry, a dwarf.Attr) *dwarf.Field

//go:linkname readerseekpc dwarf.sub_readerseekpc
func readerseekpc(r *dwarf.Reader, pc uint64) (*dwarf.Entry, error) {
	return r.SeekPC(pc)
}

//go:linkname ReaderSeekPC dwarf.sub_readerseekpc
//go:noescape
func ReaderSeekPC(r *dwarf.Reader, pc uint64) (*dwarf.Entry, error)

//go:linkname commontypecommon dwarf.sub_commontypecommon
func commontypecommon(c *dwarf.CommonType) *dwarf.CommonType {
	return c.Common()
}

//go:linkname CommonTypeCommon dwarf.sub_commontypecommon
//go:noescape
func CommonTypeCommon(c *dwarf.CommonType) *dwarf.CommonType

//go:linkname New debug/dwarf.New
//go:noescape
func New(abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) (*dwarf.Data, error)

//go:linkname dotdotdottypestring dwarf.sub_dotdotdottypestring
func dotdotdottypestring(t *dwarf.DotDotDotType) string {
	return t.String()
}

//go:linkname DotDotDotTypeString dwarf.sub_dotdotdottypestring
//go:noescape
func DotDotDotTypeString(t *dwarf.DotDotDotType) string

//go:linkname linereadertell dwarf.sub_linereadertell
func linereadertell(r *dwarf.LineReader) dwarf.LineReaderPos {
	return r.Tell()
}

//go:linkname LineReaderTell dwarf.sub_linereadertell
//go:noescape
func LineReaderTell(r *dwarf.LineReader) dwarf.LineReaderPos

//go:linkname readerbyteorder dwarf.sub_readerbyteorder
func readerbyteorder(r *dwarf.Reader) binary.ByteOrder {
	return r.ByteOrder()
}

//go:linkname ReaderByteOrder dwarf.sub_readerbyteorder
//go:noescape
func ReaderByteOrder(r *dwarf.Reader) binary.ByteOrder

//go:linkname attrstring dwarf.sub_attrstring
func attrstring(i dwarf.Attr) string {
	return i.String()
}

//go:linkname AttrString dwarf.sub_attrstring
//go:noescape
func AttrString(i dwarf.Attr) string

//go:linkname basictypebasic dwarf.sub_basictypebasic
func basictypebasic(b *dwarf.BasicType) *dwarf.BasicType {
	return b.Basic()
}

//go:linkname BasicTypeBasic dwarf.sub_basictypebasic
//go:noescape
func BasicTypeBasic(b *dwarf.BasicType) *dwarf.BasicType

//go:linkname basictypestring dwarf.sub_basictypestring
func basictypestring(t *dwarf.BasicType) string {
	return t.String()
}

//go:linkname BasicTypeString dwarf.sub_basictypestring
//go:noescape
func BasicTypeString(t *dwarf.BasicType) string

//go:linkname entryval dwarf.sub_entryval
func entryval(e *dwarf.Entry, a dwarf.Attr) interface{} {
	return e.Val(a)
}

//go:linkname EntryVal dwarf.sub_entryval
//go:noescape
func EntryVal(e *dwarf.Entry, a dwarf.Attr) interface{}

//go:linkname enumtypestring dwarf.sub_enumtypestring
func enumtypestring(t *dwarf.EnumType) string {
	return t.String()
}

//go:linkname EnumTypeString dwarf.sub_enumtypestring
//go:noescape
func EnumTypeString(t *dwarf.EnumType) string

//go:linkname structtypedefn dwarf.sub_structtypedefn
func structtypedefn(t *dwarf.StructType) string {
	return t.Defn()
}

//go:linkname StructTypeDefn dwarf.sub_structtypedefn
//go:noescape
func StructTypeDefn(t *dwarf.StructType) string

//go:linkname taggostring dwarf.sub_taggostring
func taggostring(t dwarf.Tag) string {
	return t.GoString()
}

//go:linkname TagGoString dwarf.sub_taggostring
//go:noescape
func TagGoString(t dwarf.Tag) string

//go:linkname tagstring dwarf.sub_tagstring
func tagstring(i dwarf.Tag) string {
	return i.String()
}

//go:linkname TagString dwarf.sub_tagstring
//go:noescape
func TagString(i dwarf.Tag) string

//go:linkname attrgostring dwarf.sub_attrgostring
func attrgostring(a dwarf.Attr) string {
	return a.GoString()
}

//go:linkname AttrGoString dwarf.sub_attrgostring
//go:noescape
func AttrGoString(a dwarf.Attr) string

//go:linkname commontypesize dwarf.sub_commontypesize
func commontypesize(c *dwarf.CommonType) int64 {
	return c.Size()
}

//go:linkname CommonTypeSize dwarf.sub_commontypesize
//go:noescape
func CommonTypeSize(c *dwarf.CommonType) int64

//go:linkname dataranges dwarf.sub_dataranges
func dataranges(d *dwarf.Data, e *dwarf.Entry) ([][2]uint64, error) {
	return d.Ranges(e)
}

//go:linkname DataRanges dwarf.sub_dataranges
//go:noescape
func DataRanges(d *dwarf.Data, e *dwarf.Entry) ([][2]uint64, error)

//go:linkname readeraddresssize dwarf.sub_readeraddresssize
func readeraddresssize(r *dwarf.Reader) int {
	return r.AddressSize()
}

//go:linkname ReaderAddressSize dwarf.sub_readeraddresssize
//go:noescape
func ReaderAddressSize(r *dwarf.Reader) int

//go:linkname typedeftypesize dwarf.sub_typedeftypesize
func typedeftypesize(t *dwarf.TypedefType) int64 {
	return t.Size()
}

//go:linkname TypedefTypeSize dwarf.sub_typedeftypesize
//go:noescape
func TypedefTypeSize(t *dwarf.TypedefType) int64

//go:linkname unsupportedtypestring dwarf.sub_unsupportedtypestring
func unsupportedtypestring(t *dwarf.UnsupportedType) string {
	return t.String()
}

//go:linkname UnsupportedTypeString dwarf.sub_unsupportedtypestring
//go:noescape
func UnsupportedTypeString(t *dwarf.UnsupportedType) string

//go:linkname classgostring dwarf.sub_classgostring
func classgostring(i dwarf.Class) string {
	return i.GoString()
}

//go:linkname ClassGoString dwarf.sub_classgostring
//go:noescape
func ClassGoString(i dwarf.Class) string

//go:linkname datalinereader dwarf.sub_datalinereader
func datalinereader(d *dwarf.Data, cu *dwarf.Entry) (*dwarf.LineReader, error) {
	return d.LineReader(cu)
}

//go:linkname DataLineReader dwarf.sub_datalinereader
//go:noescape
func DataLineReader(d *dwarf.Data, cu *dwarf.Entry) (*dwarf.LineReader, error)

//go:linkname linereadernext dwarf.sub_linereadernext
func linereadernext(r *dwarf.LineReader, entry *dwarf.LineEntry) error {
	return r.Next(entry)
}

//go:linkname LineReaderNext dwarf.sub_linereadernext
//go:noescape
func LineReaderNext(r *dwarf.LineReader, entry *dwarf.LineEntry) error

//go:linkname linereaderseekpc dwarf.sub_linereaderseekpc
func linereaderseekpc(r *dwarf.LineReader, pc uint64, entry *dwarf.LineEntry) error {
	return r.SeekPC(pc, entry)
}

//go:linkname LineReaderSeekPC dwarf.sub_linereaderseekpc
//go:noescape
func LineReaderSeekPC(r *dwarf.LineReader, pc uint64, entry *dwarf.LineEntry) error

//go:linkname ptrtypestring dwarf.sub_ptrtypestring
func ptrtypestring(t *dwarf.PtrType) string {
	return t.String()
}

//go:linkname PtrTypeString dwarf.sub_ptrtypestring
//go:noescape
func PtrTypeString(t *dwarf.PtrType) string

//go:linkname voidtypestring dwarf.sub_voidtypestring
func voidtypestring(t *dwarf.VoidType) string {
	return t.String()
}

//go:linkname VoidTypeString dwarf.sub_voidtypestring
//go:noescape
func VoidTypeString(t *dwarf.VoidType) string

//go:linkname arraytypesize dwarf.sub_arraytypesize
func arraytypesize(t *dwarf.ArrayType) int64 {
	return t.Size()
}

//go:linkname ArrayTypeSize dwarf.sub_arraytypesize
//go:noescape
func ArrayTypeSize(t *dwarf.ArrayType) int64
