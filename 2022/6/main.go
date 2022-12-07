package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	for scanner.Scan() {
		packet := scanner.Text()
		for i := 0; i < len(packet); i++ {
			marker := packet[i : i+4]
			chars := map[rune]struct{}{}
			for _, c := range marker {
				chars[c] = struct{}{}
			}
			if len(chars) != 4 {
				continue
			}

			fmt.Println(i + 4)
			break
		}
	}
}
