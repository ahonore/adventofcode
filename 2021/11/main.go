package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	type point struct {
		x, y int
	}
	deltas := []point{
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
	}
	var grid [10][10]int
	ct := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		for i, v := range b {
			grid[ct][i] = int(v - 48)
		}
		ct++
	}
	// for _, l := range grid {
	// 	fmt.Println(l)
	// }
	// fmt.Println()
	sumFlashes := 0
	for steps := 0; steps < 100; steps++ {
		var flashStepGrid [10][10]bool
		// inc grid by 1
		for j := 0; j < 10; j++ {
			for i := 0; i < 10; i++ {
				grid[j][i]++
			}
		}
		// flashes
		stepFlashes := 0
		oldFlashes := -1
		for {
			if oldFlashes == stepFlashes {
				// set flashers to 0
				for j := 0; j < 10; j++ {
					for i := 0; i < 10; i++ {
						if flashStepGrid[j][i] {
							grid[j][i] = 0
						}
					}
				}
				break
			}

			oldFlashes = stepFlashes
			for j := 0; j < 10; j++ {
				for i := 0; i < 10; i++ {
					if grid[j][i] > 9 && !flashStepGrid[j][i] {
						flashStepGrid[j][i] = true
						stepFlashes++
						grid[j][i] = 0
						for _, d := range deltas {
							cx := i + d.x
							cy := j + d.y
							if cx < 0 || cx >= 10 || cy < 0 || cy >= 10 {
								// not in the grid
								continue
							}

							grid[cy][cx]++
						}
					}
				}
			}
		}
		sumFlashes += stepFlashes
		// for _, l := range grid {
		// 	fmt.Println(l)
		// }
		// fmt.Println()
	}
	fmt.Println(sumFlashes)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	type point struct {
		x, y int
	}
	deltas := []point{
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
	}
	var grid [10][10]int
	ct := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		for i, v := range b {
			grid[ct][i] = int(v - 48)
		}
		ct++
	}
	// for _, l := range grid {
	// 	fmt.Println(l)
	// }
	// fmt.Println()
	steps := 0
	for {
		// inc grid by 1
		for j := 0; j < 10; j++ {
			for i := 0; i < 10; i++ {
				grid[j][i]++
			}
		}
		// flashes
		var flashStepGrid [10][10]bool
		stepFlashes := 0
		oldFlashes := -1
		for {
			if oldFlashes == stepFlashes {
				// set flashers to 0
				for j := 0; j < 10; j++ {
					for i := 0; i < 10; i++ {
						if flashStepGrid[j][i] {
							grid[j][i] = 0
						}
					}
				}
				break
			}

			oldFlashes = stepFlashes
			for j := 0; j < 10; j++ {
				for i := 0; i < 10; i++ {
					if grid[j][i] > 9 && !flashStepGrid[j][i] {
						flashStepGrid[j][i] = true
						stepFlashes++
						grid[j][i] = 0
						for _, d := range deltas {
							cx := i + d.x
							cy := j + d.y
							if cx < 0 || cx >= 10 || cy < 0 || cy >= 10 {
								// not in the grid
								continue
							}

							grid[cy][cx]++
						}
					}
				}
			}
		}
		// for _, l := range grid {
		// 	fmt.Println(l)
		// }
		// fmt.Println()

		// flashed all together?
		allFlashed := true
		for j := 0; j < 10; j++ {
			for i := 0; i < 10; i++ {
				if !flashStepGrid[j][i] {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			break
		}
		steps++
	}
	fmt.Println(steps + 1)
}

func main() {
	f1()
	f2()
}
