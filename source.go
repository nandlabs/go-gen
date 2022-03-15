package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"
)

type Source struct {
	//CodeBase include the member definitions in Base
	*CodeBase
	//dir specifies the location of the source file.
	dir string
	//pkg specifies the package name inside which the source file will be created.
	pkg string
	//name of the source file
	name string
	//structs has the list of structures in this source file.
	structs map[string]*Struct

	functions map[string]*Function

	//imports
	imports []*Import // TODO handle imports as is can be duplicate need to handle during generation
}

type Import struct {
	PkgName string
	Alias   string
}

// AddImport function will add import to the source file.
func (s *Source) AddImport(pkgName, alias string) *Source {
	i := &Import{
		PkgName: pkgName,
		Alias:   alias,
	}
	s.imports = append(s.imports, i)
	return s
}

//Struct creates a new struct for the source.
func (s *Source) Struct(name string) *Struct {
	if str, ok := s.structs[name]; ok {
		return str
	} else {
		strct := &Struct{

			name:    name,
			members: make(map[string]*members),
			source:  s,
			Base: &Base{
				GoGen:   s.GoGen,
				Comment: "",
			},
		}
		s.structs[name] = strct
		return strct
	}

}

func (s *Source) Func(name string) *Function {
	if f, ok := s.functions[name]; ok {
		return f
	} else {
		f := &Function{
			funcSignature: &funcSignature{
				fnName:   name,
				args:     nil,
				receiver: nil,
				ret:      nil,
				Base: &Base{
					GoGen:   s.GoGen,
					Comment: "",
				},
			},
			funcBody: nil,
		}
		s.functions[name] = f
		return f
	}

}

func (s *Source) Generate(writer io.Writer) {
	var buf bytes.Buffer
	s.WriteComments(&buf)

	buf.WriteString("package ")
	buf.WriteString(s.pkg)

	//TODO update the new line with const in all source files
	buf.WriteString("\n\n")
	if len(s.imports) > 0 {
		buf.WriteString("import (\n") //import open
		cleanImports(s.imports)
		for _, i := range s.imports {
			if len(i.Alias) > 0 {
				buf.WriteString(i.Alias)
				buf.WriteString(" ")
			}
			buf.WriteString(i.PkgName)
			buf.WriteString("\n")
		}
		buf.WriteString(")\n") // import close
	}

	// Prepare and write Constants

	if len(s.ConstantsSlice) > 0 {

		//for _, c := range s.Constants {
		//
		//}
	}

	// Prepare and write structs
	for n, str := range s.structs {
		//TODO update println with log statements
		fmt.Println("generating struct defn for struct name " + n)
		b := new(strings.Builder)
		str.Generate(b)
		buf.WriteString(b.String())
	}

	buf.WriteString("\n\n")
	// Prepare and write functions

	for n, fn := range s.functions {
		fmt.Println("generating function with  name " + n)
		b := new(strings.Builder)
		fn.Generate(b)
		buf.WriteString(b.String())

	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err.Error())
	}
	writer.Write(p)

}

func cleanImports(i []*Import) {

}
