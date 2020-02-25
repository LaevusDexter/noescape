package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

// TODO:
// 1) make it readable
// 2) optimize the code (?)
// 3) add wrapped types for methods

const pkgLink = "https://golang.org/pkg/"

const cssClass = "pkg-name"
const pkgXPath = `//*[contains(@class, '`+ cssClass + `')]/a`

const cssID = `manual-nav`
const indexSubPath = `dl/dd`
const definitionXPath = `//*[contains(@id, '` + cssID + `')]/` + indexSubPath + `/a`

type pkg struct {
	name string
	path string
}

var packageMap = make(map[string]string)
var currentDir string

var buildins = []string{
	"builtin",
	"unsafe",
}

// not supported packages (yet)
var filter = []string{
	"x509",
	"elf",
	"plan9obj",
	"ast",
	"types",
	"image",
	"big",
	"net",
	"os",
	"reflect",
	"runtime",
	"syscall",
	"js",
	"testing",
	"parse",
	"time",
}

func main() {
	var err error

	currentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	doc, err := htmlquery.LoadURL(pkgLink)
	if err != nil {
		panic(err)
	}

	var packages []pkg

	nodes := htmlquery.Find(doc, pkgXPath)

	for _, n := range nodes {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				pname := htmlquery.InnerText(n)
				pstart := strings.LastIndex(pname, "/")
				if pstart == -1 {
					pstart = 0
				} else {
					pstart++
				}

				pname = pname[pstart:]
				ppath := attr.Val
				
				packages = append(packages, pkg{
					name : pname,
					path : ppath,
				})

				packageMap[pname] = ppath[:len(ppath)-1]

				break
			}
		}
	}

	var wg sync.WaitGroup
	for _, p := range packages {
		if func() bool {
			for i := 0; i < len(buildins); i++ {
				if p.name == buildins[i] {
					return true
				}
			}

			for i := 0; i < len(filter); i++ {
				if p.name == filter[i] {
					return true
				}
			}

			return false
		}() {
			continue
		}

		doc, err = htmlquery.LoadURL(pkgLink + p.path)
		if err != nil {
			continue
		}

		nodes = htmlquery.Find(doc, definitionXPath)
		if len(nodes) == 0 {
			continue
		}

		funcs := make([]string, 0, 2)
		for _, n := range nodes {
			n = n.FirstChild
			line := strings.Trim(htmlquery.InnerText(n), " \t\n\r" + string('\u00A0'))

			if funcExpr.MatchString(line) {
				funcs = append(funcs, line)
			}
		}

		if len(funcs) > 0 {
			wg.Add(1)
			generate(funcs, p, &wg)
		}
	}
	
	wg.Wait()
}

var funcExpr = regexp.MustCompile(`func\s(\((.*?)\))?\s?([A-z]+)\(.*?\) ([A-z,() *]*)?`)

type function struct {
	isMethod bool
	isReturns bool
	receiver string
	receiverType string
	name    string
	originalName string
	signature string
	signatureWithReceiver string
}

