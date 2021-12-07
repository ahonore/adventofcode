package main

import (
	"fmt"
	"io"
	"log"
	"math"
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

	min, max := 0, 0
	strValues := strings.Split(strings.TrimSpace(string(b)), ",")
	values := []int{}
	for _, v := range strValues {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		if i < min {
			min = i
		}

		if i > max {
			max = i
		}

		values = append(values, i)
	}
	// fmt.Println(min, max, values)
	pos := 0
	cost := math.MaxInt
	for i := min; i <= max; i++ {
		curCost := 0
		for _, v := range values {
			curCost += int(math.Abs(float64(v - i)))
		}
		if curCost < cost {
			pos = i
			cost = curCost
		}
	}
	fmt.Println(pos, cost)
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

	min, max := 0, 0
	strValues := strings.Split(strings.TrimSpace(string(b)), ",")
	values := []int{}
	for _, v := range strValues {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		if i < min {
			min = i
		}

		if i > max {
			max = i
		}

		values = append(values, i)
	}
	// fmt.Println(min, max, values)
	pos := 0
	cost := math.MaxInt
	for i := min; i <= max; i++ {
		curCost := 0
		for _, v := range values {
			n := int(math.Abs(float64(v - i)))
			curCost += n * (n + 1) / 2
		}
		if curCost < cost {
			pos = i
			cost = curCost
		}
	}
	fmt.Println(pos, cost)
}

func main() {
	f1()
	f2()
}
