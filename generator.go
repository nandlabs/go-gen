package go_gen

import (
	"fmt"
	"io"
	"os"
	"path"
)

type Config struct {
	LineWidth int
	NewLine   string
}

type Generator interface {
	Generate(writer io.Writer)
}

type GoGen struct {
	Config  *Config
	BaseDir string
	Sources map[string]*Source
}

func getDefaultConfig() *Config {
	return &Config{
		LineWidth: 120,
		NewLine:   "\n",
	}

}

func New(baseDir string) *GoGen {
	return &GoGen{
		Config:  getDefaultConfig(),
		BaseDir: baseDir,
		Sources: make(map[string]*Source),
	}
}

func (g *GoGen) With(config *Config) *GoGen {
	g.Config = config
	return g
}

// Source  creates a new Source object,
//dir specifies the location where this source file will be created
//pkg specifies the package name of this source. If the package name structure does not exist, it will create one.
//name is the name of the source file.
func (g *GoGen) Source(dir, pkg, name string) *Source {

	if s, ok := g.Sources[getSourceId(dir, pkg, name)]; ok {
		return s
	} else {

		s := &Source{
			CodeBase: &CodeBase{
				Base:    &Base{GoGen: g},
				Comment: "",
			},
			Dir:     "",
			Pkg:     pkg,
			Name:    name,
			Structs: make(map[string]*Struct),
			Imports: nil,
		}
		g.Sources[getSourceId(dir, pkg, name)] = s
		return s
	}

}

func getSourceId(dir, pkg, name string) string {
	//TODO update this to find the right id. May be this logic will work.
	return dir + pkg + name
}

func (g *GoGen) Generate() {

	for k, v := range g.Sources {
		fmt.Println("Generating source file for key:" + k)
		sourceDir := path.Join(g.BaseDir, v.Dir)
		err := os.MkdirAll(sourceDir, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
		f, err := os.Create(path.Join(sourceDir, v.Name))
		if err != nil {
			panic(err.Error())
		}
		v.Generate(f)
	}
}

//AddExternalLib function will handle import as well as go mod entries
func (g *GoGen) AddExternalLib(name, version string) {

}
