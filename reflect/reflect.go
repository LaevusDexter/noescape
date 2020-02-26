// This file has automatically been generated on Wed Feb 26 15:50:48 +05 2020.
// DO NOT EDIT.
package reflect

import (
	"reflect"
	"unsafe"
)

//go:linkname valuelen reflect.sub_valuelen
func valuelen(v reflect.Value,) int {
	return v.Len()
}

//go:linkname ValueLen reflect.sub_valuelen
//go:noescape
func ValueLen(v reflect.Value,) int

//go:linkname valuemethod reflect.sub_valuemethod
func valuemethod(v reflect.Value, i int) reflect.Value {
	return v.Method(i)
}

//go:linkname ValueMethod reflect.sub_valuemethod
//go:noescape
func ValueMethod(v reflect.Value, i int) reflect.Value

//go:linkname valuepointer reflect.sub_valuepointer
func valuepointer(v reflect.Value,) uintptr {
	return v.Pointer()
}

//go:linkname ValuePointer reflect.sub_valuepointer
//go:noescape
func ValuePointer(v reflect.Value,) uintptr

//go:linkname AppendSlice reflect.AppendSlice
//go:noescape
func AppendSlice(s, t reflect.Value,) reflect.Value

//go:linkname valuecallslice reflect.sub_valuecallslice
func valuecallslice(v reflect.Value, in []reflect.Value,) []reflect.Value {
	return v.CallSlice(in)
}

//go:linkname ValueCallSlice reflect.sub_valuecallslice
//go:noescape
func ValueCallSlice(v reflect.Value, in []reflect.Value,) []reflect.Value

//go:linkname valuecanset reflect.sub_valuecanset
func valuecanset(v reflect.Value,) bool {
	return v.CanSet()
}

//go:linkname ValueCanSet reflect.sub_valuecanset
//go:noescape
func ValueCanSet(v reflect.Value,) bool

//go:linkname valueelem reflect.sub_valueelem
func valueelem(v reflect.Value,) reflect.Value {
	return v.Elem()
}

//go:linkname ValueElem reflect.sub_valueelem
//go:noescape
func ValueElem(v reflect.Value,) reflect.Value

//go:linkname valueindex reflect.sub_valueindex
func valueindex(v reflect.Value, i int) reflect.Value {
	return v.Index(i)
}

//go:linkname ValueIndex reflect.sub_valueindex
//go:noescape
func ValueIndex(v reflect.Value, i int) reflect.Value

//go:linkname ChanOf reflect.ChanOf
//go:noescape
func ChanOf(dir reflect.ChanDir, t reflect.Type,) reflect.Type

//go:linkname MakeChan reflect.MakeChan
//go:noescape
func MakeChan(typ reflect.Type, buffer int) reflect.Value

//go:linkname chandirstring reflect.sub_chandirstring
func chandirstring(d reflect.ChanDir,) string {
	return d.String()
}

//go:linkname ChanDirString reflect.sub_chandirstring
//go:noescape
func ChanDirString(d reflect.ChanDir,) string

//go:linkname MakeFunc reflect.MakeFunc
//go:noescape
func MakeFunc(typ reflect.Type, fn func(args []reflect.Value,) []reflect.Value) reflect.Value

//go:linkname Select reflect.Select
//go:noescape
func Select(cases []reflect.SelectCase,) (int, reflect.Value, bool)

//go:linkname valuetryrecv reflect.sub_valuetryrecv
func valuetryrecv(v reflect.Value,) (reflect.Value, bool) {
	return v.TryRecv()
}

//go:linkname ValueTryRecv reflect.sub_valuetryrecv
//go:noescape
func ValueTryRecv(v reflect.Value,) (reflect.Value, bool)

//go:linkname valuefieldbyname reflect.sub_valuefieldbyname
func valuefieldbyname(v reflect.Value, name string) reflect.Value {
	return v.FieldByName(name)
}

//go:linkname ValueFieldByName reflect.sub_valuefieldbyname
//go:noescape
func ValueFieldByName(v reflect.Value, name string) reflect.Value

