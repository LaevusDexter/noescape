// This file has automatically been generated on Wed Feb 26 15:50:35 +05 2020.
// DO NOT EDIT.
package types

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	_ "unsafe"
)

//go:linkname NewTuple go/types.NewTuple
//go:noescape
func NewTuple(x ...*types.Var,) *types.Tuple

//go:linkname typeandvaluehasok types.sub_typeandvaluehasok
func typeandvaluehasok(tv types.TypeAndValue,) bool {
	return tv.HasOk()
}

//go:linkname TypeAndValueHasOk types.sub_typeandvaluehasok
//go:noescape
func TypeAndValueHasOk(tv types.TypeAndValue,) bool

//go:linkname funcid types.sub_funcid
func funcid(obj *types.Func,) string {
	return obj.Id()
}

//go:linkname FuncId types.sub_funcid
//go:noescape
func FuncId(obj *types.Func,) string

//go:linkname labelparent types.sub_labelparent
func labelparent(obj *types.Label,) *types.Scope {
	return obj.Parent()
}

//go:linkname LabelParent types.sub_labelparent
//go:noescape
func LabelParent(obj *types.Label,) *types.Scope

//go:linkname nilparent types.sub_nilparent
func nilparent(obj *types.Nil,) *types.Scope {
	return obj.Parent()
}

//go:linkname NilParent types.sub_nilparent
//go:noescape
func NilParent(obj *types.Nil,) *types.Scope

//go:linkname NewField go/types.NewField
//go:noescape
func NewField(pos token.Pos, pkg *types.Package, name string, typ types.Type, embedded bool) *types.Var

//go:linkname constexported types.sub_constexported
func constexported(obj *types.Const,) bool {
	return obj.Exported()
}

//go:linkname ConstExported types.sub_constexported
//go:noescape
func ConstExported(obj *types.Const,) bool

//go:linkname pkgnameexported types.sub_pkgnameexported
func pkgnameexported(obj *types.PkgName,) bool {
	return obj.Exported()
}

//go:linkname PkgNameExported types.sub_pkgnameexported
//go:noescape
func PkgNameExported(obj *types.PkgName,) bool

//go:linkname scopeend types.sub_scopeend
func scopeend(s *types.Scope,) token.Pos {
	return s.End()
}

//go:linkname ScopeEnd types.sub_scopeend
//go:noescape
func ScopeEnd(s *types.Scope,) token.Pos

//go:linkname scopenames types.sub_scopenames
func scopenames(s *types.Scope,) []string {
	return s.Names()
}

//go:linkname ScopeNames types.sub_scopenames
//go:noescape
func ScopeNames(s *types.Scope,) []string

//go:linkname mapkey types.sub_mapkey
func mapkey(m *types.Map,) types.Type {
	return m.Key()
}

//go:linkname MapKey types.sub_mapkey
//go:noescape
func MapKey(m *types.Map,) types.Type

//go:linkname selectionindirect types.sub_selectionindirect
func selectionindirect(s *types.Selection,) bool {
	return s.Indirect()
}

//go:linkname SelectionIndirect types.sub_selectionindirect
//go:noescape
func SelectionIndirect(s *types.Selection,) bool

//go:linkname funcexported types.sub_funcexported
func funcexported(obj *types.Func,) bool {
	return obj.Exported()
}

//go:linkname FuncExported types.sub_funcexported
//go:noescape
func FuncExported(obj *types.Func,) bool

//go:linkname namednummethods types.sub_namednummethods
func namednummethods(t *types.Named,) int {
	return t.NumMethods()
}

//go:linkname NamedNumMethods types.sub_namednummethods
//go:noescape
func NamedNumMethods(t *types.Named,) int

//go:linkname NewSlice go/types.NewSlice
//go:noescape
func NewSlice(elem types.Type,) *types.Slice

//go:linkname constpos types.sub_constpos
func constpos(obj *types.Const,) token.Pos {
	return obj.Pos()
}

//go:linkname ConstPos types.sub_constpos
//go:noescape
func ConstPos(obj *types.Const,) token.Pos

//go:linkname interfaceembedded types.sub_interfaceembedded
func interfaceembedded(t *types.Interface, i int) *types.Named {
	return t.Embedded(i)
}

//go:linkname InterfaceEmbedded types.sub_interfaceembedded
//go:noescape
func InterfaceEmbedded(t *types.Interface, i int) *types.Named

//go:linkname nilpkg types.sub_nilpkg
func nilpkg(obj *types.Nil,) *types.Package {
	return obj.Pkg()
}

//go:linkname NilPkg types.sub_nilpkg
//go:noescape
func NilPkg(obj *types.Nil,) *types.Package

//go:linkname typenamename types.sub_typenamename
func typenamename(obj *types.TypeName,) string {
	return obj.Name()
}

//go:linkname TypeNameName types.sub_typenamename
//go:noescape
func TypeNameName(obj *types.TypeName,) string

//go:linkname infoobjectof types.sub_infoobjectof
func infoobjectof(info *types.Info, id *ast.Ident) types.Object {
	return info.ObjectOf(id)
}

//go:linkname InfoObjectOf types.sub_infoobjectof
//go:noescape
func InfoObjectOf(info *types.Info, id *ast.Ident) types.Object

//go:linkname NewLabel go/types.NewLabel
//go:noescape
func NewLabel(pos token.Pos, pkg *types.Package, name string) *types.Label

//go:linkname namedmethod types.sub_namedmethod
func namedmethod(t *types.Named, i int) *types.Func {
	return t.Method(i)
}

//go:linkname NamedMethod types.sub_namedmethod
//go:noescape
func NamedMethod(t *types.Named, i int) *types.Func

//go:linkname pkgnamepos types.sub_pkgnamepos
func pkgnamepos(obj *types.PkgName,) token.Pos {
	return obj.Pos()
}

//go:linkname PkgNamePos types.sub_pkgnamepos
//go:noescape
func PkgNamePos(obj *types.PkgName,) token.Pos

//go:linkname NewStruct go/types.NewStruct
//go:noescape
func NewStruct(fields []*types.Var, tags []string) *types.Struct

//go:linkname typeandvalueisnil types.sub_typeandvalueisnil
func typeandvalueisnil(tv types.TypeAndValue,) bool {
	return tv.IsNil()
}

//go:linkname TypeAndValueIsNil types.sub_typeandvalueisnil
//go:noescape
func TypeAndValueIsNil(tv types.TypeAndValue,) bool

//go:linkname TypeString go/types.TypeString
//go:noescape
func TypeString(typ types.Type, qf types.Qualifier,) string

//go:linkname NewFunc go/types.NewFunc
//go:noescape
func NewFunc(pos token.Pos, pkg *types.Package, name string, sig *types.Signature,) *types.Func

