// This file has automatically been generated on Wed Feb 26 15:50:40 +05 2020.
// DO NOT EDIT.
package log

import (
	"io"
	"log"
	_ "unsafe"
)

//go:linkname Prefix log.Prefix
//go:noescape
func Prefix() string

//go:linkname New log.New
//go:noescape
func New(out io.Writer, prefix string, flag int) *log.Logger

//go:linkname loggerprefix log.sub_loggerprefix
func loggerprefix(l *log.Logger,) string {
	return l.Prefix()
}

//go:linkname LoggerPrefix log.sub_loggerprefix
//go:noescape
func LoggerPrefix(l *log.Logger,) string

//go:linkname loggerwriter log.sub_loggerwriter
func loggerwriter(l *log.Logger,) io.Writer {
	return l.Writer()
}

//go:linkname LoggerWriter log.sub_loggerwriter
//go:noescape
func LoggerWriter(l *log.Logger,) io.Writer

//go:linkname loggeroutput log.sub_loggeroutput
func loggeroutput(l *log.Logger, calldepth int, s string) error {
	return l.Output(calldepth, s)
}

//go:linkname LoggerOutput log.sub_loggeroutput
//go:noescape
func LoggerOutput(l *log.Logger, calldepth int, s string) error

//go:linkname Flags log.Flags
//go:noescape
func Flags() int

//go:linkname Output log.Output
//go:noescape
func Output(calldepth int, s string) error

//go:linkname Writer log.Writer
//go:noescape
func Writer() io.Writer

//go:linkname loggerflags log.sub_loggerflags
func loggerflags(l *log.Logger,) int {
	return l.Flags()
}

//go:linkname LoggerFlags log.sub_loggerflags
//go:noescape
func LoggerFlags(l *log.Logger,) int
