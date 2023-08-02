package main

/*
This example shows a practical application for generics.
The MapValues function takes a slice of values and a function
that maps a value to a new value. It returns a new slice of
values that are the result of applying the map function to
each value in the original slice.
*/

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	// This map function takes a float64 and returns it multiplied by two.
	mapFunc := func(n float64) float64 {
		return n * 2
	}
	// We pass the map function to MapValues along with a slice of float64 values.
	result := MapValues([]float64{1, 2, 3}, mapFunc)
	fmt.Printf("result: %+v\n", result)
}
