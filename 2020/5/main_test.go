package main

import (
	"fmt"
	"testing"
)

func TestBSPOf(t *testing.T) {
	testMap := map[string]int{
		"FBFBBFFRLR": 357,
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
		"BBFFBBFLLL": 816,
	}
	for k, v := range testMap {
		r := bspOf(k[:7], rows)
		c := bspOf(k[7:], columns)
		// seat ID: multiply the row by 8, then add the column
		id := r*8 + c
		if id != v {
			fmt.Println(k, v, id)
		}
	}
}
