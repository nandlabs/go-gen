package gen

import (
	"bytes"
	"io"
)

type sigvar struct {
	name      string
	qType     string
	isPointer bool
}

type funcSignature struct {
	*Base
	fnName   string
	args     []*sigvar
	ret      []*sigvar
	receiver *sigvar
}

type funcBody struct {
	*CodeBase
}

type Function struct {
	*funcSignature
	*funcBody
}

func (f *Function) Arg(name, qType string) *Function {

	typ, isPointer := f.handleType(qType)
	p := &sigvar{
		name:      name,
		qType:     typ,
		isPointer: isPointer,
	}

	f.args = append(f.args, p)
	return f
}

func (f *Function) ReType(qType string) *Function {
	typ, isPointer := f.handleType(qType)
	f.ret = append(f.ret, &sigvar{
		name:      "",
		qType:     typ,
		isPointer: isPointer,
	})
	return f
}

func (f *Function) ReTypeVar(name, qType string) *Function {
	typ, isPointer := f.handleType(qType)
	f.ret = append(f.ret, &sigvar{
		name:      name,
		qType:     typ,
		isPointer: isPointer,
	})
	return f
}
func (f *Function) Receiver(name, qType string) *Function {
	typ, isPointer := f.handleType(qType)
	f.receiver = &sigvar{
		name:      name,
		qType:     typ,
		isPointer: isPointer,
	}
	return f
}

func (f *Function) Generate(w io.Writer) {
	var buf bytes.Buffer

	//fb := f.funcBody
	f.WriteComments(&buf)
	buf.WriteString("func ")
	if f.receiver != nil {
		buf.WriteString("(")
		buf.WriteString(f.receiver.name)
		buf.WriteString(" ")
		if f.receiver.isPointer {
			buf.WriteString("*")

		}
		buf.WriteString(f.receiver.qType)
		buf.WriteString(") ")
	}
	buf.WriteString(f.fnName)
	buf.WriteString("(")
	argLength := len(f.args)
	for i, arg := range f.args {
		buf.WriteString(arg.name)
		buf.WriteString(" ")
		if arg.isPointer {
			buf.WriteString("*")
		}
		buf.WriteString(arg.qType)
		if i < argLength-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString(") {") //Body Open
	buf.WriteString(" \n")
	buf.WriteString("}") //Body Close
	buf.WriteString(" \n")
	w.Write(buf.Bytes())

}
