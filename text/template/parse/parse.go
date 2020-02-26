// This file has automatically been generated on Wed Feb 26 15:50:55 +05 2020.
// DO NOT EDIT.
package parse

import (
	"text/template/parse"
	_ "unsafe"
)

//go:linkname boolnodecopy parse.sub_boolnodecopy
func boolnodecopy(b *parse.BoolNode,) parse.Node {
	return b.Copy()
}

//go:linkname BoolNodeCopy parse.sub_boolnodecopy
//go:noescape
func BoolNodeCopy(b *parse.BoolNode,) parse.Node

//go:linkname chainnodestring parse.sub_chainnodestring
func chainnodestring(c *parse.ChainNode,) string {
	return c.String()
}

//go:linkname ChainNodeString parse.sub_chainnodestring
//go:noescape
func ChainNodeString(c *parse.ChainNode,) string

//go:linkname listnodecopy parse.sub_listnodecopy
func listnodecopy(l *parse.ListNode,) parse.Node {
	return l.Copy()
}

//go:linkname ListNodeCopy parse.sub_listnodecopy
//go:noescape
func ListNodeCopy(l *parse.ListNode,) parse.Node

//go:linkname nilnodestring parse.sub_nilnodestring
func nilnodestring(n *parse.NilNode,) string {
	return n.String()
}

//go:linkname NilNodeString parse.sub_nilnodestring
//go:noescape
func NilNodeString(n *parse.NilNode,) string

//go:linkname pipenodecopy parse.sub_pipenodecopy
func pipenodecopy(p *parse.PipeNode,) parse.Node {
	return p.Copy()
}

//go:linkname PipeNodeCopy parse.sub_pipenodecopy
//go:noescape
func PipeNodeCopy(p *parse.PipeNode,) parse.Node

//go:linkname treeparse parse.sub_treeparse
func treeparse(t *parse.Tree, text, leftDelim, rightDelim string, treeSet map[string]*parse.Tree, funcs ...map[string]interface{}) (*parse.Tree, error) {
	return t.Parse(text, leftDelim, rightDelim, treeSet, funcs...)
}

//go:linkname TreeParse parse.sub_treeparse
//go:noescape
func TreeParse(t *parse.Tree, text, leftDelim, rightDelim string, treeSet map[string]*parse.Tree, funcs ...map[string]interface{}) (*parse.Tree, error)

//go:linkname variablenodestring parse.sub_variablenodestring
func variablenodestring(v *parse.VariableNode,) string {
	return v.String()
}

//go:linkname VariableNodeString parse.sub_variablenodestring
//go:noescape
func VariableNodeString(v *parse.VariableNode,) string

//go:linkname commandnodestring parse.sub_commandnodestring
func commandnodestring(c *parse.CommandNode,) string {
	return c.String()
}

//go:linkname CommandNodeString parse.sub_commandnodestring
//go:noescape
func CommandNodeString(c *parse.CommandNode,) string

//go:linkname NewIdentifier text/template/parse.NewIdentifier
//go:noescape
func NewIdentifier(ident string) *parse.IdentifierNode

//go:linkname ifnodecopy parse.sub_ifnodecopy
func ifnodecopy(i *parse.IfNode,) parse.Node {
	return i.Copy()
}

//go:linkname IfNodeCopy parse.sub_ifnodecopy
//go:noescape
func IfNodeCopy(i *parse.IfNode,) parse.Node

//go:linkname numbernodecopy parse.sub_numbernodecopy
func numbernodecopy(n *parse.NumberNode,) parse.Node {
	return n.Copy()
}

//go:linkname NumberNodeCopy parse.sub_numbernodecopy
//go:noescape
func NumberNodeCopy(n *parse.NumberNode,) parse.Node

//go:linkname rangenodecopy parse.sub_rangenodecopy
func rangenodecopy(r *parse.RangeNode,) parse.Node {
	return r.Copy()
}

