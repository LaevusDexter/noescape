// This file has automatically been generated on Wed Feb 26 15:50:31 +05 2020.
// DO NOT EDIT.
package flag

import (
	"flag"
	"io"
	"time"
	_ "unsafe"
)

//go:linkname UnquoteUsage flag.UnquoteUsage
//go:noescape
func UnquoteUsage(flag *flag.Flag,) (string, string)

//go:linkname NewFlagSet flag.NewFlagSet
//go:noescape
func NewFlagSet(name string, errorHandling flag.ErrorHandling,) *flag.FlagSet

//go:linkname flagsetparsed flag.sub_flagsetparsed
func flagsetparsed(f *flag.FlagSet,) bool {
	return f.Parsed()
}

//go:linkname FlagSetParsed flag.sub_flagsetparsed
//go:noescape
func FlagSetParsed(f *flag.FlagSet,) bool

//go:linkname NArg flag.NArg
//go:noescape
func NArg() int

//go:linkname flagsetnarg flag.sub_flagsetnarg
func flagsetnarg(f *flag.FlagSet,) int {
	return f.NArg()
}

//go:linkname FlagSetNArg flag.sub_flagsetnarg
//go:noescape
func FlagSetNArg(f *flag.FlagSet,) int

//go:linkname Arg flag.Arg
//go:noescape
func Arg(i int) string

//go:linkname flagsetparse flag.sub_flagsetparse
func flagsetparse(f *flag.FlagSet, arguments []string) error {
	return f.Parse(arguments)
}

//go:linkname FlagSetParse flag.sub_flagsetparse
//go:noescape
func FlagSetParse(f *flag.FlagSet, arguments []string) error

//go:linkname Parsed flag.Parsed
//go:noescape
func Parsed() bool

//go:linkname flagsetint flag.sub_flagsetint
func flagsetint(f *flag.FlagSet, name string, value int, usage string) *int {
	return f.Int(name, value, usage)
}

//go:linkname FlagSetInt flag.sub_flagsetint
//go:noescape
func FlagSetInt(f *flag.FlagSet, name string, value int, usage string) *int

//go:linkname flagsetlookup flag.sub_flagsetlookup
func flagsetlookup(f *flag.FlagSet, name string) *flag.Flag {
	return f.Lookup(name)
}

//go:linkname FlagSetLookup flag.sub_flagsetlookup
//go:noescape
func FlagSetLookup(f *flag.FlagSet, name string) *flag.Flag

//go:linkname flagsetuint flag.sub_flagsetuint
func flagsetuint(f *flag.FlagSet, name string, value uint, usage string) *uint {
	return f.Uint(name, value, usage)
}

//go:linkname FlagSetUint flag.sub_flagsetuint
//go:noescape
func FlagSetUint(f *flag.FlagSet, name string, value uint, usage string) *uint

//go:linkname Duration flag.Duration
//go:noescape
func Duration(name string, value time.Duration, usage string) *time.Duration

//go:linkname NFlag flag.NFlag
//go:noescape
func NFlag() int

//go:linkname Uint flag.Uint
//go:noescape
func Uint(name string, value uint, usage string) *uint

//go:linkname flagsetargs flag.sub_flagsetargs
func flagsetargs(f *flag.FlagSet,) []string {
	return f.Args()
}

//go:linkname FlagSetArgs flag.sub_flagsetargs
//go:noescape
func FlagSetArgs(f *flag.FlagSet,) []string

//go:linkname flagsetset flag.sub_flagsetset
func flagsetset(f *flag.FlagSet, name, value string) error {
	return f.Set(name, value)
}

//go:linkname FlagSetSet flag.sub_flagsetset
//go:noescape
func FlagSetSet(f *flag.FlagSet, name, value string) error

//go:linkname Args flag.Args
//go:noescape
func Args() []string

//go:linkname flagsetname flag.sub_flagsetname
func flagsetname(f *flag.FlagSet,) string {
	return f.Name()
}

//go:linkname FlagSetName flag.sub_flagsetname
//go:noescape
func FlagSetName(f *flag.FlagSet,) string

//go:linkname flagsetoutput flag.sub_flagsetoutput
func flagsetoutput(f *flag.FlagSet,) io.Writer {
	return f.Output()
}

//go:linkname FlagSetOutput flag.sub_flagsetoutput
//go:noescape
func FlagSetOutput(f *flag.FlagSet,) io.Writer

//go:linkname Int flag.Int
//go:noescape
func Int(name string, value int, usage string) *int

//go:linkname String flag.String
//go:noescape
func String(name string, value string, usage string) *string

//go:linkname Lookup flag.Lookup
//go:noescape
func Lookup(name string) *flag.Flag

//go:linkname flagsetbool flag.sub_flagsetbool
func flagsetbool(f *flag.FlagSet, name string, value bool, usage string) *bool {
	return f.Bool(name, value, usage)
}

//go:linkname FlagSetBool flag.sub_flagsetbool
//go:noescape
func FlagSetBool(f *flag.FlagSet, name string, value bool, usage string) *bool

//go:linkname flagsetduration flag.sub_flagsetduration
func flagsetduration(f *flag.FlagSet, name string, value time.Duration, usage string) *time.Duration {
	return f.Duration(name, value, usage)
}

//go:linkname FlagSetDuration flag.sub_flagsetduration
//go:noescape
func FlagSetDuration(f *flag.FlagSet, name string, value time.Duration, usage string) *time.Duration

//go:linkname Bool flag.Bool
//go:noescape
func Bool(name string, value bool, usage string) *bool

//go:linkname Set flag.Set
//go:noescape
func Set(name, value string) error

//go:linkname flagsetarg flag.sub_flagsetarg
func flagsetarg(f *flag.FlagSet, i int) string {
	return f.Arg(i)
}

//go:linkname FlagSetArg flag.sub_flagsetarg
//go:noescape
func FlagSetArg(f *flag.FlagSet, i int) string

//go:linkname flagseterrorhandling flag.sub_flagseterrorhandling
func flagseterrorhandling(f *flag.FlagSet,) flag.ErrorHandling {
	return f.ErrorHandling()
}

//go:linkname FlagSetErrorHandling flag.sub_flagseterrorhandling
//go:noescape
func FlagSetErrorHandling(f *flag.FlagSet,) flag.ErrorHandling

//go:linkname flagsetnflag flag.sub_flagsetnflag
func flagsetnflag(f *flag.FlagSet,) int {
	return f.NFlag()
}

//go:linkname FlagSetNFlag flag.sub_flagsetnflag
//go:noescape
func FlagSetNFlag(f *flag.FlagSet,) int

//go:linkname flagsetstring flag.sub_flagsetstring
func flagsetstring(f *flag.FlagSet, name string, value string, usage string) *string {
	return f.String(name, value, usage)
}

//go:linkname FlagSetString flag.sub_flagsetstring
//go:noescape
func FlagSetString(f *flag.FlagSet, name string, value string, usage string) *string
