package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		valueStr := strings.Fields(line)
		var values []int
		for _, s := range valueStr {
			i, _ := strconv.Atoi(s)
			values = append(values, i)
		}
		reports = append(reports, values)
	}
	var sum int
	for _, report := range reports {
		var values [][]int
		values = append(values, make([]int, len(report)))
		copy(values[0], report)
		for {
			var zero int
			var value []int
			lastValue := values[len(values)-1]
			for i := 1; i < len(lastValue); i++ {
				diff := lastValue[i] - lastValue[i-1]
				zero += diff
				value = append(value, diff)
			}
			values = append(values, value)
			if zero == 0 {
				break
			}
		}
		var curVal int
		for i := len(values) - 2; i >= 0; i-- {
			curVal += values[i][len(values[i])-1]
		}
		sum += curVal
	}
	fmt.Println(sum)
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		valueStr := strings.Fields(line)
		var values []int
		for _, s := range valueStr {
			i, _ := strconv.Atoi(s)
			values = append(values, i)
		}
		reports = append(reports, values)
	}
	var sum int
	for _, report := range reports {
		var values [][]int
		values = append(values, make([]int, len(report)))
		copy(values[0], report)
		for {
			var zero int
			var value []int
			lastValue := values[len(values)-1]
			for i := 1; i < len(lastValue); i++ {
				diff := lastValue[i] - lastValue[i-1]
				zero += diff
				value = append(value, diff)
			}
			values = append(values, value)
			if zero == 0 {
				break
			}
		}
		var curVal int
		for i := len(values) - 2; i >= 0; i-- {
			curVal = values[i][0] - curVal
		}
		sum += curVal
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
