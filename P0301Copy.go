package main

import (
	"fmt"
)

func west(c chan rune, r rune) {
	c <- r
}

func east(c chan rune, done chan bool) {
	for {
		r, more := <-c
		if more {
			fmt.Printf("%c\n", r)
		} else {
			done <- true
			return
		}
	}
}

func main() {
	rs := []rune("Hello, 世界")
	c := make(chan rune)
	done := make(chan bool)
	go east(c, done)
	for i := 0; i < len(rs); i++ {
		west(c, rs[i])
	}
	close(c)
	<- done
}