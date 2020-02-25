// This file has automatically been generated on Wed Feb 26 02:10:02 +05 2020.
// DO NOT EDIT.
package syslog

import (
	"log"
	"log/syslog"
	_ "unsafe"
)

//go:linkname NewLogger log/syslog.NewLogger
//go:noescape
func NewLogger(p syslog.Priority, logFlag int) (*log.Logger, error)

//go:linkname New log/syslog.New
//go:noescape
func New(priority syslog.Priority, tag string) (*syslog.Writer, error)

//go:linkname writerclose syslog.sub_writerclose
func writerclose(w *syslog.Writer) error {
	return w.Close()
}

//go:linkname WriterClose syslog.sub_writerclose
//go:noescape
func WriterClose(w *syslog.Writer) error

//go:linkname writerdebug syslog.sub_writerdebug
func writerdebug(w *syslog.Writer, m string) error {
	return w.Debug(m)
}

//go:linkname WriterDebug syslog.sub_writerdebug
//go:noescape
func WriterDebug(w *syslog.Writer, m string) error

//go:linkname writererr syslog.sub_writererr
func writererr(w *syslog.Writer, m string) error {
	return w.Err(m)
}

//go:linkname WriterErr syslog.sub_writererr
//go:noescape
func WriterErr(w *syslog.Writer, m string) error

//go:linkname Dial log/syslog.Dial
//go:noescape
func Dial(network, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error)

//go:linkname writeralert syslog.sub_writeralert
func writeralert(w *syslog.Writer, m string) error {
	return w.Alert(m)
}

//go:linkname WriterAlert syslog.sub_writeralert
//go:noescape
func WriterAlert(w *syslog.Writer, m string) error

//go:linkname writercrit syslog.sub_writercrit
func writercrit(w *syslog.Writer, m string) error {
	return w.Crit(m)
}

//go:linkname WriterCrit syslog.sub_writercrit
//go:noescape
func WriterCrit(w *syslog.Writer, m string) error

//go:linkname writeremerg syslog.sub_writeremerg
func writeremerg(w *syslog.Writer, m string) error {
	return w.Emerg(m)
}

//go:linkname WriterEmerg syslog.sub_writeremerg
//go:noescape
func WriterEmerg(w *syslog.Writer, m string) error

//go:linkname writerinfo syslog.sub_writerinfo
func writerinfo(w *syslog.Writer, m string) error {
	return w.Info(m)
}

//go:linkname WriterInfo syslog.sub_writerinfo
//go:noescape
func WriterInfo(w *syslog.Writer, m string) error

//go:linkname writernotice syslog.sub_writernotice
func writernotice(w *syslog.Writer, m string) error {
	return w.Notice(m)
}

//go:linkname WriterNotice syslog.sub_writernotice
//go:noescape
func WriterNotice(w *syslog.Writer, m string) error

//go:linkname writerwarning syslog.sub_writerwarning
func writerwarning(w *syslog.Writer, m string) error {
	return w.Warning(m)
}

//go:linkname WriterWarning syslog.sub_writerwarning
//go:noescape
func WriterWarning(w *syslog.Writer, m string) error

//go:linkname writerwrite syslog.sub_writerwrite
func writerwrite(w *syslog.Writer, b []byte) (int, error) {
	return w.Write(b)
}

//go:linkname WriterWrite syslog.sub_writerwrite
//go:noescape
func WriterWrite(w *syslog.Writer, b []byte) (int, error)
