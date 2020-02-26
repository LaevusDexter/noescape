// This file has automatically been generated on Wed Feb 26 15:50:45 +05 2020.
// DO NOT EDIT.
package pprof

import (
	"net/http"
	_ "net/http/pprof"
	_ "unsafe"
)

//go:linkname Handler net/http/pprof.Handler
//go:noescape
func Handler(name string) http.Handler
