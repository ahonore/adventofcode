package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main1() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	researchMap := map[int]struct{}{}
	buf := bufio.NewReader(f)
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		n, _ := strconv.Atoi(string(l))
		fmt.Println(n)
		op := 2020 - n
		if _, ok := researchMap[op]; ok {
			fmt.Printf("%d * %d = %d\n", n, op, op*n)
			return
		}
		researchMap[n] = struct{}{}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	var inputs []int
	sumMap := map[int][]int{}
	buf := bufio.NewReader(f)
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		n, _ := strconv.Atoi(string(l))
		if v, ok := sumMap[n]; ok {
			prod := n
			for _, i := range v {
				prod *= i
			}
			fmt.Println(prod)
			return
		}

		for _, v := range inputs {
			op := 2020 - n - v
			sumMap[op] = []int{n, v}
		}
		inputs = append(inputs, n)
	}
}
