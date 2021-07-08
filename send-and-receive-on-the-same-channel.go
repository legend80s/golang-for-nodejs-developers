package main

import "fmt"

var battle = make(chan string)

func warrior(name string, done chan string) {
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		fmt.Printf("send '%s' to battle\n", name)
		// I lost :-(
	}
	done <- ""
}

func main() {
	done := make(chan string)
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs {
		go warrior(l, done)
	}
	for _ = range []string{"", "", "", "", "", ""} {
		<-done
	}
}
