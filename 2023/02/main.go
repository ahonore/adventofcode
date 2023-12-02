package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

var maxColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var sumGames int
	for scanner.Scan() {
		line := scanner.Text()
		var gameId int
		fullGamesStr := strings.Split(line, ":")
		fmt.Sscanf(fullGamesStr[0], "Game %d", &gameId)
		gamesStr := strings.Split(fullGamesStr[1], ";")
		var impossibleGame bool
		for _, s := range gamesStr {
			colorsStr := strings.Split(strings.TrimSpace(s), ",")
			for _, c := range colorsStr {
				var color string
				var count int
				fmt.Sscanf(strings.TrimSpace(c), "%d %s", &count, &color)
				if count > maxColors[color] {
					impossibleGame = true
				}
			}
			if impossibleGame {
				break
			}
		}
		if !impossibleGame {
			sumGames += gameId
		}
	}
	fmt.Println(sumGames)
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var sumGames int
	for scanner.Scan() {
		line := scanner.Text()
		minColors := map[string]int{}
		var gameId int
		fullGamesStr := strings.Split(line, ":")
		fmt.Sscanf(fullGamesStr[0], "Game %d", &gameId)
		gamesStr := strings.Split(fullGamesStr[1], ";")
		for _, s := range gamesStr {
			colorsStr := strings.Split(strings.TrimSpace(s), ",")
			for _, c := range colorsStr {
				var color string
				var count int
				fmt.Sscanf(strings.TrimSpace(c), "%d %s", &count, &color)
				if count > minColors[color] {
					minColors[color] = count
				}
			}
		}
		power := 1
		for _, v := range minColors {
			power *= v
		}
		sumGames += power
	}
	fmt.Println(sumGames)
}

func main() {
	part1()
	part2()
}