func generate(funcIn []string, p pkg, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(p.name)
	
	var result []byte
	var src = []byte("package " + p.name + "\n")
	var packages = make(map[string]struct{})
	var parameters = make(map[string][]string)

	funcs := make(map[string]*function)
	for _, signature := range funcIn {
		matches := funcExpr.FindStringSubmatch(signature)

		var fn function

		if matches[1] != "" {
			fn.isMethod = true
			fn.originalName = matches[3]
			fn.signature = strings.Replace(signature, matches[1], "", 1)

			receiver := matches[2]

			sreceiver := strings.Split(matches[2], " ")
			if len(sreceiver) == 1 {
				fn.receiver = "receiverV"
				fn.receiverType = sreceiver[0]
			} else {
				fn.receiver = sreceiver[0]
				fn.receiverType = sreceiver[1]
			}

			if fn.receiverType[0] == '*' {
				fn.receiverType = fn.receiverType[1:]
			}

			fn.name = fn.receiverType + fn.originalName

			receiver = "(" + receiver
			if i := strings.IndexRune(fn.signature,'('); fn.signature[i+1] != ')' {
				receiver += ", "
			}

			fn.signatureWithReceiver = strings.Replace(fn.signature,"(", receiver, 1)
			fn.signatureWithReceiver = strings.Replace(fn.signatureWithReceiver, fn.originalName, fn.name, 1)

			src = append(src, fn.signatureWithReceiver...)
		} else {
			fn.signature = signature
			fn.name = matches[3]

			src = append(src, fn.signature...)
		}

		src = append(src, '\n')

		if matches[4] != "" {
			fn.isReturns = true
		}

		funcs[fn.name] = &fn
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	for _, f := range f.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		fname := fn.Name.Name

		if fn.Type.Params == nil {
			parameters[fname] = []string{}

			continue
		}

		var params []string
		for _, f := range fn.Type.Params.List {
			lookupType(nil, func(cur int, typ, pkg string, ext bool) {
				if ext {
					packages[pkg] = struct{}{}
				}
			}, f.Type)


			for _, n := range f.Names {
				params = append(params, n.Name)
			}
		}

		if fn.Type.Results == nil {
			parameters[fname] = params

			continue
		}

		for _, f := range fn.Type.Results.List {
			lookupType(nil, func(cur int, typ, pkg string, ext bool) {
				if ext {
					packages[pkg] = struct{}{}
				}
			}, f.Type)
		}

		parameters[fname] = params
	}

	funcOut := make([]string, 0, 2)
	for _, fn := range funcs {
		var out string

		if fn.isMethod {
			out = buildMethodSignature(p.name, fn, parameters[fn.name])
		} else {
			out = buildLinkname(fn.name, p.path[:len(p.path)-1], fn.name, true) + "\n" + fn.signature
		}

		funcOut = append(funcOut, out)
	}

	result = []byte(fmt.Sprintf(pkgTemplate,
		time.Now().Format(time.UnixDate),
		p.name,
		p.path[:len(p.path)-1],
		strings.Join(funcOut, "\n"),
	))

	fset = token.NewFileSet()
	f, err = parser.ParseFile(fset, "", result, parser.ParseComments)
	if err != nil {
		//fmt.Println(string(result))

		panic(p.path + err.Error())
	}

	unsafe := false
	for pkg := range packages {

		if pkg == "unsafe" {
			unsafe = true
		}

		if imp, ok := packageMap[pkg]; ok {
			astutil.AddImport(fset, f, imp)
		}
	}

	if !unsafe {
		// linkname requires unsafe import
		astutil.AddNamedImport(fset, f, "_", "unsafe")
	}

	n := astutil.Apply(f, nil, func(c *astutil.Cursor) bool {
		if fn, ok := c.Node().(*ast.FuncType); ok {
			if fn.Params == nil {
				return true
			}

			for i, prm := range fn.Params.List {
				var ftyp []byte
				buildType(&ftyp, prm.Type)

				cursor := -1
				lookupType(&cursor, func(cur int, typ, pkg string, ext bool) {
					if !ext && isUpper(typ) {
						var buf []byte

						buf = append(buf, ftyp[:cur]...)
						buf = append(buf, p.name...)
						buf = append(buf, '.')
						buf = append(buf, typ...)
						buf = append(buf, ftyp[cur + len(typ):]...)

						ftyp = buf

						fn.Params.List[i].Type = ast.NewIdent(string(buf))
					}
				}, prm.Type)
			}

			if fn.Results == nil {
				return true
			}

			for i, rsts := range fn.Results.List {
				var ftyp []byte
				buildType(&ftyp, rsts.Type)

				cursor := -1
				lookupType(&cursor, func(cur int, typ, pkg string, ext bool) {
					if !ext && isUpper(typ) {
						var buf []byte

						buf = append(buf, ftyp[:cur]...)
						buf = append(buf, p.name...)
						buf = append(buf, '.')
						buf = append(buf, typ...)
						buf = append(buf, ftyp[cur + len(typ):]...)

						ftyp = buf

						fn.Results.List[i].Type = ast.NewIdent(string(buf))
					}
				}, rsts.Type)
			}
		}

		return true
	})

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, n); err != nil {
		//fmt.Println(string(result))

		panic(err)
	}

	path := currentDir + "/" + p.path
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)

		return
	}

	file, err := os.Create(path + p.name + ".go")
	if err != nil {
		fmt.Println(err)

		return
	}
	defer file.Close()

	_, _ = file.Write(buf.Bytes())
}

func buildMethodSignature(pkg string, fn *function, params []string) string {
	lname := strings.ToLower(fn.name)

	var args string
	switch len(params) {
	case 0:
	case 1:
		if !fn.isMethod {
			args = params[0]
		}
	case 2:
		if fn.isMethod {
			args = params[1]

			break
		}

		fallthrough
	default:
		if fn.isMethod {
			params = params[1:]
		}

		args = strings.Join(params, ", ")
		args = args[:len(args)-2]
	}

	ret := ""
	if fn.isReturns {
		ret = "return "
	}

	return buildLinkname(lname, pkg, "sub_" + lname, false) +
		fmt.Sprintf(methodTemplate,
			strings.Replace(fn.signatureWithReceiver, fn.name, lname, 1),
			ret + fn.receiver,
			fn.originalName,
			args,
		) +
		buildLinkname(fn.name, pkg, "sub_" + lname, true) + "\n" + fn.signatureWithReceiver
}

