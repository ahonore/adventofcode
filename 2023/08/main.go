package main

import (
	"bufio"
	"fmt"
	"math/big"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type path struct {
	left  string
	right string
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Scan()
	insts := scanner.Text()
	scanner.Scan() // newline
	mapPaths := map[string]path{}
	for scanner.Scan() {
		line := scanner.Text()
		var start string
		var p path
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &start, &p.left, &p.right)
		mapPaths[start] = p
	}
	curPath := "AAA"
	var ctMoves int
	for {
		inst := insts[ctMoves%len(insts)]
		switch inst {
		case 'L':
			curPath = mapPaths[curPath].left
		case 'R':
			curPath = mapPaths[curPath].right
		default:
			fmt.Println("error")
			return
		}
		ctMoves++
		if curPath == "ZZZ" {
			break
		}
	}
	fmt.Println(ctMoves)
}

func lcm(a int, b ...int) int {
	res := big.NewInt(int64(a))
	for _, i := range b {
		v := big.NewInt(int64(i))
		res = (&big.Int{}).Div((&big.Int{}).Mul(res, v), (&big.Int{}).GCD(nil, nil, res, v))
	}
	return int(res.Int64())
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Scan()
	insts := scanner.Text()
	scanner.Scan() // newline
	mapPaths := map[string]path{}
	for scanner.Scan() {
		line := scanner.Text()
		var start string
		var p path
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &start, &p.left, &p.right)
		mapPaths[start] = p
	}
	// find all entries ending with A
	var curPaths []string
	for p := range mapPaths {
		if p[2] == 'A' {
			curPaths = append(curPaths, p)
		}
	}
	var counts []int
	for i := 0; i < len(curPaths); i++ {
		var ctMoves int
		for {
			inst := insts[ctMoves%len(insts)]
			switch inst {
			case 'L':
				curPaths[i] = mapPaths[curPaths[i]].left
			case 'R':
				curPaths[i] = mapPaths[curPaths[i]].right
			default:
				fmt.Println("error")
				return
			}
			ctMoves++
			if curPaths[i][2] == 'Z' {
				counts = append(counts, ctMoves)
				break
			}
		}
	}
	fmt.Println(lcm(counts[0], counts[1:]...))
}

func main() {
	part1()
	part2()
}