//go:linkname RangeNodeCopy parse.sub_rangenodecopy
//go:noescape
func RangeNodeCopy(r *parse.RangeNode,) parse.Node

//go:linkname withnodecopy parse.sub_withnodecopy
func withnodecopy(w *parse.WithNode,) parse.Node {
	return w.Copy()
}

//go:linkname WithNodeCopy parse.sub_withnodecopy
//go:noescape
func WithNodeCopy(w *parse.WithNode,) parse.Node

//go:linkname branchnodecopy parse.sub_branchnodecopy
func branchnodecopy(b *parse.BranchNode,) parse.Node {
	return b.Copy()
}

//go:linkname BranchNodeCopy parse.sub_branchnodecopy
//go:noescape
func BranchNodeCopy(b *parse.BranchNode,) parse.Node

//go:linkname dotnodecopy parse.sub_dotnodecopy
func dotnodecopy(d *parse.DotNode,) parse.Node {
	return d.Copy()
}

//go:linkname DotNodeCopy parse.sub_dotnodecopy
//go:noescape
func DotNodeCopy(d *parse.DotNode,) parse.Node

//go:linkname dotnodetype parse.sub_dotnodetype
func dotnodetype(d *parse.DotNode,) parse.NodeType {
	return d.Type()
}

//go:linkname DotNodeType parse.sub_dotnodetype
//go:noescape
func DotNodeType(d *parse.DotNode,) parse.NodeType

//go:linkname fieldnodestring parse.sub_fieldnodestring
func fieldnodestring(f *parse.FieldNode,) string {
	return f.String()
}

//go:linkname FieldNodeString parse.sub_fieldnodestring
//go:noescape
func FieldNodeString(f *parse.FieldNode,) string

//go:linkname nilnodecopy parse.sub_nilnodecopy
func nilnodecopy(n *parse.NilNode,) parse.Node {
	return n.Copy()
}

//go:linkname NilNodeCopy parse.sub_nilnodecopy
//go:noescape
func NilNodeCopy(n *parse.NilNode,) parse.Node

//go:linkname nodetypetype parse.sub_nodetypetype
func nodetypetype(t parse.NodeType,) parse.NodeType {
	return t.Type()
}

//go:linkname NodeTypeType parse.sub_nodetypetype
//go:noescape
func NodeTypeType(t parse.NodeType,) parse.NodeType

//go:linkname boolnodestring parse.sub_boolnodestring
func boolnodestring(b *parse.BoolNode,) string {
	return b.String()
}

//go:linkname BoolNodeString parse.sub_boolnodestring
//go:noescape
func BoolNodeString(b *parse.BoolNode,) string

//go:linkname branchnodestring parse.sub_branchnodestring
func branchnodestring(b *parse.BranchNode,) string {
	return b.String()
}

//go:linkname BranchNodeString parse.sub_branchnodestring
//go:noescape
func BranchNodeString(b *parse.BranchNode,) string

//go:linkname commandnodecopy parse.sub_commandnodecopy
func commandnodecopy(c *parse.CommandNode,) parse.Node {
	return c.Copy()
}

//go:linkname CommandNodeCopy parse.sub_commandnodecopy
//go:noescape
func CommandNodeCopy(c *parse.CommandNode,) parse.Node

//go:linkname stringnodestring parse.sub_stringnodestring
func stringnodestring(s *parse.StringNode,) string {
	return s.String()
}

//go:linkname StringNodeString parse.sub_stringnodestring
//go:noescape
func StringNodeString(s *parse.StringNode,) string

//go:linkname templatenodestring parse.sub_templatenodestring
func templatenodestring(t *parse.TemplateNode,) string {
	return t.String()
}

//go:linkname TemplateNodeString parse.sub_templatenodestring
//go:noescape
func TemplateNodeString(t *parse.TemplateNode,) string

//go:linkname variablenodecopy parse.sub_variablenodecopy
func variablenodecopy(v *parse.VariableNode,) parse.Node {
	return v.Copy()
}

