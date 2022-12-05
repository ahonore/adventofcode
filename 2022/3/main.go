package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		comp0 := line[:len(line)/2]
		comp1 := line[len(line)/2:]
		for _, c := range comp0 {
			if !strings.ContainsRune(comp1, c) {
				continue
			}

			sum += strings.IndexRune(alpha, c) + 1
			break
		}
	}
	fmt.Println(sum)
}
