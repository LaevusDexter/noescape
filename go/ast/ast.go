// This file has automatically been generated on Wed Feb 26 02:09:55 +05 2020.
// DO NOT EDIT.
package ast

import (
	"go/ast"
	"go/token"
	"io"
	"reflect"
	_ "unsafe"
)

//go:linkname deferstmtend ast.sub_deferstmtend
func deferstmtend(s *ast.DeferStmt) token.Pos {
	return s.End()
}

//go:linkname DeferStmtEnd ast.sub_deferstmtend
//go:noescape
func DeferStmtEnd(s *ast.DeferStmt) token.Pos

//go:linkname badexprpos ast.sub_badexprpos
func badexprpos(x *ast.BadExpr) token.Pos {
	return x.Pos()
}

//go:linkname BadExprPos ast.sub_badexprpos
//go:noescape
func BadExprPos(x *ast.BadExpr) token.Pos

//go:linkname blockstmtend ast.sub_blockstmtend
func blockstmtend(s *ast.BlockStmt) token.Pos {
	return s.End()
}

//go:linkname BlockStmtEnd ast.sub_blockstmtend
//go:noescape
func BlockStmtEnd(s *ast.BlockStmt) token.Pos

//go:linkname commentend ast.sub_commentend
func commentend(c *ast.Comment) token.Pos {
	return c.End()
}

//go:linkname CommentEnd ast.sub_commentend
//go:noescape
func CommentEnd(c *ast.Comment) token.Pos

//go:linkname arraytypeend ast.sub_arraytypeend
func arraytypeend(x *ast.ArrayType) token.Pos {
	return x.End()
}

//go:linkname ArrayTypeEnd ast.sub_arraytypeend
//go:noescape
func ArrayTypeEnd(x *ast.ArrayType) token.Pos

//go:linkname sendstmtend ast.sub_sendstmtend
func sendstmtend(s *ast.SendStmt) token.Pos {
	return s.End()
}

//go:linkname SendStmtEnd ast.sub_sendstmtend
//go:noescape
func SendStmtEnd(s *ast.SendStmt) token.Pos

//go:linkname switchstmtend ast.sub_switchstmtend
func switchstmtend(s *ast.SwitchStmt) token.Pos {
	return s.End()
}

//go:linkname SwitchStmtEnd ast.sub_switchstmtend
//go:noescape
func SwitchStmtEnd(s *ast.SwitchStmt) token.Pos

//go:linkname chantypepos ast.sub_chantypepos
func chantypepos(x *ast.ChanType) token.Pos {
	return x.Pos()
}

//go:linkname ChanTypePos ast.sub_chantypepos
//go:noescape
func ChanTypePos(x *ast.ChanType) token.Pos

//go:linkname funclitend ast.sub_funclitend
func funclitend(x *ast.FuncLit) token.Pos {
	return x.End()
}

//go:linkname FuncLitEnd ast.sub_funclitend
//go:noescape
func FuncLitEnd(x *ast.FuncLit) token.Pos

//go:linkname maptypepos ast.sub_maptypepos
func maptypepos(x *ast.MapType) token.Pos {
	return x.Pos()
}

//go:linkname MapTypePos ast.sub_maptypepos
//go:noescape
func MapTypePos(x *ast.MapType) token.Pos

//go:linkname FilterFile go/ast.FilterFile
//go:noescape
func FilterFile(src *ast.File, f ast.Filter) bool

//go:linkname typeswitchstmtpos ast.sub_typeswitchstmtpos
func typeswitchstmtpos(s *ast.TypeSwitchStmt) token.Pos {
	return s.Pos()
}

//go:linkname TypeSwitchStmtPos ast.sub_typeswitchstmtpos
//go:noescape
func TypeSwitchStmtPos(s *ast.TypeSwitchStmt) token.Pos

//go:linkname ellipsisend ast.sub_ellipsisend
func ellipsisend(x *ast.Ellipsis) token.Pos {
	return x.End()
}

//go:linkname EllipsisEnd ast.sub_ellipsisend
//go:noescape
func EllipsisEnd(x *ast.Ellipsis) token.Pos

//go:linkname gendeclpos ast.sub_gendeclpos
func gendeclpos(d *ast.GenDecl) token.Pos {
	return d.Pos()
}

//go:linkname GenDeclPos ast.sub_gendeclpos
//go:noescape
func GenDeclPos(d *ast.GenDecl) token.Pos

//go:linkname gostmtend ast.sub_gostmtend
func gostmtend(s *ast.GoStmt) token.Pos {
	return s.End()
}

//go:linkname GoStmtEnd ast.sub_gostmtend
//go:noescape
func GoStmtEnd(s *ast.GoStmt) token.Pos

//go:linkname indexexprend ast.sub_indexexprend
func indexexprend(x *ast.IndexExpr) token.Pos {
	return x.End()
}

//go:linkname IndexExprEnd ast.sub_indexexprend
//go:noescape
func IndexExprEnd(x *ast.IndexExpr) token.Pos

