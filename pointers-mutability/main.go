// In Go, only constants are immutable. However because arguments are passed by value (copying the arguments),
// a function receiving a value argument and mutating it, won't mutate the original value.
package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	// Output without pointer [me := Artist{Name: "Shivalik", Genre: "Jazz", Songs: 52}]
	// Shivalik released their 53th song
	// Shivalik has total of 52 songs
	//
	// Output with pointer [me := &Artist{Name: "Shivalik", Genre: "Jazz", Songs: 52}]
	// Shivalik released their 53th song
	// Shivalik has total of 53 songs
	me := &Artist{Name: "Shivalik", Genre: "Jazz", Songs: 52}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has total of %d songs\n", me.Name, me.Songs)
}
