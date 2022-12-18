package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		number = append(number, rand.Intn(1000))
	}
	fmt.Println("before: ", number)
	for i := 0; i < len(number); i++ {
		for j := 0; j < len(number)-i-1; j++ {
			if number[j] > number[j+1] {
				number[j], number[j+1] = number[j+1], number[j]
			}
		}

	}
	fmt.Println("after:", number)
}
