// There are two reasons to use a pointer receiver. First, to
// avoid copying the value on each method call (more efficient
// if the value type is a large struct).
// The other reason why you might want to use a pointer is so
// that the method can modify the value that its receiver points to.
//
// Remember that Go passes everything by value, meaning that when Greeting()
// is defined on the value type, every time you call Greeting(), you are copying
// the User struct. Instead when using a pointer, only the pointer is copied (cheap).
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
