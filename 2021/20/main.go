package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func f1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	txt := scanner.Text()
	imgEnAlgo := make([]int, len(txt))
	for i, c := range txt {
		if c == '#' {
			imgEnAlgo[i] = 1
		}
	}
	scanner.Scan()
	// newline
	scanner.Text()
	inputImg := [][]int{}
	for scanner.Scan() {
		txt := scanner.Text()
		line := []int{}
		for _, c := range txt {
			v := 0
			if c == '#' {
				v = 1
			}

			line = append(line, v)
			// fmt.Print(v)
		}
		inputImg = append(inputImg, line)
		// fmt.Println()
	}
	delta := []struct{ x, y int }{
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
		{x: -1, y: 0},
		{x: 0, y: 0},
		{x: 1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: 1},
		{x: 1, y: 1},
	}
	shiftLenY := 4
	shiftLenX := 4
	for steps := 0; steps < 2; steps++ {
		// dim increased by len/2 in each direction, to compute on borders
		outputImg := make([][]int, len(inputImg)+2*shiftLenY)
		for j := 0; j < len(outputImg); j++ {
			outputLine := make([]int, len(inputImg[0])+2*shiftLenX)
			for i := 0; i < len(outputImg[0]); i++ {
				num := 0
				for _, d := range delta {
					x := i + d.x - shiftLenX
					y := j + d.y - shiftLenY
					num <<= 1
					if x < 0 || x >= len(inputImg[0]) || y < 0 || y >= len(inputImg) {
						continue
					}

					num |= inputImg[y][x]
				}
				outputLine[i] = imgEnAlgo[num]
			}
			outputImg[j] = outputLine
		}
		inputImg = outputImg
		shiftLenY = 0
		shiftLenX = 0
	}
	// count lits
	sum := 0
	for j := 2; j < len(inputImg)-2; j++ {
		for i := 2; i < len(inputImg[0])-2; i++ {
			// fmt.Print(inputImg[j][i])
			sum += inputImg[j][i]
		}
		// fmt.Println()
	}
	return sum
}

func f2(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	txt := scanner.Text()
	imgEnAlgo := make([]int, len(txt))
	for i, c := range txt {
		if c == '#' {
			imgEnAlgo[i] = 1
		}
	}
	scanner.Scan()
	// newline
	scanner.Text()
	inputImg := [][]int{}
	for scanner.Scan() {
		txt := scanner.Text()
		line := []int{}
		for _, c := range txt {
			v := 0
			if c == '#' {
				v = 1
			}

			line = append(line, v)
			// fmt.Print(v)
		}
		inputImg = append(inputImg, line)
		// fmt.Println()
	}
	delta := []struct{ x, y int }{
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
		{x: -1, y: 0},
		{x: 0, y: 0},
		{x: 1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: 1},
		{x: 1, y: 1},
	}
	shiftLenY := 100
	shiftLenX := 100
	for steps := 0; steps < 50; steps++ {
		// dim increased by len/2 in each direction, to compute on borders
		outputImg := make([][]int, len(inputImg)+2*shiftLenY)
		for j := 0; j < len(outputImg); j++ {
			outputLine := make([]int, len(inputImg[0])+2*shiftLenX)
			for i := 0; i < len(outputImg[0]); i++ {
				num := 0
				for _, d := range delta {
					x := i + d.x - shiftLenX
					y := j + d.y - shiftLenY
					num <<= 1
					if x < 0 || x >= len(inputImg[0]) || y < 0 || y >= len(inputImg) {
						continue
					}

					num |= inputImg[y][x]
				}
				outputLine[i] = imgEnAlgo[num]
			}
			outputImg[j] = outputLine
		}
		inputImg = outputImg
		shiftLenY = 0
		shiftLenX = 0
	}
	// count lits
	sum := 0
	for j := 50; j < len(inputImg)-50; j++ {
		for i := 50; i < len(inputImg[0])-50; i++ {
			// fmt.Print(inputImg[j][i])
			sum += inputImg[j][i]
		}
		// fmt.Println()
	}
	return sum
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
