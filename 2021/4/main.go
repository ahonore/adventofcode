package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gridValue struct {
	n     int
	drawn bool
}

type grid struct {
	values [][]gridValue
}

func (g *grid) AddNumLine(num []int) {
	line := []gridValue{}
	for _, i := range num {
		line = append(line, gridValue{
			n: i,
		})
	}
	g.values = append(g.values, line)
}

func (g *grid) checkWin(x, y int) bool {
	bL := true
	bC := true
	for i := 0; i < len(g.values); i++ {
		bL = bL && g.values[y][i].drawn
		bC = bC && g.values[i][x].drawn
	}
	return bC || bL
}

// returns true if won
func (g *grid) NextDrawNumber(n int) bool {
	for j := 0; j < len(g.values); j++ {
		for i := 0; i < len(g.values); i++ {
			if g.values[j][i].n == n {
				g.values[j][i].drawn = true
				if g.checkWin(i, j) {
					return true
				}
			}
		}
	}
	return false
}

func (g *grid) ComputeSum() int {
	v := 0
	for j := 0; j < len(g.values); j++ {
		for i := 0; i < len(g.values); i++ {
			if !g.values[j][i].drawn {
				v += g.values[j][i].n
			}
		}
	}
	return v
}

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// reading draw numbers
	scanner.Scan()
	txt := scanner.Text()
	stxt := strings.Split(txt, ",")
	drawNum := []int{}
	for _, n := range stxt {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}

		drawNum = append(drawNum, i)
	}

	scanner.Scan()
	txt = scanner.Text()
	if txt != "" {
		log.Fatal("must be a nl")
	}

	grids := []grid{}
	var g grid
	for {
		if !scanner.Scan() {
			grids = append(grids, g)
			break
		}

		txt = scanner.Text()
		if txt == "" {
			grids = append(grids, g)
			g = grid{}
			continue
		}

		stxt = strings.Fields(txt)
		line := []int{}
		for _, n := range stxt {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}

			line = append(line, i)
		}
		g.AddNumLine(line)
	}

	// fmt.Println(len(grids), grids)

	// draw
	for _, i := range drawNum {
		for _, gn := range grids {
			if gn.NextDrawNumber(i) {
				sum := gn.ComputeSum()
				fmt.Println(i, sum, i*sum)
				return
			}
		}
	}
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// reading draw numbers
	scanner.Scan()
	txt := scanner.Text()
	stxt := strings.Split(txt, ",")
	drawNum := []int{}
	for _, n := range stxt {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}

		drawNum = append(drawNum, i)
	}

	scanner.Scan()
	txt = scanner.Text()
	if txt != "" {
		log.Fatal("must be a nl")
	}

	grids := []grid{}
	var g grid
	for {
		if !scanner.Scan() {
			grids = append(grids, g)
			break
		}

		txt = scanner.Text()
		if txt == "" {
			grids = append(grids, g)
			g = grid{}
			continue
		}

		stxt = strings.Fields(txt)
		line := []int{}
		for _, n := range stxt {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}

			line = append(line, i)
		}
		g.AddNumLine(line)
	}

	// draw
	for _, i := range drawNum {
		idxGridsToRemove := []int{}
		for id, gn := range grids {
			if gn.NextDrawNumber(i) {
				idxGridsToRemove = append(idxGridsToRemove, id)
			}
		}
		if len(grids) == 1 && len(idxGridsToRemove) == 1 {
			fmt.Println(grids)
			sum := grids[0].ComputeSum()
			fmt.Println(i, sum, i*sum)
			return
		}

		for c := len(idxGridsToRemove) - 1; c >= 0; c-- {
			id := idxGridsToRemove[c]
			grids = append(grids[:id], grids[id+1:]...)
		}
	}
}

func main() {
	f1()
	f2()
}
