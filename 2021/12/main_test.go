package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_f1(t *testing.T) {
	for name, test := range map[string]struct {
		inputFilename         string
		expectedPaths         []string
		expectedNumberOfPaths int
	}{
		"test1": {
			inputFilename: "test.txt",
			expectedPaths: []string{
				"start,A,b,A,c,A,end",
				"start,A,b,A,end",
				"start,A,b,end",
				"start,A,c,A,b,A,end",
				"start,A,c,A,b,end",
				"start,A,c,A,end",
				"start,A,end",
				"start,b,A,c,A,end",
				"start,b,A,end",
				"start,b,end",
			},
			expectedNumberOfPaths: 10,
		},
		"test2": {
			inputFilename: "test2.txt",
			expectedPaths: []string{
				"start,HN,dc,HN,end",
				"start,HN,dc,HN,kj,HN,end",
				"start,HN,dc,end",
				"start,HN,dc,kj,HN,end",
				"start,HN,end",
				"start,HN,kj,HN,dc,HN,end",
				"start,HN,kj,HN,dc,end",
				"start,HN,kj,HN,end",
				"start,HN,kj,dc,HN,end",
				"start,HN,kj,dc,end",
				"start,dc,HN,end",
				"start,dc,HN,kj,HN,end",
				"start,dc,end",
				"start,dc,kj,HN,end",
				"start,kj,HN,dc,HN,end",
				"start,kj,HN,dc,end",
				"start,kj,HN,end",
				"start,kj,dc,HN,end",
				"start,kj,dc,end",
			},
			expectedNumberOfPaths: 19,
		},
		"test3": {
			inputFilename:         "test3.txt",
			expectedNumberOfPaths: 226,
		},
	} {
		test := test
		_ = test
		t.Run(name, func(t *testing.T) {
			f, err := os.Open(test.inputFilename)
			require.NoError(t, err)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			connectedCaves := []string{}
			for scanner.Scan() {
				connectedCaves = append(connectedCaves, scanner.Text())
			}

			paths := f1(connectedCaves)
			assert.Equal(t, test.expectedNumberOfPaths, len(paths))
			if len(test.expectedPaths) > 0 {
				assert.ElementsMatch(t, test.expectedPaths, paths)
			}
		})
	}
}

func Test_f2(t *testing.T) {
	for name, test := range map[string]struct {
		inputFilename         string
		expectedPaths         []string
		expectedNumberOfPaths int
	}{
		"test1": {
			inputFilename: "test.txt",
			expectedPaths: []string{
				"start,A,b,A,b,A,c,A,end",
				"start,A,b,A,b,A,end",
				"start,A,b,A,b,end",
				"start,A,b,A,c,A,b,A,end",
				"start,A,b,A,c,A,b,end",
				"start,A,b,A,c,A,c,A,end",
				"start,A,b,A,c,A,end",
				"start,A,b,A,end",
				"start,A,b,d,b,A,c,A,end",
				"start,A,b,d,b,A,end",
				"start,A,b,d,b,end",
				"start,A,b,end",
				"start,A,c,A,b,A,b,A,end",
				"start,A,c,A,b,A,b,end",
				"start,A,c,A,b,A,c,A,end",
				"start,A,c,A,b,A,end",
				"start,A,c,A,b,d,b,A,end",
				"start,A,c,A,b,d,b,end",
				"start,A,c,A,b,end",
				"start,A,c,A,c,A,b,A,end",
				"start,A,c,A,c,A,b,end",
				"start,A,c,A,c,A,end",
				"start,A,c,A,end",
				"start,A,end",
				"start,b,A,b,A,c,A,end",
				"start,b,A,b,A,end",
				"start,b,A,b,end",
				"start,b,A,c,A,b,A,end",
				"start,b,A,c,A,b,end",
				"start,b,A,c,A,c,A,end",
				"start,b,A,c,A,end",
				"start,b,A,end",
				"start,b,d,b,A,c,A,end",
				"start,b,d,b,A,end",
				"start,b,d,b,end",
				"start,b,end",
			},
			expectedNumberOfPaths: 36,
		},
		"test2": {
			inputFilename:         "test2.txt",
			expectedNumberOfPaths: 103,
		},
		"test3": {
			inputFilename:         "test3.txt",
			expectedNumberOfPaths: 3509,
		},
	} {
		test := test
		_ = test
		t.Run(name, func(t *testing.T) {
			f, err := os.Open(test.inputFilename)
			require.NoError(t, err)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			connectedCaves := []string{}
			for scanner.Scan() {
				connectedCaves = append(connectedCaves, scanner.Text())
			}

			paths := f2(connectedCaves)
			assert.Equal(t, test.expectedNumberOfPaths, len(paths))
			if len(test.expectedPaths) > 0 {
				assert.ElementsMatch(t, test.expectedPaths, paths)
			}
		})
	}
}
