package main

import "fmt"

/*
This example shows that its possible to create a custom map
type that can be used with any type that is comparable.
This is because the valid type for a map key is any comparable.
For the value type in this example, we specify that it must be
either an int or a string.
*/

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	m := make(CustomMap[int, string])
	m[3] = "three"
	fmt.Printf("map: %+v\n", m)
}
