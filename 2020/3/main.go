package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	var grid []string
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		grid = append(grid, string(l))
	}

	prodTrees := 1
	for _, dir := range []struct {
		x, y int
	}{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	} {
		trees := 0
		x, y := 0, 0
		for y < len(grid) {
			if grid[y][x] == '#' {
				trees++
			}

			x = (x + dir.x + len(grid[y])) % len(grid[y])
			y += dir.y
		}
		prodTrees *= trees
	}
	fmt.Println(prodTrees)
}
