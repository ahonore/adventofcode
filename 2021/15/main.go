package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

func f1(r io.Reader) int {
	grid := [][]int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		line := make([]int, len(txt))
		for i, c := range txt {
			line[i] = int(c - 48)
		}
		grid = append(grid, line)
	}

	type point struct {
		x, y int
	}
	type node struct {
		p         point
		cost      int
		heuristic int
	}
	target := point{x: len(grid[0]) - 1, y: len(grid) - 1}
	closedList := map[point]struct{}{}
	openList := []node{{p: point{x: 0, y: 0}, cost: 0, heuristic: 0}}
	res := 0
	for {
		current := openList[0]
		openList = openList[1:]
		if current.p == target {
			res = current.cost
			break
		}

		for _, d := range []point{
			{x: 0, y: -1},
			{x: 0, y: 1},
			{x: -1, y: 0},
			{x: 1, y: 0},
		} {
			newPoint := point{x: current.p.x + d.x, y: current.p.y + d.y}
			if newPoint.x < 0 || newPoint.x >= len(grid[0]) || newPoint.y < 0 || newPoint.y >= len(grid) {
				continue
			}

			if _, ok := closedList[newPoint]; ok {
				continue
			}

			cost := current.cost + grid[newPoint.y][newPoint.x]
			found := false
			for _, pol := range openList {
				if pol.p == newPoint && pol.cost <= cost {
					found = true
					break
				}
			}
			if found {
				continue
			}

			dist := int(math.Abs(float64(newPoint.x - target.x)))
			dist += int(math.Abs(float64(newPoint.y - target.y)))
			openList = append(openList, node{p: newPoint, cost: cost, heuristic: cost + dist})
		}
		closedList[current.p] = struct{}{}
		sort.Slice(openList, func(i, j int) bool { return openList[i].heuristic < openList[j].heuristic })
	}
	return res
}

type extendedGrid struct {
	grid [][]int
}

func (eg extendedGrid) Cost(x, y int) int {
	gridNumX := x / len(eg.grid[0])
	gridNumY := y / len(eg.grid)
	costX := x % len(eg.grid[0])
	costY := y % len(eg.grid)
	cost := eg.grid[costY][costX] + gridNumX + gridNumY
	if cost > 9 {
		cost = cost%10 + 1
	}

	return cost
}

func (eg extendedGrid) Len() (int, int) {
	return len(eg.grid[0]) * 5, len(eg.grid) * 5
}

func f2(r io.Reader) int {
	cave := [][]int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		line := make([]int, len(txt))
		for i, c := range txt {
			line[i] = int(c - 48)
		}
		cave = append(cave, line)
	}
	g := extendedGrid{grid: cave}
	lenX, lenY := g.Len()

	type point struct {
		x, y int
	}
	type node struct {
		p         point
		cost      int
		heuristic int
	}
	target := point{x: lenX - 1, y: lenY - 1}
	closedList := map[point]struct{}{}
	openList := []node{{p: point{x: 0, y: 0}, cost: 0, heuristic: 0}}
	res := 0
	for {
		current := openList[0]
		openList = openList[1:]
		if current.p == target {
			res = current.cost
			break
		}

		for _, d := range []point{
			{x: 0, y: -1},
			{x: 0, y: 1},
			{x: -1, y: 0},
			{x: 1, y: 0},
		} {
			newPoint := point{x: current.p.x + d.x, y: current.p.y + d.y}
			if newPoint.x < 0 || newPoint.x >= lenX || newPoint.y < 0 || newPoint.y >= lenY {
				continue
			}

			if _, ok := closedList[newPoint]; ok {
				continue
			}

			cost := current.cost + g.Cost(newPoint.x, newPoint.y)
			found := false
			for _, pol := range openList {
				if pol.p == newPoint && pol.cost <= cost {
					found = true
					break
				}
			}
			if found {
				continue
			}

			dist := int(math.Abs(float64(newPoint.x - target.x)))
			dist += int(math.Abs(float64(newPoint.y - target.y)))
			openList = append(openList, node{p: newPoint, cost: cost, heuristic: cost + dist})
		}
		closedList[current.p] = struct{}{}
		sort.Slice(openList, func(i, j int) bool { return openList[i].heuristic < openList[j].heuristic })
	}
	return res
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
