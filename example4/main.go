package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type UserData interface {
	constraints.Ordered | []byte | []rune
}

type User[T UserData] struct {
	ID   int
	Name string
	Data T
}

func main() {
	u := User[int]{
		ID:   0,
		Name: "John Doe",
		Data: 123,
	}
	fmt.Printf("user: %+v\n", u)
}
