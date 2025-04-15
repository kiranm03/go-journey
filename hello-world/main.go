package main

import (
	"fmt"
	"strings"
)

type studentScore struct {
	name string
	score int
}

func main() {

	scores := []studentScore {
		{"Dent", 50},
		{"Arthur", 75},
		{"Kiran", 85},
		addStudent("John", 65),
	}

	printReport(scores)
}

func addStudent(name string, score int) studentScore {
	return studentScore{ name, score }
}

func printReport(scores []studentScore) {
	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 20))

	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.score)
	}
}