//go:linkname switchstmtpos ast.sub_switchstmtpos
func switchstmtpos(s *ast.SwitchStmt) token.Pos {
	return s.Pos()
}

//go:linkname SwitchStmtPos ast.sub_switchstmtpos
//go:noescape
func SwitchStmtPos(s *ast.SwitchStmt) token.Pos

//go:linkname FileExports go/ast.FileExports
//go:noescape
func FileExports(src *ast.File) bool

//go:linkname basiclitend ast.sub_basiclitend
func basiclitend(x *ast.BasicLit) token.Pos {
	return x.End()
}

//go:linkname BasicLitEnd ast.sub_basiclitend
//go:noescape
func BasicLitEnd(x *ast.BasicLit) token.Pos

//go:linkname commentmapstring ast.sub_commentmapstring
func commentmapstring(cmap ast.CommentMap) string {
	return cmap.String()
}

//go:linkname CommentMapString ast.sub_commentmapstring
//go:noescape
func CommentMapString(cmap ast.CommentMap) string

//go:linkname NotNilFilter go/ast.NotNilFilter
//go:noescape
func NotNilFilter(_ string, v reflect.Value) bool

//go:linkname functypepos ast.sub_functypepos
func functypepos(x *ast.FuncType) token.Pos {
	return x.Pos()
}

//go:linkname FuncTypePos ast.sub_functypepos
//go:noescape
func FuncTypePos(x *ast.FuncType) token.Pos

//go:linkname gostmtpos ast.sub_gostmtpos
func gostmtpos(s *ast.GoStmt) token.Pos {
	return s.Pos()
}

//go:linkname GoStmtPos ast.sub_gostmtpos
//go:noescape
func GoStmtPos(s *ast.GoStmt) token.Pos

//go:linkname ifstmtend ast.sub_ifstmtend
func ifstmtend(s *ast.IfStmt) token.Pos {
	return s.End()
}

//go:linkname IfStmtEnd ast.sub_ifstmtend
//go:noescape
func IfStmtEnd(s *ast.IfStmt) token.Pos

//go:linkname typeswitchstmtend ast.sub_typeswitchstmtend
func typeswitchstmtend(s *ast.TypeSwitchStmt) token.Pos {
	return s.End()
}

//go:linkname TypeSwitchStmtEnd ast.sub_typeswitchstmtend
//go:noescape
func TypeSwitchStmtEnd(s *ast.TypeSwitchStmt) token.Pos

//go:linkname blockstmtpos ast.sub_blockstmtpos
func blockstmtpos(s *ast.BlockStmt) token.Pos {
	return s.Pos()
}

//go:linkname BlockStmtPos ast.sub_blockstmtpos
//go:noescape
func BlockStmtPos(s *ast.BlockStmt) token.Pos

//go:linkname keyvalueexprpos ast.sub_keyvalueexprpos
func keyvalueexprpos(x *ast.KeyValueExpr) token.Pos {
	return x.Pos()
}

//go:linkname KeyValueExprPos ast.sub_keyvalueexprpos
//go:noescape
func KeyValueExprPos(x *ast.KeyValueExpr) token.Pos

//go:linkname fieldlistend ast.sub_fieldlistend
func fieldlistend(f *ast.FieldList) token.Pos {
	return f.End()
}

//go:linkname FieldListEnd ast.sub_fieldlistend
//go:noescape
func FieldListEnd(f *ast.FieldList) token.Pos

//go:linkname fieldlistnumfields ast.sub_fieldlistnumfields
func fieldlistnumfields(f *ast.FieldList) int {
	return f.NumFields()
}

//go:linkname FieldListNumFields ast.sub_fieldlistnumfields
//go:noescape
func FieldListNumFields(f *ast.FieldList) int

//go:linkname identisexported ast.sub_identisexported
func identisexported(id *ast.Ident) bool {
	return id.IsExported()
}

//go:linkname IdentIsExported ast.sub_identisexported
//go:noescape
func IdentIsExported(id *ast.Ident) bool

//go:linkname parenexprend ast.sub_parenexprend
func parenexprend(x *ast.ParenExpr) token.Pos {
	return x.End()
}

//go:linkname ParenExprEnd ast.sub_parenexprend
//go:noescape
func ParenExprEnd(x *ast.ParenExpr) token.Pos

//go:linkname rangestmtpos ast.sub_rangestmtpos
func rangestmtpos(s *ast.RangeStmt) token.Pos {
	return s.Pos()
}

//go:linkname RangeStmtPos ast.sub_rangestmtpos
//go:noescape
func RangeStmtPos(s *ast.RangeStmt) token.Pos

//go:linkname assignstmtend ast.sub_assignstmtend
func assignstmtend(s *ast.AssignStmt) token.Pos {
	return s.End()
}

//go:linkname AssignStmtEnd ast.sub_assignstmtend
//go:noescape
func AssignStmtEnd(s *ast.AssignStmt) token.Pos

//go:linkname callexprpos ast.sub_callexprpos
func callexprpos(x *ast.CallExpr) token.Pos {
	return x.Pos()
}

