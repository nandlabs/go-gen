package gen

type Interface struct {
	name      string
	functions []funcSignature
}

func NewInterface(name string) Interface {

	if name != "" {

		return Interface{
			name:      name,
			functions: nil,
		}

	} else {
		panic("Interface name cannot be empty!")
	}

}

func (i Interface) AddFunc(name string) Interface {

	return i
}
