package go_gen

import (
	"bytes"
	"io"
)

type Struct struct {
	*CodeBase
	Name    string
	Members map[string]*Member
	source  *Source
}

type Member struct {
	*CodeBase
	Name      string
	Typ       string
	IsPointer bool
	Tags      map[string]string
}

func (s *Struct) AddMember(name, typ string, isPointer bool) *Member {
	//TODO Add check for presence of valid type
	// TODO Type Imports needs to be checked

	if v, ok := s.Members[name]; ok {
		return v
	} else {
		m := &Member{
			Name:      name,
			Typ:       typ,
			IsPointer: isPointer,
			Tags:      make(map[string]string),
		}
		//TODO add  the import for struct reference from another package.
		s.Members[name] = m
		return m
	}

}

func (m *Member) AddTag(name, val string) *Member {
	m.Tags[name] = val
	return m
}

func (s Struct) Generate(w io.Writer) {
	var buf bytes.Buffer
	if s.Comment != "" {
		buf.WriteString("//")
		buf.WriteString(s.Comment)
		buf.WriteString("\n")

	}
	buf.WriteString("type ")
	buf.WriteString(s.Name + " struct { \n")
	for k, v := range s.Members {
		buf.WriteString(k)
		if v.IsPointer {
			buf.WriteString(" *")
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(v.Typ)
		l := len(v.Tags)
		if l > 0 {
			buf.WriteString(" `") //Tag Start
			for tn, tv := range v.Tags {
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
