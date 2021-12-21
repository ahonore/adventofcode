package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_add(t *testing.T) {
	for name, test := range map[string]struct {
		tree0          string
		tree1          string
		expectedOutput string
	}{
		"test0": {
			tree0:          "[[[[4,3],4],4],[7,[[8,4],9]]]",
			tree1:          "[1,1]",
			expectedOutput: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	} {
		test := test
		t.Run(name, func(t *testing.T) {
			n0, _ := parseNumber([]byte(test.tree0))
			n1, _ := parseNumber([]byte(test.tree1))
			n := n0.add(n1)
			assert.Equal(t, test.expectedOutput, n.String())
		})
	}

}

func TestNode_magnitude(t *testing.T) {
	for name, test := range map[string]struct {
		tree           *node
		expectedOutput int
	}{
		"test0": {
			// [[1,2],[[3,4],5]]
			tree: &node{
				left: &node{
					left:  &node{value: 1},
					right: &node{value: 2},
				},
				right: &node{
					left: &node{
						left:  &node{value: 3},
						right: &node{value: 4},
					},
					right: &node{value: 5},
				},
			},
			expectedOutput: 143,
		},
		"test1": {
			// [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
			tree: &node{
				left: &node{
					left: &node{
						left: &node{
							left:  &node{value: 0},
							right: &node{value: 7},
						},
						right: &node{value: 4},
					},
					right: &node{
						left: &node{
							left:  &node{value: 7},
							right: &node{value: 8},
						},
						right: &node{
							left:  &node{value: 6},
							right: &node{value: 0},
						},
					},
				},
				right: &node{
					left:  &node{value: 8},
					right: &node{value: 1},
				},
			},
			expectedOutput: 1384,
		},
	} {
		test := test
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, test.tree.magnitude())
		})
	}
}