//go:linkname CallExprPos ast.sub_callexprpos
//go:noescape
func CallExprPos(x *ast.CallExpr) token.Pos

//go:linkname compositelitpos ast.sub_compositelitpos
func compositelitpos(x *ast.CompositeLit) token.Pos {
	return x.Pos()
}

//go:linkname CompositeLitPos ast.sub_compositelitpos
//go:noescape
func CompositeLitPos(x *ast.CompositeLit) token.Pos

//go:linkname objectpos ast.sub_objectpos
func objectpos(obj *ast.Object) token.Pos {
	return obj.Pos()
}

//go:linkname ObjectPos ast.sub_objectpos
//go:noescape
func ObjectPos(obj *ast.Object) token.Pos

//go:linkname unaryexprpos ast.sub_unaryexprpos
func unaryexprpos(x *ast.UnaryExpr) token.Pos {
	return x.Pos()
}

//go:linkname UnaryExprPos ast.sub_unaryexprpos
//go:noescape
func UnaryExprPos(x *ast.UnaryExpr) token.Pos

//go:linkname typeassertexprend ast.sub_typeassertexprend
func typeassertexprend(x *ast.TypeAssertExpr) token.Pos {
	return x.End()
}

//go:linkname TypeAssertExprEnd ast.sub_typeassertexprend
//go:noescape
func TypeAssertExprEnd(x *ast.TypeAssertExpr) token.Pos

//go:linkname typeassertexprpos ast.sub_typeassertexprpos
func typeassertexprpos(x *ast.TypeAssertExpr) token.Pos {
	return x.Pos()
}

//go:linkname TypeAssertExprPos ast.sub_typeassertexprpos
//go:noescape
func TypeAssertExprPos(x *ast.TypeAssertExpr) token.Pos

//go:linkname valuespecend ast.sub_valuespecend
func valuespecend(s *ast.ValueSpec) token.Pos {
	return s.End()
}

//go:linkname ValueSpecEnd ast.sub_valuespecend
//go:noescape
func ValueSpecEnd(s *ast.ValueSpec) token.Pos

//go:linkname FilterDecl go/ast.FilterDecl
//go:noescape
func FilterDecl(decl ast.Decl, f ast.Filter) bool

//go:linkname funclitpos ast.sub_funclitpos
func funclitpos(x *ast.FuncLit) token.Pos {
	return x.Pos()
}

//go:linkname FuncLitPos ast.sub_funclitpos
//go:noescape
func FuncLitPos(x *ast.FuncLit) token.Pos

//go:linkname selectorexprpos ast.sub_selectorexprpos
func selectorexprpos(x *ast.SelectorExpr) token.Pos {
	return x.Pos()
}

//go:linkname SelectorExprPos ast.sub_selectorexprpos
//go:noescape
func SelectorExprPos(x *ast.SelectorExpr) token.Pos

//go:linkname identstring ast.sub_identstring
func identstring(id *ast.Ident) string {
	return id.String()
}

//go:linkname IdentString ast.sub_identstring
//go:noescape
func IdentString(id *ast.Ident) string

//go:linkname badstmtpos ast.sub_badstmtpos
func badstmtpos(s *ast.BadStmt) token.Pos {
	return s.Pos()
}

//go:linkname BadStmtPos ast.sub_badstmtpos
//go:noescape
func BadStmtPos(s *ast.BadStmt) token.Pos

//go:linkname binaryexprend ast.sub_binaryexprend
func binaryexprend(x *ast.BinaryExpr) token.Pos {
	return x.End()
}

//go:linkname BinaryExprEnd ast.sub_binaryexprend
//go:noescape
func BinaryExprEnd(x *ast.BinaryExpr) token.Pos

//go:linkname fileend ast.sub_fileend
func fileend(f *ast.File) token.Pos {
	return f.End()
}

//go:linkname FileEnd ast.sub_fileend
//go:noescape
func FileEnd(f *ast.File) token.Pos

//go:linkname commentmapcomments ast.sub_commentmapcomments
func commentmapcomments(cmap ast.CommentMap) []*ast.CommentGroup {
	return cmap.Comments()
}

//go:linkname CommentMapComments ast.sub_commentmapcomments
//go:noescape
func CommentMapComments(cmap ast.CommentMap) []*ast.CommentGroup

//go:linkname ellipsispos ast.sub_ellipsispos
func ellipsispos(x *ast.Ellipsis) token.Pos {
	return x.Pos()
}

//go:linkname EllipsisPos ast.sub_ellipsispos
//go:noescape
func EllipsisPos(x *ast.Ellipsis) token.Pos

//go:linkname typespecend ast.sub_typespecend
func typespecend(s *ast.TypeSpec) token.Pos {
	return s.End()
}

//go:linkname TypeSpecEnd ast.sub_typespecend
//go:noescape
func TypeSpecEnd(s *ast.TypeSpec) token.Pos

//go:linkname incdecstmtpos ast.sub_incdecstmtpos
func incdecstmtpos(s *ast.IncDecStmt) token.Pos {
	return s.Pos()
}

