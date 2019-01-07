// To define methods on a type you don’t “own”, you
// need to define an alias for the type you want to extend
package main

import (
	"fmt"
	"strings"
)

type MyStr string

func (s MyStr) UpperCase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyStr("shivalik").UpperCase())
}
