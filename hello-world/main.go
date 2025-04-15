package main

import (
	"fmt"
	"strings"
)

func main() {

	type score struct {
		name string
		score int
	}

	scores := []score {
		{"Dent", 50},
		{"Arthur", 75},
		{"Kiran", 85},
	}

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 20))

	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.score)
	}
}