//go:linkname funcparent types.sub_funcparent
func funcparent(obj *types.Func,) *types.Scope {
	return obj.Parent()
}

//go:linkname FuncParent types.sub_funcparent
//go:noescape
func FuncParent(obj *types.Func,) *types.Scope

//go:linkname infotypeof types.sub_infotypeof
func infotypeof(info *types.Info, e ast.Expr) types.Type {
	return info.TypeOf(e)
}

//go:linkname InfoTypeOf types.sub_infotypeof
//go:noescape
func InfoTypeOf(info *types.Info, e ast.Expr) types.Type

//go:linkname NewMethodSet go/types.NewMethodSet
//go:noescape
func NewMethodSet(T types.Type,) *types.MethodSet

//go:linkname pointerunderlying types.sub_pointerunderlying
func pointerunderlying(p *types.Pointer,) types.Type {
	return p.Underlying()
}

//go:linkname PointerUnderlying types.sub_pointerunderlying
//go:noescape
func PointerUnderlying(p *types.Pointer,) types.Type

//go:linkname sliceunderlying types.sub_sliceunderlying
func sliceunderlying(s *types.Slice,) types.Type {
	return s.Underlying()
}

//go:linkname SliceUnderlying types.sub_sliceunderlying
//go:noescape
func SliceUnderlying(s *types.Slice,) types.Type

//go:linkname varisfield types.sub_varisfield
func varisfield(obj *types.Var,) bool {
	return obj.IsField()
}

//go:linkname VarIsField types.sub_varisfield
//go:noescape
func VarIsField(obj *types.Var,) bool

//go:linkname basicstring types.sub_basicstring
func basicstring(b *types.Basic,) string {
	return b.String()
}

//go:linkname BasicString types.sub_basicstring
//go:noescape
func BasicString(b *types.Basic,) string

//go:linkname builtinpos types.sub_builtinpos
func builtinpos(obj *types.Builtin,) token.Pos {
	return obj.Pos()
}

//go:linkname BuiltinPos types.sub_builtinpos
//go:noescape
func BuiltinPos(obj *types.Builtin,) token.Pos

//go:linkname labelid types.sub_labelid
func labelid(obj *types.Label,) string {
	return obj.Id()
}

//go:linkname LabelId types.sub_labelid
//go:noescape
func LabelId(obj *types.Label,) string

//go:linkname pkgnamename types.sub_pkgnamename
func pkgnamename(obj *types.PkgName,) string {
	return obj.Name()
}

//go:linkname PkgNameName types.sub_pkgnamename
//go:noescape
func PkgNameName(obj *types.PkgName,) string

//go:linkname MissingMethod go/types.MissingMethod
//go:noescape
func MissingMethod(V types.Type, T *types.Interface, static bool) (*types.Func, bool)

//go:linkname funcname types.sub_funcname
func funcname(obj *types.Func,) string {
	return obj.Name()
}

//go:linkname FuncName types.sub_funcname
//go:noescape
func FuncName(obj *types.Func,) string

//go:linkname scopeparent types.sub_scopeparent
func scopeparent(s *types.Scope,) *types.Scope {
	return s.Parent()
}

//go:linkname ScopeParent types.sub_scopeparent
//go:noescape
func ScopeParent(s *types.Scope,) *types.Scope

//go:linkname varstring types.sub_varstring
func varstring(obj *types.Var,) string {
	return obj.String()
}

//go:linkname VarString types.sub_varstring
//go:noescape
func VarString(obj *types.Var,) string

//go:linkname chanelem types.sub_chanelem
func chanelem(c *types.Chan,) types.Type {
	return c.Elem()
}

//go:linkname ChanElem types.sub_chanelem
//go:noescape
func ChanElem(c *types.Chan,) types.Type

//go:linkname errorerror types.sub_errorerror
func errorerror(err types.Error,) string {
	return err.Error()
}

//go:linkname ErrorError types.sub_errorerror
//go:noescape
func ErrorError(err types.Error,) string

//go:linkname packagepath types.sub_packagepath
func packagepath(pkg *types.Package,) string {
	return pkg.Path()
}

//go:linkname PackagePath types.sub_packagepath
//go:noescape
func PackagePath(pkg *types.Package,) string

//go:linkname scopepos types.sub_scopepos
func scopepos(s *types.Scope,) token.Pos {
	return s.Pos()
}

//go:linkname ScopePos types.sub_scopepos
//go:noescape
func ScopePos(s *types.Scope,) token.Pos

//go:linkname NewChecker go/types.NewChecker
//go:noescape
func NewChecker(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info,) *types.Checker

//go:linkname pointerstring types.sub_pointerstring
func pointerstring(p *types.Pointer,) string {
	return p.String()
}

//go:linkname PointerString types.sub_pointerstring
//go:noescape
func PointerString(p *types.Pointer,) string

//go:linkname SizesFor go/types.SizesFor
//go:noescape
func SizesFor(compiler, arch string) types.Sizes

//go:linkname varpkg types.sub_varpkg
func varpkg(obj *types.Var,) *types.Package {
	return obj.Pkg()
}

//go:linkname VarPkg types.sub_varpkg
//go:noescape
func VarPkg(obj *types.Var,) *types.Package

//go:linkname chanunderlying types.sub_chanunderlying
func chanunderlying(c *types.Chan,) types.Type {
	return c.Underlying()
}

//go:linkname ChanUnderlying types.sub_chanunderlying
//go:noescape
func ChanUnderlying(c *types.Chan,) types.Type

//go:linkname methodsetlen types.sub_methodsetlen
func methodsetlen(s *types.MethodSet,) int {
	return s.Len()
}

//go:linkname MethodSetLen types.sub_methodsetlen
//go:noescape
func MethodSetLen(s *types.MethodSet,) int

//go:linkname scopelen types.sub_scopelen
func scopelen(s *types.Scope,) int {
	return s.Len()
}

//go:linkname ScopeLen types.sub_scopelen
//go:noescape
func ScopeLen(s *types.Scope,) int

//go:linkname NewArray go/types.NewArray
//go:noescape
func NewArray(elem types.Type, len int64) *types.Array

//go:linkname chandir types.sub_chandir
func chandir(c *types.Chan,) types.ChanDir {
	return c.Dir()
}

//go:linkname ChanDir types.sub_chandir
//go:noescape
func ChanDir(c *types.Chan,) types.ChanDir

//go:linkname packageimports types.sub_packageimports
func packageimports(pkg *types.Package,) []*types.Package {
	return pkg.Imports()
}

//go:linkname PackageImports types.sub_packageimports
//go:noescape
func PackageImports(pkg *types.Package,) []*types.Package

//go:linkname Implements go/types.Implements
//go:noescape
func Implements(V types.Type, T *types.Interface,) bool