//go:linkname VariableNodeCopy parse.sub_variablenodecopy
//go:noescape
func VariableNodeCopy(v *parse.VariableNode,) parse.Node

//go:linkname actionnodestring parse.sub_actionnodestring
func actionnodestring(a *parse.ActionNode,) string {
	return a.String()
}

//go:linkname ActionNodeString parse.sub_actionnodestring
//go:noescape
func ActionNodeString(a *parse.ActionNode,) string

//go:linkname fieldnodecopy parse.sub_fieldnodecopy
func fieldnodecopy(f *parse.FieldNode,) parse.Node {
	return f.Copy()
}

//go:linkname FieldNodeCopy parse.sub_fieldnodecopy
//go:noescape
func FieldNodeCopy(f *parse.FieldNode,) parse.Node

//go:linkname identifiernodecopy parse.sub_identifiernodecopy
func identifiernodecopy(i *parse.IdentifierNode,) parse.Node {
	return i.Copy()
}

//go:linkname IdentifierNodeCopy parse.sub_identifiernodecopy
//go:noescape
func IdentifierNodeCopy(i *parse.IdentifierNode,) parse.Node

//go:linkname identifiernodesetpos parse.sub_identifiernodesetpos
func identifiernodesetpos(i *parse.IdentifierNode, pos parse.Pos,) *parse.IdentifierNode {
	return i.SetPos(pos)
}

//go:linkname IdentifierNodeSetPos parse.sub_identifiernodesetpos
//go:noescape
func IdentifierNodeSetPos(i *parse.IdentifierNode, pos parse.Pos,) *parse.IdentifierNode

//go:linkname identifiernodestring parse.sub_identifiernodestring
func identifiernodestring(i *parse.IdentifierNode,) string {
	return i.String()
}

//go:linkname IdentifierNodeString parse.sub_identifiernodestring
//go:noescape
func IdentifierNodeString(i *parse.IdentifierNode,) string

//go:linkname New text/template/parse.New
//go:noescape
func New(name string, funcs ...map[string]interface{}) *parse.Tree

//go:linkname treeerrorcontext parse.sub_treeerrorcontext
func treeerrorcontext(t *parse.Tree, n parse.Node,) (string, string) {
	return t.ErrorContext(n)
}

//go:linkname TreeErrorContext parse.sub_treeerrorcontext
//go:noescape
func TreeErrorContext(t *parse.Tree, n parse.Node,) (string, string)

//go:linkname chainnodecopy parse.sub_chainnodecopy
func chainnodecopy(c *parse.ChainNode,) parse.Node {
	return c.Copy()
}

//go:linkname ChainNodeCopy parse.sub_chainnodecopy
//go:noescape
func ChainNodeCopy(c *parse.ChainNode,) parse.Node

//go:linkname listnodecopylist parse.sub_listnodecopylist
func listnodecopylist(l *parse.ListNode,) *parse.ListNode {
	return l.CopyList()
}

//go:linkname ListNodeCopyList parse.sub_listnodecopylist
//go:noescape
func ListNodeCopyList(l *parse.ListNode,) *parse.ListNode

//go:linkname listnodestring parse.sub_listnodestring
func listnodestring(l *parse.ListNode,) string {
	return l.String()
}

//go:linkname ListNodeString parse.sub_listnodestring
//go:noescape
func ListNodeString(l *parse.ListNode,) string

//go:linkname pipenodecopypipe parse.sub_pipenodecopypipe
func pipenodecopypipe(p *parse.PipeNode,) *parse.PipeNode {
	return p.CopyPipe()
}

//go:linkname PipeNodeCopyPipe parse.sub_pipenodecopypipe
//go:noescape
func PipeNodeCopyPipe(p *parse.PipeNode,) *parse.PipeNode

