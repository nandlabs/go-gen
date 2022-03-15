package gen

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

//Base type defn to be used by all top level types
type Base struct {
	GoGen   *GoGen
	Comment string
}

type CodeBase struct {
	*Base
	Variables      map[string]*Var
	Constants      map[string]*Const
	ConstantsSlice []*Const
}

type Var struct {
	*Base
	Name      string
	QType     string
	Value     interface{}
	IsPointer bool
}

type Const struct {
	*Base
	Name  string
	QType string
	Value interface{}
}

//AddComment function to add comment to  the type using this Base Type
func (b *Base) AddComment(c string) {

	// TODO Add logic to break lines on max width from Conf

	var buf bytes.Buffer
	buf.WriteString(b.Comment)
	s := bufio.NewScanner(strings.NewReader(c))
	for s.Scan() {
		buf.WriteString("// ")
		buf.WriteString(s.Text())
		buf.WriteString("\n")
	}
	b.Comment = buf.String()
	b.Comment = b.Comment[:len(b.Comment)-1]

}

func (b *Base) WriteComments(sw io.StringWriter) {
	if b.Comment != "" {
		sw.WriteString(b.Comment)
		sw.WriteString("\n")
	}

}

func (cb *CodeBase) Var(n string, isPointer bool) *Var {

	if cb.Variables == nil {
		cb.Variables = make(map[string]*Var)
	}

	if v, ok := cb.Variables[n]; ok {
		return v
	} else {
		v = &Var{

			Base:      cb.Base,
			Name:      n,
			Value:     nil,
			IsPointer: isPointer,
		}
		cb.Variables[n] = v
		return v
	}

}

func (v *Var) Typ(qType string) *Var {

	v.QType, _ = v.handleType(qType)
	return v
}

// Val This method accepts either value of this type or a Map if the qName resolves to a structType with multiple
//members. The map can be nested if the member itself is of type struct
//TODO Add documentation example
func (v *Var) Val(i interface{}) *Var {
	v.Value = i
	return v
}

//Const Create or retrieve const defn.
// A slice of Const  maintained to preserve the order to handle cases with iota value
func (cb *CodeBase) Const(n string) *Const {

	if cb.Variables == nil {
		cb.Variables = make(map[string]*Var)
	}

	if v, ok := cb.Constants[n]; ok {
		return v
	} else {
		v = &Const{
			Base:  cb.Base,
			Name:  n,
			Value: nil,
		}
		cb.Constants[n] = v
		cb.ConstantsSlice = append(cb.ConstantsSlice, v)
		return v
	}

}

//Typ This method sets the type of the constant.
// Its optional and will have
func (c *Const) Typ(qType string) *Const {

	c.QType, _ = c.handleType(qType)
	return c
}

func (c *Const) Val(i interface{}) *Const {
	c.Value = i
	return c
}

func (b *Base) handleType(qName string) (string, bool) {
	//TODO this function handles the qName imports.
	name := strings.TrimSpace(qName)
	isPointer := strings.HasPrefix(qName, "*")
	if isPointer {
		name = qName[1:]
	}

	return name, isPointer
}
