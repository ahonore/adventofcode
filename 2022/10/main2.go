package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	X int
}

func (c *cpu) init() {
	c.X = 1
}

var instCycles = map[string]int{
	"noop": 1,
	"addx": 2,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	// var cycles int = 1
	c := &cpu{}
	c.init()
	crtID := 0
	var crt [240]byte
	for scanner.Scan() {
		inst := strings.Fields(scanner.Text())
		curInstCycles := instCycles[inst[0]]
		for i := 0; i < curInstCycles-1; i++ {
			spriteMiddleID := c.X - crtID%40
			crt[crtID%240] = '.'
			switch spriteMiddleID {
			case -1, 0, 1:
				crt[crtID%240] = '#'
			}

			// fmt.Println(string(crt[:40]))
			crtID++
			// cycles++
		}
		spriteMiddleID := c.X - crtID%40
		crt[crtID%240] = '.'
		switch spriteMiddleID {
		case -1, 0, 1:
			crt[crtID%240] = '#'
		}

		// fmt.Println(string(crt[:40]))
		if inst[0] == "addx" {
			v, _ := strconv.Atoi(inst[1])
			c.X += v
		}

		crtID++
		// cycles++
	}
	for i := 0; i < 6; i++ {
		fmt.Println(string(crt[i*40 : (i+1)*40]))
	}
}