//go:linkname templatenodecopy parse.sub_templatenodecopy
func templatenodecopy(t *parse.TemplateNode,) parse.Node {
	return t.Copy()
}

//go:linkname TemplateNodeCopy parse.sub_templatenodecopy
//go:noescape
func TemplateNodeCopy(t *parse.TemplateNode,) parse.Node

//go:linkname textnodecopy parse.sub_textnodecopy
func textnodecopy(t *parse.TextNode,) parse.Node {
	return t.Copy()
}

//go:linkname TextNodeCopy parse.sub_textnodecopy
//go:noescape
func TextNodeCopy(t *parse.TextNode,) parse.Node

//go:linkname treecopy parse.sub_treecopy
func treecopy(t *parse.Tree,) *parse.Tree {
	return t.Copy()
}

//go:linkname TreeCopy parse.sub_treecopy
//go:noescape
func TreeCopy(t *parse.Tree,) *parse.Tree

//go:linkname Parse text/template/parse.Parse
//go:noescape
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]interface{}) (map[string]*parse.Tree, error)

//go:linkname dotnodestring parse.sub_dotnodestring
func dotnodestring(d *parse.DotNode,) string {
	return d.String()
}

//go:linkname DotNodeString parse.sub_dotnodestring
//go:noescape
func DotNodeString(d *parse.DotNode,) string

//go:linkname nilnodetype parse.sub_nilnodetype
func nilnodetype(n *parse.NilNode,) parse.NodeType {
	return n.Type()
}

//go:linkname NilNodeType parse.sub_nilnodetype
//go:noescape
func NilNodeType(n *parse.NilNode,) parse.NodeType

//go:linkname numbernodestring parse.sub_numbernodestring
func numbernodestring(n *parse.NumberNode,) string {
	return n.String()
}

//go:linkname NumberNodeString parse.sub_numbernodestring
//go:noescape
func NumberNodeString(n *parse.NumberNode,) string

//go:linkname pipenodestring parse.sub_pipenodestring
func pipenodestring(p *parse.PipeNode,) string {
	return p.String()
}

//go:linkname PipeNodeString parse.sub_pipenodestring
//go:noescape
func PipeNodeString(p *parse.PipeNode,) string

//go:linkname stringnodecopy parse.sub_stringnodecopy
func stringnodecopy(s *parse.StringNode,) parse.Node {
	return s.Copy()
}

//go:linkname StringNodeCopy parse.sub_stringnodecopy
//go:noescape
func StringNodeCopy(s *parse.StringNode,) parse.Node

//go:linkname textnodestring parse.sub_textnodestring
func textnodestring(t *parse.TextNode,) string {
	return t.String()
}

//go:linkname TextNodeString parse.sub_textnodestring
//go:noescape
func TextNodeString(t *parse.TextNode,) string

//go:linkname IsEmptyTree text/template/parse.IsEmptyTree
//go:noescape
func IsEmptyTree(n parse.Node,) bool

//go:linkname actionnodecopy parse.sub_actionnodecopy
func actionnodecopy(a *parse.ActionNode,) parse.Node {
	return a.Copy()
}

//go:linkname ActionNodeCopy parse.sub_actionnodecopy
//go:noescape
func ActionNodeCopy(a *parse.ActionNode,) parse.Node

//go:linkname identifiernodesettree parse.sub_identifiernodesettree
func identifiernodesettree(i *parse.IdentifierNode, t *parse.Tree,) *parse.IdentifierNode {
	return i.SetTree(t)
}

//go:linkname IdentifierNodeSetTree parse.sub_identifiernodesettree
//go:noescape
func IdentifierNodeSetTree(i *parse.IdentifierNode, t *parse.Tree,) *parse.IdentifierNode

//go:linkname posposition parse.sub_posposition
func posposition(p parse.Pos,) parse.Pos {
	return p.Position()
}

//go:linkname PosPosition parse.sub_posposition
//go:noescape
func PosPosition(p parse.Pos,) parse.Pos
