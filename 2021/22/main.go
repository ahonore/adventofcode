package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type cubeInst struct {
	isOn   bool
	x0, x1 int
	y0, y1 int
	z0, z1 int
}

type cube struct {
	x, y, z int
}

func f1(r io.Reader) int {
	bootSequence := []cubeInst{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		cubeInst := cubeInst{}
		var lit string
		fmt.Sscanf(scanner.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d", &lit, &cubeInst.x0, &cubeInst.x1, &cubeInst.y0, &cubeInst.y1, &cubeInst.z0, &cubeInst.z1)
		if lit == "on" {
			cubeInst.isOn = true
		}

		if cubeInst.x0 < -50 ||
			cubeInst.y0 < -50 ||
			cubeInst.z0 < -50 ||
			cubeInst.x1 > 50 ||
			cubeInst.y1 > 50 ||
			cubeInst.z1 > 50 {
			continue
		}

		bootSequence = append(bootSequence, cubeInst)
	}

	cubes := map[cube]struct{}{}
	for _, ci := range bootSequence {
		for k := ci.z0; k <= ci.z1; k++ {
			for j := ci.y0; j <= ci.y1; j++ {
				for i := ci.x0; i <= ci.x1; i++ {
					if ci.isOn {
						cubes[cube{x: i, y: j, z: k}] = struct{}{}
					} else {
						delete(cubes, cube{x: i, y: j, z: k})
					}
				}
			}
		}
	}
	return len(cubes)
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
	// fmt.Println(f2(bytes.NewReader(b)))
}
