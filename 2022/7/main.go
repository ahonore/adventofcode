package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	name      string
	parentDir *dir
	subDirs   []*dir
	files     []*file
	size      int // size of cur dir
	fullSize  int // size of cur dir and sub dirs
}

func computeSize(tree *dir) int {
	var sum int
	for _, sd := range tree.subDirs {
		sum += computeSize(sd)
	}
	var curDirSize int
	for _, f := range tree.files {
		curDirSize += f.size
	}
	sum += curDirSize
	tree.size = curDirSize
	tree.fullSize = sum
	return sum
}

func findSize(tree *dir) int {
	var sum int
	for _, sd := range tree.subDirs {
		sum += findSize(sd)
	}
	if tree.fullSize <= 100000 {
		sum += tree.fullSize
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	root := &dir{
		name: "/",
	}
	curPath := root
	for scanner.Scan() {
		cmd := strings.Fields(scanner.Text())
		// not a cmd? so listing cur dir
		if cmd[0] != "$" {
			if cmd[0] == "dir" {
				curPath.subDirs = append(curPath.subDirs, &dir{
					name:      cmd[1],
					parentDir: curPath,
				})
				continue
			}

			i, _ := strconv.Atoi(cmd[0])
			curPath.files = append(curPath.files, &file{
				name: cmd[1],
				size: i,
			})
			continue
		}

		// listing cur dir? see next lines
		if cmd[1] == "ls" {
			continue
		}

		// cmd is "cd"?
		if cmd[1] != "cd" {
			fmt.Printf("unknown cmd: %s\n", cmd[0])
			continue
		}

		switch cmd[2] {
		case "..":
			curPath = curPath.parentDir
		case "/":
			curPath = root
		default: // go to dir
			var subDir *dir
			for _, sd := range curPath.subDirs {
				if sd.name == cmd[2] {
					subDir = sd
					break
				}
			}

			if subDir == nil {
				subDir = &dir{
					name:      cmd[2],
					parentDir: curPath,
				}
				curPath.subDirs = append(curPath.subDirs, subDir)
			}

			curPath = subDir
		}
	}

	computeSize(root)
	fmt.Println(findSize(root))
}
