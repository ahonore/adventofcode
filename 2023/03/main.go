package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	_ "embed"
)

//go:embed input.txt
var input string

type point struct {
	x int
	y int
}

var delta = []point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

type gridData struct {
	value  *int
	symbol *rune
}

var reNum = regexp.MustCompile(`(\d)+`)

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	grid := [][]gridData{}
	for scanner.Scan() {
		line := scanner.Text()
		gridLine := make([]gridData, len(line))
		// find digits
		for _, idx := range reNum.FindAllStringIndex(line, -1) {
			i, _ := strconv.Atoi(line[idx[0]:idx[1]])
			for x := idx[0]; x < idx[1]; x++ {
				gridLine[x].value = &i
			}
		}
		// find symbols
		for x, r := range line {
			r := r
			if unicode.IsDigit(r) || (r == '.') {
				continue
			}

			gridLine[x].symbol = &r
		}
		grid = append(grid, gridLine)
	}
	uniqValues := map[*int]struct{}{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].symbol == nil {
				continue
			}

			// found a symbol, look around
			for _, d := range delta {
				xx := x + d.x
				yy := y + d.y
				if xx < 0 || xx >= len(grid[y]) || yy < 0 || yy >= len(grid) {
					continue
				}

				if grid[yy][xx].value != nil {
					uniqValues[grid[yy][xx].value] = struct{}{}
				}
			}
		}
	}
	var sum int
	for k := range uniqValues {
		sum += *k
	}
	fmt.Println(sum)
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	grid := [][]gridData{}
	for scanner.Scan() {
		line := scanner.Text()
		gridLine := make([]gridData, len(line))
		// find digits
		for _, idx := range reNum.FindAllStringIndex(line, -1) {
			i, _ := strconv.Atoi(line[idx[0]:idx[1]])
			for x := idx[0]; x < idx[1]; x++ {
				gridLine[x].value = &i
			}
		}
		// find symbols
		for x, r := range line {
			r := r
			if r != '*' {
				continue
			}

			gridLine[x].symbol = &r
		}
		grid = append(grid, gridLine)
	}
	var sum int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].symbol == nil {
				continue
			}

			// found symbol '*', look around for only 2 close numbers
			uniqValues := map[*int]struct{}{}
			for _, d := range delta {
				xx := x + d.x
				yy := y + d.y
				if xx < 0 || xx >= len(grid[y]) || yy < 0 || yy >= len(grid) {
					continue
				}

				if grid[yy][xx].value == nil {
					continue
				}

				uniqValues[grid[yy][xx].value] = struct{}{}
			}
			if len(uniqValues) != 2 {
				continue
			}

			mult := 1
			for k := range uniqValues {
				mult *= *k
			}
			sum += mult
		}
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
