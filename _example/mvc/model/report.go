package model

type ModelI interface  {
	Print() string
}

type ModelOne struct {}
type ModelTwo struct {}

func (mo *ModelOne) Print() string {
	return "One"
}

func (mt *ModelTwo) Print() string {
	return"Two"
}