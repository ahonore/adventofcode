package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var (
	openToClose = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	closeToOpen = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	illegalCharValues = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	illegalChars := []rune{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		stack := []rune{} // we add at 0 and remove from 0
		for _, r := range txt {
			if _, ok := openToClose[r]; ok {
				stack = append([]rune{r}, stack...)
				continue
			}

			if len(stack) == 0 {
				illegalChars = append(illegalChars, r)
				break
			}

			closeChar := closeToOpen[r]
			if stack[0] != closeChar {
				illegalChars = append(illegalChars, r)
				break
			}

			stack = stack[1:]
		}
	}
	sum := 0
	for _, r := range illegalChars {
		sum += illegalCharValues[r]
	}
	fmt.Println(sum)
}

var (
	closeCharPointValues = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	closingSequenceScores := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		stack := []rune{} // we add at 0 and remove from 0
		foundIllegalChar := false
		for _, r := range txt {
			if _, ok := openToClose[r]; ok {
				stack = append([]rune{r}, stack...)
				continue
			}

			closeChar := closeToOpen[r]
			if stack[0] != closeChar {
				foundIllegalChar = true
				break
			}

			stack = stack[1:]
		}
		if foundIllegalChar || len(stack) == 0 {
			continue
		}

		// must fix chars
		score := 0
		for len(stack) != 0 {
			r := stack[0]
			rc := openToClose[r]
			v := closeCharPointValues[rc]
			score *= 5
			score += v
			stack = stack[1:]
		}
		closingSequenceScores = append(closingSequenceScores, score)
	}
	sort.Ints(closingSequenceScores)
	fmt.Println(closingSequenceScores[len(closingSequenceScores)/2])
}

func main() {
	f1()
	f2()
}
