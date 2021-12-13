package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func f1CopyMap(m map[string]struct{}) map[string]struct{} {
	dst := map[string]struct{}{}
	for k := range m {
		dst[k] = struct{}{}
	}
	return dst
}

func f1VisitCaves(currentPath []string, alreadyVisitedCaves map[string]struct{}, readOnlyConnectedCaves map[string][]string) []string {
	if currentPath[len(currentPath)-1] == "end" {
		return []string{strings.Join(currentPath, ",")}
	}

	paths := []string{}
	for _, cave := range readOnlyConnectedCaves[currentPath[len(currentPath)-1]] {
		if _, ok := alreadyVisitedCaves[cave]; ok && unicode.IsLower(rune(cave[0])) {
			// small cave yet visited
			continue
		}

		newAlreadyVisitedCaves := f1CopyMap(alreadyVisitedCaves)
		newAlreadyVisitedCaves[cave] = struct{}{}
		newPath := make([]string, len(currentPath))
		copy(newPath, currentPath)
		newPath = append(newPath, cave)
		paths = append(paths, f1VisitCaves(newPath, newAlreadyVisitedCaves, readOnlyConnectedCaves)...)
	}
	return paths
}

func f1(caves []string) []string {
	cavesMap := map[string][]string{}
	for _, s := range caves {
		nodes := strings.Split(s, "-")
		cavesMap[nodes[0]] = append(cavesMap[nodes[0]], nodes[1])
		cavesMap[nodes[1]] = append(cavesMap[nodes[1]], nodes[0])
	}
	path := []string{"start"}
	visitedSmallCaves := map[string]struct{}{
		"start": {},
	}
	return f1VisitCaves(
		path,
		visitedSmallCaves,
		cavesMap,
	)
}

func f2CopyMap(m map[string]int) map[string]int {
	dst := map[string]int{}
	for k, v := range m {
		dst[k] = v
	}
	return dst
}

func f2VisitCaves(currentPath []string, oneSmallCaveVisitedTwice bool, alreadyVisitedCaves map[string]int, readOnlyConnectedCaves map[string][]string) []string {
	if currentPath[len(currentPath)-1] == "end" {
		return []string{strings.Join(currentPath, ",")}
	}

	paths := []string{}
	for _, cave := range readOnlyConnectedCaves[currentPath[len(currentPath)-1]] {
		if cave == "start" {
			continue
		}

		newOneSmallCaveVisitedTwice := oneSmallCaveVisitedTwice
		v := alreadyVisitedCaves[cave]
		switch v {
		case 1:
			if unicode.IsLower(rune(cave[0])) {
				if oneSmallCaveVisitedTwice {
					continue
				}

				newOneSmallCaveVisitedTwice = true
			}
		case 2:
			if unicode.IsLower(rune(cave[0])) {
				continue
			}
		}

		newAlreadyVisitedCaves := f2CopyMap(alreadyVisitedCaves)
		newAlreadyVisitedCaves[cave]++
		newPath := make([]string, len(currentPath))
		copy(newPath, currentPath)
		newPath = append(newPath, cave)
		paths = append(paths, f2VisitCaves(newPath, newOneSmallCaveVisitedTwice, newAlreadyVisitedCaves, readOnlyConnectedCaves)...)
	}
	return paths
}

func f2(caves []string) []string {
	cavesMap := map[string][]string{}
	for _, s := range caves {
		nodes := strings.Split(s, "-")
		cavesMap[nodes[0]] = append(cavesMap[nodes[0]], nodes[1])
		cavesMap[nodes[1]] = append(cavesMap[nodes[1]], nodes[0])
	}
	path := []string{"start"}
	visitedSmallCaves := map[string]int{
		"start": 1,
	}
	return f2VisitCaves(
		path,
		false,
		visitedSmallCaves,
		cavesMap,
	)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	connectedCaves := []string{}
	for scanner.Scan() {
		connectedCaves = append(connectedCaves, scanner.Text())
	}

	fmt.Println(len(f1(connectedCaves)))
	fmt.Println(len(f2(connectedCaves)))
}
