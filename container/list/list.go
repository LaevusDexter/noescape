// This file has automatically been generated on Wed Feb 26 02:09:42 +05 2020.
// DO NOT EDIT.
package list

import (
	"container/list"
	_ "unsafe"
)

//go:linkname elementnext list.sub_elementnext
func elementnext(e *list.Element) *list.Element {
	return e.Next()
}

//go:linkname ElementNext list.sub_elementnext
//go:noescape
func ElementNext(e *list.Element) *list.Element

//go:linkname elementprev list.sub_elementprev
func elementprev(e *list.Element) *list.Element {
	return e.Prev()
}

//go:linkname ElementPrev list.sub_elementprev
//go:noescape
func ElementPrev(e *list.Element) *list.Element

//go:linkname New container/list.New
//go:noescape
func New() *list.List

//go:linkname listback list.sub_listback
func listback(l *list.List) *list.Element {
	return l.Back()
}

//go:linkname ListBack list.sub_listback
//go:noescape
func ListBack(l *list.List) *list.Element

//go:linkname listinit list.sub_listinit
func listinit(l *list.List) *list.List {
	return l.Init()
}

//go:linkname ListInit list.sub_listinit
//go:noescape
func ListInit(l *list.List) *list.List

//go:linkname listinsertafter list.sub_listinsertafter
func listinsertafter(l *list.List, v interface{}, mark *list.Element) *list.Element {
	return l.InsertAfter(v, mark)
}

//go:linkname ListInsertAfter list.sub_listinsertafter
//go:noescape
func ListInsertAfter(l *list.List, v interface{}, mark *list.Element) *list.Element

//go:linkname listpushback list.sub_listpushback
func listpushback(l *list.List, v interface{}) *list.Element {
	return l.PushBack(v)
}

//go:linkname ListPushBack list.sub_listpushback
//go:noescape
func ListPushBack(l *list.List, v interface{}) *list.Element

//go:linkname listpushfront list.sub_listpushfront
func listpushfront(l *list.List, v interface{}) *list.Element {
	return l.PushFront(v)
}

//go:linkname ListPushFront list.sub_listpushfront
//go:noescape
func ListPushFront(l *list.List, v interface{}) *list.Element

//go:linkname listremove list.sub_listremove
func listremove(l *list.List, e *list.Element) interface{} {
	return l.Remove(e)
}

//go:linkname ListRemove list.sub_listremove
//go:noescape
func ListRemove(l *list.List, e *list.Element) interface{}

//go:linkname listfront list.sub_listfront
func listfront(l *list.List) *list.Element {
	return l.Front()
}

//go:linkname ListFront list.sub_listfront
//go:noescape
func ListFront(l *list.List) *list.Element

//go:linkname listinsertbefore list.sub_listinsertbefore
func listinsertbefore(l *list.List, v interface{}, mark *list.Element) *list.Element {
	return l.InsertBefore(v, mark)
}

//go:linkname ListInsertBefore list.sub_listinsertbefore
//go:noescape
func ListInsertBefore(l *list.List, v interface{}, mark *list.Element) *list.Element

//go:linkname listlen list.sub_listlen
func listlen(l *list.List) int {
	return l.Len()
}

//go:linkname ListLen list.sub_listlen
//go:noescape
func ListLen(l *list.List) int