//go:linkname valueinterface reflect.sub_valueinterface
func valueinterface(v reflect.Value,) interface{} {
	return v.Interface()
}

//go:linkname ValueInterface reflect.sub_valueinterface
//go:noescape
func ValueInterface(v reflect.Value,) interface{}

//go:linkname valueerrorerror reflect.sub_valueerrorerror
func valueerrorerror(e *reflect.ValueError,) string {
	return e.Error()
}

//go:linkname ValueErrorError reflect.sub_valueerrorerror
//go:noescape
func ValueErrorError(e *reflect.ValueError,) string

//go:linkname mapitervalue reflect.sub_mapitervalue
func mapitervalue(it *reflect.MapIter,) reflect.Value {
	return it.Value()
}

//go:linkname MapIterValue reflect.sub_mapitervalue
//go:noescape
func MapIterValue(it *reflect.MapIter,) reflect.Value

//go:linkname SliceOf reflect.SliceOf
//go:noescape
func SliceOf(t reflect.Type,) reflect.Type

//go:linkname Append reflect.Append
//go:noescape
func Append(s reflect.Value, x ...reflect.Value,) reflect.Value

//go:linkname valueisnil reflect.sub_valueisnil
func valueisnil(v reflect.Value,) bool {
	return v.IsNil()
}

//go:linkname ValueIsNil reflect.sub_valueisnil
//go:noescape
func ValueIsNil(v reflect.Value,) bool

//go:linkname valueoverflowcomplex reflect.sub_valueoverflowcomplex
func valueoverflowcomplex(v reflect.Value, x complex128) bool {
	return v.OverflowComplex(x)
}

//go:linkname ValueOverflowComplex reflect.sub_valueoverflowcomplex
//go:noescape
func ValueOverflowComplex(v reflect.Value, x complex128) bool

//go:linkname NewAt reflect.NewAt
//go:noescape
func NewAt(typ reflect.Type, p unsafe.Pointer) reflect.Value

//go:linkname valueconvert reflect.sub_valueconvert
func valueconvert(v reflect.Value, t reflect.Type,) reflect.Value {
	return v.Convert(t)
}

//go:linkname ValueConvert reflect.sub_valueconvert
//go:noescape
func ValueConvert(v reflect.Value, t reflect.Type,) reflect.Value

//go:linkname valuefloat reflect.sub_valuefloat
func valuefloat(v reflect.Value,) float64 {
	return v.Float()
}

//go:linkname ValueFloat reflect.sub_valuefloat
//go:noescape
func ValueFloat(v reflect.Value,) float64

//go:linkname valuemethodbyname reflect.sub_valuemethodbyname
func valuemethodbyname(v reflect.Value, name string) reflect.Value {
	return v.MethodByName(name)
}

//go:linkname ValueMethodByName reflect.sub_valuemethodbyname
//go:noescape
func ValueMethodByName(v reflect.Value, name string) reflect.Value

//go:linkname valuefieldbynamefunc reflect.sub_valuefieldbynamefunc
func valuefieldbynamefunc(v reflect.Value, match func(string) bool) reflect.Value {
	return v.FieldByNameFunc(match)
}

//go:linkname ValueFieldByNameFunc reflect.sub_valuefieldbynamefunc
//go:noescape
func ValueFieldByNameFunc(v reflect.Value, match func(string) bool) reflect.Value

//go:linkname valueisvalid reflect.sub_valueisvalid
func valueisvalid(v reflect.Value,) bool {
	return v.IsValid()
}

//go:linkname ValueIsValid reflect.sub_valueisvalid
//go:noescape
func ValueIsValid(v reflect.Value,) bool

//go:linkname Indirect reflect.Indirect
//go:noescape
func Indirect(v reflect.Value,) reflect.Value

//go:linkname MakeMap reflect.MakeMap
//go:noescape
func MakeMap(typ reflect.Type,) reflect.Value

//go:linkname New reflect.New
//go:noescape
func New(typ reflect.Type,) reflect.Value

