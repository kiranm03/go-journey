package main

import (
	"fmt"
	"strings"
)

func main() {
	name, score := "Kiran", 85

	students := []string{"Dent", "Arthur"}
	scores := []int{50, 75}

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(name, ":", score)
	fmt.Println(students[0], ":", scores[0])
	fmt.Println(students[1], ":", scores[1])
}