//go:linkname namedobj types.sub_namedobj
func namedobj(t *types.Named,) *types.TypeName {
	return t.Obj()
}

//go:linkname NamedObj types.sub_namedobj
//go:noescape
func NamedObj(t *types.Named,) *types.TypeName

//go:linkname labelname types.sub_labelname
func labelname(obj *types.Label,) string {
	return obj.Name()
}

//go:linkname LabelName types.sub_labelname
//go:noescape
func LabelName(obj *types.Label,) string

//go:linkname selectionindex types.sub_selectionindex
func selectionindex(s *types.Selection,) []int {
	return s.Index()
}

//go:linkname SelectionIndex types.sub_selectionindex
//go:noescape
func SelectionIndex(s *types.Selection,) []int

//go:linkname sliceelem types.sub_sliceelem
func sliceelem(s *types.Slice,) types.Type {
	return s.Elem()
}

//go:linkname SliceElem types.sub_sliceelem
//go:noescape
func SliceElem(s *types.Slice,) types.Type

//go:linkname typeandvalueassignable types.sub_typeandvalueassignable
func typeandvalueassignable(tv types.TypeAndValue,) bool {
	return tv.Assignable()
}

//go:linkname TypeAndValueAssignable types.sub_typeandvalueassignable
//go:noescape
func TypeAndValueAssignable(tv types.TypeAndValue,) bool

//go:linkname varexported types.sub_varexported
func varexported(obj *types.Var,) bool {
	return obj.Exported()
}

//go:linkname VarExported types.sub_varexported
//go:noescape
func VarExported(obj *types.Var,) bool

//go:linkname builtinpkg types.sub_builtinpkg
func builtinpkg(obj *types.Builtin,) *types.Package {
	return obj.Pkg()
}

//go:linkname BuiltinPkg types.sub_builtinpkg
//go:noescape
func BuiltinPkg(obj *types.Builtin,) *types.Package

//go:linkname signaturestring types.sub_signaturestring
func signaturestring(s *types.Signature,) string {
	return s.String()
}

//go:linkname SignatureString types.sub_signaturestring
//go:noescape
func SignatureString(s *types.Signature,) string

//go:linkname selectiontype types.sub_selectiontype
func selectiontype(s *types.Selection,) types.Type {
	return s.Type()
}

//go:linkname SelectionType types.sub_selectiontype
//go:noescape
func SelectionType(s *types.Selection,) types.Type

//go:linkname interfacemethod types.sub_interfacemethod
func interfacemethod(t *types.Interface, i int) *types.Func {
	return t.Method(i)
}

//go:linkname InterfaceMethod types.sub_interfacemethod
//go:noescape
func InterfaceMethod(t *types.Interface, i int) *types.Func

//go:linkname labeltype types.sub_labeltype
func labeltype(obj *types.Label,) types.Type {
	return obj.Type()
}

//go:linkname LabelType types.sub_labeltype
//go:noescape
func LabelType(obj *types.Label,) types.Type

//go:linkname builtinparent types.sub_builtinparent
func builtinparent(obj *types.Builtin,) *types.Scope {
	return obj.Parent()
}

//go:linkname BuiltinParent types.sub_builtinparent
//go:noescape
func BuiltinParent(obj *types.Builtin,) *types.Scope

//go:linkname checkerfiles types.sub_checkerfiles
func checkerfiles(check *types.Checker, files []*ast.File) error {
	return check.Files(files)
}

//go:linkname CheckerFiles types.sub_checkerfiles
//go:noescape
func CheckerFiles(check *types.Checker, files []*ast.File) error

//go:linkname typenameisalias types.sub_typenameisalias
func typenameisalias(obj *types.TypeName,) bool {
	return obj.IsAlias()
}

//go:linkname TypeNameIsAlias types.sub_typenameisalias
//go:noescape
func TypeNameIsAlias(obj *types.TypeName,) bool

//go:linkname arraystring types.sub_arraystring
func arraystring(a *types.Array,) string {
	return a.String()
}

//go:linkname ArrayString types.sub_arraystring
//go:noescape
func ArrayString(a *types.Array,) string

//go:linkname basickind types.sub_basickind
func basickind(b *types.Basic,) types.BasicKind {
	return b.Kind()
}

//go:linkname BasicKind types.sub_basickind
//go:noescape
func BasicKind(b *types.Basic,) types.BasicKind

//go:linkname interfacenumembeddeds types.sub_interfacenumembeddeds
func interfacenumembeddeds(t *types.Interface,) int {
	return t.NumEmbeddeds()
}

//go:linkname InterfaceNumEmbeddeds types.sub_interfacenumembeddeds
//go:noescape
func InterfaceNumEmbeddeds(t *types.Interface,) int

//go:linkname interfacenumexplicitmethods types.sub_interfacenumexplicitmethods
func interfacenumexplicitmethods(t *types.Interface,) int {
	return t.NumExplicitMethods()
}

//go:linkname InterfaceNumExplicitMethods types.sub_interfacenumexplicitmethods
//go:noescape
func InterfaceNumExplicitMethods(t *types.Interface,) int

//go:linkname labelexported types.sub_labelexported
func labelexported(obj *types.Label,) bool {
	return obj.Exported()
}

//go:linkname LabelExported types.sub_labelexported
//go:noescape
func LabelExported(obj *types.Label,) bool

//go:linkname packagescope types.sub_packagescope
func packagescope(pkg *types.Package,) *types.Scope {
	return pkg.Scope()
}

//go:linkname PackageScope types.sub_packagescope
//go:noescape
func PackageScope(pkg *types.Package,) *types.Scope

//go:linkname typenamepos types.sub_typenamepos
func typenamepos(obj *types.TypeName,) token.Pos {
	return obj.Pos()
}

//go:linkname TypeNamePos types.sub_typenamepos
//go:noescape
func TypeNamePos(obj *types.TypeName,) token.Pos

//go:linkname constid types.sub_constid
func constid(obj *types.Const,) string {
	return obj.Id()
}

//go:linkname ConstId types.sub_constid
//go:noescape
func ConstId(obj *types.Const,) string

//go:linkname nilpos types.sub_nilpos
func nilpos(obj *types.Nil,) token.Pos {
	return obj.Pos()
}

//go:linkname NilPos types.sub_nilpos
//go:noescape
func NilPos(obj *types.Nil,) token.Pos

//go:linkname pkgnameid types.sub_pkgnameid
func pkgnameid(obj *types.PkgName,) string {
	return obj.Id()
}

//go:linkname PkgNameId types.sub_pkgnameid
//go:noescape
func PkgNameId(obj *types.PkgName,) string

