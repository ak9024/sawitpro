// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type Estate struct {
	Id     string
	Width  int
	Length int
}

type EstateTree struct {
	Id       string
	EstateID string
	X        int
	Y        int
	Height   int
}
