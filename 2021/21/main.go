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
	return (p-1+move)%10 + 1
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

type game struct {
	pos1, pos2     int
	score1, score2 int
}

// a mapping between a triple dice draw and number of universes it creates
var diceDrawUniverses = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func f2(r io.Reader) (int64, int64) {
	scanner := bufio.NewScanner(r)
	// player 1
	scanner.Scan()
	var p1 int
	fmt.Sscanf(scanner.Text(), "Player 1 starting position: %d", &p1)
	// player 2
	scanner.Scan()
	var p2 int
	fmt.Sscanf(scanner.Text(), "Player 2 starting position: %d", &p2)

	curGames := map[game]int64{
		{pos1: p1, pos2: p2}: 1,
	}
	p1Wins, p2Wins := int64(0), int64(0)
	for len(curGames) > 0 {
		// creating the new generation
		newCurGames := map[game]int64{}
		for g, count := range curGames {
			// p1
			for p1Draw, p1Univ := range diceDrawUniverses {
				newCount1 := count * p1Univ
				pos1 := playerMove(g.pos1, p1Draw)
				score1 := g.score1 + pos1
				if score1 >= 21 {
					p1Wins += newCount1
					continue
				}

				// p2
				for p2Draw, p2Univ := range diceDrawUniverses {
					newCount2 := newCount1 * p2Univ
					pos2 := playerMove(g.pos2, p2Draw)
					score2 := g.score2 + pos2
					if score2 >= 21 {
						p2Wins += newCount2
						continue
					}

					newCurGames[game{
						pos1:   pos1,
						pos2:   pos2,
						score1: score1,
						score2: score2,
					}] += newCount2
				}
			}
		}
		curGames = newCurGames
	}
	return p1Wins, p2Wins
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1(bytes.NewReader(b)))
	fmt.Println(f2(bytes.NewReader(b)))
}
