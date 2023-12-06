package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type convMap struct {
	destStart int
	srcStart  int
	length    int
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// seeds
	var seeds []int
	scanner.Scan()
	line := scanner.Text()
	seedsStr := strings.TrimSpace(strings.Split(line, ":")[1])
	for _, v := range strings.Split(seedsStr, " ") {
		i, _ := strconv.Atoi(v)
		seeds = append(seeds, i)
	}
	scanner.Scan() // bypass newline

	// maps
	var convMaps [][]convMap
	for scanner.Scan() { // map name string
		var currConvMaps []convMap
		for {
			scan := scanner.Scan()
			if !scan {
				convMaps = append(convMaps, currConvMaps)
				break
			}
			line := scanner.Text()
			if line == "" {
				convMaps = append(convMaps, currConvMaps)
				break
			}
			var curConvMap convMap
			fmt.Sscan(line, &curConvMap.destStart, &curConvMap.srcStart, &curConvMap.length)
			currConvMaps = append(currConvMaps, curConvMap)
		}
	}

	minLocation := math.MaxInt
	for _, seed := range seeds {
		convValue := seed
		for _, convMap := range convMaps {
			for _, rangeMap := range convMap {
				if convValue >= rangeMap.srcStart && convValue < (rangeMap.srcStart+rangeMap.length) {
					convValue = convValue - rangeMap.srcStart + rangeMap.destStart
					break
				}
			}
		}
		if convValue < minLocation {
			minLocation = convValue
		}
	}
	fmt.Println(minLocation)
}

type seedRange struct {
	start  int
	length int
}

func findMinLocation(sr seedRange, idTable, idMap int, convMap [][]convMap) int {
	if idTable >= len(convMap) {
		return sr.start
	}

	if idMap >= len(convMap[idTable]) {
		return findMinLocation(sr, idTable+1, 0, convMap)
	}

	curConvMap := convMap[idTable][idMap]
	if (sr.start + sr.length) <= curConvMap.srcStart { // seed range lower than the conv map line
		return findMinLocation(sr, idTable, idMap+1, convMap)
	}

	if sr.start >= (curConvMap.srcStart + curConvMap.length) { // seed range higher than the convMap line
		return findMinLocation(sr, idTable, idMap+1, convMap)
	}

	maxStart := sr.start
	if maxStart < curConvMap.srcStart {
		maxStart = curConvMap.srcStart
	}

	minEnd := sr.start + sr.length
	if minEnd > (curConvMap.srcStart + curConvMap.length) {
		minEnd = curConvMap.srcStart + curConvMap.length
	}

	minLocation := findMinLocation(
		seedRange{
			start:  maxStart + curConvMap.destStart - curConvMap.srcStart,
			length: minEnd - maxStart,
		},
		idTable+1,
		0,
		convMap,
	)

	// low part
	lowPart := maxStart - sr.start
	if lowPart > 0 {
		l := findMinLocation(
			seedRange{
				start:  sr.start,
				length: lowPart,
			},
			idTable,
			idMap+1,
			convMap,
		)
		if l < minLocation {
			minLocation = l
		}
	}
	// high part
	highPart := sr.start + sr.length - minEnd
	if highPart > 0 {
		l := findMinLocation(
			seedRange{
				start:  minEnd,
				length: highPart,
			},
			idTable,
			idMap+1,
			convMap,
		)
		if l < minLocation {
			minLocation = l
		}
	}

	return minLocation
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// seeds
	var seeds []seedRange
	scanner.Scan()
	line := scanner.Text()
	seedsStr := strings.TrimSpace(strings.Split(line, ":")[1])
	seedsData := strings.Split(seedsStr, " ")
	for id := 0; id < len(seedsData); id += 2 {
		i0, _ := strconv.Atoi(seedsData[id])
		i1, _ := strconv.Atoi(seedsData[id+1])
		seeds = append(seeds, seedRange{
			start:  i0,
			length: i1,
		})
	}
	scanner.Scan() // bypass newline

	// maps
	var convMaps [][]convMap
	for scanner.Scan() { // map name string
		var currConvMaps []convMap
		for {
			scan := scanner.Scan()
			if !scan {
				convMaps = append(convMaps, currConvMaps)
				break
			}
			line := scanner.Text()
			if line == "" {
				convMaps = append(convMaps, currConvMaps)
				break
			}
			var curConvMap convMap
			fmt.Sscan(line, &curConvMap.destStart, &curConvMap.srcStart, &curConvMap.length)
			currConvMaps = append(currConvMaps, curConvMap)
		}
	}

	minLocation := math.MaxInt
	for _, sr := range seeds {
		v := findMinLocation(sr, 0, 0, convMaps)
		if v < minLocation {
			minLocation = v
		}
	}
	fmt.Println(minLocation)
}

func main() {
	part1()
	part2()
}
