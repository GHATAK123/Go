package main

import (
	"fmt"
)

func main() {
	// Initializing the values
	// of the author structure
	result := author{"Sona", "CSE", 203, 34000}
	// Creating and initializing
	// the anonymous structure
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}

	// Creating a structure
	// with anonymous fields
	type student struct {
		int
		string
		float64
	}
	value := student{123, "Bud", 8900.23}
	fmt.Println(value)
	// Display the anonymous structure
	fmt.Println(Element)
	fmt.Println("Welcome Gofers")
	a, b, c, d := returnMany(4)
	fmt.Println(a, b, c, d)
	w, x, y, z := calculate(12, 2)
	fmt.Println(w, x, y, z)
	t1, t2, t3, t4 := result.show()
	fmt.Println(t1, t2, t3, t4)
	// Infinite For Loop
	// for {
	// 	fmt.Printf("It will print Infinite times.\n")
	// }
	res := func(a, b int) int { // Anonymous function
		return a * b
	}(4, 5)
	fmt.Println(res)

}
