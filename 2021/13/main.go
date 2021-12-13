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

	type point struct {
		x, y int
	}
	dots := map[point]struct{}{}
	scanner := bufio.NewScanner(f)
	// dots
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}

		var p point
		if _, err := fmt.Sscanf(scanner.Text(), "%d,%d", &p.x, &p.y); err != nil {
			log.Fatal(err)
		}
		dots[p] = struct{}{}
	}
	type fold struct {
		axis  rune
		value int
	}
	folds := []fold{}
	// folds
	for scanner.Scan() {
		var fv fold
		if _, err := fmt.Sscanf(scanner.Text(), "fold along %c=%d", &fv.axis, &fv.value); err != nil {
			log.Fatal(err)
		}

		folds = append(folds, fv)
	}
	// foldings
	fv := folds[0]
	// list of keys
	dotsList := make([]point, 0, len(dots))
	for d := range dots {
		dotsList = append(dotsList, d)
	}
	if fv.axis == 'x' {
		for _, d := range dotsList {
			if d.x <= fv.value {
				continue
			}

			delete(dots, d)
			dots[point{
				x: 2*fv.value - d.x,
				y: d.y,
			}] = struct{}{}
		}
	} else {
		for _, d := range dotsList {
			if d.y <= fv.value {
				continue
			}

			delete(dots, d)
			dots[point{
				x: d.x,
				y: 2*fv.value - d.y,
			}] = struct{}{}
		}
	}
	fmt.Println(len(dots))
}

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	type point struct {
		x, y int
	}
	dots := map[point]struct{}{}
	scanner := bufio.NewScanner(f)
	// dots
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}

		var p point
		if _, err := fmt.Sscanf(scanner.Text(), "%d,%d", &p.x, &p.y); err != nil {
			log.Fatal(err)
		}
		dots[p] = struct{}{}
	}
	type fold struct {
		axis  rune
		value int
	}
	folds := []fold{}
	// folds
	for scanner.Scan() {
		var fv fold
		if _, err := fmt.Sscanf(scanner.Text(), "fold along %c=%d", &fv.axis, &fv.value); err != nil {
			log.Fatal(err)
		}

		folds = append(folds, fv)
	}
	// foldings
	for _, fv := range folds {
		// list of keys
		dotsList := make([]point, 0, len(dots))
		for d := range dots {
			dotsList = append(dotsList, d)
		}
		if fv.axis == 'x' {
			for _, d := range dotsList {
				if d.x <= fv.value {
					continue
				}

				delete(dots, d)
				dots[point{
					x: 2*fv.value - d.x,
					y: d.y,
				}] = struct{}{}
			}
		} else {
			for _, d := range dotsList {
				if d.y <= fv.value {
					continue
				}

				delete(dots, d)
				dots[point{
					x: d.x,
					y: 2*fv.value - d.y,
				}] = struct{}{}
			}
		}
	}
	// find max x and max y
	xMax, yMax := 0, 0
	for d := range dots {
		if d.x > xMax {
			xMax = d.x
		}

		if d.y > yMax {
			yMax = d.y
		}
	}
	for j := 0; j <= yMax; j++ {
		for i := 0; i <= xMax; i++ {
			if _, ok := dots[point{
				x: i,
				y: j,
			}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	f1()
	f2()
}
