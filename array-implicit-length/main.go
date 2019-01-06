package main

import "fmt"

func main() {
	a := [...]string{"hello", "my", "name", "is", "shivalik"}
	fmt.Println(a)
	fmt.Printf("%s\n", a)
	fmt.Printf("%q\n", a)
}