//go:linkname Zero reflect.Zero
//go:noescape
func Zero(typ reflect.Type,) reflect.Value

//go:linkname valuecomplex reflect.sub_valuecomplex
func valuecomplex(v reflect.Value,) complex128 {
	return v.Complex()
}

//go:linkname ValueComplex reflect.sub_valuecomplex
//go:noescape
func ValueComplex(v reflect.Value,) complex128

//go:linkname valuecanaddr reflect.sub_valuecanaddr
func valuecanaddr(v reflect.Value,) bool {
	return v.CanAddr()
}

//go:linkname ValueCanAddr reflect.sub_valuecanaddr
//go:noescape
func ValueCanAddr(v reflect.Value,) bool

//go:linkname valuenumfield reflect.sub_valuenumfield
func valuenumfield(v reflect.Value,) int {
	return v.NumField()
}

//go:linkname ValueNumField reflect.sub_valuenumfield
//go:noescape
func ValueNumField(v reflect.Value,) int

//go:linkname valuerecv reflect.sub_valuerecv
func valuerecv(v reflect.Value,) (reflect.Value, bool) {
	return v.Recv()
}

//go:linkname ValueRecv reflect.sub_valuerecv
//go:noescape
func ValueRecv(v reflect.Value,) (reflect.Value, bool)

//go:linkname valueslice reflect.sub_valueslice
func valueslice(v reflect.Value, i, j int) reflect.Value {
	return v.Slice(i, j)
}

//go:linkname ValueSlice reflect.sub_valueslice
//go:noescape
func ValueSlice(v reflect.Value, i, j int) reflect.Value

//go:linkname valueuint reflect.sub_valueuint
func valueuint(v reflect.Value,) uint64 {
	return v.Uint()
}

//go:linkname ValueUint reflect.sub_valueuint
//go:noescape
func ValueUint(v reflect.Value,) uint64

//go:linkname valuemapindex reflect.sub_valuemapindex
func valuemapindex(v reflect.Value, key reflect.Value,) reflect.Value {
	return v.MapIndex(key)
}

//go:linkname ValueMapIndex reflect.sub_valuemapindex
//go:noescape
func ValueMapIndex(v reflect.Value, key reflect.Value,) reflect.Value

//go:linkname valuemapkeys reflect.sub_valuemapkeys
func valuemapkeys(v reflect.Value,) []reflect.Value {
	return v.MapKeys()
}

//go:linkname ValueMapKeys reflect.sub_valuemapkeys
//go:noescape
func ValueMapKeys(v reflect.Value,) []reflect.Value

//go:linkname valueoverflowuint reflect.sub_valueoverflowuint
func valueoverflowuint(v reflect.Value, x uint64) bool {
	return v.OverflowUint(x)
}

//go:linkname ValueOverflowUint reflect.sub_valueoverflowuint
//go:noescape
func ValueOverflowUint(v reflect.Value, x uint64) bool

//go:linkname DeepEqual reflect.DeepEqual
//go:noescape
func DeepEqual(x, y interface{}) bool

//go:linkname MakeMapWithSize reflect.MakeMapWithSize
//go:noescape
func MakeMapWithSize(typ reflect.Type, n int) reflect.Value

//go:linkname MakeSlice reflect.MakeSlice
//go:noescape
func MakeSlice(typ reflect.Type, len, cap int) reflect.Value

//go:linkname ValueOf reflect.ValueOf
//go:noescape
func ValueOf(i interface{}) reflect.Value

//go:linkname valuekind reflect.sub_valuekind
func valuekind(v reflect.Value,) reflect.Kind {
	return v.Kind()
}

//go:linkname ValueKind reflect.sub_valuekind
//go:noescape
func ValueKind(v reflect.Value,) reflect.Kind

//go:linkname valuetype reflect.sub_valuetype
func valuetype(v reflect.Value,) reflect.Type {
	return v.Type()
}

//go:linkname ValueType reflect.sub_valuetype
//go:noescape
func ValueType(v reflect.Value,) reflect.Type

