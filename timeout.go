package main

import (
	"fmt"
	"sync"
	"time"
)

var println = fmt.Println

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	start := time.Now()

	println("Now", start)

	time.AfterFunc(time.Second, func() {
		fmt.Println("after 1s")

		println("Now", time.Now())
		println("Now", time.Since(start))

		defer wg.Done()
	})

	wg.Wait()
}
