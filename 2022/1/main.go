package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var calories, curCalories, maxCalories int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if curCalories > maxCalories {
				maxCalories = curCalories
			}

			curCalories = 0
			continue
		}
		fmt.Sscan(scanner.Text(), &calories)
		curCalories += calories
	}
	fmt.Println(maxCalories)
}
