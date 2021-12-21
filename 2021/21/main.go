package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

var deterministicDice100 = 0

func nextDiceRoll() int {
	ret := deterministicDice100
	deterministicDice100 = (deterministicDice100 + 1) % 100
	return ret + 1
}

func playerMove(p int, move int) int {
	p--
	p = (p + move) % 10
	p++
	return p
}

func f1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	// player 1
	scanner.Scan()
	var p1 int
	fmt.Sscanf(scanner.Text(), "Player 1 starting position: %d", &p1)
	// player 2
	scanner.Scan()
	var p2 int
	fmt.Sscanf(scanner.Text(), "Player 2 starting position: %d", &p2)

	var scoreP1, scoreP2, countRolls int
	p1Win := false
	for {
		rolls := nextDiceRoll() + nextDiceRoll() + nextDiceRoll()
		countRolls += 3
		p1 = playerMove(p1, rolls)
		scoreP1 += p1
		if scoreP1 >= 1000 {
			p1Win = true
			break
		}

		rolls = nextDiceRoll() + nextDiceRoll() + nextDiceRoll()
		countRolls += 3
		p2 = playerMove(p2, rolls)
		scoreP2 += p2
		if scoreP2 >= 1000 {
			break
		}
	}
	finalCount := 0
	if p1Win {
		finalCount = scoreP2 * countRolls
	} else {
		finalCount = scoreP1 * countRolls
	}

	return finalCount
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1(bytes.NewReader(b)))
	// fmt.Println(f2(bytes.NewReader(b)))
}
