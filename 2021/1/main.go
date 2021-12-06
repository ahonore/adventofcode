package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	ct := 0
	iOld := math.MaxInt64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if iOld < i {
			ct++
		}
		iOld = i
	}
	fmt.Println(ct)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	id := 0
	measures := []int{}
	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		measures = append(measures, m)
		if id-1 >= 0 {
			measures[len(measures)-2] += m
		}

		if id-2 >= 0 {
			measures[len(measures)-3] += m
		}
		id++
	}

	ct := 0
	for i := 1; i < len(measures); i++ {
		if measures[i-1] < measures[i] {
			ct++
		}
	}
	fmt.Println(ct)
}

func main() {
	f1()
	f2()
}
