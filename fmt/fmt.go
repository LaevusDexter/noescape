// This file has automatically been generated on Wed Feb 26 02:09:54 +05 2020.
// DO NOT EDIT.
package fmt

import (
	_ "fmt"
	"io"
	_ "unsafe"
)

//go:linkname Fprintf fmt.Fprintf
//go:noescape
func Fprintf(w io.Writer, format string, a ...interface{}) (int, error)

//go:linkname Printf fmt.Printf
//go:noescape
func Printf(format string, a ...interface{}) (int, error)

//go:linkname Println fmt.Println
//go:noescape
func Println(a ...interface{}) (int, error)

//go:linkname Sscanf fmt.Sscanf
//go:noescape
func Sscanf(str string, format string, a ...interface{}) (int, error)

//go:linkname Fprintln fmt.Fprintln
//go:noescape
func Fprintln(w io.Writer, a ...interface{}) (int, error)

//go:linkname Fscan fmt.Fscan
//go:noescape
func Fscan(r io.Reader, a ...interface{}) (int, error)

//go:linkname Fscanf fmt.Fscanf
//go:noescape
func Fscanf(r io.Reader, format string, a ...interface{}) (int, error)

//go:linkname Sscanln fmt.Sscanln
//go:noescape
func Sscanln(str string, a ...interface{}) (int, error)

//go:linkname Errorf fmt.Errorf
//go:noescape
func Errorf(format string, a ...interface{}) error

//go:linkname Fprint fmt.Fprint
//go:noescape
func Fprint(w io.Writer, a ...interface{}) (int, error)

//go:linkname Fscanln fmt.Fscanln
//go:noescape
func Fscanln(r io.Reader, a ...interface{}) (int, error)

//go:linkname Print fmt.Print
//go:noescape
func Print(a ...interface{}) (int, error)

//go:linkname Scan fmt.Scan
//go:noescape
func Scan(a ...interface{}) (int, error)

//go:linkname Sprintf fmt.Sprintf
//go:noescape
func Sprintf(format string, a ...interface{}) string

//go:linkname Scanf fmt.Scanf
//go:noescape
func Scanf(format string, a ...interface{}) (int, error)

//go:linkname Scanln fmt.Scanln
//go:noescape
func Scanln(a ...interface{}) (int, error)

//go:linkname Sprint fmt.Sprint
//go:noescape
func Sprint(a ...interface{}) string

//go:linkname Sprintln fmt.Sprintln
//go:noescape
func Sprintln(a ...interface{}) string

//go:linkname Sscan fmt.Sscan
//go:noescape
func Sscan(str string, a ...interface{}) (int, error)
