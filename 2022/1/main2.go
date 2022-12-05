package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var calories, curCalories int
	var caloriesPerElf []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			caloriesPerElf = append(caloriesPerElf, curCalories)
			curCalories = 0
			continue
		}

		fmt.Sscan(scanner.Text(), &calories)
		curCalories += calories
	}
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))
	fmt.Println(caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2])
}
