package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rows    = 128
	columns = 8
)

func bspOf(s string, max int) int {
	rowMin, rowMax := 0, max
	for _, c := range s {
		med := (rowMax + rowMin) / 2
		switch c {
		case 'F', 'L':
			rowMax = med
		case 'B', 'R':
			rowMin = med
		default:
			panic(c)
		}
		// fmt.Println(rowMin, rowMax)
	}
	return rowMin
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	ct := 0
	seat, id := "", 0
	buf := bufio.NewReader(f)
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		entry := string(l)
		r := bspOf(entry[:7], rows)
		c := bspOf(entry[7:], columns)
		// seat ID: multiply the row by 8, then add the column
		curID := r*8 + c
		if curID > id {
			seat = entry
			id = curID
		}
		ct++
	}
	fmt.Println(seat, id)
}