//go:linkname IncDecStmtPos ast.sub_incdecstmtpos
//go:noescape
func IncDecStmtPos(s *ast.IncDecStmt) token.Pos

//go:linkname NewObj go/ast.NewObj
//go:noescape
func NewObj(kind ast.ObjKind, name string) *ast.Object

//go:linkname packageend ast.sub_packageend
func packageend(p *ast.Package) token.Pos {
	return p.End()
}

//go:linkname PackageEnd ast.sub_packageend
//go:noescape
func PackageEnd(p *ast.Package) token.Pos

//go:linkname structtypeend ast.sub_structtypeend
func structtypeend(x *ast.StructType) token.Pos {
	return x.End()
}

//go:linkname StructTypeEnd ast.sub_structtypeend
//go:noescape
func StructTypeEnd(x *ast.StructType) token.Pos

//go:linkname incdecstmtend ast.sub_incdecstmtend
func incdecstmtend(s *ast.IncDecStmt) token.Pos {
	return s.End()
}

//go:linkname IncDecStmtEnd ast.sub_incdecstmtend
//go:noescape
func IncDecStmtEnd(s *ast.IncDecStmt) token.Pos

//go:linkname maptypeend ast.sub_maptypeend
func maptypeend(x *ast.MapType) token.Pos {
	return x.End()
}

//go:linkname MapTypeEnd ast.sub_maptypeend
//go:noescape
func MapTypeEnd(x *ast.MapType) token.Pos

//go:linkname sliceexprpos ast.sub_sliceexprpos
func sliceexprpos(x *ast.SliceExpr) token.Pos {
	return x.Pos()
}

//go:linkname SliceExprPos ast.sub_sliceexprpos
//go:noescape
func SliceExprPos(x *ast.SliceExpr) token.Pos

//go:linkname ifstmtpos ast.sub_ifstmtpos
func ifstmtpos(s *ast.IfStmt) token.Pos {
	return s.Pos()
}

//go:linkname IfStmtPos ast.sub_ifstmtpos
//go:noescape
func IfStmtPos(s *ast.IfStmt) token.Pos

//go:linkname indexexprpos ast.sub_indexexprpos
func indexexprpos(x *ast.IndexExpr) token.Pos {
	return x.Pos()
}

//go:linkname IndexExprPos ast.sub_indexexprpos
//go:noescape
func IndexExprPos(x *ast.IndexExpr) token.Pos

//go:linkname arraytypepos ast.sub_arraytypepos
func arraytypepos(x *ast.ArrayType) token.Pos {
	return x.Pos()
}

//go:linkname ArrayTypePos ast.sub_arraytypepos
//go:noescape
func ArrayTypePos(x *ast.ArrayType) token.Pos

//go:linkname commentpos ast.sub_commentpos
func commentpos(c *ast.Comment) token.Pos {
	return c.Pos()
}

//go:linkname CommentPos ast.sub_commentpos
//go:noescape
func CommentPos(c *ast.Comment) token.Pos

//go:linkname commentmapupdate ast.sub_commentmapupdate
func commentmapupdate(cmap ast.CommentMap, old, new ast.Node) ast.Node {
	return cmap.Update(old, new)
}

//go:linkname CommentMapUpdate ast.sub_commentmapupdate
//go:noescape
func CommentMapUpdate(cmap ast.CommentMap, old, new ast.Node) ast.Node

//go:linkname declstmtpos ast.sub_declstmtpos
func declstmtpos(s *ast.DeclStmt) token.Pos {
	return s.Pos()
}

//go:linkname DeclStmtPos ast.sub_declstmtpos
//go:noescape
func DeclStmtPos(s *ast.DeclStmt) token.Pos

//go:linkname labeledstmtend ast.sub_labeledstmtend
func labeledstmtend(s *ast.LabeledStmt) token.Pos {
	return s.End()
}

//go:linkname LabeledStmtEnd ast.sub_labeledstmtend
//go:noescape
func LabeledStmtEnd(s *ast.LabeledStmt) token.Pos

//go:linkname selectstmtend ast.sub_selectstmtend
func selectstmtend(s *ast.SelectStmt) token.Pos {
	return s.End()
}

//go:linkname SelectStmtEnd ast.sub_selectstmtend
//go:noescape
func SelectStmtEnd(s *ast.SelectStmt) token.Pos

//go:linkname sendstmtpos ast.sub_sendstmtpos
func sendstmtpos(s *ast.SendStmt) token.Pos {
	return s.Pos()
}

//go:linkname SendStmtPos ast.sub_sendstmtpos
//go:noescape
func SendStmtPos(s *ast.SendStmt) token.Pos

//go:linkname starexprend ast.sub_starexprend
func starexprend(x *ast.StarExpr) token.Pos {
	return x.End()
}

//go:linkname StarExprEnd ast.sub_starexprend
//go:noescape
func StarExprEnd(x *ast.StarExpr) token.Pos

//go:linkname branchstmtpos ast.sub_branchstmtpos
func branchstmtpos(s *ast.BranchStmt) token.Pos {
	return s.Pos()
}

