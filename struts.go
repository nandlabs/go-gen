package gen

import (
	"bytes"
	"io"
)

type Struct struct {
	*Base
	name    string
	members map[string]*members
	source  *Source
}

type members struct {
	*Base
	name      string
	qType     string
	isPointer bool
	tags      map[string]string
}

func (s *Struct) AddMember(name, typ string, isPointer bool) *members {
	//TODO Add check for presence of valid type
	// TODO Type imports needs to be checked

	if v, ok := s.members[name]; ok {
		return v
	} else {
		m := &members{
			name:      name,
			qType:     typ,
			isPointer: isPointer,
			tags:      make(map[string]string),
			Base: &Base{
				GoGen:   s.GoGen,
				Comment: "",
			},
		}
		//TODO add  the import for struct reference from another package.
		s.members[name] = m
		return m
	}

}

func (m *members) AddTag(name, val string) *members {
	m.tags[name] = val
	return m
}

func (s Struct) Generate(w io.Writer) {
	var buf bytes.Buffer
	s.WriteComments(&buf)
	buf.WriteString("type ")
	buf.WriteString(s.name + " struct { \n")
	for k, v := range s.members {
		buf.WriteString(k)
		if v.isPointer {
			buf.WriteString(" *")
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(v.qType)
		l := len(v.tags)
		if l > 0 {
			buf.WriteString(" `") //Tag Start
			for tn, tv := range v.tags {
				buf.WriteString(tn)
				buf.WriteString(":\"")
				buf.WriteString(tv)
				l--
				if l != 0 {
					buf.WriteString("\" ")
				} else {
					buf.WriteString("\"")
				}

			}
			buf.WriteString("`") // Tag End
		}
		buf.WriteString("\n")
	}
	buf.WriteString("\n}")

	w.Write(buf.Bytes())
}