//go:linkname structtaglookup reflect.sub_structtaglookup
func structtaglookup(tag reflect.StructTag, key string) (string, bool) {
	return tag.Lookup(key)
}

//go:linkname StructTagLookup reflect.sub_structtaglookup
//go:noescape
func StructTagLookup(tag reflect.StructTag, key string) (string, bool)

//go:linkname valuecap reflect.sub_valuecap
func valuecap(v reflect.Value,) int {
	return v.Cap()
}

//go:linkname ValueCap reflect.sub_valuecap
//go:noescape
func ValueCap(v reflect.Value,) int

//go:linkname valueoverflowfloat reflect.sub_valueoverflowfloat
func valueoverflowfloat(v reflect.Value, x float64) bool {
	return v.OverflowFloat(x)
}

//go:linkname ValueOverflowFloat reflect.sub_valueoverflowfloat
//go:noescape
func ValueOverflowFloat(v reflect.Value, x float64) bool

//go:linkname PtrTo reflect.PtrTo
//go:noescape
func PtrTo(t reflect.Type,) reflect.Type

//go:linkname StructOf reflect.StructOf
//go:noescape
func StructOf(fields []reflect.StructField,) reflect.Type

//go:linkname valuebytes reflect.sub_valuebytes
func valuebytes(v reflect.Value,) []byte {
	return v.Bytes()
}

//go:linkname ValueBytes reflect.sub_valuebytes
//go:noescape
func ValueBytes(v reflect.Value,) []byte

//go:linkname FuncOf reflect.FuncOf
//go:noescape
func FuncOf(in, out []reflect.Type, variadic bool) reflect.Type

//go:linkname MapOf reflect.MapOf
//go:noescape
func MapOf(key, elem reflect.Type,) reflect.Type

//go:linkname valuefieldbyindex reflect.sub_valuefieldbyindex
func valuefieldbyindex(v reflect.Value, index []int) reflect.Value {
	return v.FieldByIndex(index)
}

//go:linkname ValueFieldByIndex reflect.sub_valuefieldbyindex
//go:noescape
func ValueFieldByIndex(v reflect.Value, index []int) reflect.Value

//go:linkname valueinterfacedata reflect.sub_valueinterfacedata
func valueinterfacedata(v reflect.Value,) [2]uintptr {
	return v.InterfaceData()
}

//go:linkname ValueInterfaceData reflect.sub_valueinterfacedata
//go:noescape
func ValueInterfaceData(v reflect.Value,) [2]uintptr

//go:linkname valueunsafeaddr reflect.sub_valueunsafeaddr
func valueunsafeaddr(v reflect.Value,) uintptr {
	return v.UnsafeAddr()
}

//go:linkname ValueUnsafeAddr reflect.sub_valueunsafeaddr
//go:noescape
func ValueUnsafeAddr(v reflect.Value,) uintptr

//go:linkname valueint reflect.sub_valueint
func valueint(v reflect.Value,) int64 {
	return v.Int()
}

//go:linkname ValueInt reflect.sub_valueint
//go:noescape
func ValueInt(v reflect.Value,) int64

//go:linkname valuestring reflect.sub_valuestring
func valuestring(v reflect.Value,) string {
	return v.String()
}

//go:linkname ValueString reflect.sub_valuestring
//go:noescape
func ValueString(v reflect.Value,) string

//go:linkname Swapper reflect.Swapper
//go:noescape
func Swapper(slice interface{}) func(i, j int)

//go:linkname kindstring reflect.sub_kindstring
func kindstring(k reflect.Kind,) string {
	return k.String()
}

//go:linkname KindString reflect.sub_kindstring
//go:noescape
func KindString(k reflect.Kind,) string

//go:linkname mapiterkey reflect.sub_mapiterkey
func mapiterkey(it *reflect.MapIter,) reflect.Value {
	return it.Key()
}

//go:linkname MapIterKey reflect.sub_mapiterkey
//go:noescape
func MapIterKey(it *reflect.MapIter,) reflect.Value

