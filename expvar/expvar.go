// This file has automatically been generated on Wed Feb 26 02:09:54 +05 2020.
// DO NOT EDIT.
package expvar

import (
	"expvar"
	"net/http"
	_ "unsafe"
)

//go:linkname NewFloat expvar.NewFloat
//go:noescape
func NewFloat(name string) *expvar.Float

//go:linkname floatvalue expvar.sub_floatvalue
func floatvalue(v *expvar.Float) float64 {
	return v.Value()
}

//go:linkname FloatValue expvar.sub_floatvalue
//go:noescape
func FloatValue(v *expvar.Float) float64

//go:linkname stringstring expvar.sub_stringstring
func stringstring(v *expvar.String) string {
	return v.String()
}

//go:linkname StringString expvar.sub_stringstring
//go:noescape
func StringString(v *expvar.String) string

//go:linkname intvalue expvar.sub_intvalue
func intvalue(v *expvar.Int) int64 {
	return v.Value()
}

//go:linkname IntValue expvar.sub_intvalue
//go:noescape
func IntValue(v *expvar.Int) int64

//go:linkname mapstring expvar.sub_mapstring
func mapstring(v *expvar.Map) string {
	return v.String()
}

//go:linkname MapString expvar.sub_mapstring
//go:noescape
func MapString(v *expvar.Map) string

//go:linkname Handler expvar.Handler
//go:noescape
func Handler() http.Handler

//go:linkname floatstring expvar.sub_floatstring
func floatstring(v *expvar.Float) string {
	return v.String()
}

//go:linkname FloatString expvar.sub_floatstring
//go:noescape
func FloatString(v *expvar.Float) string

//go:linkname funcvalue expvar.sub_funcvalue
func funcvalue(f expvar.Func) interface{} {
	return f.Value()
}

//go:linkname FuncValue expvar.sub_funcvalue
//go:noescape
func FuncValue(f expvar.Func) interface{}

//go:linkname funcstring expvar.sub_funcstring
func funcstring(f expvar.Func) string {
	return f.String()
}

//go:linkname FuncString expvar.sub_funcstring
//go:noescape
func FuncString(f expvar.Func) string

//go:linkname intstring expvar.sub_intstring
func intstring(v *expvar.Int) string {
	return v.String()
}

//go:linkname IntString expvar.sub_intstring
//go:noescape
func IntString(v *expvar.Int) string

//go:linkname NewMap expvar.NewMap
//go:noescape
func NewMap(name string) *expvar.Map

//go:linkname NewString expvar.NewString
//go:noescape
func NewString(name string) *expvar.String

//go:linkname stringvalue expvar.sub_stringvalue
func stringvalue(v *expvar.String) string {
	return v.Value()
}

//go:linkname StringValue expvar.sub_stringvalue
//go:noescape
func StringValue(v *expvar.String) string

//go:linkname Get expvar.Get
//go:noescape
func Get(name string) expvar.Var

//go:linkname NewInt expvar.NewInt
//go:noescape
func NewInt(name string) *expvar.Int

//go:linkname mapget expvar.sub_mapget
func mapget(v *expvar.Map, key string) expvar.Var {
	return v.Get(key)
}

//go:linkname MapGet expvar.sub_mapget
//go:noescape
func MapGet(v *expvar.Map, key string) expvar.Var

//go:linkname mapinit expvar.sub_mapinit
func mapinit(v *expvar.Map) *expvar.Map {
	return v.Init()
}

//go:linkname MapInit expvar.sub_mapinit
//go:noescape
func MapInit(v *expvar.Map) *expvar.Map
