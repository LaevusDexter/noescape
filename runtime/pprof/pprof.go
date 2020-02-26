// This file has automatically been generated on Wed Feb 26 15:50:50 +05 2020.
// DO NOT EDIT.
package pprof

import (
	"context"
	"io"
	"runtime/pprof"
	_ "unsafe"
)

//go:linkname WithLabels runtime/pprof.WithLabels
//go:noescape
func WithLabels(ctx context.Context, labels pprof.LabelSet,) context.Context

//go:linkname Profiles runtime/pprof.Profiles
//go:noescape
func Profiles() []*pprof.Profile

//go:linkname StartCPUProfile runtime/pprof.StartCPUProfile
//go:noescape
func StartCPUProfile(w io.Writer) error

//go:linkname Label runtime/pprof.Label
//go:noescape
func Label(ctx context.Context, key string) (string, bool)

//go:linkname WriteHeapProfile runtime/pprof.WriteHeapProfile
//go:noescape
func WriteHeapProfile(w io.Writer) error

//go:linkname Labels runtime/pprof.Labels
//go:noescape
func Labels(args ...string) pprof.LabelSet

//go:linkname Lookup runtime/pprof.Lookup
//go:noescape
func Lookup(name string) *pprof.Profile

//go:linkname NewProfile runtime/pprof.NewProfile
//go:noescape
func NewProfile(name string) *pprof.Profile

//go:linkname profilecount pprof.sub_profilecount
func profilecount(p *pprof.Profile,) int {
	return p.Count()
}

//go:linkname ProfileCount pprof.sub_profilecount
//go:noescape
func ProfileCount(p *pprof.Profile,) int

//go:linkname profilename pprof.sub_profilename
func profilename(p *pprof.Profile,) string {
	return p.Name()
}

//go:linkname ProfileName pprof.sub_profilename
//go:noescape
func ProfileName(p *pprof.Profile,) string

//go:linkname ForLabels runtime/pprof.ForLabels
//go:noescape
func ForLabels(ctx context.Context, f func(key, value string) bool)

//go:linkname profilewriteto pprof.sub_profilewriteto
func profilewriteto(p *pprof.Profile, w io.Writer, debug int) error {
	return p.WriteTo(w, debug)
}

//go:linkname ProfileWriteTo pprof.sub_profilewriteto
//go:noescape
func ProfileWriteTo(p *pprof.Profile, w io.Writer, debug int) error