//go:linkname BranchStmtPos ast.sub_branchstmtpos
//go:noescape
func BranchStmtPos(s *ast.BranchStmt) token.Pos

//go:linkname interfacetypepos ast.sub_interfacetypepos
func interfacetypepos(x *ast.InterfaceType) token.Pos {
	return x.Pos()
}

//go:linkname InterfaceTypePos ast.sub_interfacetypepos
//go:noescape
func InterfaceTypePos(x *ast.InterfaceType) token.Pos

//go:linkname scopelookup ast.sub_scopelookup
func scopelookup(s *ast.Scope, name string) *ast.Object {
	return s.Lookup(name)
}

//go:linkname ScopeLookup ast.sub_scopelookup
//go:noescape
func ScopeLookup(s *ast.Scope, name string) *ast.Object

//go:linkname emptystmtend ast.sub_emptystmtend
func emptystmtend(s *ast.EmptyStmt) token.Pos {
	return s.End()
}

//go:linkname EmptyStmtEnd ast.sub_emptystmtend
//go:noescape
func EmptyStmtEnd(s *ast.EmptyStmt) token.Pos

//go:linkname fieldpos ast.sub_fieldpos
func fieldpos(f *ast.Field) token.Pos {
	return f.Pos()
}

//go:linkname FieldPos ast.sub_fieldpos
//go:noescape
func FieldPos(f *ast.Field) token.Pos

//go:linkname forstmtend ast.sub_forstmtend
func forstmtend(s *ast.ForStmt) token.Pos {
	return s.End()
}

//go:linkname ForStmtEnd ast.sub_forstmtend
//go:noescape
func ForStmtEnd(s *ast.ForStmt) token.Pos

//go:linkname Print go/ast.Print
//go:noescape
func Print(fset *token.FileSet, x interface{}) error

//go:linkname commentgrouppos ast.sub_commentgrouppos
func commentgrouppos(g *ast.CommentGroup) token.Pos {
	return g.Pos()
}

//go:linkname CommentGroupPos ast.sub_commentgrouppos
//go:noescape
func CommentGroupPos(g *ast.CommentGroup) token.Pos

//go:linkname NewCommentMap go/ast.NewCommentMap
//go:noescape
func NewCommentMap(fset *token.FileSet, node ast.Node, comments []*ast.CommentGroup) ast.CommentMap

//go:linkname selectorexprend ast.sub_selectorexprend
func selectorexprend(x *ast.SelectorExpr) token.Pos {
	return x.End()
}

//go:linkname SelectorExprEnd ast.sub_selectorexprend
//go:noescape
func SelectorExprEnd(x *ast.SelectorExpr) token.Pos

//go:linkname callexprend ast.sub_callexprend
func callexprend(x *ast.CallExpr) token.Pos {
	return x.End()
}

//go:linkname CallExprEnd ast.sub_callexprend
//go:noescape
func CallExprEnd(x *ast.CallExpr) token.Pos

//go:linkname importspecpos ast.sub_importspecpos
func importspecpos(s *ast.ImportSpec) token.Pos {
	return s.Pos()
}

//go:linkname ImportSpecPos ast.sub_importspecpos
//go:noescape
func ImportSpecPos(s *ast.ImportSpec) token.Pos

//go:linkname returnstmtpos ast.sub_returnstmtpos
func returnstmtpos(s *ast.ReturnStmt) token.Pos {
	return s.Pos()
}

//go:linkname ReturnStmtPos ast.sub_returnstmtpos
//go:noescape
func ReturnStmtPos(s *ast.ReturnStmt) token.Pos

//go:linkname exprstmtpos ast.sub_exprstmtpos
func exprstmtpos(s *ast.ExprStmt) token.Pos {
	return s.Pos()
}

//go:linkname ExprStmtPos ast.sub_exprstmtpos
//go:noescape
func ExprStmtPos(s *ast.ExprStmt) token.Pos

//go:linkname filepos ast.sub_filepos
func filepos(f *ast.File) token.Pos {
	return f.Pos()
}

//go:linkname FilePos ast.sub_filepos
//go:noescape
func FilePos(f *ast.File) token.Pos

//go:linkname valuespecpos ast.sub_valuespecpos
func valuespecpos(s *ast.ValueSpec) token.Pos {
	return s.Pos()
}

//go:linkname ValueSpecPos ast.sub_valuespecpos
//go:noescape
func ValueSpecPos(s *ast.ValueSpec) token.Pos

//go:linkname baddeclend ast.sub_baddeclend
func baddeclend(d *ast.BadDecl) token.Pos {
	return d.End()
}

//go:linkname BadDeclEnd ast.sub_baddeclend
//go:noescape
func BadDeclEnd(d *ast.BadDecl) token.Pos

//go:linkname basiclitpos ast.sub_basiclitpos
func basiclitpos(x *ast.BasicLit) token.Pos {
	return x.Pos()
}

