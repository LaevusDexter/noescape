// This file has automatically been generated on Wed Feb 26 02:10:00 +05 2020.
// DO NOT EDIT.
package template

import (
	"html/template"
	"io"
	"text/template/parse"
	_ "unsafe"
)

//go:linkname JSEscaper html/template.JSEscaper
//go:noescape
func JSEscaper(args ...interface{}) string

//go:linkname Must html/template.Must
//go:noescape
func Must(t *template.Template, err error) *template.Template

//go:linkname templateparsefiles template.sub_templateparsefiles
func templateparsefiles(t *template.Template, filenames ...string) (*template.Template, error) {
	return t.ParseFiles(filenames)
}

//go:linkname TemplateParseFiles template.sub_templateparsefiles
//go:noescape
func TemplateParseFiles(t *template.Template, filenames ...string) (*template.Template, error)

//go:linkname HTMLEscaper html/template.HTMLEscaper
//go:noescape
func HTMLEscaper(args ...interface{}) string

//go:linkname JSEscapeString html/template.JSEscapeString
//go:noescape
func JSEscapeString(s string) string

//go:linkname templatetemplates template.sub_templatetemplates
func templatetemplates(t *template.Template) []*template.Template {
	return t.Templates()
}

//go:linkname TemplateTemplates template.sub_templatetemplates
//go:noescape
func TemplateTemplates(t *template.Template) []*template.Template

//go:linkname ParseFiles html/template.ParseFiles
//go:noescape
func ParseFiles(filenames ...string) (*template.Template, error)

//go:linkname templatenew template.sub_templatenew
func templatenew(t *template.Template, name string) *template.Template {
	return t.New(name)
}

//go:linkname TemplateNew template.sub_templatenew
//go:noescape
func TemplateNew(t *template.Template, name string) *template.Template

//go:linkname errorerror template.sub_errorerror
func errorerror(e *template.Error) string {
	return e.Error()
}

//go:linkname ErrorError template.sub_errorerror
//go:noescape
func ErrorError(e *template.Error) string

//go:linkname New html/template.New
//go:noescape
func New(name string) *template.Template

//go:linkname ParseGlob html/template.ParseGlob
//go:noescape
func ParseGlob(pattern string) (*template.Template, error)

//go:linkname templateaddparsetree template.sub_templateaddparsetree
func templateaddparsetree(t *template.Template, name string, tree *parse.Tree) (*template.Template, error) {
	return t.AddParseTree(name, tree)
}

//go:linkname TemplateAddParseTree template.sub_templateaddparsetree
//go:noescape
func TemplateAddParseTree(t *template.Template, name string, tree *parse.Tree) (*template.Template, error)

//go:linkname templatedelims template.sub_templatedelims
func templatedelims(t *template.Template, left, right string) *template.Template {
	return t.Delims(left, right)
}

//go:linkname TemplateDelims template.sub_templatedelims
//go:noescape
func TemplateDelims(t *template.Template, left, right string) *template.Template

//go:linkname templatefuncs template.sub_templatefuncs
func templatefuncs(t *template.Template, funcMap template.FuncMap) *template.Template {
	return t.Funcs(funcMap)
}

//go:linkname TemplateFuncs template.sub_templatefuncs
//go:noescape
func TemplateFuncs(t *template.Template, funcMap template.FuncMap) *template.Template

//go:linkname HTMLEscapeString html/template.HTMLEscapeString
//go:noescape
func HTMLEscapeString(s string) string

//go:linkname IsTrue html/template.IsTrue
//go:noescape
func IsTrue(val interface{}) bool

//go:linkname templateparseglob template.sub_templateparseglob
func templateparseglob(t *template.Template, pattern string) (*template.Template, error) {
	return t.ParseGlob(pattern)
}

//go:linkname TemplateParseGlob template.sub_templateparseglob
//go:noescape
func TemplateParseGlob(t *template.Template, pattern string) (*template.Template, error)

//go:linkname templatedefinedtemplates template.sub_templatedefinedtemplates
func templatedefinedtemplates(t *template.Template) string {
	return t.DefinedTemplates()
}

//go:linkname TemplateDefinedTemplates template.sub_templatedefinedtemplates
//go:noescape
func TemplateDefinedTemplates(t *template.Template) string

//go:linkname templateexecute template.sub_templateexecute
func templateexecute(t *template.Template, wr io.Writer, data interface{}) error {
	return t.Execute(wr, data)
}

//go:linkname TemplateExecute template.sub_templateexecute
//go:noescape
func TemplateExecute(t *template.Template, wr io.Writer, data interface{}) error

//go:linkname templateexecutetemplate template.sub_templateexecutetemplate
func templateexecutetemplate(t *template.Template, wr io.Writer, name string, data interface{}) error {
	return t.ExecuteTemplate(wr, name, data)
}

//go:linkname TemplateExecuteTemplate template.sub_templateexecutetemplate
//go:noescape
func TemplateExecuteTemplate(t *template.Template, wr io.Writer, name string, data interface{}) error

//go:linkname templatelookup template.sub_templatelookup
func templatelookup(t *template.Template, name string) *template.Template {
	return t.Lookup(name)
}

//go:linkname TemplateLookup template.sub_templatelookup
//go:noescape
func TemplateLookup(t *template.Template, name string) *template.Template

//go:linkname templatename template.sub_templatename
func templatename(t *template.Template) string {
	return t.Name()
}

//go:linkname TemplateName template.sub_templatename
//go:noescape
func TemplateName(t *template.Template) string

//go:linkname templateoption template.sub_templateoption
func templateoption(t *template.Template, opt ...string) *template.Template {
	return t.Option(opt)
}

//go:linkname TemplateOption template.sub_templateoption
//go:noescape
func TemplateOption(t *template.Template, opt ...string) *template.Template

//go:linkname URLQueryEscaper html/template.URLQueryEscaper
//go:noescape
func URLQueryEscaper(args ...interface{}) string

//go:linkname templateclone template.sub_templateclone
func templateclone(t *template.Template) (*template.Template, error) {
	return t.Clone()
}

//go:linkname TemplateClone template.sub_templateclone
//go:noescape
func TemplateClone(t *template.Template) (*template.Template, error)

//go:linkname templateparse template.sub_templateparse
func templateparse(t *template.Template, text string) (*template.Template, error) {
	return t.Parse(text)
}

//go:linkname TemplateParse template.sub_templateparse
//go:noescape
func TemplateParse(t *template.Template, text string) (*template.Template, error)