func buildLinkname(name, path, simbol string, noescape bool) string {
	pragma := ""
	if noescape {
		pragma = "\n" + "//go:noescape"
	}

	return fmt.Sprintf(linknameTemplate, name, path, simbol) + pragma
}

const methodTemplate = `
%s {
	%s.%s(%s)
}
`

const linknameTemplate = `
//go:linkname %s %s.%s`

const pkgTemplate = `// This file has automatically been generated on %s.
// DO NOT EDIT.
package %s

import "%s"

%s
`

func lookupType(cur *int, visit func(cur int, typ, pkg string, ext bool), expr ast.Expr) {
	cursor := 0
	if cur != nil {
		*cur++

		cursor = *cur
	}

	switch x := expr.(type) {
	case *ast.ArrayType:
		lookupType(cur, visit, x.Elt)
	case *ast.ChanType:
		lookupType(cur, visit, x.Value)
	case *ast.FuncType:
		if x.Params == nil {
			return
		}
		for _, p := range x.Params.List {
			lookupType(cur, visit, p.Type)
		}

		if x.Results == nil {
			return
		}

		for _, p := range x.Results.List {
			lookupType(cur, visit, p.Type)
		}
	case *ast.StructType:
		if x.Fields == nil {
			return
		}

		for _, f := range x.Fields.List {
			lookupType(cur, visit, f.Type)
		}
	case *ast.MapType:
		lookupType(cur, visit, x.Key)
		lookupType(cur, visit, x.Value)
	case *ast.StarExpr:
		lookupType(cur, visit, x.X)
	case *ast.SelectorExpr:
		var pkg []byte
		buildType(&pkg, x.X)

		visit(cursor, x.Sel.Name, string(pkg), true)
	case *ast.Ident:
		visit(cursor, x.Name, "", false)
	}
}

func buildType(buf *[]byte, expr ast.Expr) {
	switch x := expr.(type) {
	case *ast.ArrayType:
		*buf = append(*buf,'[')
		if x.Len != nil {
			buildType(buf, x.Len)
		}

		*buf = append(*buf,']')
		buildType(buf, x.Elt)
	case *ast.ChanType:
		switch x.Dir {
		case ast.SEND:
			*buf = append(*buf,"chan<- "...)
		case ast.RECV:
			*buf = append(*buf,"<-chan "...)
		default:
			*buf = append(*buf,"chan "...)
		}

		buildType(buf, x.Value)
	case *ast.FuncType:
		*buf = append(*buf, "func"...)

		*buf = append(*buf,'(')
		writeFieldList(buf, x.Params, ", ")
		*buf = append(*buf, ')')

		res := x.Results
		n := res.NumFields()
		if n == 0 {
			return
		}

		*buf = append(*buf, ' ')
		if n == 1 && len(res.List[0].Names) == 0 {
			buildType(buf, res.List[0].Type)

			return
		}

		*buf = append(*buf, '(')
		writeFieldList(buf, res, ", ")
		*buf = append(*buf, ')')
	case *ast.StructType:
		*buf = append(*buf, "struct{"...)

		if x.Fields != nil {
			writeFieldList(buf, x.Fields, "; ")
		}

		*buf = append(*buf, '}')
	case *ast.MapType:
		*buf = append(*buf,"map["...)
		buildType(buf, x.Key)

		*buf = append(*buf, ']')
		buildType(buf, x.Value)
	case *ast.InterfaceType:
		*buf = append(*buf, "interface{}"...)
	case *ast.SelectorExpr:
		buildType(buf, x.X)
		*buf = append(*buf,'.')
		*buf = append(*buf, x.Sel.Name...)
	case *ast.StarExpr:
		*buf = append(*buf,'*')
		buildType(buf, x.X)
	case *ast.Ident:
		*buf = append(*buf, x.Name...)
	}
}

func writeFieldList(buf *[]byte, fields *ast.FieldList, sep string) {
	for i, f := range fields.List {
		if i > 0 {
			*buf = append(*buf, sep...)
		}

		for i, name := range f.Names {
			if i > 0 {
				*buf = append(*buf, ", "...)
			}

			*buf = append(*buf, name.Name...)
		}

		if len(f.Names) > 0 {
			*buf = append(*buf, ' ')
		}

		buildType(buf, f.Type)
	}
}

func isUpper(str string) bool {
	if r, _ := utf8.DecodeRuneInString(str); unicode.IsUpper(r) {
		return true
	}

	return false
}