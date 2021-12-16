package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_f1(t *testing.T) {
	for name, test := range map[string]struct {
		inputText      string
		expectedOutput int
	}{
		"test0": {
			inputText:      "D2FE28",
			expectedOutput: 6,
		},
		"test1": {
			inputText:      "8A004A801A8002F478",
			expectedOutput: 16,
		},
		"test2": {
			inputText:      "620080001611562C8802118E34",
			expectedOutput: 12,
		},
		"test3": {
			inputText:      "C0015000016115A2E0802F182340",
			expectedOutput: 23,
		},
		"test4": {
			inputText:      "A0016C880162017C3686B18A3D4780",
			expectedOutput: 31,
		},
	} {
		test := test
		_ = test
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, f1(bytes.NewReader([]byte(test.inputText))))
		})
	}
}

func Test_f2(t *testing.T) {
	for name, test := range map[string]struct {
		inputText      string
		expectedOutput int
	}{
		"test0": {
			inputText:      "C200B40A82",
			expectedOutput: 3,
		},
		"test1": {
			inputText:      "04005AC33890",
			expectedOutput: 54,
		},
		"test2": {
			inputText:      "880086C3E88112",
			expectedOutput: 7,
		},
		"test3": {
			inputText:      "CE00C43D881120",
			expectedOutput: 9,
		},
		"test4": {
			inputText:      "D8005AC2A8F0",
			expectedOutput: 1,
		},
		"test5": {
			inputText:      "F600BC2D8F",
			expectedOutput: 0,
		},
		"test6": {
			inputText:      "9C005AC2F8F0",
			expectedOutput: 0,
		},
		"test7": {
			inputText:      "9C0141080250320F1802104A08",
			expectedOutput: 1,
		},
	} {
		test := test
		_ = test
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, f2(bytes.NewReader([]byte(test.inputText))))
		})
	}
}
