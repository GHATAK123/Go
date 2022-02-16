package main

import "fmt"

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