//go:linkname BasicLitPos ast.sub_basiclitpos
//go:noescape
func BasicLitPos(x *ast.BasicLit) token.Pos

//go:linkname binaryexprpos ast.sub_binaryexprpos
func binaryexprpos(x *ast.BinaryExpr) token.Pos {
	return x.Pos()
}

//go:linkname BinaryExprPos ast.sub_binaryexprpos
//go:noescape
func BinaryExprPos(x *ast.BinaryExpr) token.Pos

//go:linkname fieldlistpos ast.sub_fieldlistpos
func fieldlistpos(f *ast.FieldList) token.Pos {
	return f.Pos()
}

//go:linkname FieldListPos ast.sub_fieldlistpos
//go:noescape
func FieldListPos(f *ast.FieldList) token.Pos

//go:linkname MergePackageFiles go/ast.MergePackageFiles
//go:noescape
func MergePackageFiles(pkg *ast.Package, mode ast.MergeMode) *ast.File

//go:linkname funcdeclend ast.sub_funcdeclend
func funcdeclend(d *ast.FuncDecl) token.Pos {
	return d.End()
}

//go:linkname FuncDeclEnd ast.sub_funcdeclend
//go:noescape
func FuncDeclEnd(d *ast.FuncDecl) token.Pos

//go:linkname interfacetypeend ast.sub_interfacetypeend
func interfacetypeend(x *ast.InterfaceType) token.Pos {
	return x.End()
}

//go:linkname InterfaceTypeEnd ast.sub_interfacetypeend
//go:noescape
func InterfaceTypeEnd(x *ast.InterfaceType) token.Pos

//go:linkname scopestring ast.sub_scopestring
func scopestring(s *ast.Scope) string {
	return s.String()
}

//go:linkname ScopeString ast.sub_scopestring
//go:noescape
func ScopeString(s *ast.Scope) string

//go:linkname badexprend ast.sub_badexprend
func badexprend(x *ast.BadExpr) token.Pos {
	return x.End()
}

//go:linkname BadExprEnd ast.sub_badexprend
//go:noescape
func BadExprEnd(x *ast.BadExpr) token.Pos

//go:linkname caseclauseend ast.sub_caseclauseend
func caseclauseend(s *ast.CaseClause) token.Pos {
	return s.End()
}

//go:linkname CaseClauseEnd ast.sub_caseclauseend
//go:noescape
func CaseClauseEnd(s *ast.CaseClause) token.Pos

//go:linkname fieldend ast.sub_fieldend
func fieldend(f *ast.Field) token.Pos {
	return f.End()
}

//go:linkname FieldEnd ast.sub_fieldend
//go:noescape
func FieldEnd(f *ast.Field) token.Pos

//go:linkname NewIdent go/ast.NewIdent
//go:noescape
func NewIdent(name string) *ast.Ident

//go:linkname identpos ast.sub_identpos
func identpos(x *ast.Ident) token.Pos {
	return x.Pos()
}

//go:linkname IdentPos ast.sub_identpos
//go:noescape
func IdentPos(x *ast.Ident) token.Pos

//go:linkname rangestmtend ast.sub_rangestmtend
func rangestmtend(s *ast.RangeStmt) token.Pos {
	return s.End()
}

//go:linkname RangeStmtEnd ast.sub_rangestmtend
//go:noescape
func RangeStmtEnd(s *ast.RangeStmt) token.Pos

//go:linkname scopeinsert ast.sub_scopeinsert
func scopeinsert(s *ast.Scope, obj *ast.Object) *ast.Object {
	return s.Insert(obj)
}

//go:linkname ScopeInsert ast.sub_scopeinsert
//go:noescape
func ScopeInsert(s *ast.Scope, obj *ast.Object) *ast.Object

//go:linkname commentgrouptext ast.sub_commentgrouptext
func commentgrouptext(g *ast.CommentGroup) string {
	return g.Text()
}

//go:linkname CommentGroupText ast.sub_commentgrouptext
//go:noescape
func CommentGroupText(g *ast.CommentGroup) string

//go:linkname deferstmtpos ast.sub_deferstmtpos
func deferstmtpos(s *ast.DeferStmt) token.Pos {
	return s.Pos()
}

//go:linkname DeferStmtPos ast.sub_deferstmtpos
//go:noescape
func DeferStmtPos(s *ast.DeferStmt) token.Pos

//go:linkname funcdeclpos ast.sub_funcdeclpos
func funcdeclpos(d *ast.FuncDecl) token.Pos {
	return d.Pos()
}

//go:linkname FuncDeclPos ast.sub_funcdeclpos
//go:noescape
func FuncDeclPos(d *ast.FuncDecl) token.Pos

//go:linkname labeledstmtpos ast.sub_labeledstmtpos
func labeledstmtpos(s *ast.LabeledStmt) token.Pos {
	return s.Pos()
}

//go:linkname LabeledStmtPos ast.sub_labeledstmtpos
//go:noescape
func LabeledStmtPos(s *ast.LabeledStmt) token.Pos

//go:linkname packagepos ast.sub_packagepos
func packagepos(p *ast.Package) token.Pos {
	return p.Pos()
}

