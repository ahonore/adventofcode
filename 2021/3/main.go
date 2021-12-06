package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// counting bits to 1
	var binOneCounters []int
	ctLines := 0
	for scanner.Scan() {
		txt := scanner.Text()
		ctLines++
		// first time -> init counters
		if binOneCounters == nil {
			binOneCounters = make([]int, len(txt))
		}

		for id, c := range txt {
			if c == '1' {
				binOneCounters[id]++
			}
		}
	}
	bitFieldPivot := 1<<len(binOneCounters) - 1
	gamma := 0
	for id, i := range binOneCounters {
		if i > (ctLines - i) {
			gamma |= 1 << (len(binOneCounters) - 1 - id)
		}
	}
	epsilon := bitFieldPivot - gamma
	fmt.Println("gamma", gamma, "epsilon", epsilon, gamma*epsilon)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// counting bits to 1
	numbers := []string{}
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	oxygenGen := rec(0, numbers, true)
	co2Scrubber := rec(0, numbers, false)
	fmt.Println("oxygenGen", oxygenGen, "co2Scrubber", co2Scrubber, oxygenGen*co2Scrubber)
}

func rec(id int, numbers []string, higherCount bool) int {
	if len(numbers) == 1 {
		i, err := strconv.ParseInt(numbers[0], 2, 32)
		if err != nil {
			log.Fatal(err)
		}

		return int(i)
	}

	ctOne := 0
	ctLines := 0
	for _, s := range numbers {
		if s[id] == '1' {
			ctOne++
		}

		ctLines++
	}
	var charToKeep byte
	if higherCount {
		if ctOne >= (ctLines - ctOne) {
			charToKeep = '1'
		} else {
			charToKeep = '0'
		}
	} else if ctOne >= (ctLines - ctOne) {
		charToKeep = '0'
	} else {
		charToKeep = '1'
	}

	newNumbers := []string{}
	for _, s := range numbers {
		if s[id] == charToKeep {
			newNumbers = append(newNumbers, s)
		}
	}
	return rec(id+1, newNumbers, higherCount)
}

func main() {
	f1()
	f2()
}
