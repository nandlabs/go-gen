package go_gen

//Base type defn to be used by all top level types
type Base struct {
	GoGen *GoGen
}

type CodeBase struct {
	*Base
	Comment string
}

//AddComment function to add comment to  the type using this Base Type
func (cb *CodeBase) AddComment(c string) {
	// TODO Add logic to break lines on max width from Conf
	if cb.Comment != "" {
		cb.Comment += "\n"
		cb.Comment += c
	} else {
		cb.Comment = c
	}
}

func (cb *CodeBase) HandleType(qName string) {

}
func (cb *CodeBase) HandleFunction(qName string) {

}
