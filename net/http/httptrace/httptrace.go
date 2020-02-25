// This file has automatically been generated on Wed Feb 26 02:10:07 +05 2020.
// DO NOT EDIT.
package httptrace

import (
	"context"
	"net/http/httptrace"
	_ "unsafe"
)

//go:linkname WithClientTrace net/http/httptrace.WithClientTrace
//go:noescape
func WithClientTrace(ctx context.Context, trace *httptrace.ClientTrace) context.Context

//go:linkname ContextClientTrace net/http/httptrace.ContextClientTrace
//go:noescape
func ContextClientTrace(ctx context.Context) *httptrace.ClientTrace
