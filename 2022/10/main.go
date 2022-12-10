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

var markCycles = map[int]struct{}{
	20:  {},
	60:  {},
	100: {},
	140: {},
	180: {},
	220: {},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var cycles int = 1
	c := &cpu{}
	c.init()
	var signalStrengths int
	for scanner.Scan() {
		inst := strings.Fields(scanner.Text())
		curInstCycles := instCycles[inst[0]]
		for i := 0; i < curInstCycles-1; i++ {
			if _, ok := markCycles[cycles]; ok {
				signalStrengths += cycles * c.X
			}
			cycles++
		}
		if _, ok := markCycles[cycles]; ok {
			signalStrengths += cycles * c.X
		}

		if inst[0] == "addx" {
			v, _ := strconv.Atoi(inst[1])
			c.X += v
		}

		cycles++
	}
	fmt.Println(signalStrengths)
}
