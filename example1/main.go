package main

/*
This example shows how generics can be used to
create a generic function that can be used with
different types.
*/

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
Add adds two values of the same type.
The type must implement the Ordered interface.
The ordered interface contains the following types:
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- float32
- float64
*/
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {
	result := Add(1.2, 2.0)
	fmt.Printf("result: %+v\n", result)
}
