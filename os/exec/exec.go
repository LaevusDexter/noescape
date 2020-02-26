// This file has automatically been generated on Wed Feb 26 15:50:47 +05 2020.
// DO NOT EDIT.
package exec

import (
	"context"
	"io"
	"os/exec"
	_ "unsafe"
)

//go:linkname cmdstdoutpipe exec.sub_cmdstdoutpipe
func cmdstdoutpipe(c *exec.Cmd,) (io.ReadCloser, error) {
	return c.StdoutPipe()
}

//go:linkname CmdStdoutPipe exec.sub_cmdstdoutpipe
//go:noescape
func CmdStdoutPipe(c *exec.Cmd,) (io.ReadCloser, error)

//go:linkname cmdcombinedoutput exec.sub_cmdcombinedoutput
func cmdcombinedoutput(c *exec.Cmd,) ([]byte, error) {
	return c.CombinedOutput()
}

//go:linkname CmdCombinedOutput exec.sub_cmdcombinedoutput
//go:noescape
func CmdCombinedOutput(c *exec.Cmd,) ([]byte, error)

//go:linkname cmdrun exec.sub_cmdrun
func cmdrun(c *exec.Cmd,) error {
	return c.Run()
}

//go:linkname CmdRun exec.sub_cmdrun
//go:noescape
func CmdRun(c *exec.Cmd,) error

//go:linkname cmdstart exec.sub_cmdstart
func cmdstart(c *exec.Cmd,) error {
	return c.Start()
}

//go:linkname CmdStart exec.sub_cmdstart
//go:noescape
func CmdStart(c *exec.Cmd,) error

//go:linkname cmdstdinpipe exec.sub_cmdstdinpipe
func cmdstdinpipe(c *exec.Cmd,) (io.WriteCloser, error) {
	return c.StdinPipe()
}

//go:linkname CmdStdinPipe exec.sub_cmdstdinpipe
//go:noescape
func CmdStdinPipe(c *exec.Cmd,) (io.WriteCloser, error)

//go:linkname errorunwrap exec.sub_errorunwrap
func errorunwrap(e *exec.Error,) error {
	return e.Unwrap()
}

//go:linkname ErrorUnwrap exec.sub_errorunwrap
//go:noescape
func ErrorUnwrap(e *exec.Error,) error

//go:linkname Command os/exec.Command
//go:noescape
func Command(name string, arg ...string) *exec.Cmd

//go:linkname CommandContext os/exec.CommandContext
//go:noescape
func CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd

//go:linkname cmdoutput exec.sub_cmdoutput
func cmdoutput(c *exec.Cmd,) ([]byte, error) {
	return c.Output()
}

//go:linkname CmdOutput exec.sub_cmdoutput
//go:noescape
func CmdOutput(c *exec.Cmd,) ([]byte, error)

//go:linkname cmdstderrpipe exec.sub_cmdstderrpipe
func cmdstderrpipe(c *exec.Cmd,) (io.ReadCloser, error) {
	return c.StderrPipe()
}

//go:linkname CmdStderrPipe exec.sub_cmdstderrpipe
//go:noescape
func CmdStderrPipe(c *exec.Cmd,) (io.ReadCloser, error)

//go:linkname exiterrorerror exec.sub_exiterrorerror
func exiterrorerror(e *exec.ExitError,) string {
	return e.Error()
}

//go:linkname ExitErrorError exec.sub_exiterrorerror
//go:noescape
func ExitErrorError(e *exec.ExitError,) string

//go:linkname LookPath os/exec.LookPath
//go:noescape
func LookPath(file string) (string, error)

//go:linkname cmdstring exec.sub_cmdstring
func cmdstring(c *exec.Cmd,) string {
	return c.String()
}

//go:linkname CmdString exec.sub_cmdstring
//go:noescape
func CmdString(c *exec.Cmd,) string

//go:linkname cmdwait exec.sub_cmdwait
func cmdwait(c *exec.Cmd,) error {
	return c.Wait()
}

//go:linkname CmdWait exec.sub_cmdwait
//go:noescape
func CmdWait(c *exec.Cmd,) error

//go:linkname errorerror exec.sub_errorerror
func errorerror(e *exec.Error,) string {
	return e.Error()
}

//go:linkname ErrorError exec.sub_errorerror
//go:noescape
func ErrorError(e *exec.Error,) string