//go:linkname CheckExpr go/types.CheckExpr
//go:noescape
func CheckExpr(fset *token.FileSet, pkg *types.Package, pos token.Pos, expr ast.Expr, info *types.Info,) error

//go:linkname interfaceembeddedtype types.sub_interfaceembeddedtype
func interfaceembeddedtype(t *types.Interface, i int) types.Type {
	return t.EmbeddedType(i)
}

//go:linkname InterfaceEmbeddedType types.sub_interfaceembeddedtype
//go:noescape
func InterfaceEmbeddedType(t *types.Interface, i int) types.Type

//go:linkname selectionkind types.sub_selectionkind
func selectionkind(s *types.Selection,) types.SelectionKind {
	return s.Kind()
}

//go:linkname SelectionKind types.sub_selectionkind
//go:noescape
func SelectionKind(s *types.Selection,) types.SelectionKind

//go:linkname selectionstring types.sub_selectionstring
func selectionstring(s *types.Selection,) string {
	return s.String()
}

//go:linkname SelectionString types.sub_selectionstring
//go:noescape
func SelectionString(s *types.Selection,) string

//go:linkname tupleat types.sub_tupleat
func tupleat(t *types.Tuple, i int) *types.Var {
	return t.At(i)
}

//go:linkname TupleAt types.sub_tupleat
//go:noescape
func TupleAt(t *types.Tuple, i int) *types.Var

//go:linkname ObjectString go/types.ObjectString
//go:noescape
func ObjectString(obj types.Object, qf types.Qualifier,) string

//go:linkname builtinstring types.sub_builtinstring
func builtinstring(obj *types.Builtin,) string {
	return obj.String()
}

//go:linkname BuiltinString types.sub_builtinstring
//go:noescape
func BuiltinString(obj *types.Builtin,) string

//go:linkname interfaceempty types.sub_interfaceempty
func interfaceempty(t *types.Interface,) bool {
	return t.Empty()
}

//go:linkname InterfaceEmpty types.sub_interfaceempty
//go:noescape
func InterfaceEmpty(t *types.Interface,) bool

//go:linkname pkgnamepkg types.sub_pkgnamepkg
func pkgnamepkg(obj *types.PkgName,) *types.Package {
	return obj.Pkg()
}

//go:linkname PkgNamePkg types.sub_pkgnamepkg
//go:noescape
func PkgNamePkg(obj *types.PkgName,) *types.Package

//go:linkname Eval go/types.Eval
//go:noescape
func Eval(fset *token.FileSet, pkg *types.Package, pos token.Pos, expr string) (types.TypeAndValue, error)

//go:linkname NewParam go/types.NewParam
//go:noescape
func NewParam(pos token.Pos, pkg *types.Package, name string, typ types.Type,) *types.Var

//go:linkname interfaceexplicitmethod types.sub_interfaceexplicitmethod
func interfaceexplicitmethod(t *types.Interface, i int) *types.Func {
	return t.ExplicitMethod(i)
}

//go:linkname InterfaceExplicitMethod types.sub_interfaceexplicitmethod
//go:noescape
func InterfaceExplicitMethod(t *types.Interface, i int) *types.Func

//go:linkname scopelookup types.sub_scopelookup
func scopelookup(s *types.Scope, name string) types.Object {
	return s.Lookup(name)
}

//go:linkname ScopeLookup types.sub_scopelookup
//go:noescape
func ScopeLookup(s *types.Scope, name string) types.Object

//go:linkname pkgnamestring types.sub_pkgnamestring
func pkgnamestring(obj *types.PkgName,) string {
	return obj.String()
}

//go:linkname PkgNameString types.sub_pkgnamestring
//go:noescape
func PkgNameString(obj *types.PkgName,) string

//go:linkname arrayunderlying types.sub_arrayunderlying
func arrayunderlying(a *types.Array,) types.Type {
	return a.Underlying()
}

//go:linkname ArrayUnderlying types.sub_arrayunderlying
//go:noescape
func ArrayUnderlying(a *types.Array,) types.Type

//go:linkname basicinfo types.sub_basicinfo
func basicinfo(b *types.Basic,) types.BasicInfo {
	return b.Info()
}

//go:linkname BasicInfo types.sub_basicinfo
//go:noescape
func BasicInfo(b *types.Basic,) types.BasicInfo

//go:linkname chanstring types.sub_chanstring
func chanstring(c *types.Chan,) string {
	return c.String()
}

//go:linkname ChanString types.sub_chanstring
//go:noescape
func ChanString(c *types.Chan,) string

//go:linkname funcpkg types.sub_funcpkg
func funcpkg(obj *types.Func,) *types.Package {
	return obj.Pkg()
}

//go:linkname FuncPkg types.sub_funcpkg
//go:noescape
func FuncPkg(obj *types.Func,) *types.Package

//go:linkname methodsetat types.sub_methodsetat
func methodsetat(s *types.MethodSet, i int) *types.Selection {
	return s.At(i)
}

//go:linkname MethodSetAt types.sub_methodsetat
//go:noescape
func MethodSetAt(s *types.MethodSet, i int) *types.Selection

//go:linkname RelativeTo go/types.RelativeTo
//go:noescape
func RelativeTo(pkg *types.Package,) types.Qualifier

//go:linkname typenametype types.sub_typenametype
func typenametype(obj *types.TypeName,) types.Type {
	return obj.Type()
}

//go:linkname TypeNameType types.sub_typenametype
//go:noescape
func TypeNameType(obj *types.TypeName,) types.Type

//go:linkname basicunderlying types.sub_basicunderlying
func basicunderlying(b *types.Basic,) types.Type {
	return b.Underlying()
}

//go:linkname BasicUnderlying types.sub_basicunderlying
//go:noescape
func BasicUnderlying(b *types.Basic,) types.Type

//go:linkname NewMap go/types.NewMap
//go:noescape
func NewMap(key, elem types.Type,) *types.Map

//go:linkname nilexported types.sub_nilexported
func nilexported(obj *types.Nil,) bool {
	return obj.Exported()
}

//go:linkname NilExported types.sub_nilexported
//go:noescape
func NilExported(obj *types.Nil,) bool

//go:linkname pkgnameparent types.sub_pkgnameparent
func pkgnameparent(obj *types.PkgName,) *types.Scope {
	return obj.Parent()
}

//go:linkname PkgNameParent types.sub_pkgnameparent
//go:noescape
func PkgNameParent(obj *types.PkgName,) *types.Scope

//go:linkname varparent types.sub_varparent
func varparent(obj *types.Var,) *types.Scope {
	return obj.Parent()
}

//go:linkname VarParent types.sub_varparent
//go:noescape
func VarParent(obj *types.Var,) *types.Scope

//go:linkname signatureunderlying types.sub_signatureunderlying
func signatureunderlying(s *types.Signature,) types.Type {
	return s.Underlying()
}

