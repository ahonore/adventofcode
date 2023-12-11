package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type point struct {
	x, y int
}

// moves allowed from starting point
var deltas = []point{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func expandUniverse(grid [][]byte) [][]byte {
	var rowsId []int
	for i := 0; i < len(grid); i++ {
		if bytes.Contains(grid[i], []byte{'#'}) {
			continue
		}
		rowsId = append(rowsId, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rowsId)))
	for _, i := range rowsId {
		grid = append(grid[:i+1], grid[i:]...)
	}
	var columnsId []int
	for i := 0; i < len(grid[0]); i++ {
		colId := -1
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '#' {
				colId = i
				break
			}
		}
		if colId != -1 {
			continue
		}
		columnsId = append(columnsId, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(columnsId)))
	for _, id := range columnsId {
		for i := 0; i < len(grid); i++ {
			grid[i] = append(grid[i][:id+1], grid[i][id:]...)
		}
	}
	return grid
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var grid [][]byte
	for scanner.Scan() {
		gridLine := []byte(scanner.Text())
		grid = append(grid, gridLine)
	}
	grid = expandUniverse(grid)
	// find galaxies
	var galaxies []point
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == '.' {
				continue
			}
			galaxies = append(galaxies, point{i, j})
		}
	}
	var sumDist int
	for _, g := range galaxies {
		pointsMemory := map[point]bool{
			g: true,
		}
		curPts := []point{g}
		grid[g.y][g.x] = '.'
		var ctDist int
		for len(curPts) > 0 {
			var newPts []point
			ctDist++
			for _, p := range curPts {
				for _, d := range deltas {
					curPt := point{p.x + d.x, p.y + d.y}
					if curPt.x < 0 || curPt.x >= len(grid[0]) || curPt.y < 0 || curPt.y >= len(grid) {
						continue
					}
					if pointsMemory[curPt] {
						continue
					}
					if grid[curPt.y][curPt.x] == '#' {
						sumDist += ctDist
					}
					pointsMemory[curPt] = true
					newPts = append(newPts, curPt)
				}
			}
			curPts = newPts
		}
	}
	fmt.Println(sumDist)
}

func findGalaxiesAfterExpandingUniverse(grid [][]byte, shift int) []point {
	// find galaxies
	var galaxies []point
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == '.' {
				continue
			}
			galaxies = append(galaxies, point{i, j})
		}
	}
	// find empty rows
	var rowsId []int
	for i := 0; i < len(grid); i++ {
		if bytes.Contains(grid[i], []byte{'#'}) {
			continue
		}
		rowsId = append(rowsId, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rowsId)))
	for _, i := range rowsId {
		for g := 0; g < len(galaxies); g++ {
			if galaxies[g].y >= i {
				galaxies[g].y += shift - 1
			}
		}
	}
	// find empty columns
	var columnsId []int
	for i := 0; i < len(grid[0]); i++ {
		colId := -1
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '#' {
				colId = i
				break
			}
		}
		if colId != -1 {
			continue
		}
		columnsId = append(columnsId, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(columnsId)))
	for _, id := range columnsId {
		for g := 0; g < len(galaxies); g++ {
			if galaxies[g].x >= id {
				galaxies[g].x += shift - 1
			}
		}
	}
	return galaxies
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var grid [][]byte
	for scanner.Scan() {
		gridLine := []byte(scanner.Text())
		grid = append(grid, gridLine)
	}
	galaxies := findGalaxiesAfterExpandingUniverse(grid, 1000000)
	var sumDist int
	for i := 0; i < len(galaxies)-1; i++ {
		for d := i + 1; d < len(galaxies); d++ {
			sumDist += int(math.Abs(float64(galaxies[d].x)-float64(galaxies[i].x)) + math.Abs(float64(galaxies[d].y)-float64(galaxies[i].y)))
		}
	}
	fmt.Println(sumDist)
}

func main() {
	part1()
	part2()
}
