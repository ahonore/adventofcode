package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	_ "embed"
)

//go:embed input.txt
var input string

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		var sb strings.Builder
		for _, c := range line {
			if unicode.IsDigit(c) {
				sb.WriteRune(c)
			}
		}
		s := sb.String()
		i, _ := strconv.Atoi(string(s[0]) + string(s[len(s)-1]))
		sum += i
	}
	fmt.Println("part1", sum)
}

var words2Digit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		// first occ
		id := -1
		var firstDigit int
		curId := -1
		for w, d := range words2Digit {
			curId = strings.Index(line, w)
			if curId >= 0 && (curId < id || id == -1) {
				id = curId
				firstDigit = d
			}
		}
		// last occ
		id = -1
		var lastDigit int
		curId = -1
		for w, d := range words2Digit {
			curId = strings.LastIndex(line, w)
			if curId >= 0 && (curId > id || id == -1) {
				id = curId
				lastDigit = d
			}
		}
		i := firstDigit*10 + lastDigit
		sum += i
	}
	fmt.Println("part2", sum)
}

func main() {
	part1()
	part2()
}