//go:linkname SignatureUnderlying types.sub_signatureunderlying
//go:noescape
func SignatureUnderlying(s *types.Signature,) types.Type

//go:linkname typeandvalueisvoid types.sub_typeandvalueisvoid
func typeandvalueisvoid(tv types.TypeAndValue,) bool {
	return tv.IsVoid()
}

//go:linkname TypeAndValueIsVoid types.sub_typeandvalueisvoid
//go:noescape
func TypeAndValueIsVoid(tv types.TypeAndValue,) bool

//go:linkname interfacecomplete types.sub_interfacecomplete
func interfacecomplete(t *types.Interface,) *types.Interface {
	return t.Complete()
}

//go:linkname InterfaceComplete types.sub_interfacecomplete
//go:noescape
func InterfaceComplete(t *types.Interface,) *types.Interface

//go:linkname tuplestring types.sub_tuplestring
func tuplestring(t *types.Tuple,) string {
	return t.String()
}

//go:linkname TupleString types.sub_tuplestring
//go:noescape
func TupleString(t *types.Tuple,) string

//go:linkname mapelem types.sub_mapelem
func mapelem(m *types.Map,) types.Type {
	return m.Elem()
}

//go:linkname MapElem types.sub_mapelem
//go:noescape
func MapElem(m *types.Map,) types.Type

//go:linkname packagename types.sub_packagename
func packagename(pkg *types.Package,) string {
	return pkg.Name()
}

//go:linkname PackageName types.sub_packagename
//go:noescape
func PackageName(pkg *types.Package,) string

//go:linkname structnumfields types.sub_structnumfields
func structnumfields(s *types.Struct,) int {
	return s.NumFields()
}

//go:linkname StructNumFields types.sub_structnumfields
//go:noescape
func StructNumFields(s *types.Struct,) int

//go:linkname builtinid types.sub_builtinid
func builtinid(obj *types.Builtin,) string {
	return obj.Id()
}

//go:linkname BuiltinId types.sub_builtinid
//go:noescape
func BuiltinId(obj *types.Builtin,) string

//go:linkname funcstring types.sub_funcstring
func funcstring(obj *types.Func,) string {
	return obj.String()
}

//go:linkname FuncString types.sub_funcstring
//go:noescape
func FuncString(obj *types.Func,) string

//go:linkname functype types.sub_functype
func functype(obj *types.Func,) types.Type {
	return obj.Type()
}

//go:linkname FuncType types.sub_functype
//go:noescape
func FuncType(obj *types.Func,) types.Type

//go:linkname signatureparams types.sub_signatureparams
func signatureparams(s *types.Signature,) *types.Tuple {
	return s.Params()
}

//go:linkname SignatureParams types.sub_signatureparams
//go:noescape
func SignatureParams(s *types.Signature,) *types.Tuple

//go:linkname funcscope types.sub_funcscope
func funcscope(obj *types.Func,) *types.Scope {
	return obj.Scope()
}

//go:linkname FuncScope types.sub_funcscope
//go:noescape
func FuncScope(obj *types.Func,) *types.Scope

//go:linkname scopestring types.sub_scopestring
func scopestring(s *types.Scope,) string {
	return s.String()
}

//go:linkname ScopeString types.sub_scopestring
//go:noescape
func ScopeString(s *types.Scope,) string

//go:linkname typenameid types.sub_typenameid
func typenameid(obj *types.TypeName,) string {
	return obj.Id()
}

//go:linkname TypeNameId types.sub_typenameid
//go:noescape
func TypeNameId(obj *types.TypeName,) string

//go:linkname varname types.sub_varname
func varname(obj *types.Var,) string {
	return obj.Name()
}

//go:linkname VarName types.sub_varname
//go:noescape
func VarName(obj *types.Var,) string

//go:linkname labelstring types.sub_labelstring
func labelstring(obj *types.Label,) string {
	return obj.String()
}

//go:linkname LabelString types.sub_labelstring
//go:noescape
func LabelString(obj *types.Label,) string

//go:linkname namedunderlying types.sub_namedunderlying
func namedunderlying(t *types.Named,) types.Type {
	return t.Underlying()
}

//go:linkname NamedUnderlying types.sub_namedunderlying
//go:noescape
func NamedUnderlying(t *types.Named,) types.Type

//go:linkname nilname types.sub_nilname
func nilname(obj *types.Nil,) string {
	return obj.Name()
}

//go:linkname NilName types.sub_nilname
//go:noescape
func NilName(obj *types.Nil,) string

//go:linkname packagestring types.sub_packagestring
func packagestring(pkg *types.Package,) string {
	return pkg.String()
}

//go:linkname PackageString types.sub_packagestring
//go:noescape
func PackageString(pkg *types.Package,) string

//go:linkname NewPkgName go/types.NewPkgName
//go:noescape
func NewPkgName(pos token.Pos, pkg *types.Package, name string, imported *types.Package,) *types.PkgName

//go:linkname structfield types.sub_structfield
func structfield(s *types.Struct, i int) *types.Var {
	return s.Field(i)
}

//go:linkname StructField types.sub_structfield
//go:noescape
func StructField(s *types.Struct, i int) *types.Var

//go:linkname Default go/types.Default
//go:noescape
func Default(typ types.Type,) types.Type

//go:linkname typenamepkg types.sub_typenamepkg
func typenamepkg(obj *types.TypeName,) *types.Package {
	return obj.Pkg()
}

//go:linkname TypeNamePkg types.sub_typenamepkg
//go:noescape
func TypeNamePkg(obj *types.TypeName,) *types.Package

//go:linkname IdenticalIgnoreTags go/types.IdenticalIgnoreTags
//go:noescape
func IdenticalIgnoreTags(x, y types.Type,) bool

//go:linkname constpkg types.sub_constpkg
func constpkg(obj *types.Const,) *types.Package {
	return obj.Pkg()
}

//go:linkname ConstPkg types.sub_constpkg
//go:noescape
func ConstPkg(obj *types.Const,) *types.Package

//go:linkname scopecontains types.sub_scopecontains
func scopecontains(s *types.Scope, pos token.Pos) bool {
	return s.Contains(pos)
}

//go:linkname ScopeContains types.sub_scopecontains
//go:noescape
func ScopeContains(s *types.Scope, pos token.Pos) bool

//go:linkname vartype types.sub_vartype
func vartype(obj *types.Var,) types.Type {
	return obj.Type()
}

//go:linkname VarType types.sub_vartype
//go:noescape
func VarType(obj *types.Var,) types.Type

//go:linkname scopeinsert types.sub_scopeinsert
func scopeinsert(s *types.Scope, obj types.Object,) types.Object {
	return s.Insert(obj)
}

