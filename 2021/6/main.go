package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	strValues := strings.Split(strings.TrimSpace(string(b)), ",")
	// list of numbers of fish by index (index = age)
	values := make([]int, 9)
	for _, v := range strValues {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		values[i]++
	}
	// fmt.Println(values)
	for gen := 0; gen < 80; gen++ {
		newValues := make([]int, 9)
		for i := 1; i < 9; i++ {
			newValues[i-1] = values[i]
		}
		newValues[6] += values[0]
		newValues[8] += values[0]
		values = newValues
		// fmt.Println(values)
	}
	sum := 0
	for _, i := range values {
		sum += i
	}
	fmt.Println(sum)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	strValues := strings.Split(strings.TrimSpace(string(b)), ",")
	// list of numbers of fish by index (index = age)
	values := make([]int, 9)
	for _, v := range strValues {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		values[i]++
	}
	// fmt.Println(values)
	for gen := 0; gen < 256; gen++ {
		newValues := make([]int, 9)
		for i := 1; i < 9; i++ {
			newValues[i-1] = values[i]
		}
		newValues[6] += values[0]
		newValues[8] += values[0]
		values = newValues
		// fmt.Println(values)
	}
	sum := 0
	for _, i := range values {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	f1()
	f2()
}
