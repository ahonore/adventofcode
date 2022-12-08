package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct{ x, y int }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var trees [][]int
	for scanner.Scan() {
		var treeLine []int
		for _, c := range scanner.Text() {
			treeLine = append(treeLine, int(c)-48)
		}
		trees = append(trees, treeLine)
		// fmt.Println(treeLine)
	}
	var score int
	for j := 0; j < len(trees); j++ {
		for i := 0; i < len(trees[j]); i++ {
			curTree := trees[j][i]
			var curScore int = 1
			for _, d := range []coord{
				{x: -1, y: 0},
				{x: 0, y: -1},
				{x: 0, y: 1},
				{x: 1, y: 0},
			} {
				nt := coord{
					x: i + d.x,
					y: j + d.y,
				}
				var ctTrees int
				for {
					if nt.x < 0 || nt.x >= len(trees[j]) {
						break
					}

					if nt.y < 0 || nt.y >= len(trees) {
						break
					}

					ctTrees++
					if trees[nt.y][nt.x] >= curTree {
						break
					}

					nt.x += d.x
					nt.y += d.y
				}
				curScore *= ctTrees
			}
			if curScore > score {
				score = curScore
			}
		}
	}
	fmt.Println(score)
}
