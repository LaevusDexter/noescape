// This file has automatically been generated on Wed Feb 26 02:10:13 +05 2020.
// DO NOT EDIT.
package trace

import (
	"context"
	"io"
	"runtime/trace"
	_ "unsafe"
)

//go:linkname IsEnabled runtime/trace.IsEnabled
//go:noescape
func IsEnabled() bool

//go:linkname Start runtime/trace.Start
//go:noescape
func Start(w io.Writer) error

//go:linkname StartRegion runtime/trace.StartRegion
//go:noescape
func StartRegion(ctx context.Context, regionType string) *trace.Region

//go:linkname NewTask runtime/trace.NewTask
//go:noescape
func NewTask(pctx context.Context, taskType string) (context.Context, *trace.Task)
