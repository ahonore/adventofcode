package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func hash(s string) int {
	var h int
	for _, c := range s {
		h += int(c)
		h *= 17
		h %= 256
	}
	return h
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var insts []string
	for scanner.Scan() {
		insts = strings.Split(scanner.Text(), ",")
	}
	var sum int
	for _, inst := range insts {
		sum += hash(inst)
	}
	fmt.Println(sum)
}

type lens struct {
	label    string
	focalLen int
}

type box struct {
	labelMap map[string]*lens
	content  []*lens
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var insts []string
	for scanner.Scan() {
		insts = strings.Split(scanner.Text(), ",")
	}
	var boxes [256]*box
	for _, inst := range insts {
		var label string
		var h, num int
		var op byte
		if id := strings.Index(inst, "-"); id >= 0 {
			label = inst[:id]
			h = hash(label)
			op = '-'
		} else if id := strings.Index(inst, "="); id >= 0 {
			label = inst[:id]
			h = hash(label)
			op = '='
			num, _ = strconv.Atoi(inst[id+1:])
		}

		switch op {
		case '-':
			box := boxes[h]
			if box == nil {
				continue
			}
			l, ok := box.labelMap[label]
			if !ok {
				continue
			}
			for i := 0; i < len(box.content); i++ {
				if box.content[i] == l {
					delete(box.labelMap, label)
					box.content = append(box.content[:i], box.content[i+1:]...)
					if len(box.content) == 0 {
						boxes[h] = nil
					}
					break
				}
			}
		case '=':
			b := boxes[h]
			if b == nil {
				l := &lens{
					label:    label,
					focalLen: num,
				}
				boxes[h] = &box{
					content:  []*lens{l},
					labelMap: map[string]*lens{label: l},
				}
				continue
			}
			l, ok := b.labelMap[label]
			if !ok { // add new lens
				l = &lens{
					label:    label,
					focalLen: num,
				}
				b.content = append(b.content, l)
				b.labelMap[label] = l
				continue
			}
			l.focalLen = num
		default:
			fmt.Println("error")
			return
		}
	}
	var power int
	for ib, b := range boxes {
		if b == nil {
			continue
		}
		for il, l := range b.content {
			power += (ib + 1) * (il + 1) * l.focalLen
		}
	}
	fmt.Println(power)
}

func main() {
	part1()
	part2()
}
