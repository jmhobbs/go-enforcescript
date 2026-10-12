package ast

type Class struct {
	Identifier string
	Body       Block
}

func (c Class) IsClassOrFunction() {}
