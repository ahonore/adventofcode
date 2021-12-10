package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func f1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ones, fours, sevens, eights := 0, 0, 0, 0
	for scanner.Scan() {
		var digits [10]string
		var outputs [4]string
		fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&digits[0], &digits[1], &digits[2], &digits[3], &digits[4], &digits[5], &digits[6], &digits[7], &digits[8], &digits[9],
			&outputs[0], &outputs[1], &outputs[2], &outputs[3])
		for _, v := range outputs {
			switch len(v) {
			case 2:
				ones++
			case 4:
				fours++
			case 3:
				sevens++
			case 7:
				eights++
			}
		}
	}
	fmt.Println(ones + fours + sevens + eights)
}

// func f2() {
// 	f, err := os.Open("test.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	ones, fours, sevens, eights := 0, 0, 0, 0
// 	for scanner.Scan() {
// 		var display [7]string // up, upleft, upright, middle, bottomleft, bottomright, bottom
// 		for i := 0; i < len(display); i++ {
// 			display[i] = "abcdefg" // init all possibilities per segment
// 		}
// 		digits := make([]string, 10)
// 		var outputs [4]string
// 		fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
// 			&digits[0], &digits[1], &digits[2], &digits[3], &digits[4], &digits[5], &digits[6], &digits[7], &digits[8], &digits[9],
// 			&outputs[0], &outputs[1], &outputs[2], &outputs[3])
// 		// find one
// 		for _, v := range digits {
// 			if len(v) == 2 {
// 				display[2] = v
// 				display[5] = v
// 			}
// 		}
// 		// find seven
// 		for _, v := range digits {
// 			if len(v) == 3 {
// 				display[2] = v
// 			}
// 		}
// 	}
// 	fmt.Println(ones + fours + sevens + eights)
// }

func main() {
	f1()
	// f2()
}