//go:linkname ScopeInsert types.sub_scopeinsert
//go:noescape
func ScopeInsert(s *types.Scope, obj types.Object,) types.Object

//go:linkname typenameparent types.sub_typenameparent
func typenameparent(obj *types.TypeName,) *types.Scope {
	return obj.Parent()
}

//go:linkname TypeNameParent types.sub_typenameparent
//go:noescape
func TypeNameParent(obj *types.TypeName,) *types.Scope

//go:linkname labelpos types.sub_labelpos
func labelpos(obj *types.Label,) token.Pos {
	return obj.Pos()
}

//go:linkname LabelPos types.sub_labelpos
//go:noescape
func LabelPos(obj *types.Label,) token.Pos

//go:linkname NewConst go/types.NewConst
//go:noescape
func NewConst(pos token.Pos, pkg *types.Package, name string, typ types.Type, val constant.Value) *types.Const

//go:linkname pkgnameimported types.sub_pkgnameimported
func pkgnameimported(obj *types.PkgName,) *types.Package {
	return obj.Imported()
}

//go:linkname PkgNameImported types.sub_pkgnameimported
//go:noescape
func PkgNameImported(obj *types.PkgName,) *types.Package

//go:linkname NewTypeName go/types.NewTypeName
//go:noescape
func NewTypeName(pos token.Pos, pkg *types.Package, name string, typ types.Type,) *types.TypeName

//go:linkname ConvertibleTo go/types.ConvertibleTo
//go:noescape
func ConvertibleTo(V, T types.Type,) bool

//go:linkname conststring types.sub_conststring
func conststring(obj *types.Const,) string {
	return obj.String()
}

//go:linkname ConstString types.sub_conststring
//go:noescape
func ConstString(obj *types.Const,) string

//go:linkname scopechild types.sub_scopechild
func scopechild(s *types.Scope, i int) *types.Scope {
	return s.Child(i)
}

//go:linkname ScopeChild types.sub_scopechild
//go:noescape
func ScopeChild(s *types.Scope, i int) *types.Scope

//go:linkname slicestring types.sub_slicestring
func slicestring(s *types.Slice,) string {
	return s.String()
}

//go:linkname SliceString types.sub_slicestring
//go:noescape
func SliceString(s *types.Slice,) string

//go:linkname structunderlying types.sub_structunderlying
func structunderlying(s *types.Struct,) types.Type {
	return s.Underlying()
}

//go:linkname StructUnderlying types.sub_structunderlying
//go:noescape
func StructUnderlying(s *types.Struct,) types.Type

//go:linkname builtinname types.sub_builtinname
func builtinname(obj *types.Builtin,) string {
	return obj.Name()
}

//go:linkname BuiltinName types.sub_builtinname
//go:noescape
func BuiltinName(obj *types.Builtin,) string

//go:linkname interfacenummethods types.sub_interfacenummethods
func interfacenummethods(t *types.Interface,) int {
	return t.NumMethods()
}

//go:linkname InterfaceNumMethods types.sub_interfacenummethods
//go:noescape
func InterfaceNumMethods(t *types.Interface,) int

//go:linkname selectionobj types.sub_selectionobj
func selectionobj(s *types.Selection,) types.Object {
	return s.Obj()
}

//go:linkname SelectionObj types.sub_selectionobj
//go:noescape
func SelectionObj(s *types.Selection,) types.Object

//go:linkname nilid types.sub_nilid
func nilid(obj *types.Nil,) string {
	return obj.Id()
}

//go:linkname NilId types.sub_nilid
//go:noescape
func NilId(obj *types.Nil,) string

//go:linkname stdsizesoffsetsof types.sub_stdsizesoffsetsof
func stdsizesoffsetsof(s *types.StdSizes, fields []*types.Var,) []int64 {
	return s.Offsetsof(fields)
}

//go:linkname StdSizesOffsetsof types.sub_stdsizesoffsetsof
//go:noescape
func StdSizesOffsetsof(s *types.StdSizes, fields []*types.Var,) []int64

//go:linkname varembedded types.sub_varembedded
func varembedded(obj *types.Var,) bool {
	return obj.Embedded()
}

//go:linkname VarEmbedded types.sub_varembedded
//go:noescape
func VarEmbedded(obj *types.Var,) bool

//go:linkname interfacestring types.sub_interfacestring
func interfacestring(t *types.Interface,) string {
	return t.String()
}

//go:linkname InterfaceString types.sub_interfacestring
//go:noescape
func InterfaceString(t *types.Interface,) string

//go:linkname niltype types.sub_niltype
func niltype(obj *types.Nil,) types.Type {
	return obj.Type()
}

//go:linkname NilType types.sub_niltype
//go:noescape
func NilType(obj *types.Nil,) types.Type

//go:linkname selectionrecv types.sub_selectionrecv
func selectionrecv(s *types.Selection,) types.Type {
	return s.Recv()
}

//go:linkname SelectionRecv types.sub_selectionrecv
//go:noescape
func SelectionRecv(s *types.Selection,) types.Type

//go:linkname varid types.sub_varid
func varid(obj *types.Var,) string {
	return obj.Id()
}

//go:linkname VarId types.sub_varid
//go:noescape
func VarId(obj *types.Var,) string

//go:linkname builtintype types.sub_builtintype
func builtintype(obj *types.Builtin,) types.Type {
	return obj.Type()
}

//go:linkname BuiltinType types.sub_builtintype
//go:noescape
func BuiltinType(obj *types.Builtin,) types.Type

//go:linkname NewInterfaceType go/types.NewInterfaceType
//go:noescape
func NewInterfaceType(methods []*types.Func, embeddeds []types.Type,) *types.Interface

//go:linkname mapstring types.sub_mapstring
func mapstring(m *types.Map,) string {
	return m.String()
}

//go:linkname MapString types.sub_mapstring
//go:noescape
func MapString(m *types.Map,) string

//go:linkname varpos types.sub_varpos
func varpos(obj *types.Var,) token.Pos {
	return obj.Pos()
}

//go:linkname VarPos types.sub_varpos
//go:noescape
func VarPos(obj *types.Var,) token.Pos

//go:linkname AssignableTo go/types.AssignableTo
//go:noescape
func AssignableTo(V, T types.Type,) bool

//go:linkname arrayelem types.sub_arrayelem
func arrayelem(a *types.Array,) types.Type {
	return a.Elem()
}

//go:linkname ArrayElem types.sub_arrayelem
//go:noescape
func ArrayElem(a *types.Array,) types.Type

//go:linkname NewInterface go/types.NewInterface
//go:noescape
func NewInterface(methods []*types.Func, embeddeds []*types.Named,) *types.Interface

