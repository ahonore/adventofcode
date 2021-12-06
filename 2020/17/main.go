package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y, z int
}

const cycles = 6

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	cubes := map[point]bool{}
	// fill the map (at z = 0)
	buf := bufio.NewReader(f)
	j := 0
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		entry := string(l)
		for i, c := range entry {
			v := false
			if c == '#' {
				v = true
			}

			cubes[point{x: i, y: j, z: 0}] = v
		}
		j++
	}

	for cycle := 0; cycle < cycles; cycle++ {
		newCubes := map[point]bool{}
		for p, enabled := range cubes {
			c := 0
			for _, k := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					for _, i := range []int{-1, 0, 1} {
						if i == 0 && j == 0 && k == 0 {
							continue
						}

						pNeighbour := point{
							x: p.x + i,
							y: p.y + j,
							z: p.z + k,
						}
						enabledNeighbour, ok := cubes[pNeighbour]
						// if !ok {
						// 	enabledNeighbour = false
						// }

						if ok {
							if enabledNeighbour {
								c++
							}

						}

						// newCubes[pNeighbour] = enabledNeighbour
					}
				}
			}
			if enabled && c != 2 && c != 3 {
				newCubes[p] = false
			}

			if !enabled && c == 3 {
				newCubes[p] = true
			}
		}
		cubes = newCubes
	}

	// count enabled cubes
	c := 0
	for _, enabled := range cubes {
		if enabled {
			c++
		}
	}
	fmt.Println(c)
}
