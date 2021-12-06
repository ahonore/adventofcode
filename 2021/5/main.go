package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	vents := map[string]int{}
	for scanner.Scan() {
		var x0, x1, y0, y1 int
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		if x0 == x1 {
			if y0 > y1 {
				y0, y1 = y1, y0
			}

			for i := y0; i <= y1; i++ {
				vents[fmt.Sprintf("%d %d", x0, i)]++
			}
		} else if y0 == y1 {
			if x0 > x1 {
				x0, x1 = x1, x0
			}

			for i := x0; i <= x1; i++ {
				vents[fmt.Sprintf("%d %d", i, y0)]++
			}
		}
	}

	overlaps := 0
	for _, v := range vents {
		if v >= 2 {
			overlaps++
		}
	}
	fmt.Println("overlaps", overlaps)
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	vents := map[string]int{}
	for scanner.Scan() {
		var x0, x1, y0, y1 int
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		if x0 == x1 {
			if y0 > y1 {
				y0, y1 = y1, y0
			}

			for i := y0; i <= y1; i++ {
				vents[fmt.Sprintf("%d %d", x0, i)]++
			}
		} else if y0 == y1 {
			if x0 > x1 {
				x0, x1 = x1, x0
			}

			for i := x0; i <= x1; i++ {
				vents[fmt.Sprintf("%d %d", i, y0)]++
			}
		} else {
			// diagonal
			yInc := 1
			if y0 > y1 {
				yInc = -1
			}

			xInc := 1
			if x0 > x1 {
				xInc = -1
			}

			x := x0
			y := y0
			for {
				if x == x1 && y == y1 {
					vents[fmt.Sprintf("%d %d", x, y)]++
					break
				}

				vents[fmt.Sprintf("%d %d", x, y)]++
				x += xInc
				y += yInc
			}
		}
	}

	overlaps := 0
	for _, v := range vents {
		if v >= 2 {
			overlaps++
		}
	}
	fmt.Println("overlaps", overlaps)
}

func main() {
	f1()
	f2()
}
