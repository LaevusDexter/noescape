// This file has automatically been generated on Wed Feb 26 02:10:09 +05 2020.
// DO NOT EDIT.
package os

import (
	"os"
	"syscall"
	"time"
	_ "unsafe"
)

//go:linkname processstateexitcode os.sub_processstateexitcode
func processstateexitcode(p *os.ProcessState) int {
	return p.ExitCode()
}

//go:linkname ProcessStateExitCode os.sub_processstateexitcode
//go:noescape
func ProcessStateExitCode(p *os.ProcessState) int

//go:linkname Getpagesize os.Getpagesize
//go:noescape
func Getpagesize() int

//go:linkname Getpid os.Getpid
//go:noescape
func Getpid() int

//go:linkname Truncate os.Truncate
//go:noescape
func Truncate(name string, size int64) error

//go:linkname filesync os.sub_filesync
func filesync(f *os.File) error {
	return f.Sync()
}

//go:linkname FileSync os.sub_filesync
//go:noescape
func FileSync(f *os.File) error

//go:linkname filewriteat os.sub_filewriteat
func filewriteat(f *os.File, b []byte, off int64) (int, error) {
	return f.WriteAt(b, off)
}

//go:linkname FileWriteAt os.sub_filewriteat
//go:noescape
func FileWriteAt(f *os.File, b []byte, off int64) (int, error)

//go:linkname Stat os.Stat
//go:noescape
func Stat(name string) (os.FileInfo, error)

//go:linkname FindProcess os.FindProcess
//go:noescape
func FindProcess(pid int) (*os.Process, error)

//go:linkname processstatepid os.sub_processstatepid
func processstatepid(p *os.ProcessState) int {
	return p.Pid()
}

//go:linkname ProcessStatePid os.sub_processstatepid
//go:noescape
func ProcessStatePid(p *os.ProcessState) int

//go:linkname Getuid os.Getuid
//go:noescape
func Getuid() int

//go:linkname NewFile os.NewFile
//go:noescape
func NewFile(fd uintptr, name string) *os.File

//go:linkname filechown os.sub_filechown
func filechown(f *os.File, uid, gid int) error {
	return f.Chown(uid, gid)
}

//go:linkname FileChown os.sub_filechown
//go:noescape
func FileChown(f *os.File, uid, gid int) error

//go:linkname filename os.sub_filename
func filename(f *os.File) string {
	return f.Name()
}

//go:linkname FileName os.sub_filename
//go:noescape
func FileName(f *os.File) string

//go:linkname filesyscallconn os.sub_filesyscallconn
func filesyscallconn(f *os.File) (syscall.RawConn, error) {
	return f.SyscallConn()
}

//go:linkname FileSyscallConn os.sub_filesyscallconn
//go:noescape
func FileSyscallConn(f *os.File) (syscall.RawConn, error)

//go:linkname processstatesystemtime os.sub_processstatesystemtime
func processstatesystemtime(p *os.ProcessState) time.Duration {
	return p.SystemTime()
}

//go:linkname ProcessStateSystemTime os.sub_processstatesystemtime
//go:noescape
func ProcessStateSystemTime(p *os.ProcessState) time.Duration

//go:linkname syscallerrorunwrap os.sub_syscallerrorunwrap
func syscallerrorunwrap(e *os.SyscallError) error {
	return e.Unwrap()
}

//go:linkname SyscallErrorUnwrap os.sub_syscallerrorunwrap
//go:noescape
func SyscallErrorUnwrap(e *os.SyscallError) error

//go:linkname Executable os.Executable
//go:noescape
func Executable() (string, error)

//go:linkname ExpandEnv os.ExpandEnv
//go:noescape
func ExpandEnv(s string) string

//go:linkname Getenv os.Getenv
//go:noescape
func Getenv(key string) string

//go:linkname filesetreaddeadline os.sub_filesetreaddeadline
func filesetreaddeadline(f *os.File, t time.Time) error {
	return f.SetReadDeadline(t)
}

//go:linkname FileSetReadDeadline os.sub_filesetreaddeadline
//go:noescape
func FileSetReadDeadline(f *os.File, t time.Time) error

//go:linkname filewritestring os.sub_filewritestring
func filewritestring(f *os.File, s string) (int, error) {
	return f.WriteString(s)
}

//go:linkname FileWriteString os.sub_filewritestring
//go:noescape
func FileWriteString(f *os.File, s string) (int, error)

//go:linkname processsignal os.sub_processsignal
func processsignal(p *os.Process, sig os.Signal) error {
	return p.Signal(sig)
}

