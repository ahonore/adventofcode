package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	heightMap := [][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		values := []int{}
		for _, v := range b {
			i := int(v - 48) // ascii to num conv
			values = append(values, i)
		}
		heightMap = append(heightMap, values)
	}

	lowPoints := []int{}
	for j := 0; j < len(heightMap); j++ {
		for i := 0; i < len(heightMap[0]); i++ {
			v := heightMap[j][i]
			if (j-1) >= 0 && heightMap[j-1][i] <= v {
				continue
			}

			if (j+1) < len(heightMap) && heightMap[j+1][i] <= v {
				continue
			}

			if (i-1) >= 0 && heightMap[j][i-1] <= v {
				continue
			}

			if (i+1) < len(heightMap[0]) && heightMap[j][i+1] <= v {
				continue
			}

			lowPoints = append(lowPoints, v+1)
		}
	}
	sum := 0
	for _, v := range lowPoints {
		sum += v
	}
	fmt.Println(sum)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	heightMap := [][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := []byte(scanner.Text())
		values := []int{}
		for _, v := range b {
			i := int(v - 48) // ascii to num conv
			values = append(values, i)
		}
		heightMap = append(heightMap, values)
	}

	type point struct {
		x, y int
	}
	lowPoints := []point{}
	for j := 0; j < len(heightMap); j++ {
		for i := 0; i < len(heightMap[0]); i++ {
			v := heightMap[j][i]
			if (j-1) >= 0 && heightMap[j-1][i] <= v {
				continue
			}

			if (j+1) < len(heightMap) && heightMap[j+1][i] <= v {
				continue
			}

			if (i-1) >= 0 && heightMap[j][i-1] <= v {
				continue
			}

			if (i+1) < len(heightMap[0]) && heightMap[j][i+1] <= v {
				continue
			}

			lowPoints = append(lowPoints, point{
				x: i,
				y: j,
			})
		}
	}
	// delta points
	deltaPoints := [4]point{
		{
			x: 0,
			y: -1,
		},
		{
			x: 0,
			y: 1,
		},
		{
			x: -1,
			y: 0,
		},
		{
			x: 1,
			y: 0,
		},
	}
	largestBasins := make([]int, len(lowPoints))
	for id, v := range lowPoints {
		basin := map[point]struct{}{
			v: {},
		}
		oldBasinLen := 0
		for {
			if oldBasinLen == len(basin) {
				break
			}

			oldBasinLen = len(basin)
			// select new points
			newPoints := []point{}
			for p := range basin {
				for _, dp := range deltaPoints {
					cp := point{
						x: p.x + dp.x,
						y: p.y + dp.y,
					}
					if cp.x >= 0 && cp.x < len(heightMap[0]) && cp.y >= 0 && cp.y < len(heightMap) && heightMap[cp.y][cp.x] < 9 {
						newPoints = append(newPoints, cp)
					}
				}
			}
			// fmt.Println(newPoints)
			for _, p := range newPoints {
				basin[p] = struct{}{}
			}
		}
		largestBasins[id] = len(basin)
		// fmt.Println(basin)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(largestBasins)))
	// fmt.Println(largestBasins)
	fmt.Println(largestBasins[0] * largestBasins[1] * largestBasins[2])
}

func main() {
	f1()
	f2()
}
