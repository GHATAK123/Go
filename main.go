package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome Gofers")
	a, b, c, d := returnMany(4)
	fmt.Println(a, b, c, d)
}
