package main

import (
	"bufio"
	"fmt"
	"os"
)

const size int = 14

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	for scanner.Scan() {
		packet := scanner.Text()
		for i := 0; i < len(packet); i++ {
			marker := packet[i : i+size]
			chars := map[rune]struct{}{}
			for _, c := range marker {
				chars[c] = struct{}{}
			}
			if len(chars) != size {
				continue
			}

			fmt.Println(i + size)
			break
		}
	}
}
