package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Taking Input In Go
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tell me about you?")
	intro, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("This will print in last because of defer: ", intro)
	// String To Int conversion
	rating := "45"
	numRating, err := strconv.ParseInt(rating, 10, 64)
	if err != nil {
		log.Fatal(err)
	} else {
		numRating += 1
		fmt.Println("numRating", numRating)
	}
	// Handling Time In Go
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02-01-2006 15:04:05 Monday"))
	createdTime := time.Date(2022, time.February, 16, 10, 10, 10, 10, time.UTC)
	fmt.Println(createdTime.Format("02-01-2006 15:04:05 Monday"))
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
	res := func(a, b int) int { // Anonymous function In Go
		return a * b
	}(4, 5)
	fmt.Println(res)

	if num := 5; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
	// Random Number Generator
	dice := random()
	fmt.Println(dice)

	Variadic := func(values ...int) int { // Variadic Functions in Go
		total := 0
		for _, v := range values {
			total += v
		}
		return total
	}(4, 5, 4, 7, 8, 6, 4)
	fmt.Println(Variadic)
	// File Write In Go
	writeFile(intro)
	// Read File In GO
	readFile("./writeFile.txt")
	// Defer in Go
	myDefer()

}
