// This file has automatically been generated on Wed Feb 26 02:10:01 +05 2020.
// DO NOT EDIT.
package suffixarray

import (
	"index/suffixarray"
	"io"
	"regexp"
	_ "unsafe"
)

//go:linkname indexlookup suffixarray.sub_indexlookup
func indexlookup(x *suffixarray.Index, s []byte, n int) []int {
	return x.Lookup(s, n)
}

//go:linkname IndexLookup suffixarray.sub_indexlookup
//go:noescape
func IndexLookup(x *suffixarray.Index, s []byte, n int) []int

//go:linkname indexread suffixarray.sub_indexread
func indexread(x *suffixarray.Index, r io.Reader) error {
	return x.Read(r)
}

//go:linkname IndexRead suffixarray.sub_indexread
//go:noescape
func IndexRead(x *suffixarray.Index, r io.Reader) error

//go:linkname indexwrite suffixarray.sub_indexwrite
func indexwrite(x *suffixarray.Index, w io.Writer) error {
	return x.Write(w)
}

//go:linkname IndexWrite suffixarray.sub_indexwrite
//go:noescape
func IndexWrite(x *suffixarray.Index, w io.Writer) error

//go:linkname New index/suffixarray.New
//go:noescape
func New(data []byte) *suffixarray.Index

//go:linkname indexbytes suffixarray.sub_indexbytes
func indexbytes(x *suffixarray.Index) []byte {
	return x.Bytes()
}

//go:linkname IndexBytes suffixarray.sub_indexbytes
//go:noescape
func IndexBytes(x *suffixarray.Index) []byte

//go:linkname indexfindallindex suffixarray.sub_indexfindallindex
func indexfindallindex(x *suffixarray.Index, r *regexp.Regexp, n int) [][]int {
	return x.FindAllIndex(r, n)
}

//go:linkname IndexFindAllIndex suffixarray.sub_indexfindallindex
//go:noescape
func IndexFindAllIndex(x *suffixarray.Index, r *regexp.Regexp, n int) [][]int