//go:linkname ProcessSignal os.sub_processsignal
//go:noescape
func ProcessSignal(p *os.Process, sig os.Signal) error

//go:linkname processstatesys os.sub_processstatesys
func processstatesys(p *os.ProcessState) interface{} {
	return p.Sys()
}

//go:linkname ProcessStateSys os.sub_processstatesys
//go:noescape
func ProcessStateSys(p *os.ProcessState) interface{}

//go:linkname Create os.Create
//go:noescape
func Create(name string) (*os.File, error)

//go:linkname Getegid os.Getegid
//go:noescape
func Getegid() int

//go:linkname Getppid os.Getppid
//go:noescape
func Getppid() int

//go:linkname Link os.Link
//go:noescape
func Link(oldname, newname string) error

//go:linkname patherrortimeout os.sub_patherrortimeout
func patherrortimeout(e *os.PathError) bool {
	return e.Timeout()
}

//go:linkname PathErrorTimeout os.sub_patherrortimeout
//go:noescape
func PathErrorTimeout(e *os.PathError) bool

//go:linkname processstatesysusage os.sub_processstatesysusage
func processstatesysusage(p *os.ProcessState) interface{} {
	return p.SysUsage()
}

//go:linkname ProcessStateSysUsage os.sub_processstatesysusage
//go:noescape
func ProcessStateSysUsage(p *os.ProcessState) interface{}

//go:linkname StartProcess os.StartProcess
//go:noescape
func StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)

//go:linkname IsTimeout os.IsTimeout
//go:noescape
func IsTimeout(err error) bool

//go:linkname Mkdir os.Mkdir
//go:noescape
func Mkdir(name string, perm os.FileMode) error

//go:linkname Readlink os.Readlink
//go:noescape
func Readlink(name string) (string, error)

//go:linkname Remove os.Remove
//go:noescape
func Remove(name string) error

//go:linkname Rename os.Rename
//go:noescape
func Rename(oldpath, newpath string) error

//go:linkname filereaddir os.sub_filereaddir
func filereaddir(f *os.File, n int) ([]os.FileInfo, error) {
	return f.Readdir(n)
}

//go:linkname FileReaddir os.sub_filereaddir
//go:noescape
func FileReaddir(f *os.File, n int) ([]os.FileInfo, error)

//go:linkname Lstat os.Lstat
//go:noescape
func Lstat(name string) (os.FileInfo, error)

//go:linkname Chmod os.Chmod
//go:noescape
func Chmod(name string, mode os.FileMode) error

//go:linkname fileread os.sub_fileread
func fileread(f *os.File, b []byte) (int, error) {
	return f.Read(b)
}

//go:linkname FileRead os.sub_fileread
//go:noescape
func FileRead(f *os.File, b []byte) (int, error)

//go:linkname filesetdeadline os.sub_filesetdeadline
func filesetdeadline(f *os.File, t time.Time) error {
	return f.SetDeadline(t)
}

//go:linkname FileSetDeadline os.sub_filesetdeadline
//go:noescape
func FileSetDeadline(f *os.File, t time.Time) error

//go:linkname filewrite os.sub_filewrite
func filewrite(f *os.File, b []byte) (int, error) {
	return f.Write(b)
}

//go:linkname FileWrite os.sub_filewrite
//go:noescape
func FileWrite(f *os.File, b []byte) (int, error)

//go:linkname processstatestring os.sub_processstatestring
func processstatestring(p *os.ProcessState) string {
	return p.String()
}

//go:linkname ProcessStateString os.sub_processstatestring
//go:noescape
func ProcessStateString(p *os.ProcessState) string

//go:linkname Hostname os.Hostname
//go:noescape
func Hostname() (string, error)

//go:linkname filemodeisdir os.sub_filemodeisdir
func filemodeisdir(m os.FileMode) bool {
	return m.IsDir()
}

//go:linkname FileModeIsDir os.sub_filemodeisdir
//go:noescape
func FileModeIsDir(m os.FileMode) bool

//go:linkname Chdir os.Chdir
//go:noescape
func Chdir(dir string) error

//go:linkname IsPermission os.IsPermission
//go:noescape
func IsPermission(err error) bool

//go:linkname filemodeisregular os.sub_filemodeisregular
func filemodeisregular(m os.FileMode) bool {
	return m.IsRegular()
}

//go:linkname FileModeIsRegular os.sub_filemodeisregular
//go:noescape
func FileModeIsRegular(m os.FileMode) bool

//go:linkname filemodestring os.sub_filemodestring
func filemodestring(m os.FileMode) string {
	return m.String()
}

//go:linkname FileModeString os.sub_filemodestring
//go:noescape
func FileModeString(m os.FileMode) string

