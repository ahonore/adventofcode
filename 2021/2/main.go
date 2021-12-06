package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	horizontalPos := 0
	depth := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var cmd string
		var i int
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &cmd, &i)
		if err != nil {
			log.Fatal(err)
		}

		switch cmd {
		case "forward":
			horizontalPos += i
		case "down":
			depth += i
		case "up":
			depth -= i
		}
	}

	fmt.Println(horizontalPos * depth)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	horizontalPos := 0
	depth := 0
	aim := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var cmd string
		var i int
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &cmd, &i)
		if err != nil {
			log.Fatal(err)
		}

		switch cmd {
		case "forward":
			horizontalPos += i
			depth += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		}
	}

	fmt.Println(horizontalPos * depth)
}

func main() {
	f1()
	f2()
}
