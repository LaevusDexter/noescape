// This file has automatically been generated on Wed Feb 26 15:50:37 +05 2020.
// DO NOT EDIT.
package html

import (
	_ "html"
	_ "unsafe"
)

//go:linkname EscapeString html.EscapeString
//go:noescape
func EscapeString(s string) string

//go:linkname UnescapeString html.UnescapeString
//go:noescape
func UnescapeString(s string) string