//go:linkname PackagePos ast.sub_packagepos
//go:noescape
func PackagePos(p *ast.Package) token.Pos

//go:linkname typespecpos ast.sub_typespecpos
func typespecpos(s *ast.TypeSpec) token.Pos {
	return s.Pos()
}

//go:linkname TypeSpecPos ast.sub_typespecpos
//go:noescape
func TypeSpecPos(s *ast.TypeSpec) token.Pos

//go:linkname Inspect go/ast.Inspect
//go:noescape
func Inspect(node ast.Node, f func(ast.Node) bool)

//go:linkname IsExported go/ast.IsExported
//go:noescape
func IsExported(name string) bool

//go:linkname compositelitend ast.sub_compositelitend
func compositelitend(x *ast.CompositeLit) token.Pos {
	return x.End()
}

//go:linkname CompositeLitEnd ast.sub_compositelitend
//go:noescape
func CompositeLitEnd(x *ast.CompositeLit) token.Pos

//go:linkname Fprint go/ast.Fprint
//go:noescape
func Fprint(w io.Writer, fset *token.FileSet, x interface{}, f ast.FieldFilter) error

//go:linkname baddeclpos ast.sub_baddeclpos
func baddeclpos(d *ast.BadDecl) token.Pos {
	return d.Pos()
}

//go:linkname BadDeclPos ast.sub_baddeclpos
//go:noescape
func BadDeclPos(d *ast.BadDecl) token.Pos

//go:linkname commentgroupend ast.sub_commentgroupend
func commentgroupend(g *ast.CommentGroup) token.Pos {
	return g.End()
}

//go:linkname CommentGroupEnd ast.sub_commentgroupend
//go:noescape
func CommentGroupEnd(g *ast.CommentGroup) token.Pos

//go:linkname identend ast.sub_identend
func identend(x *ast.Ident) token.Pos {
	return x.End()
}

//go:linkname IdentEnd ast.sub_identend
//go:noescape
func IdentEnd(x *ast.Ident) token.Pos

//go:linkname sliceexprend ast.sub_sliceexprend
func sliceexprend(x *ast.SliceExpr) token.Pos {
	return x.End()
}

//go:linkname SliceExprEnd ast.sub_sliceexprend
//go:noescape
func SliceExprEnd(x *ast.SliceExpr) token.Pos

//go:linkname structtypepos ast.sub_structtypepos
func structtypepos(x *ast.StructType) token.Pos {
	return x.Pos()
}

//go:linkname StructTypePos ast.sub_structtypepos
//go:noescape
func StructTypePos(x *ast.StructType) token.Pos

//go:linkname caseclausepos ast.sub_caseclausepos
func caseclausepos(s *ast.CaseClause) token.Pos {
	return s.Pos()
}

//go:linkname CaseClausePos ast.sub_caseclausepos
//go:noescape
func CaseClausePos(s *ast.CaseClause) token.Pos

//go:linkname emptystmtpos ast.sub_emptystmtpos
func emptystmtpos(s *ast.EmptyStmt) token.Pos {
	return s.Pos()
}

//go:linkname EmptyStmtPos ast.sub_emptystmtpos
//go:noescape
func EmptyStmtPos(s *ast.EmptyStmt) token.Pos

//go:linkname starexprpos ast.sub_starexprpos
func starexprpos(x *ast.StarExpr) token.Pos {
	return x.Pos()
}

//go:linkname StarExprPos ast.sub_starexprpos
//go:noescape
func StarExprPos(x *ast.StarExpr) token.Pos

//go:linkname NewPackage go/ast.NewPackage
//go:noescape
func NewPackage(fset *token.FileSet, files map[string]*ast.File, importer ast.Importer, universe *ast.Scope) (*ast.Package, error)

//go:linkname parenexprpos ast.sub_parenexprpos
func parenexprpos(x *ast.ParenExpr) token.Pos {
	return x.Pos()
}

//go:linkname ParenExprPos ast.sub_parenexprpos
//go:noescape
func ParenExprPos(x *ast.ParenExpr) token.Pos

//go:linkname badstmtend ast.sub_badstmtend
func badstmtend(s *ast.BadStmt) token.Pos {
	return s.End()
}

//go:linkname BadStmtEnd ast.sub_badstmtend
//go:noescape
func BadStmtEnd(s *ast.BadStmt) token.Pos

//go:linkname branchstmtend ast.sub_branchstmtend
func branchstmtend(s *ast.BranchStmt) token.Pos {
	return s.End()
}

//go:linkname BranchStmtEnd ast.sub_branchstmtend
//go:noescape
func BranchStmtEnd(s *ast.BranchStmt) token.Pos

//go:linkname keyvalueexprend ast.sub_keyvalueexprend
func keyvalueexprend(x *ast.KeyValueExpr) token.Pos {
	return x.End()
}

//go:linkname KeyValueExprEnd ast.sub_keyvalueexprend
//go:noescape
func KeyValueExprEnd(x *ast.KeyValueExpr) token.Pos

