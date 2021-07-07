package main

import (
	"fmt"
	"time"
)

var println = fmt.Println

func cb(i int) {
	println(i)
}

func main() {
	ticker := time.NewTicker(time.Second)
	i := 0

	for range ticker.C {
		cb(i)

		if i == 3 {
			ticker.Stop()

			break
		}

		i += 1
	}

	// timer := time.NewTimer(time.Second)

	// start := time.Now()
	// println(start)

	// for range timer.C {
	// 	println(time.Since(start))
	// 	break
	// }
}
