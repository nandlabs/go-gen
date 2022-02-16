package go_gen

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Source struct {
	//CodeBase include the member definitions in Base
	*CodeBase
	//Dir specifies the location of the source file.
	Dir string
	//Pkg specifies the package name inside which the source file will be created.
	Pkg string
	//Name of the source file
	Name string
	//Structs has the list of structures in this source file.
	Structs map[string]*Struct
	//Imports
	Imports []*Import // TODO imports as is can be duplicate need to handle during generation
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
	s.Imports = append(s.Imports, i)
	return s
}

//Struct creates a new struct for the source.
func (s *Source) Struct(name string) *Struct {
	if str, ok := s.Structs[name]; ok {
		return str
	} else {
		strct := &Struct{
			CodeBase: &CodeBase{
				Base:    s.Base,
				Comment: "",
			},
			Name:    name,
			Members: make(map[string]*Member),
			source:  s,
		}
		s.Structs[name] = strct
		return strct
	}

}

func (s Source) AddVar(qName, typ string, isPointer bool, val interface{}) {

}

func (s *Source) Generate(writer io.Writer) {
	var buf bytes.Buffer
	if s.Comment != "" {
		buf.WriteString("//" + s.Comment + "\n\n")
	}
	buf.WriteString("package ")
	buf.WriteString(s.Pkg)

	//TODO update the new line with const in all source files
	buf.WriteString("\n\n")
	if len(s.Imports) > 0 {
		buf.WriteString("import (\n") //import open
		cleanImports(s.Imports)
		for _, i := range s.Imports {
			if len(i.Alias) > 0 {
				buf.WriteString(i.Alias)
				buf.WriteString(" ")
			}
			buf.WriteString(i.PkgName)
			buf.WriteString("\n")
		}
		buf.WriteString(")\n") // import close
	}

	for n, str := range s.Structs {
		//TODO update println with log statements
		fmt.Println("generating struct defn for struct name " + n)
		b := new(strings.Builder)
		str.Generate(b)
		buf.WriteString(b.String())
	}

	writer.Write(buf.Bytes())

}

func cleanImports(i []*Import) {

}
