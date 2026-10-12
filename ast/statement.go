package ast

type Statement struct{}

type ClassOrFunction interface {
	IsClassOrFunction()
}
