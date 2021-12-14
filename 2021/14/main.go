package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// read template at first line
	scanner.Scan()
	template := scanner.Text()
	// read newline
	scanner.Scan()
	scanner.Text()
	// read rules
	rules := map[string]byte{}
	for scanner.Scan() {
		var rule string
		var b byte
		if _, err := fmt.Sscanf(scanner.Text(), "%s -> %c", &rule, &b); err != nil {
			log.Fatal(err)
		}

		rules[rule] = b
	}
	// run template
	for step := 0; step < 10; step++ {
		var sb strings.Builder
		for i := 1; i < len(template); i++ {
			b := rules[template[i-1:i+1]]
			sb.WriteByte(template[i-1])
			sb.WriteByte(b)
		}
		sb.WriteByte(template[len(template)-1])
		template = sb.String()
	}
	// count
	countChars := map[rune]int{}
	for _, c := range template {
		countChars[c]++
	}
	// find min max
	min := math.MaxInt
	max := 0
	for _, v := range countChars {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// read template at first line
	scanner.Scan()
	template := scanner.Text()
	// read newline
	scanner.Scan()
	scanner.Text()
	// read rules
	rules := map[string]byte{}
	for scanner.Scan() {
		var rule string
		var b byte
		if _, err := fmt.Sscanf(scanner.Text(), "%s -> %c", &rule, &b); err != nil {
			log.Fatal(err)
		}

		rules[rule] = b
	}
	// template init countPairs and countChars
	countPairs := map[string]int64{}
	countChars := map[byte]int64{}
	for i := 1; i < len(template); i++ {
		countPairs[template[i-1:i+1]]++
		countChars[template[i-1]]++
	}
	countChars[template[len(template)-1]]++
	// run steps
	for step := 0; step < 40; step++ {
		newCountPairs := map[string]int64{}
		for pair, count := range countPairs {
			b := rules[pair]
			countChars[b] += count // new char added count times
			newPair0 := string([]byte{pair[0], b})
			newCountPairs[newPair0] += count
			newPair1 := string([]byte{b, pair[1]})
			newCountPairs[newPair1] += count
		}
		countPairs = newCountPairs
	}
	// find min max
	min := int64(math.MaxInt64)
	max := int64(0)
	for _, v := range countChars {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func main() {
	f1()
	f2()
}