//go:linkname ArrayOf reflect.ArrayOf
//go:noescape
func ArrayOf(count int, elem reflect.Type,) reflect.Type

//go:linkname valuecaninterface reflect.sub_valuecaninterface
func valuecaninterface(v reflect.Value,) bool {
	return v.CanInterface()
}

//go:linkname ValueCanInterface reflect.sub_valuecaninterface
//go:noescape
func ValueCanInterface(v reflect.Value,) bool

//go:linkname valueoverflowint reflect.sub_valueoverflowint
func valueoverflowint(v reflect.Value, x int64) bool {
	return v.OverflowInt(x)
}

//go:linkname ValueOverflowInt reflect.sub_valueoverflowint
//go:noescape
func ValueOverflowInt(v reflect.Value, x int64) bool

//go:linkname Copy reflect.Copy
//go:noescape
func Copy(dst, src reflect.Value,) int

//go:linkname mapiternext reflect.sub_mapiternext
func mapiternext(it *reflect.MapIter,) bool {
	return it.Next()
}

//go:linkname MapIterNext reflect.sub_mapiternext
//go:noescape
func MapIterNext(it *reflect.MapIter,) bool

//go:linkname structtagget reflect.sub_structtagget
func structtagget(tag reflect.StructTag, key string) string {
	return tag.Get(key)
}

//go:linkname StructTagGet reflect.sub_structtagget
//go:noescape
func StructTagGet(tag reflect.StructTag, key string) string

//go:linkname valuecall reflect.sub_valuecall
func valuecall(v reflect.Value, in []reflect.Value,) []reflect.Value {
	return v.Call(in)
}

//go:linkname ValueCall reflect.sub_valuecall
//go:noescape
func ValueCall(v reflect.Value, in []reflect.Value,) []reflect.Value

//go:linkname valuemaprange reflect.sub_valuemaprange
func valuemaprange(v reflect.Value,) *reflect.MapIter {
	return v.MapRange()
}

//go:linkname ValueMapRange reflect.sub_valuemaprange
//go:noescape
func ValueMapRange(v reflect.Value,) *reflect.MapIter

//go:linkname TypeOf reflect.TypeOf
//go:noescape
func TypeOf(i interface{}) reflect.Type

//go:linkname valueiszero reflect.sub_valueiszero
func valueiszero(v reflect.Value,) bool {
	return v.IsZero()
}

//go:linkname ValueIsZero reflect.sub_valueiszero
//go:noescape
func ValueIsZero(v reflect.Value,) bool

//go:linkname valuetrysend reflect.sub_valuetrysend
func valuetrysend(v reflect.Value, x reflect.Value,) bool {
	return v.TrySend(x)
}

//go:linkname ValueTrySend reflect.sub_valuetrysend
//go:noescape
func ValueTrySend(v reflect.Value, x reflect.Value,) bool

//go:linkname valueaddr reflect.sub_valueaddr
func valueaddr(v reflect.Value,) reflect.Value {
	return v.Addr()
}

//go:linkname ValueAddr reflect.sub_valueaddr
//go:noescape
func ValueAddr(v reflect.Value,) reflect.Value

//go:linkname valuebool reflect.sub_valuebool
func valuebool(v reflect.Value,) bool {
	return v.Bool()
}

//go:linkname ValueBool reflect.sub_valuebool
//go:noescape
func ValueBool(v reflect.Value,) bool

//go:linkname valuefield reflect.sub_valuefield
func valuefield(v reflect.Value, i int) reflect.Value {
	return v.Field(i)
}

//go:linkname ValueField reflect.sub_valuefield
//go:noescape
func ValueField(v reflect.Value, i int) reflect.Value

//go:linkname valuenummethod reflect.sub_valuenummethod
func valuenummethod(v reflect.Value,) int {
	return v.NumMethod()
}

//go:linkname ValueNumMethod reflect.sub_valuenummethod
//go:noescape
func ValueNumMethod(v reflect.Value,) int
