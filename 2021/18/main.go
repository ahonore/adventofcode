package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
}

func parseNumber(b []byte) (*node, []byte) {
	if b[0] != '[' {
		log.Fatalf("expected '[' but was '%c'", b[0])
	}

	n := &node{}
	// left part
	switch {
	case b[1] >= 48 && b[1] < 58:
		n.left = &node{value: int(b[1]) - 48}
		n.left.parent = n
		// fmt.Print(n.left.value, " ")
		b = b[2:]
	case b[1] == '[':
		n.left, b = parseNumber(b[1:])
		n.left.parent = n
	default:
		log.Fatalf("not expected '%c'", b[1])
	}
	if b[0] != ',' {
		log.Fatalf("expected ',' but was '%c'", b[0])
	}

	// right part
	switch {
	case b[1] >= 48 && b[1] < 58:
		n.right = &node{value: int(b[1]) - 48}
		n.right.parent = n
		// fmt.Print(n.right.value, " ")
		b = b[2:]
	case b[1] == '[':
		n.right, b = parseNumber(b[1:])
		n.right.parent = n
	default:
		log.Fatalf("not expected '%c'", b[1])
	}
	if b[0] != ']' {
		log.Fatalf("expected ']' but was '%c'", b[0])
	}

	return n, b[1:]
}

func (n *node) isPair() bool {
	return n.left != nil || n.right != nil
}

func (n *node) add(right *node) *node {
	addNode := n.concat(right)
	addNode.reduce()
	return addNode
}

func (n *node) concat(right *node) *node {
	root := &node{
		left:  n,
		right: right,
	}
	root.left.parent = root
	root.right.parent = root
	return root
}

func (n *node) propagateExplodeUp(fromChild *node, left bool, value int) bool {
	if left {
		if !n.left.isPair() {
			n.left.value += value
			return true
		}
	} else {
		if !n.right.isPair() {
			n.right.value += value
			return true
		}
	}

	ret := false
	if n.parent != nil {
		ret = n.parent.propagateExplodeUp(n, left, value)
	} else if left && n.right == fromChild {
		ret = n.left.propagateExplodeDown(left, value)
	} else if !left && n.left == fromChild {
		ret = n.left.propagateExplodeDown(left, value)
	}

	return ret
}

func (n *node) propagateExplodeDown(left bool, value int) bool {
	if left {
		if !n.left.isPair() {
			n.left.value += value
			return true
		}

		return n.left.propagateExplodeDown(left, value)
	} else {
		if !n.right.isPair() {
			n.right.value += value
			return true
		}

		return n.right.propagateExplodeDown(left, value)
	}
}

func (n *node) reduceRec(depth int) bool {
	if depth == 4 {
		if n.isPair() {
			leftValue := n.left.value
			rightValue := n.right.value
			// pair from the left
			if n.parent.left == n {
				prev := n.parent
				cur := n.parent
				for cur != nil {
					cur = cur.parent
					// if I come from the left, continue up
					if cur.left == prev {
						prev = cur
						cur = cur.
					}
					// a value to the left ?
					if cur.left != nil && cur.left.isPair() {

					}
					// find rightmost value

				}
			}
			n.parent.propagateExplodeUp(n, true, leftValue)

			// explode right
			n.parent.propagateExplodeUp(n, false, rightValue)

			n.value = 0
			n.left = nil
			n.right = nil
			return true
		}
	}

	if !n.isPair() && n.value >= 10 {
		// split value
		n.left = &node{value: n.value / 2}
		n.right = &node{value: n.value - n.left.value}
		n.value = 0
		return true
	}

	if n.isPair() {
		ret := false
		if n.left != nil {
			ret = n.left.reduceRec(depth + 1)
		}

		if n.right != nil && !ret {
			ret = n.right.reduceRec(depth + 1)
		}

		return ret
	}

	return false
}

func (n *node) reduce() {
	fmt.Println(n.String())
	for n.reduceRec(0) {
		fmt.Println(n.String())
	}
}

func (n node) magnitude() int {
	if n.left == nil && n.right == nil {
		return n.value
	}

	mag := 0
	if n.left != nil {
		mag = 3 * n.left.magnitude()
	}

	if n.right != nil {
		mag += 2 * n.right.magnitude()
	}

	return mag
}

func (n node) String() string {
	if !n.isPair() {
		return strconv.Itoa(n.value)
	}

	var sb strings.Builder
	sb.WriteByte('[')
	sb.WriteString(n.left.String())
	sb.WriteByte(',')
	sb.WriteString(n.right.String())
	sb.WriteByte(']')
	return sb.String()
}

func f1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	acc, _ := parseNumber([]byte(scanner.Text()))
	fmt.Println(acc.String())
	for scanner.Scan() {
		n, _ := parseNumber([]byte(scanner.Text()))
		fmt.Println(n.String())
		acc = acc.add(n)
		acc.reduce()
	}
	fmt.Println(acc.String())
	return acc.magnitude()
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1(bytes.NewReader(b)))
	// fmt.Println(f2(bytes.NewReader(b)))
}