//go:linkname syscallerrorerror os.sub_syscallerrorerror
func syscallerrorerror(e *os.SyscallError) string {
	return e.Error()
}

//go:linkname SyscallErrorError os.sub_syscallerrorerror
//go:noescape
func SyscallErrorError(e *os.SyscallError) string

//go:linkname Expand os.Expand
//go:noescape
func Expand(s string, mapping func(string) string) string

//go:linkname UserCacheDir os.UserCacheDir
//go:noescape
func UserCacheDir() (string, error)

//go:linkname UserHomeDir os.UserHomeDir
//go:noescape
func UserHomeDir() (string, error)

//go:linkname filechmod os.sub_filechmod
func filechmod(f *os.File, mode os.FileMode) error {
	return f.Chmod(mode)
}

//go:linkname FileChmod os.sub_filechmod
//go:noescape
func FileChmod(f *os.File, mode os.FileMode) error

//go:linkname linkerrorerror os.sub_linkerrorerror
func linkerrorerror(e *os.LinkError) string {
	return e.Error()
}

//go:linkname LinkErrorError os.sub_linkerrorerror
//go:noescape
func LinkErrorError(e *os.LinkError) string

//go:linkname patherrorunwrap os.sub_patherrorunwrap
func patherrorunwrap(e *os.PathError) error {
	return e.Unwrap()
}

//go:linkname PathErrorUnwrap os.sub_patherrorunwrap
//go:noescape
func PathErrorUnwrap(e *os.PathError) error

//go:linkname processstatesuccess os.sub_processstatesuccess
func processstatesuccess(p *os.ProcessState) bool {
	return p.Success()
}

//go:linkname ProcessStateSuccess os.sub_processstatesuccess
//go:noescape
func ProcessStateSuccess(p *os.ProcessState) bool

//go:linkname Environ os.Environ
//go:noescape
func Environ() []string

//go:linkname LookupEnv os.LookupEnv
//go:noescape
func LookupEnv(key string) (string, bool)

//go:linkname NewSyscallError os.NewSyscallError
//go:noescape
func NewSyscallError(syscall string, err error) error

//go:linkname fileclose os.sub_fileclose
func fileclose(f *os.File) error {
	return f.Close()
}

//go:linkname FileClose os.sub_fileclose
//go:noescape
func FileClose(f *os.File) error

//go:linkname filefd os.sub_filefd
func filefd(f *os.File) uintptr {
	return f.Fd()
}

//go:linkname FileFd os.sub_filefd
//go:noescape
func FileFd(f *os.File) uintptr

//go:linkname processstateexited os.sub_processstateexited
func processstateexited(p *os.ProcessState) bool {
	return p.Exited()
}

//go:linkname ProcessStateExited os.sub_processstateexited
//go:noescape
func ProcessStateExited(p *os.ProcessState) bool

//go:linkname filereadat os.sub_filereadat
func filereadat(f *os.File, b []byte, off int64) (int, error) {
	return f.ReadAt(b, off)
}

//go:linkname FileReadAt os.sub_filereadat
//go:noescape
func FileReadAt(f *os.File, b []byte, off int64) (int, error)

//go:linkname Chown os.Chown
//go:noescape
func Chown(name string, uid, gid int) error

//go:linkname IsExist os.IsExist
//go:noescape
func IsExist(err error) bool

//go:linkname IsNotExist os.IsNotExist
//go:noescape
func IsNotExist(err error) bool

//go:linkname Lchown os.Lchown
//go:noescape
func Lchown(name string, uid, gid int) error

//go:linkname RemoveAll os.RemoveAll
//go:noescape
func RemoveAll(path string) error

//go:linkname Setenv os.Setenv
//go:noescape
func Setenv(key, value string) error

//go:linkname Symlink os.Symlink
//go:noescape
func Symlink(oldname, newname string) error

//go:linkname fileseek os.sub_fileseek
func fileseek(f *os.File, offset int64, whence int) (int64, error) {
	return f.Seek(offset, whence)
}

//go:linkname FileSeek os.sub_fileseek
//go:noescape
func FileSeek(f *os.File, offset int64, whence int) (int64, error)

//go:linkname processwait os.sub_processwait
func processwait(p *os.Process) (*os.ProcessState, error) {
	return p.Wait()
}

//go:linkname ProcessWait os.sub_processwait
//go:noescape
func ProcessWait(p *os.Process) (*os.ProcessState, error)

//go:linkname processkill os.sub_processkill
func processkill(p *os.Process) error {
	return p.Kill()
}

//go:linkname ProcessKill os.sub_processkill
//go:noescape
func ProcessKill(p *os.Process) error

