// This file has automatically been generated on Wed Feb 26 02:10:14 +05 2020.
// DO NOT EDIT.
package strconv

import (
	"strconv"
	_ "unsafe"
)

//go:linkname AppendQuoteRuneToGraphic strconv.AppendQuoteRuneToGraphic
//go:noescape
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte

//go:linkname AppendQuoteToGraphic strconv.AppendQuoteToGraphic
//go:noescape
func AppendQuoteToGraphic(dst []byte, s string) []byte

//go:linkname FormatInt strconv.FormatInt
//go:noescape
func FormatInt(i int64, base int) string

//go:linkname ParseInt strconv.ParseInt
//go:noescape
func ParseInt(s string, base int, bitSize int) (int64, error)

//go:linkname QuoteRuneToGraphic strconv.QuoteRuneToGraphic
//go:noescape
func QuoteRuneToGraphic(r rune) string

//go:linkname numerrorerror strconv.sub_numerrorerror
func numerrorerror(e *strconv.NumError) string {
	return e.Error()
}

//go:linkname NumErrorError strconv.sub_numerrorerror
//go:noescape
func NumErrorError(e *strconv.NumError) string

//go:linkname AppendQuote strconv.AppendQuote
//go:noescape
func AppendQuote(dst []byte, s string) []byte

//go:linkname AppendQuoteRune strconv.AppendQuoteRune
//go:noescape
func AppendQuoteRune(dst []byte, r rune) []byte

//go:linkname AppendQuoteRuneToASCII strconv.AppendQuoteRuneToASCII
//go:noescape
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte

//go:linkname QuoteRuneToASCII strconv.QuoteRuneToASCII
//go:noescape
func QuoteRuneToASCII(r rune) string

//go:linkname UnquoteChar strconv.UnquoteChar
//go:noescape
func UnquoteChar(s string, quote byte) (rune, bool, string, error)

//go:linkname FormatFloat strconv.FormatFloat
//go:noescape
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

//go:linkname ParseFloat strconv.ParseFloat
//go:noescape
func ParseFloat(s string, bitSize int) (float64, error)

//go:linkname ParseUint strconv.ParseUint
//go:noescape
func ParseUint(s string, base int, bitSize int) (uint64, error)

//go:linkname QuoteRune strconv.QuoteRune
//go:noescape
func QuoteRune(r rune) string

//go:linkname AppendFloat strconv.AppendFloat
//go:noescape
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte

//go:linkname CanBackquote strconv.CanBackquote
//go:noescape
func CanBackquote(s string) bool

//go:linkname IsGraphic strconv.IsGraphic
//go:noescape
func IsGraphic(r rune) bool

//go:linkname ParseBool strconv.ParseBool
//go:noescape
func ParseBool(str string) (bool, error)

//go:linkname QuoteToASCII strconv.QuoteToASCII
//go:noescape
func QuoteToASCII(s string) string

//go:linkname AppendUint strconv.AppendUint
//go:noescape
func AppendUint(dst []byte, i uint64, base int) []byte

//go:linkname FormatBool strconv.FormatBool
//go:noescape
func FormatBool(b bool) string

//go:linkname IsPrint strconv.IsPrint
//go:noescape
func IsPrint(r rune) bool

//go:linkname Itoa strconv.Itoa
//go:noescape
func Itoa(i int) string

//go:linkname QuoteToGraphic strconv.QuoteToGraphic
//go:noescape
func QuoteToGraphic(s string) string

//go:linkname AppendBool strconv.AppendBool
//go:noescape
func AppendBool(dst []byte, b bool) []byte

//go:linkname AppendInt strconv.AppendInt
//go:noescape
func AppendInt(dst []byte, i int64, base int) []byte

//go:linkname FormatUint strconv.FormatUint
//go:noescape
func FormatUint(i uint64, base int) string

//go:linkname Quote strconv.Quote
//go:noescape
func Quote(s string) string

//go:linkname AppendQuoteToASCII strconv.AppendQuoteToASCII
//go:noescape
func AppendQuoteToASCII(dst []byte, s string) []byte

//go:linkname Atoi strconv.Atoi
//go:noescape
func Atoi(s string) (int, error)

//go:linkname Unquote strconv.Unquote
//go:noescape
func Unquote(s string) (string, error)

//go:linkname numerrorunwrap strconv.sub_numerrorunwrap
func numerrorunwrap(e *strconv.NumError) error {
	return e.Unwrap()
}

//go:linkname NumErrorUnwrap strconv.sub_numerrorunwrap
//go:noescape
func NumErrorUnwrap(e *strconv.NumError) error
