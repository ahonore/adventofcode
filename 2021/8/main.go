package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
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

const (
	segmentUp = iota
	segmentUpLeft
	segmentUpRight
	segmentMiddle
	segmentBottomLeft
	segmentBottomRight
	segmentBottom
)

func f2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// mapping of string to uint8 bitfield: bit 0->'a', bit 1->'b' and so on
		var display [7]uint8 // up, upleft, upright, middle, bottomleft, bottomright, bottom
		for i := 0; i < len(display); i++ {
			display[i] = 0x7f // init all possibilities per segment
		}
		var strDigits [10]string
		var digits [10]uint8
		var strOutputs [4]string
		var outputs [4]uint8
		var values [10]uint8
		fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&strDigits[0], &strDigits[1], &strDigits[2], &strDigits[3], &strDigits[4], &strDigits[5], &strDigits[6], &strDigits[7], &strDigits[8], &strDigits[9],
			&strOutputs[0], &strOutputs[1], &strOutputs[2], &strOutputs[3])
		for i := 0; i < 10; i++ {
			for _, r := range strDigits[i] {
				digits[i] |= 1 << (r - 97)
			}
		}
		for i := 0; i < 4; i++ {
			for _, r := range strOutputs[i] {
				outputs[i] |= 1 << (r - 97)
			}
		}
		// find 1 and 8
		for _, v := range digits {
			if bits.OnesCount8(v) == 2 {
				values[1] = v
				display[segmentUpRight] = v
				display[segmentBottomRight] = v
				continue
			}

			if bits.OnesCount8(v) == 7 {
				values[8] = v
				continue
			}
		}
		// find 7
		for _, v := range digits {
			if bits.OnesCount8(v) == 3 {
				values[7] = v
				v1 := values[1]
				display[segmentUp] = v ^ v1
				break
			}
		}
		// find 4
		for _, v := range digits {
			if bits.OnesCount8(v) == 4 {
				values[4] = v
				v1 := values[1]
				display[segmentUpLeft] = (0x7f ^ v1) & v
				display[segmentMiddle] = (0x7f ^ v1) & v
				break
			}
		}
		// find 3
		for _, v := range digits {
			if bits.OnesCount8(v) == 5 {
				v1 := values[1]
				// must include segments of one
				if bits.OnesCount8(v&v1) != 2 {
					continue
				}

				values[3] = v
				v7 := values[7]
				midAndBot := v & (0x7f ^ v7)
				display[segmentMiddle] &= midAndBot
				display[segmentBottom] = display[segmentMiddle] ^ midAndBot
				display[segmentUpLeft] ^= display[segmentMiddle]
				break
			}
		}
		// find 5
		for _, v := range digits {
			// must include segments of up and upleft
			if bits.OnesCount8(v) == 5 && bits.OnesCount8(v&(display[segmentUp]|display[segmentUpLeft])) == 2 {
				values[5] = v
				display[segmentBottomRight] &= v
				display[segmentUpRight] ^= display[segmentBottomRight]
				break
			}
		}
		// set segment bottom left
		display[segmentBottomLeft] ^= display[segmentUp] | display[segmentUpLeft] | display[segmentUpRight] | display[segmentMiddle] | display[segmentBottomRight] | display[segmentBottom]
		// find 2, 6, 9, 0
		for _, v := range digits {
			if v == (display[segmentUp] | display[segmentUpRight] | display[segmentUpLeft] | display[segmentBottomLeft] | display[segmentBottomRight] | display[segmentBottom]) {
				values[0] = v
				continue
			}

			if v == (display[segmentUp] | display[segmentUpRight] | display[segmentMiddle] | display[segmentBottomLeft] | display[segmentBottom]) {
				values[2] = v
				continue
			}

			if v == (display[segmentUp] | display[segmentUpLeft] | display[segmentMiddle] | display[segmentBottomLeft] | display[segmentBottomRight] | display[segmentBottom]) {
				values[6] = v
				continue
			}

			if v == (display[segmentUp] | display[segmentUpLeft] | display[segmentUpRight] | display[segmentMiddle] | display[segmentBottomRight] | display[segmentBottom]) {
				values[9] = v
				continue
			}
		}
		num := 0
		for _, o := range outputs {
			for id, v := range values {
				if v == o {
					num = num*10 + id
				}
			}
		}
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	f1()
	f2()
}