//go:linkname NewPackage go/types.NewPackage
//go:noescape
func NewPackage(path, name string) *types.Package

//go:linkname NewPointer go/types.NewPointer
//go:noescape
func NewPointer(elem types.Type,) *types.Pointer

//go:linkname builtinexported types.sub_builtinexported
func builtinexported(obj *types.Builtin,) bool {
	return obj.Exported()
}

//go:linkname BuiltinExported types.sub_builtinexported
//go:noescape
func BuiltinExported(obj *types.Builtin,) bool

//go:linkname LookupFieldOrMethod go/types.LookupFieldOrMethod
//go:noescape
func LookupFieldOrMethod(T types.Type, addressable bool, pkg *types.Package, name string) (types.Object, []int, bool)

//go:linkname tupleunderlying types.sub_tupleunderlying
func tupleunderlying(t *types.Tuple,) types.Type {
	return t.Underlying()
}

//go:linkname TupleUnderlying types.sub_tupleunderlying
//go:noescape
func TupleUnderlying(t *types.Tuple,) types.Type

//go:linkname typeandvalueistype types.sub_typeandvalueistype
func typeandvalueistype(tv types.TypeAndValue,) bool {
	return tv.IsType()
}

//go:linkname TypeAndValueIsType types.sub_typeandvalueistype
//go:noescape
func TypeAndValueIsType(tv types.TypeAndValue,) bool

//go:linkname basicname types.sub_basicname
func basicname(b *types.Basic,) string {
	return b.Name()
}

//go:linkname BasicName types.sub_basicname
//go:noescape
func BasicName(b *types.Basic,) string

//go:linkname signaturerecv types.sub_signaturerecv
func signaturerecv(s *types.Signature,) *types.Var {
	return s.Recv()
}

//go:linkname SignatureRecv types.sub_signaturerecv
//go:noescape
func SignatureRecv(s *types.Signature,) *types.Var

//go:linkname typeandvalueaddressable types.sub_typeandvalueaddressable
func typeandvalueaddressable(tv types.TypeAndValue,) bool {
	return tv.Addressable()
}

//go:linkname TypeAndValueAddressable types.sub_typeandvalueaddressable
//go:noescape
func TypeAndValueAddressable(tv types.TypeAndValue,) bool

//go:linkname packagecomplete types.sub_packagecomplete
func packagecomplete(pkg *types.Package,) bool {
	return pkg.Complete()
}

//go:linkname PackageComplete types.sub_packagecomplete
//go:noescape
func PackageComplete(pkg *types.Package,) bool

//go:linkname NewVar go/types.NewVar
//go:noescape
func NewVar(pos token.Pos, pkg *types.Package, name string, typ types.Type,) *types.Var

//go:linkname funcfullname types.sub_funcfullname
func funcfullname(obj *types.Func,) string {
	return obj.FullName()
}

//go:linkname FuncFullName types.sub_funcfullname
//go:noescape
func FuncFullName(obj *types.Func,) string

//go:linkname scopenumchildren types.sub_scopenumchildren
func scopenumchildren(s *types.Scope,) int {
	return s.NumChildren()
}

//go:linkname ScopeNumChildren types.sub_scopenumchildren
//go:noescape
func ScopeNumChildren(s *types.Scope,) int

//go:linkname structstring types.sub_structstring
func structstring(s *types.Struct,) string {
	return s.String()
}

//go:linkname StructString types.sub_structstring
//go:noescape
func StructString(s *types.Struct,) string

//go:linkname structtag types.sub_structtag
func structtag(s *types.Struct, i int) string {
	return s.Tag(i)
}

//go:linkname StructTag types.sub_structtag
//go:noescape
func StructTag(s *types.Struct, i int) string

//go:linkname typenameexported types.sub_typenameexported
func typenameexported(obj *types.TypeName,) bool {
	return obj.Exported()
}

//go:linkname TypeNameExported types.sub_typenameexported
//go:noescape
func TypeNameExported(obj *types.TypeName,) bool

//go:linkname constparent types.sub_constparent
func constparent(obj *types.Const,) *types.Scope {
	return obj.Parent()
}

//go:linkname ConstParent types.sub_constparent
//go:noescape
func ConstParent(obj *types.Const,) *types.Scope

//go:linkname signatureresults types.sub_signatureresults
func signatureresults(s *types.Signature,) *types.Tuple {
	return s.Results()
}

//go:linkname SignatureResults types.sub_signatureresults
//go:noescape
func SignatureResults(s *types.Signature,) *types.Tuple

//go:linkname tuplelen types.sub_tuplelen
func tuplelen(t *types.Tuple,) int {
	return t.Len()
}

//go:linkname TupleLen types.sub_tuplelen
//go:noescape
func TupleLen(t *types.Tuple,) int

//go:linkname typeandvalueisvalue types.sub_typeandvalueisvalue
func typeandvalueisvalue(tv types.TypeAndValue,) bool {
	return tv.IsValue()
}

//go:linkname TypeAndValueIsValue types.sub_typeandvalueisvalue
//go:noescape
func TypeAndValueIsValue(tv types.TypeAndValue,) bool

//go:linkname NewChan go/types.NewChan
//go:noescape
func NewChan(dir types.ChanDir, elem types.Type,) *types.Chan

//go:linkname initializerstring types.sub_initializerstring
func initializerstring(init *types.Initializer,) string {
	return init.String()
}

//go:linkname InitializerString types.sub_initializerstring
//go:noescape
func InitializerString(init *types.Initializer,) string

//go:linkname NewSignature go/types.NewSignature
//go:noescape
func NewSignature(recv *types.Var, params, results *types.Tuple, variadic bool) *types.Signature

//go:linkname ExprString go/types.ExprString
//go:noescape
func ExprString(x ast.Expr) string

//go:linkname consttype types.sub_consttype
func consttype(obj *types.Const,) types.Type {
	return obj.Type()
}

//go:linkname ConstType types.sub_consttype
//go:noescape
func ConstType(obj *types.Const,) types.Type

//go:linkname typeandvalueisbuiltin types.sub_typeandvalueisbuiltin
func typeandvalueisbuiltin(tv types.TypeAndValue,) bool {
	return tv.IsBuiltin()
}

//go:linkname TypeAndValueIsBuiltin types.sub_typeandvalueisbuiltin
//go:noescape
func TypeAndValueIsBuiltin(tv types.TypeAndValue,) bool

//go:linkname AssertableTo go/types.AssertableTo
//go:noescape
func AssertableTo(V *types.Interface, T types.Type,) bool

//go:linkname NewNamed go/types.NewNamed
//go:noescape
func NewNamed(obj *types.TypeName, underlying types.Type, methods []*types.Func,) *types.Named

//go:linkname NewScope go/types.NewScope
//go:noescape
func NewScope(parent *types.Scope, pos, end token.Pos, comment string) *types.Scope

