// This file has automatically been generated on Wed Feb 26 02:10:11 +05 2020.
// DO NOT EDIT.
package syntax

import (
	"regexp/syntax"
	_ "unsafe"
)

//go:linkname regexpmaxcap syntax.sub_regexpmaxcap
func regexpmaxcap(re *syntax.Regexp) int {
	return re.MaxCap()
}

//go:linkname RegexpMaxCap syntax.sub_regexpmaxcap
//go:noescape
func RegexpMaxCap(re *syntax.Regexp) int

//go:linkname errorcodestring syntax.sub_errorcodestring
func errorcodestring(e syntax.ErrorCode) string {
	return e.String()
}

//go:linkname ErrorCodeString syntax.sub_errorcodestring
//go:noescape
func ErrorCodeString(e syntax.ErrorCode) string

//go:linkname instmatchrunepos syntax.sub_instmatchrunepos
func instmatchrunepos(i *syntax.Inst, r rune) int {
	return i.MatchRunePos(r)
}

//go:linkname InstMatchRunePos syntax.sub_instmatchrunepos
//go:noescape
func InstMatchRunePos(i *syntax.Inst, r rune) int

//go:linkname progstartcond syntax.sub_progstartcond
func progstartcond(p *syntax.Prog) syntax.EmptyOp {
	return p.StartCond()
}

//go:linkname ProgStartCond syntax.sub_progstartcond
//go:noescape
func ProgStartCond(p *syntax.Prog) syntax.EmptyOp

//go:linkname progstring syntax.sub_progstring
func progstring(p *syntax.Prog) string {
	return p.String()
}

//go:linkname ProgString syntax.sub_progstring
//go:noescape
func ProgString(p *syntax.Prog) string

//go:linkname Parse regexp/syntax.Parse
//go:noescape
func Parse(s string, flags syntax.Flags) (*syntax.Regexp, error)

//go:linkname regexpcapnames syntax.sub_regexpcapnames
func regexpcapnames(re *syntax.Regexp) []string {
	return re.CapNames()
}

//go:linkname RegexpCapNames syntax.sub_regexpcapnames
//go:noescape
func RegexpCapNames(re *syntax.Regexp) []string

//go:linkname regexpsimplify syntax.sub_regexpsimplify
func regexpsimplify(re *syntax.Regexp) *syntax.Regexp {
	return re.Simplify()
}

//go:linkname RegexpSimplify syntax.sub_regexpsimplify
//go:noescape
func RegexpSimplify(re *syntax.Regexp) *syntax.Regexp

//go:linkname regexpstring syntax.sub_regexpstring
func regexpstring(re *syntax.Regexp) string {
	return re.String()
}

//go:linkname RegexpString syntax.sub_regexpstring
//go:noescape
func RegexpString(re *syntax.Regexp) string

//go:linkname EmptyOpContext regexp/syntax.EmptyOpContext
//go:noescape
func EmptyOpContext(r1, r2 rune) syntax.EmptyOp

//go:linkname instmatchemptywidth syntax.sub_instmatchemptywidth
func instmatchemptywidth(i *syntax.Inst, before rune, after rune) bool {
	return i.MatchEmptyWidth(before, after)
}

//go:linkname InstMatchEmptyWidth syntax.sub_instmatchemptywidth
//go:noescape
func InstMatchEmptyWidth(i *syntax.Inst, before rune, after rune) bool

//go:linkname inststring syntax.sub_inststring
func inststring(i *syntax.Inst) string {
	return i.String()
}

//go:linkname InstString syntax.sub_inststring
//go:noescape
func InstString(i *syntax.Inst) string

//go:linkname instopstring syntax.sub_instopstring
func instopstring(i syntax.InstOp) string {
	return i.String()
}

//go:linkname InstOpString syntax.sub_instopstring
//go:noescape
func InstOpString(i syntax.InstOp) string

//go:linkname opstring syntax.sub_opstring
func opstring(i syntax.Op) string {
	return i.String()
}

//go:linkname OpString syntax.sub_opstring
//go:noescape
func OpString(i syntax.Op) string

//go:linkname progprefix syntax.sub_progprefix
func progprefix(p *syntax.Prog) (string, bool) {
	return p.Prefix()
}

//go:linkname ProgPrefix syntax.sub_progprefix
//go:noescape
func ProgPrefix(p *syntax.Prog) (string, bool)

//go:linkname IsWordChar regexp/syntax.IsWordChar
//go:noescape
func IsWordChar(r rune) bool

//go:linkname Compile regexp/syntax.Compile
//go:noescape
func Compile(re *syntax.Regexp) (*syntax.Prog, error)

//go:linkname regexpequal syntax.sub_regexpequal
func regexpequal(x *syntax.Regexp, y *syntax.Regexp) bool {
	return x.Equal(y)
}

//go:linkname RegexpEqual syntax.sub_regexpequal
//go:noescape
func RegexpEqual(x *syntax.Regexp, y *syntax.Regexp) bool

//go:linkname errorerror syntax.sub_errorerror
func errorerror(e *syntax.Error) string {
	return e.Error()
}

//go:linkname ErrorError syntax.sub_errorerror
//go:noescape
func ErrorError(e *syntax.Error) string

//go:linkname instmatchrune syntax.sub_instmatchrune
func instmatchrune(i *syntax.Inst, r rune) bool {
	return i.MatchRune(r)
}

//go:linkname InstMatchRune syntax.sub_instmatchrune
//go:noescape
func InstMatchRune(i *syntax.Inst, r rune) bool
