# Go Generics

The following are some example use cases for go generics.
## Example 1
Generics can be used to create a generic function.
```go
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {
	result := Add(1.2, 2.0)
	fmt.Printf("result: %+v\n", result)
}
```
## Example 2
Type aliases can be used in generics.
```go
func Add[T ~int](a, b T) T {
	return a + b
}

func main() {
	a := UserID(1)
	b := UserID(2)
	result := Add(a,b)
	fmt.Printf("result: %+v\n", result)
}
```
### Example 3
Generics can be used in combination with function parameters.
```go
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
```
## Example 4
Generics can be used in combination with structs to create generic structs.
```go
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
```
## Example 5
Generics can be used with maps to create custom maps.
```go
type CustomMap[T comparable, V int | string] map[T]V

func main() {
	m := make(CustomMap[int, string])
	m[3] = "three"
	fmt.Printf("map: %+v\n", m)
}
```
See example subfolders for more in-depth explanations.
