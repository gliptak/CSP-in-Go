package main

import (
	"fmt"
)

func west(c chan <- rune, r rune) {
	c <- r
}

func east(c <- chan rune, done chan bool) {
	for {
		r, more := <-c
		if more {
			if r != rune('*') {
				fmt.Printf("%c", r)
			} else {
				r1, more1 := <-c
				if more1 {
					if r1 != rune('*') {
						fmt.Printf("%c", r)
						fmt.Printf("%c", r1)
					} else {
						fmt.Printf("^")
					}
				} else {
					fmt.Printf("%c", r)
					done <- true
					return
				}
			}
		} else {
			done <- true
			return
		}
	}
}

func main() {
	rs := []rune("Hel*lo**World***")
	c := make(chan rune)
	done := make(chan bool)
	go east(c, done)
	for i := 0; i < len(rs); i++ {
		west(c, rs[i])
	}
	close(c)
	<- done
}