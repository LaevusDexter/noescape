// This file has automatically been generated on Wed Feb 26 15:50:55 +05 2020.
// DO NOT EDIT.
package unicode

import (
	"unicode"
	_ "unsafe"
)

//go:linkname ToLower unicode.ToLower
//go:noescape
func ToLower(r rune) rune

//go:linkname specialcasetotitle unicode.sub_specialcasetotitle
func specialcasetotitle(special unicode.SpecialCase, r rune) rune {
	return special.ToTitle(r)
}

//go:linkname SpecialCaseToTitle unicode.sub_specialcasetotitle
//go:noescape
func SpecialCaseToTitle(special unicode.SpecialCase, r rune) rune

//go:linkname IsNumber unicode.IsNumber
//go:noescape
func IsNumber(r rune) bool

//go:linkname IsOneOf unicode.IsOneOf
//go:noescape
func IsOneOf(ranges []*unicode.RangeTable, r rune) bool

//go:linkname IsPunct unicode.IsPunct
//go:noescape
func IsPunct(r rune) bool

//go:linkname To unicode.To
//go:noescape
func To(_case int, r rune) rune

//go:linkname ToTitle unicode.ToTitle
//go:noescape
func ToTitle(r rune) rune

//go:linkname specialcasetoupper unicode.sub_specialcasetoupper
func specialcasetoupper(special unicode.SpecialCase, r rune) rune {
	return special.ToUpper(r)
}

//go:linkname SpecialCaseToUpper unicode.sub_specialcasetoupper
//go:noescape
func SpecialCaseToUpper(special unicode.SpecialCase, r rune) rune

//go:linkname IsControl unicode.IsControl
//go:noescape
func IsControl(r rune) bool

//go:linkname IsDigit unicode.IsDigit
//go:noescape
func IsDigit(r rune) bool

//go:linkname IsSpace unicode.IsSpace
//go:noescape
func IsSpace(r rune) bool

//go:linkname IsUpper unicode.IsUpper
//go:noescape
func IsUpper(r rune) bool

//go:linkname IsSymbol unicode.IsSymbol
//go:noescape
func IsSymbol(r rune) bool

//go:linkname IsTitle unicode.IsTitle
//go:noescape
func IsTitle(r rune) bool

//go:linkname SimpleFold unicode.SimpleFold
//go:noescape
func SimpleFold(r rune) rune

//go:linkname ToUpper unicode.ToUpper
//go:noescape
func ToUpper(r rune) rune

//go:linkname In unicode.In
//go:noescape
func In(r rune, ranges ...*unicode.RangeTable,) bool

//go:linkname IsLetter unicode.IsLetter
//go:noescape
func IsLetter(r rune) bool

//go:linkname IsLower unicode.IsLower
//go:noescape
func IsLower(r rune) bool

//go:linkname IsMark unicode.IsMark
//go:noescape
func IsMark(r rune) bool

//go:linkname specialcasetolower unicode.sub_specialcasetolower
func specialcasetolower(special unicode.SpecialCase, r rune) rune {
	return special.ToLower(r)
}

//go:linkname SpecialCaseToLower unicode.sub_specialcasetolower
//go:noescape
func SpecialCaseToLower(special unicode.SpecialCase, r rune) rune

//go:linkname Is unicode.Is
//go:noescape
func Is(rangeTab *unicode.RangeTable, r rune) bool

//go:linkname IsGraphic unicode.IsGraphic
//go:noescape
func IsGraphic(r rune) bool

//go:linkname IsPrint unicode.IsPrint
//go:noescape
func IsPrint(r rune) bool
