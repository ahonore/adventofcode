package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	valuesGameResultOppChoice := map[byte]map[byte]int{
		// must lose
		'X': {
			'A': 3,
			'B': 1,
			'C': 2,
		},
		// must draw
		'Y': {
			'A': 3 + 1,
			'B': 3 + 2,
			'C': 3 + 3,
		},
		// must win
		'Z': {
			'A': 6 + 2,
			'B': 6 + 3,
			'C': 6 + 1,
		},
	}
	var totalScore int
	for scanner.Scan() {
		gameTurn := scanner.Text()
		totalScore += valuesGameResultOppChoice[gameTurn[2]][gameTurn[0]]
	}
	fmt.Println(totalScore)
}