//go:linkname Getgid os.Getgid
//go:noescape
func Getgid() int

//go:linkname Getwd os.Getwd
//go:noescape
func Getwd() (string, error)

//go:linkname IsPathSeparator os.IsPathSeparator
//go:noescape
func IsPathSeparator(c uint8) bool

//go:linkname UserConfigDir os.UserConfigDir
//go:noescape
func UserConfigDir() (string, error)

//go:linkname OpenFile os.OpenFile
//go:noescape
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)

//go:linkname filechdir os.sub_filechdir
func filechdir(f *os.File) error {
	return f.Chdir()
}

//go:linkname FileChdir os.sub_filechdir
//go:noescape
func FileChdir(f *os.File) error

//go:linkname filemodeperm os.sub_filemodeperm
func filemodeperm(m os.FileMode) os.FileMode {
	return m.Perm()
}

//go:linkname FileModePerm os.sub_filemodeperm
//go:noescape
func FileModePerm(m os.FileMode) os.FileMode

//go:linkname processrelease os.sub_processrelease
func processrelease(p *os.Process) error {
	return p.Release()
}

//go:linkname ProcessRelease os.sub_processrelease
//go:noescape
func ProcessRelease(p *os.Process) error

//go:linkname patherrorerror os.sub_patherrorerror
func patherrorerror(e *os.PathError) string {
	return e.Error()
}

//go:linkname PathErrorError os.sub_patherrorerror
//go:noescape
func PathErrorError(e *os.PathError) string

//go:linkname Chtimes os.Chtimes
//go:noescape
func Chtimes(name string, atime time.Time, mtime time.Time) error

//go:linkname Getgroups os.Getgroups
//go:noescape
func Getgroups() ([]int, error)

//go:linkname MkdirAll os.MkdirAll
//go:noescape
func MkdirAll(path string, perm os.FileMode) error

//go:linkname Unsetenv os.Unsetenv
//go:noescape
func Unsetenv(key string) error

//go:linkname filesetwritedeadline os.sub_filesetwritedeadline
func filesetwritedeadline(f *os.File, t time.Time) error {
	return f.SetWriteDeadline(t)
}

//go:linkname FileSetWriteDeadline os.sub_filesetwritedeadline
//go:noescape
func FileSetWriteDeadline(f *os.File, t time.Time) error

//go:linkname filetruncate os.sub_filetruncate
func filetruncate(f *os.File, size int64) error {
	return f.Truncate(size)
}

//go:linkname FileTruncate os.sub_filetruncate
//go:noescape
func FileTruncate(f *os.File, size int64) error

//go:linkname linkerrorunwrap os.sub_linkerrorunwrap
func linkerrorunwrap(e *os.LinkError) error {
	return e.Unwrap()
}

//go:linkname LinkErrorUnwrap os.sub_linkerrorunwrap
//go:noescape
func LinkErrorUnwrap(e *os.LinkError) error

//go:linkname processstateusertime os.sub_processstateusertime
func processstateusertime(p *os.ProcessState) time.Duration {
	return p.UserTime()
}

//go:linkname ProcessStateUserTime os.sub_processstateusertime
//go:noescape
func ProcessStateUserTime(p *os.ProcessState) time.Duration

//go:linkname Geteuid os.Geteuid
//go:noescape
func Geteuid() int

//go:linkname Pipe os.Pipe
//go:noescape
func Pipe() (*os.File, *os.File, error)

//go:linkname SameFile os.SameFile
//go:noescape
func SameFile(fi1, fi2 os.FileInfo) bool

//go:linkname TempDir os.TempDir
//go:noescape
func TempDir() string

//go:linkname Open os.Open
//go:noescape
func Open(name string) (*os.File, error)

//go:linkname filereaddirnames os.sub_filereaddirnames
func filereaddirnames(f *os.File, n int) ([]string, error) {
	return f.Readdirnames(n)
}

//go:linkname FileReaddirnames os.sub_filereaddirnames
//go:noescape
func FileReaddirnames(f *os.File, n int) ([]string, error)

//go:linkname filestat os.sub_filestat
func filestat(f *os.File) (os.FileInfo, error) {
	return f.Stat()
}

//go:linkname FileStat os.sub_filestat
//go:noescape
func FileStat(f *os.File) (os.FileInfo, error)

//go:linkname syscallerrortimeout os.sub_syscallerrortimeout
func syscallerrortimeout(e *os.SyscallError) bool {
	return e.Timeout()
}

//go:linkname SyscallErrorTimeout os.sub_syscallerrortimeout
//go:noescape
func SyscallErrorTimeout(e *os.SyscallError) bool
