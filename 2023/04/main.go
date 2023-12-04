package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	space := regexp.MustCompile(`\s+`)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		fullGamesStr := strings.Split(line, ":")
		numStr := strings.Split(fullGamesStr[1], "|")
		winningNumbers := map[int]bool{}
		for _, s := range strings.Split(space.ReplaceAllString(strings.TrimSpace(numStr[0]), " "), " ") {
			i, _ := strconv.Atoi(s)
			winningNumbers[i] = true
		}
		ownNumbers := map[int]bool{}
		for _, s := range strings.Split(space.ReplaceAllString(strings.TrimSpace(numStr[1]), " "), " ") {
			i, _ := strconv.Atoi(s)
			ownNumbers[i] = true
		}
		var ct int
		for k := range winningNumbers {
			if ownNumbers[k] {
				ct++
			}
		}
		if ct > 0 {
			sum += 1 << (ct - 1)
		}
	}
	fmt.Println(sum)
}

type card struct {
	id             int
	ownNumbers     map[int]bool
	winningNumbers map[int]bool
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	space := regexp.MustCompile(`\s+`)
	var cards []card
	for scanner.Scan() {
		line := scanner.Text()
		fullGamesStr := strings.Split(line, ":")
		var cardId int
		fmt.Sscanf(fullGamesStr[0], "Card %d", &cardId)
		numStr := strings.Split(fullGamesStr[1], "|")
		winningNumbers := map[int]bool{}
		for _, s := range strings.Split(space.ReplaceAllString(strings.TrimSpace(numStr[0]), " "), " ") {
			i, _ := strconv.Atoi(s)
			winningNumbers[i] = true
		}
		ownNumbers := map[int]bool{}
		for _, s := range strings.Split(space.ReplaceAllString(strings.TrimSpace(numStr[1]), " "), " ") {
			i, _ := strconv.Atoi(s)
			ownNumbers[i] = true
		}
		cards = append(cards, card{
			id:             cardId,
			ownNumbers:     ownNumbers,
			winningNumbers: winningNumbers,
		})
	}
	var ct int
	for ct < len(cards) {
		// count wins
		var ctWins int
		c := cards[ct]
		for k := range c.winningNumbers {
			if c.ownNumbers[k] {
				ctWins++
			}
		}
		if ctWins > 0 {
			// stack num of wins
			cards = append(cards, cards[c.id:c.id+ctWins]...)
		}

		ct++
	}
	fmt.Println(ct)
}

func main() {
	part1()
	part2()
}
