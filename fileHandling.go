package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeFile(content string) {
	file, err := os.Create("./writeFile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	defer file.Close()

}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("This file contains", string(data))
}
