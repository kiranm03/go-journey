package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "ERROR", "log level")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}