//go:linkname unaryexprend ast.sub_unaryexprend
func unaryexprend(x *ast.UnaryExpr) token.Pos {
	return x.End()
}

//go:linkname UnaryExprEnd ast.sub_unaryexprend
//go:noescape
func UnaryExprEnd(x *ast.UnaryExpr) token.Pos

//go:linkname assignstmtpos ast.sub_assignstmtpos
func assignstmtpos(s *ast.AssignStmt) token.Pos {
	return s.Pos()
}

//go:linkname AssignStmtPos ast.sub_assignstmtpos
//go:noescape
func AssignStmtPos(s *ast.AssignStmt) token.Pos

//go:linkname returnstmtend ast.sub_returnstmtend
func returnstmtend(s *ast.ReturnStmt) token.Pos {
	return s.End()
}

//go:linkname ReturnStmtEnd ast.sub_returnstmtend
//go:noescape
func ReturnStmtEnd(s *ast.ReturnStmt) token.Pos

//go:linkname NewScope go/ast.NewScope
//go:noescape
func NewScope(outer *ast.Scope) *ast.Scope

//go:linkname forstmtpos ast.sub_forstmtpos
func forstmtpos(s *ast.ForStmt) token.Pos {
	return s.Pos()
}

//go:linkname ForStmtPos ast.sub_forstmtpos
//go:noescape
func ForStmtPos(s *ast.ForStmt) token.Pos

//go:linkname gendeclend ast.sub_gendeclend
func gendeclend(d *ast.GenDecl) token.Pos {
	return d.End()
}

//go:linkname GenDeclEnd ast.sub_gendeclend
//go:noescape
func GenDeclEnd(d *ast.GenDecl) token.Pos

//go:linkname objkindstring ast.sub_objkindstring
func objkindstring(kind ast.ObjKind) string {
	return kind.String()
}

//go:linkname ObjKindString ast.sub_objkindstring
//go:noescape
func ObjKindString(kind ast.ObjKind) string

//go:linkname PackageExports go/ast.PackageExports
//go:noescape
func PackageExports(pkg *ast.Package) bool

//go:linkname commclausepos ast.sub_commclausepos
func commclausepos(s *ast.CommClause) token.Pos {
	return s.Pos()
}

//go:linkname CommClausePos ast.sub_commclausepos
//go:noescape
func CommClausePos(s *ast.CommClause) token.Pos

//go:linkname commentmapfilter ast.sub_commentmapfilter
func commentmapfilter(cmap ast.CommentMap, node ast.Node) ast.CommentMap {
	return cmap.Filter(node)
}

//go:linkname CommentMapFilter ast.sub_commentmapfilter
//go:noescape
func CommentMapFilter(cmap ast.CommentMap, node ast.Node) ast.CommentMap

//go:linkname declstmtend ast.sub_declstmtend
func declstmtend(s *ast.DeclStmt) token.Pos {
	return s.End()
}

//go:linkname DeclStmtEnd ast.sub_declstmtend
//go:noescape
func DeclStmtEnd(s *ast.DeclStmt) token.Pos

//go:linkname exprstmtend ast.sub_exprstmtend
func exprstmtend(s *ast.ExprStmt) token.Pos {
	return s.End()
}

//go:linkname ExprStmtEnd ast.sub_exprstmtend
//go:noescape
func ExprStmtEnd(s *ast.ExprStmt) token.Pos

//go:linkname functypeend ast.sub_functypeend
func functypeend(x *ast.FuncType) token.Pos {
	return x.End()
}

//go:linkname FuncTypeEnd ast.sub_functypeend
//go:noescape
func FuncTypeEnd(x *ast.FuncType) token.Pos

//go:linkname importspecend ast.sub_importspecend
func importspecend(s *ast.ImportSpec) token.Pos {
	return s.End()
}

//go:linkname ImportSpecEnd ast.sub_importspecend
//go:noescape
func ImportSpecEnd(s *ast.ImportSpec) token.Pos

//go:linkname selectstmtpos ast.sub_selectstmtpos
func selectstmtpos(s *ast.SelectStmt) token.Pos {
	return s.Pos()
}

//go:linkname SelectStmtPos ast.sub_selectstmtpos
//go:noescape
func SelectStmtPos(s *ast.SelectStmt) token.Pos

//go:linkname FilterPackage go/ast.FilterPackage
//go:noescape
func FilterPackage(pkg *ast.Package, f ast.Filter) bool

//go:linkname chantypeend ast.sub_chantypeend
func chantypeend(x *ast.ChanType) token.Pos {
	return x.End()
}

//go:linkname ChanTypeEnd ast.sub_chantypeend
//go:noescape
func ChanTypeEnd(x *ast.ChanType) token.Pos

//go:linkname commclauseend ast.sub_commclauseend
func commclauseend(s *ast.CommClause) token.Pos {
	return s.End()
}

//go:linkname CommClauseEnd ast.sub_commclauseend
//go:noescape
func CommClauseEnd(s *ast.CommClause) token.Pos
