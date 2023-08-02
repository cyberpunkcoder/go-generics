package main

/*
This example shows that a tilde can be used to
use any type that is an alias for an underlying type.
*/

import "fmt"

type UserID int

/*
 Here we use the tilde to specify that the type
 can be any type that is an alias for an int.
 */
func Add[T ~int](a, b T) T {
	return a + b
}

func main() {
	a := UserID(1)
	b := UserID(2)
	result := Add(a,b)
	fmt.Printf("result: %+v\n", result)
}