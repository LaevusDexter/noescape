// This file has automatically been generated on Wed Feb 26 02:10:09 +05 2020.
// DO NOT EDIT.
package signal

import (
	"os"
	_ "os/signal"
	_ "unsafe"
)

//go:linkname Ignored os/signal.Ignored
//go:noescape
func Ignored(sig os.Signal) bool
