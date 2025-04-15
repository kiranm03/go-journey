package main

import (
	"fmt"
	"strings"
)

func main() {
	name, score := "Kiran", 85

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(name, ":", score)
}