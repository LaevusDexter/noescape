// This file has automatically been generated on Wed Feb 26 02:10:16 +05 2020.
// DO NOT EDIT.
package tabwriter

import (
	"io"
	"text/tabwriter"
	_ "unsafe"
)

//go:linkname NewWriter text/tabwriter.NewWriter
//go:noescape
func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *tabwriter.Writer

//go:linkname writerflush tabwriter.sub_writerflush
func writerflush(b *tabwriter.Writer) error {
	return b.Flush()
}

//go:linkname WriterFlush tabwriter.sub_writerflush
//go:noescape
func WriterFlush(b *tabwriter.Writer) error

//go:linkname writerinit tabwriter.sub_writerinit
func writerinit(b *tabwriter.Writer, output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *tabwriter.Writer {
	return b.Init(output, minwidth, tabwidth, padding, padchar, flags)
}

//go:linkname WriterInit tabwriter.sub_writerinit
//go:noescape
func WriterInit(b *tabwriter.Writer, output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *tabwriter.Writer

//go:linkname writerwrite tabwriter.sub_writerwrite
func writerwrite(b *tabwriter.Writer, buf []byte) (int, error) {
	return b.Write(buf)
}

//go:linkname WriterWrite tabwriter.sub_writerwrite
//go:noescape
func WriterWrite(b *tabwriter.Writer, buf []byte) (int, error)
