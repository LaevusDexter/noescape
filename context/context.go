// This file has automatically been generated on Wed Feb 26 02:09:43 +05 2020.
// DO NOT EDIT.
package context

import (
	"context"
	"time"
	_ "unsafe"
)

//go:linkname Background context.Background
//go:noescape
func Background() context.Context

//go:linkname TODO context.TODO
//go:noescape
func TODO() context.Context

//go:linkname WithValue context.WithValue
//go:noescape
func WithValue(parent context.Context, key, val interface{}) context.Context

//go:linkname WithCancel context.WithCancel
//go:noescape
func WithCancel(parent context.Context) (context.Context, context.CancelFunc)

//go:linkname WithDeadline context.WithDeadline
//go:noescape
func WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc)

//go:linkname WithTimeout context.WithTimeout
//go:noescape
func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
