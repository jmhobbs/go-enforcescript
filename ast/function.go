package ast

type Function struct {
	ReturnType uint8
	Identifier string
	Arguments  []Argument
	Body       Block
}

func (f Function) IsClassOrFunction() {}
