package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(m *sync.Mutex) {
	if true {
		m.Lock()
		// defer only gets called after the function increment completely ends
		// hence the scope of the lock goes outside the if conditional and x = x + 1
		// would work as if it is surrounded by the lock
		// https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01
		defer m.Unlock()
	}
	x = x + 1

}

func main() {
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		go increment(&m)
	}
	time.Sleep(time.Second)
	fmt.Println("final value of x", x)
}
