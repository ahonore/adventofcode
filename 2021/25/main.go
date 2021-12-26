package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func f1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	grid := [][]int{}
	for scanner.Scan() {
		line := []int{}
		for _, c := range scanner.Text() {
			v := 0
			switch c {
			case '>':
				v = 1
			case 'v':
				v = 2
			}
			line = append(line, v)
		}
		grid = append(grid, line)
	}
	steps := 0
	for {
		hasMoved := false
		// east-facing
		for j := 0; j < len(grid); j++ {
			moves := make([]bool, len(grid[0]))
			for i := 0; i < len(grid[0]); i++ {
				if grid[j][i] == 1 && grid[j][(i+1)%len(grid[0])] == 0 {
					moves[i] = true
				}
			}
			for i := 0; i < len(grid[0]); i++ {
				if grid[j][i] == 1 && moves[i] {
					hasMoved = true
					grid[j][(i+1)%len(grid[0])] = 1
					grid[j][i] = 0
					i++
				}
			}
		}
		// south-facing
		for i := 0; i < len(grid[0]); i++ {
			moves := make([]bool, len(grid))
			for j := 0; j < len(grid); j++ {
				if grid[j][i] == 2 && grid[(j+1)%len(grid)][i] == 0 {
					moves[j] = true
				}
			}
			for j := 0; j < len(grid); j++ {
				if grid[j][i] == 2 && moves[j] {
					hasMoved = true
					grid[(j+1)%len(grid)][i] = 2
					grid[j][i] = 0
					j++
				}
			}
		}
		steps++
		if !hasMoved {
			break
		}
	}
	return steps
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
}
