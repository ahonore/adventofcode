package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opcode  string
	operand int
}

func main() {
	listInst := []instruction{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inst := scanner.Text()
		instParts := strings.Split(inst, " ")
		i, err := strconv.Atoi(instParts[1])
		if err != nil {
			log.Fatal(err)
		}

		listInst = append(listInst, instruction{
			opcode:  instParts[0],
			operand: i,
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	acc := 0
	pc := 0
	loopChecker := map[int]struct{}{}
	loop := false
	for !loop {
		loopChecker[pc] = struct{}{}
		inst := listInst[pc]
		switch inst.opcode {
		case "acc":
			acc += inst.operand
			pc++
		case "jmp":
			pc += inst.operand
		case "nop":
			fallthrough
		default:
			pc++
		}
		_, loop = loopChecker[pc]
	}
	fmt.Println(acc)
}
