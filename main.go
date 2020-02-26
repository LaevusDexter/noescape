package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"go/ast"
	"go/parser"
	"go/printer"
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
// 4) figure out what causes "image" package errors
// 5) figure out why does comma appear after function parameters

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
	"syscall",
	"js",
	"syslog",
	"testing",
}

// not supported packages (yet)
var filter = []string{
	"image",
}

// versionFilter
var _ = []string {
	"tls", // 1.14
	"json", // 1.14
	"dwarf", // 1.14
	"textproto", // 1.14
	"strconv", // 1.14
	"http", // 1.14
	"maphash", // 1.14
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
			go generate(funcs, p, &wg)
		}
	}
	
	wg.Wait()
}

var funcExpr = regexp.MustCompile(`func\s(\((.*?)\))?\s?([A-z]+)\(.*?\) ([A-z,() \[\]*]*)?`)

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

	pimp := p.path[:len(p.path)-1]

	fmt.Println(pimp, "in process...")
	defer fmt.Println(pimp, "just got generated.")
	
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

				receiver = fn.receiver + " " + receiver

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
		for i, f := range fn.Type.Params.List {
			cursor := 0
			lookupType(&cursor, func(cur int, typ, pkg string, ext bool) {
				if ext {
					packages[pkg] = struct{}{}
				}
			}, f.Type)

			if i == len(fn.Type.Params.List) - 1 && len(f.Names) == 1 {
				var buf []byte
				buildType(&buf, f.Type)

				if len(buf) > 3 && buf[0] == '.' && buf[1] == '.' && buf[2] == '.' {
					params = append(params, f.Names[0].Name + "...")

					continue
				}
			}

			for _, n := range f.Names {
				params = append(params, n.Name)
			}
		}

		fun, ok := funcs[fname]
		if !ok {
			fun = new(function)
		}

		if fn.Type.Results == nil {
			parameters[fname] = params

			continue
		}

		fun.isReturns = true

		for _, f := range fn.Type.Results.List {
			cursor := 0
			lookupType(&cursor, func(cur int, typ, pkg string, ext bool) {
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

	self := false
	n := astutil.Apply(f, nil, func(c *astutil.Cursor) bool {
		if fn, ok := c.Node().(*ast.FuncType); ok {
			if fn.Params == nil {
				return true
			}

			for i, prm := range fn.Params.List {
				var ftyp []byte
				buildType(&ftyp, prm.Type)

				cursor := 0
				lookupType(&cursor, func(cur int, typ, pkg string, ext bool) {
					if !ext && isUpper(typ) {
						self = true

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

				cursor := 0
				lookupType(&cursor, func(cur int, typ, pkg string, ext bool) {
					if !ext && isUpper(typ) {
						self = true

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


				var buf []byte
				if len(rsts.Names) > 1 && len(fn.Results.List) == 1 {
					buf = append(buf, '(')
				}

				for c := len(rsts.Names)-1; c > 0; c-- {
					buf = append(buf, ftyp...)
					buf = append(buf, ", "...)
					/*rsts := *rsts
					rsts.Type = ast.NewIdent(string(ftyp))
					rsts.Names = nil

					fn.Results.List = append(fn.Results.List[:i], append([]*ast.Field{
						&rsts,
					}, fn.Results.List[i:]...)...)*/
				}

				if len(rsts.Names) > 1 {
					buf = append(buf, ftyp...)
					if len(fn.Results.List) == 1 {
						buf = append(buf, ')')
					}

					fn.Results.List[i].Type = ast.NewIdent(string(buf))
				}

				rsts.Names = nil
			}
		}

		return true
	})

	if self {
		astutil.AddImport(fset, f, pimp)
	} else {
		// for linkname
		astutil.AddNamedImport(fset, f, "_", pimp)
	}

	ast.SortImports(fset, f)

	var buf bytes.Buffer

	config := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	err2 := config.Fprint(&buf, fset, n)
	if err2 != nil {

		panic(p.path + err2.Error())
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

		args = strings.Join(params, ", ")
	default:
		if fn.isMethod {
			params = params[1:]
		}

		args = strings.Join(params, ", ")
		if args[len(args)-2:] == ", " {
			args = args[:len(args)-2]
		}
	}

	return buildLinkname(lname, pkg, "sub_" + lname, false) +
		buildMethodWrapper(fn, lname, args) +
		buildLinkname(fn.name, pkg, "sub_" + lname, true) + "\n" + fn.signatureWithReceiver
}

func buildMethodWrapper(fn *function, lname, args string) string {
	ret := ""
	if fn.isReturns {
		ret = "return " + fn.receiver
	} else {
		ret = fn.receiver
	}

	return fmt.Sprintf(methodTemplate,
		strings.Replace(fn.signatureWithReceiver, fn.name, lname, 1),
		ret,
		fn.originalName,
		args,
	)
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

%s
`

func lookupType(cur *int, visit func(cur int, typ, pkg string, ext bool), expr ast.Expr) {
	switch x := expr.(type) {
	case *ast.ArrayType:
		*cur += len("[")
		if x.Len != nil {
			lookupType(cur, visit, x.Len)
		}
		*cur += len("]")
		lookupType(cur, visit, x.Elt)
	case *ast.ChanType:
		switch x.Dir {
		case ast.SEND:
			*cur += len("chan<- ")
		case ast.RECV:
			*cur += len("<-chan ")
		default:
			*cur += len("chan ")
		}

		lookupType(cur, visit, x.Value)
	case *ast.FuncType:
		*cur += len("func(")
		calcFieldList(cur, visit, x.Params,", ")
		*cur += len(")")

		res := x.Results
		n := res.NumFields()
		if n == 0 {
			return
		}

		*cur += len(" ")
		if n == 1 && len(res.List[0].Names) == 0 {
			lookupType(cur, visit, res.List[0].Type)

			return
		}

		*cur += len("(")
		calcFieldList(cur, visit, x.Results,", ")
		*cur += len(")")
	case *ast.StructType:
		*cur += len("struct{")
		if x.Fields == nil {
			*cur += len("}")

			return
		}

		calcFieldList(cur, visit, x.Fields,"; ")

		*cur += len("}")
	case *ast.MapType:
		*cur += len("map[")
		lookupType(cur, visit, x.Key)
		*cur += len("]")
		lookupType(cur, visit, x.Value)
	case *ast.Ellipsis:
		*cur += len( "...")
		if x.Elt != nil {
			lookupType(cur, visit, x.Elt)
		}
	case *ast.StarExpr:
		*cur += len("*")
		lookupType(cur, visit, x.X)
	case *ast.InterfaceType:
		*cur = len( "interface{}")
	case *ast.SelectorExpr:
		var pkg []byte
		buildType(&pkg, x.X)

		visit(*cur, x.Sel.Name, string(pkg), true)

		*cur += len(pkg) +  len(".") + len(x.Sel.Name)
	case *ast.Ident:
		visit(*cur, x.Name, "", false)

		*cur += len(x.Name)
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
	case *ast.Ellipsis:
		*buf = append(*buf, "..."...)
		if x.Elt != nil {
			buildType(buf, x.Elt)
		}
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

func calcFieldList(cur *int, visit func(cur int, typ, pkg string, ext bool), fields *ast.FieldList, sep string) {
	for i, f := range fields.List {
		if i > 0 {
			*cur = len(sep)
		}

		for i, name := range f.Names {
			if i > 0 {
				*cur = len(", ")
			}

			*cur = len(name.Name)
		}

		if len(f.Names) > 0 {
			*cur = len( " ")
		}

		lookupType(cur, visit, f.Type)
	}
}

func isUpper(str string) bool {
	if r, _ := utf8.DecodeRuneInString(str); unicode.IsUpper(r) {
		return true
	}

	return false
}