package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	valid := 0
	buf := bufio.NewReader(f)
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		var min, max int
		var c string
		var pwd string
		s := string(l)
		fmt.Sscanf(s, "%d-%d %1s: %s", &min, &max, &c, &pwd)
		r := rune(c[0])
		ct := 0
		for _, v := range pwd {
			if v == r {
				ct++
			}
		}
		if ct >= min && ct <= max {
			valid++
		}
	}
	fmt.Println(valid)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	valid := 0
	buf := bufio.NewReader(f)
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		var min, max int
		var sc string
		var pwd string
		s := string(l)
		fmt.Sscanf(s, "%d-%d %1s: %s", &min, &max, &sc, &pwd)
		c := sc[0]
		p0 := pwd[min-1] == c
		p1 := pwd[max-1] == c
		if (p0 && !p1) || (p1 && !p0) {
			valid++
		}
	}
	fmt.Println(valid)
}
