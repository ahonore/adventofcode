package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var regexOnlyDigitsLine = regexp.MustCompile(`^[ 0-9]+$`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var stasks [][]byte
	split := 0
	// stacks
	for scanner.Scan() {
		line := scanner.Text()
		if split == 0 {
			split = (len(line) + 1) / 4
			stasks = make([][]byte, split)
		}

		bLine := []byte(line)
		if regexOnlyDigitsLine.Match(bLine) {
			break
		}

		bLine = append(bLine, ' ')
		for i := 0; i < split; i++ {
			stackValue := bLine[:4]
			bLine = bLine[4:]
			if stackValue[1] == ' ' {
				continue
			}

			stasks[i] = append([]byte{stackValue[1]}, stasks[i]...)
		}
	}
	// 1 empty line
	scanner.Scan()
	// moves
	for scanner.Scan() {
		var n, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &n, &from, &to)
		stasks[to-1] = append(stasks[to-1], stasks[from-1][len(stasks[from-1])-n:len(stasks[from-1])]...)
		stasks[from-1] = stasks[from-1][:len(stasks[from-1])-n]
	}
	for _, s := range stasks {
		fmt.Print(string(s[len(s)-1]))
	}
	fmt.Println()
}
