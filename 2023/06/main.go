package main

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

// formula: d = (t-x)x
// where:
// - t: the max time allowed
// - x: the time the button is held
// - d: distance
// solutions: x in [ (t-sqrt(t²-4*d))/2, (t+sqrt(t²-4*d))/2 ]

type boatRun struct {
	time     int
	distance int
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	space := regexp.MustCompile(`\s+`)
	// times
	scanner.Scan()
	line := scanner.Text()
	str := strings.Split(line, ":")
	timesStr := strings.Split(space.ReplaceAllString(strings.TrimSpace(str[1]), " "), " ")
	// distances
	scanner.Scan()
	line = scanner.Text()
	str = strings.Split(line, ":")
	distancesStr := strings.Split(space.ReplaceAllString(strings.TrimSpace(str[1]), " "), " ")

	var boatRuns []boatRun
	for i := 0; i < len(timesStr); i++ {
		t, _ := strconv.Atoi(timesStr[i])
		d, _ := strconv.Atoi(distancesStr[i])
		boatRuns = append(boatRuns, boatRun{
			time:     t,
			distance: d,
		})
	}

	wins := 1
	for _, br := range boatRuns {
		t := float64(br.time)
		d := float64(br.distance)
		x0 := (t - math.Sqrt(t*t-4*d)) / 2.0
		x1 := (t + math.Sqrt(t*t-4*d)) / 2.0
		if x0 > x1 {
			x0, x1 = x1, x0
		}
		x0 = math.Ceil(x0)
		if (t-x0)*x0 == d {
			x0++
		}
		x1 = math.Floor(x1)
		if (t-x1)*x1 == d {
			x1--
		}
		s0 := int(x0)
		s1 := int(x1)
		curWins := s1 - s0 + 1
		wins *= curWins
	}
	fmt.Println(wins)
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	// time
	scanner.Scan()
	line := scanner.Text()
	str := strings.Split(line, ":")
	ti, _ := strconv.Atoi(strings.ReplaceAll(str[1], " ", ""))
	// distance
	scanner.Scan()
	line = scanner.Text()
	str = strings.Split(line, ":")
	di, _ := strconv.Atoi(strings.ReplaceAll(str[1], " ", ""))

	t := float64(ti)
	d := float64(di)
	x0 := (t - math.Sqrt(t*t-4*d)) / 2.0
	x1 := (t + math.Sqrt(t*t-4*d)) / 2.0
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	x0 = math.Ceil(x0)
	if (t-x0)*x0 == d {
		x0++
	}
	x1 = math.Floor(x1)
	if (t-x1)*x1 == d {
		x1--
	}
	s0 := int(x0)
	s1 := int(x1)
	wins := s1 - s0 + 1
	fmt.Println(wins)
}

func main() {
	part1()
	part2()
}
