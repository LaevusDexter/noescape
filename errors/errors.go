// This file has automatically been generated on Wed Feb 26 02:09:54 +05 2020.
// DO NOT EDIT.
package errors

import (
	_ "errors"
	_ "unsafe"
)

//go:linkname As errors.As
//go:noescape
func As(err error, target interface{}) bool

//go:linkname Is errors.Is
//go:noescape
func Is(err, target error) bool

//go:linkname New errors.New
//go:noescape
func New(text string) error

//go:linkname Unwrap errors.Unwrap
//go:noescape
func Unwrap(err error) error
