// This file has automatically been generated on Wed Feb 26 15:50:51 +05 2020.
// DO NOT EDIT.
package sort

import (
	"sort"
	_ "unsafe"
)

//go:linkname intslicelen sort.sub_intslicelen
func intslicelen(p sort.IntSlice,) int {
	return p.Len()
}

//go:linkname IntSliceLen sort.sub_intslicelen
//go:noescape
func IntSliceLen(p sort.IntSlice,) int

//go:linkname intslicesearch sort.sub_intslicesearch
func intslicesearch(p sort.IntSlice, x int) int {
	return p.Search(x)
}

//go:linkname IntSliceSearch sort.sub_intslicesearch
//go:noescape
func IntSliceSearch(p sort.IntSlice, x int) int

//go:linkname stringsliceless sort.sub_stringsliceless
func stringsliceless(p sort.StringSlice, i, j int) bool {
	return p.Less(i, j)
}

//go:linkname StringSliceLess sort.sub_stringsliceless
//go:noescape
func StringSliceLess(p sort.StringSlice, i, j int) bool

//go:linkname IntsAreSorted sort.IntsAreSorted
//go:noescape
func IntsAreSorted(a []int) bool

//go:linkname Search sort.Search
//go:noescape
func Search(n int, f func(int) bool) int

//go:linkname SearchStrings sort.SearchStrings
//go:noescape
func SearchStrings(a []string, x string) int

//go:linkname SliceIsSorted sort.SliceIsSorted
//go:noescape
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool

//go:linkname SliceStable sort.SliceStable
//go:noescape
func SliceStable(slice interface{}, less func(i, j int) bool)

//go:linkname intsliceless sort.sub_intsliceless
func intsliceless(p sort.IntSlice, i, j int) bool {
	return p.Less(i, j)
}

//go:linkname IntSliceLess sort.sub_intsliceless
//go:noescape
func IntSliceLess(p sort.IntSlice, i, j int) bool

//go:linkname stringslicesearch sort.sub_stringslicesearch
func stringslicesearch(p sort.StringSlice, x string) int {
	return p.Search(x)
}

//go:linkname StringSliceSearch sort.sub_stringslicesearch
//go:noescape
func StringSliceSearch(p sort.StringSlice, x string) int

//go:linkname SearchInts sort.SearchInts
//go:noescape
func SearchInts(a []int, x int) int

//go:linkname Slice sort.Slice
//go:noescape
func Slice(slice interface{}, less func(i, j int) bool)

//go:linkname StringsAreSorted sort.StringsAreSorted
//go:noescape
func StringsAreSorted(a []string) bool

//go:linkname float64slicelen sort.sub_float64slicelen
func float64slicelen(p sort.Float64Slice,) int {
	return p.Len()
}

//go:linkname Float64SliceLen sort.sub_float64slicelen
//go:noescape
func Float64SliceLen(p sort.Float64Slice,) int

//go:linkname float64sliceless sort.sub_float64sliceless
func float64sliceless(p sort.Float64Slice, i, j int) bool {
	return p.Less(i, j)
}

//go:linkname Float64SliceLess sort.sub_float64sliceless
//go:noescape
func Float64SliceLess(p sort.Float64Slice, i, j int) bool

//go:linkname float64slicesearch sort.sub_float64slicesearch
func float64slicesearch(p sort.Float64Slice, x float64) int {
	return p.Search(x)
}

//go:linkname Float64SliceSearch sort.sub_float64slicesearch
//go:noescape
func Float64SliceSearch(p sort.Float64Slice, x float64) int

//go:linkname Reverse sort.Reverse
//go:noescape
func Reverse(data sort.Interface,) sort.Interface

//go:linkname IsSorted sort.IsSorted
//go:noescape
func IsSorted(data sort.Interface,) bool

//go:linkname stringslicelen sort.sub_stringslicelen
func stringslicelen(p sort.StringSlice,) int {
	return p.Len()
}

//go:linkname StringSliceLen sort.sub_stringslicelen
//go:noescape
func StringSliceLen(p sort.StringSlice,) int
