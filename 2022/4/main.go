package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var ct, min0, max0, min1, max1 int
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &min0, &max0, &min1, &max1)
		// fmt.Println(min0, max0, min1, max1)
		if (min0 <= min1 && max0 >= max1) || (min0 >= min1 && max0 <= max1) {
			ct++
		}
	}
	fmt.Println(ct)
}
