// This file has automatically been generated on Wed Feb 26 02:09:55 +05 2020.
// DO NOT EDIT.
package build

import (
	"go/build"
	_ "unsafe"
)

//go:linkname packageiscommand build.sub_packageiscommand
func packageiscommand(p *build.Package) bool {
	return p.IsCommand()
}

//go:linkname PackageIsCommand build.sub_packageiscommand
//go:noescape
func PackageIsCommand(p *build.Package) bool

//go:linkname ArchChar go/build.ArchChar
//go:noescape
func ArchChar(goarch string) (string, error)

//go:linkname IsLocalImport go/build.IsLocalImport
//go:noescape
func IsLocalImport(path string) bool

//go:linkname contextimport build.sub_contextimport
func contextimport(ctxt *build.Context, path string, srcDir string, mode build.ImportMode) (*build.Package, error) {
	return ctxt.Import(path, srcDir, mode)
}

//go:linkname ContextImport build.sub_contextimport
//go:noescape
func ContextImport(ctxt *build.Context, path string, srcDir string, mode build.ImportMode) (*build.Package, error)

//go:linkname contextimportdir build.sub_contextimportdir
func contextimportdir(ctxt *build.Context, dir string, mode build.ImportMode) (*build.Package, error) {
	return ctxt.ImportDir(dir, mode)
}

//go:linkname ContextImportDir build.sub_contextimportdir
//go:noescape
func ContextImportDir(ctxt *build.Context, dir string, mode build.ImportMode) (*build.Package, error)

//go:linkname contextmatchfile build.sub_contextmatchfile
func contextmatchfile(ctxt *build.Context, dir, name string) (bool, error) {
	return ctxt.MatchFile(dir, name)
}

//go:linkname ContextMatchFile build.sub_contextmatchfile
//go:noescape
func ContextMatchFile(ctxt *build.Context, dir, name string) (bool, error)

//go:linkname Import go/build.Import
//go:noescape
func Import(path, srcDir string, mode build.ImportMode) (*build.Package, error)

//go:linkname contextsrcdirs build.sub_contextsrcdirs
func contextsrcdirs(ctxt *build.Context) []string {
	return ctxt.SrcDirs()
}

//go:linkname ContextSrcDirs build.sub_contextsrcdirs
//go:noescape
func ContextSrcDirs(ctxt *build.Context) []string

//go:linkname multiplepackageerrorerror build.sub_multiplepackageerrorerror
func multiplepackageerrorerror(e *build.MultiplePackageError) string {
	return e.Error()
}

//go:linkname MultiplePackageErrorError build.sub_multiplepackageerrorerror
//go:noescape
func MultiplePackageErrorError(e *build.MultiplePackageError) string

//go:linkname nogoerrorerror build.sub_nogoerrorerror
func nogoerrorerror(e *build.NoGoError) string {
	return e.Error()
}

//go:linkname NoGoErrorError build.sub_nogoerrorerror
//go:noescape
func NoGoErrorError(e *build.NoGoError) string

//go:linkname ImportDir go/build.ImportDir
//go:noescape
func ImportDir(dir string, mode build.ImportMode) (*build.Package, error)
