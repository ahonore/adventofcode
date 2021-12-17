package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func f1(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var xMin, xMax, yMin, yMax int
	fmt.Sscanf(string(b), "target area: x=%d..%d, y=%d..%d", &xMin, &xMax, &yMin, &yMax)
	// max x velocity = xMin/2
	velocityXMax := xMin / 2
	// max y velocity = yMin
	velocityYMax := int(math.Abs(float64(yMin)))
	// find yBest
	yBest := 0
	for j := 1; j < velocityYMax; j++ {
		for i := 1; i < velocityXMax; i++ {
			curYBest := 0
			velX, velY := i, j
			x, y := 0, 0
			for {
				x += velX
				y += velY
				if y > curYBest {
					curYBest = y
				}

				// inside square?
				if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
					if curYBest > yBest {
						yBest = curYBest
					}
					break
				}

				// ouside square?
				if x > xMax || y < yMin {
					break
				}

				if velX > 0 {
					velX--
				} else if velX < 0 {
					velX++
				}

				velY--
			}
		}
	}

	return yBest
}

func f2(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var xMin, xMax, yMin, yMax int
	fmt.Sscanf(string(b), "target area: x=%d..%d, y=%d..%d", &xMin, &xMax, &yMin, &yMax)
	maxValue := int(math.Abs(float64(yMin)))
	if maxValue < xMax {
		maxValue = xMax
	}

	velocities := 0
	for j := -maxValue; j <= maxValue; j++ {
		for i := 1; i <= maxValue; i++ {
			velX, velY := i, j
			x, y := 0, 0
			for {
				x += velX
				y += velY
				// inside square?
				if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
					velocities++
					break
				}

				// ouside square?
				if x > xMax || y < yMin {
					break
				}

				if velX > 0 {
					velX--
				} else if velX < 0 {
					velX++
				}

				velY--
			}
		}
	}

	return velocities
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