//go:linkname scopeinnermost types.sub_scopeinnermost
func scopeinnermost(s *types.Scope, pos token.Pos) *types.Scope {
	return s.Innermost(pos)
}

//go:linkname ScopeInnermost types.sub_scopeinnermost
//go:noescape
func ScopeInnermost(s *types.Scope, pos token.Pos) *types.Scope

//go:linkname stdsizessizeof types.sub_stdsizessizeof
func stdsizessizeof(s *types.StdSizes, T types.Type,) int64 {
	return s.Sizeof(T)
}

//go:linkname StdSizesSizeof types.sub_stdsizessizeof
//go:noescape
func StdSizesSizeof(s *types.StdSizes, T types.Type,) int64

//go:linkname Id go/types.Id
//go:noescape
func Id(pkg *types.Package, name string) string

//go:linkname constval types.sub_constval
func constval(obj *types.Const,) constant.Value {
	return obj.Val()
}

//go:linkname ConstVal types.sub_constval
//go:noescape
func ConstVal(obj *types.Const,) constant.Value

//go:linkname nilstring types.sub_nilstring
func nilstring(obj *types.Nil,) string {
	return obj.String()
}

//go:linkname NilString types.sub_nilstring
//go:noescape
func NilString(obj *types.Nil,) string

//go:linkname scopelookupparent types.sub_scopelookupparent
func scopelookupparent(s *types.Scope, name string, pos token.Pos) (*types.Scope, types.Object,) {
	return s.LookupParent(name, pos)
}

//go:linkname ScopeLookupParent types.sub_scopelookupparent
//go:noescape
func ScopeLookupParent(s *types.Scope, name string, pos token.Pos) (*types.Scope, types.Object,)

//go:linkname signaturevariadic types.sub_signaturevariadic
func signaturevariadic(s *types.Signature,) bool {
	return s.Variadic()
}

//go:linkname SignatureVariadic types.sub_signaturevariadic
//go:noescape
func SignatureVariadic(s *types.Signature,) bool

//go:linkname stdsizesalignof types.sub_stdsizesalignof
func stdsizesalignof(s *types.StdSizes, T types.Type,) int64 {
	return s.Alignof(T)
}

//go:linkname StdSizesAlignof types.sub_stdsizesalignof
//go:noescape
func StdSizesAlignof(s *types.StdSizes, T types.Type,) int64

//go:linkname typenamestring types.sub_typenamestring
func typenamestring(obj *types.TypeName,) string {
	return obj.String()
}

//go:linkname TypeNameString types.sub_typenamestring
//go:noescape
func TypeNameString(obj *types.TypeName,) string

//go:linkname IsInterface go/types.IsInterface
//go:noescape
func IsInterface(typ types.Type,) bool

//go:linkname labelpkg types.sub_labelpkg
func labelpkg(obj *types.Label,) *types.Package {
	return obj.Pkg()
}

//go:linkname LabelPkg types.sub_labelpkg
//go:noescape
func LabelPkg(obj *types.Label,) *types.Package

//go:linkname mapunderlying types.sub_mapunderlying
func mapunderlying(m *types.Map,) types.Type {
	return m.Underlying()
}

//go:linkname MapUnderlying types.sub_mapunderlying
//go:noescape
func MapUnderlying(m *types.Map,) types.Type

//go:linkname Comparable go/types.Comparable
//go:noescape
func Comparable(T types.Type,) bool

//go:linkname methodsetlookup types.sub_methodsetlookup
func methodsetlookup(s *types.MethodSet, pkg *types.Package, name string) *types.Selection {
	return s.Lookup(pkg, name)
}

//go:linkname MethodSetLookup types.sub_methodsetlookup
//go:noescape
func MethodSetLookup(s *types.MethodSet, pkg *types.Package, name string) *types.Selection

//go:linkname configcheck types.sub_configcheck
func configcheck(conf *types.Config, path string, fset *token.FileSet, files []*ast.File, info *types.Info,) (*types.Package, error) {
	return conf.Check(path, fset, files, info)
}

//go:linkname ConfigCheck types.sub_configcheck
//go:noescape
func ConfigCheck(conf *types.Config, path string, fset *token.FileSet, files []*ast.File, info *types.Info,) (*types.Package, error)

//go:linkname constname types.sub_constname
func constname(obj *types.Const,) string {
	return obj.Name()
}

//go:linkname ConstName types.sub_constname
//go:noescape
func ConstName(obj *types.Const,) string

//go:linkname funcpos types.sub_funcpos
func funcpos(obj *types.Func,) token.Pos {
	return obj.Pos()
}

//go:linkname FuncPos types.sub_funcpos
//go:noescape
func FuncPos(obj *types.Func,) token.Pos

//go:linkname interfaceunderlying types.sub_interfaceunderlying
func interfaceunderlying(t *types.Interface,) types.Type {
	return t.Underlying()
}

//go:linkname InterfaceUnderlying types.sub_interfaceunderlying
//go:noescape
func InterfaceUnderlying(t *types.Interface,) types.Type

//go:linkname namedstring types.sub_namedstring
func namedstring(t *types.Named,) string {
	return t.String()
}

//go:linkname NamedString types.sub_namedstring
//go:noescape
func NamedString(t *types.Named,) string

//go:linkname pkgnametype types.sub_pkgnametype
func pkgnametype(obj *types.PkgName,) types.Type {
	return obj.Type()
}

//go:linkname PkgNameType types.sub_pkgnametype
//go:noescape
func PkgNameType(obj *types.PkgName,) types.Type

//go:linkname Identical go/types.Identical
//go:noescape
func Identical(x, y types.Type,) bool

//go:linkname arraylen types.sub_arraylen
func arraylen(a *types.Array,) int64 {
	return a.Len()
}

//go:linkname ArrayLen types.sub_arraylen
//go:noescape
func ArrayLen(a *types.Array,) int64

//go:linkname methodsetstring types.sub_methodsetstring
func methodsetstring(s *types.MethodSet,) string {
	return s.String()
}

//go:linkname MethodSetString types.sub_methodsetstring
//go:noescape
func MethodSetString(s *types.MethodSet,) string

//go:linkname pointerelem types.sub_pointerelem
func pointerelem(p *types.Pointer,) types.Type {
	return p.Elem()
}

//go:linkname PointerElem types.sub_pointerelem
//go:noescape
func PointerElem(p *types.Pointer,) types.Type

//go:linkname varanonymous types.sub_varanonymous
func varanonymous(obj *types.Var,) bool {
	return obj.Anonymous()
}

//go:linkname VarAnonymous types.sub_varanonymous
//go:noescape
func VarAnonymous(obj *types.Var,) bool
