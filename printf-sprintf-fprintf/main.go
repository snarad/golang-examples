package main

import (
	"fmt"
	"os"
)

func main() {
	// `Printf`, which prints the
	// formatted string to `os.Stdout`.
	fmt.Printf("Hello World\n")

	// `Sprintf` formats and returns a string without
	// printing it anywhere.
	s := fmt.Sprintf("Name is %s", "Shivalik")
	fmt.Println(s)

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
