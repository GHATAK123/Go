package main

import (
	"math/rand"
	"time"
)

func random() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(6) + 1
	return randomNumber

}
