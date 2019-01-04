package main

import "fmt"

const (
	Pi    = 3.14
	Truth = false
)

func main() {
	const Gretting = "ハローワールド"
	action := func() { // a variable can contain any type, including functions
		fmt.Println(Gretting)
		fmt.Println(Pi)
		fmt.Println(Truth)
	}
	action()
}
