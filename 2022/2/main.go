package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	values := map[rune]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	var totalScore int
	for scanner.Scan() {
		gameTurn := scanner.Text()

		switch gameTurn {
		case "A X", "B Y", "C Z":
			totalScore += 3 + values[rune(gameTurn[2])]

		case "A Z", "C Y", "B X":
			totalScore += 0 + values[rune(gameTurn[2])]

		case "A Y", "B Z", "C X":
			totalScore += 6 + values[rune(gameTurn[2])]
		}
	}
	fmt.Println(totalScore)
}
