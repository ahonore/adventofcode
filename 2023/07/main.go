package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	_ "embed"

	"gonum.org/v1/gonum/stat/combin"
)

//go:embed input.txt
var input string

const handCount = 5

var cardWeightMap = map[byte]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}

type hand struct {
	cards [handCount]byte
	bid   int
}

type byHandStrength []hand

func (a byHandStrength) Len() int      { return len(a) }
func (a byHandStrength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byHandStrength) Less(i, j int) bool {
	h0 := a[i].cards
	var h0Value [13]int // count among 13 cards
	for _, b := range h0 {
		h0Value[cardWeightMap[b]]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(h0Value[:])))
	var vh0 int
	for id, v := range h0Value {
		if v == 0 {
			vh0 = vh0 * int(math.Pow(10, float64(handCount-id)))
			break
		}
		vh0 = vh0*10 + v
	}
	h1 := a[j].cards
	var h1Value [13]int // count among 13 cards
	for _, b := range h1 {
		h1Value[cardWeightMap[b]]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(h1Value[:])))
	var vh1 int
	for id, v := range h1Value {
		if v == 0 {
			vh1 = vh1 * int(math.Pow(10, float64(handCount-id)))
			break
		}
		vh1 = vh1*10 + v
	}

	if vh0 != vh1 {
		return vh0 < vh1
	}

	for c := 0; c < handCount; c++ {
		cw0 := cardWeightMap[h0[c]]
		cw1 := cardWeightMap[h1[c]]
		if cw0 == cw1 {
			continue
		}

		return cw0 < cw1
	}

	return true // must never happen
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Fields(line)
		i, _ := strconv.Atoi(str[1])
		var cards [handCount]byte
		copy(cards[:], []byte(str[0]))
		hands = append(hands, hand{
			cards: cards,
			bid:   i,
		})
	}
	sort.Sort(byHandStrength(hands))
	var sum int
	for id, h := range hands {
		sum += (id + 1) * h.bid
	}
	fmt.Println(sum)
}

var cardWeightMap2 = map[byte]int{'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12}

type hand2 struct {
	originalCards [handCount]byte
	bestType      int
	bid           int
}

// type to value
// Five of a kind  -> 10000
// Four of a kind  -> 01001
// Full house      -> 00110
// Three of a kind -> 00102
// Two pair        -> 00021
// One pair        -> 00013
// High card       -> 00005
func findType(hand [handCount]byte) int {
	ctMap := map[byte]int{}
	for _, c := range hand {
		ctMap[c]++
	}
	var typ int
	for _, ct := range ctMap {
		switch ct {
		case 1:
			typ += 1
		case 2:
			typ += 10
		case 3:
			typ += 100
		case 4:
			typ += 1000
		case 5:
			typ += 10000
		}
	}
	return typ
}

func findBestType(hand [handCount]byte) int {
	// count jokers and fill the accuTable
	var countJokers int
	handCardsCount := map[byte]int{} // count among 13 cards
	// do not put jokers at first
	for _, b := range hand {
		if b == 'J' {
			countJokers++
			continue
		}
		handCardsCount[b]++
	}

	if countJokers == 0 {
		return findType(hand)
	}

	// make list of all cards in hand, add the joker as a possibility
	cardsList := []byte{'J'}
	for c := range handCardsCount {
		cardsList = append(cardsList, c)
	}

	var combiList []int
	for i := 0; i < countJokers; i++ {
		combiList = append(combiList, len(cardsList))
	}

	var bestType int
	for _, cIds := range combin.Cartesian(combiList) {
		var s string
		for card, ct := range handCardsCount {
			s += strings.Repeat(string(card), ct)
		}
		for _, id := range cIds {
			s += string(cardsList[id])
		}
		var newHand [handCount]byte
		copy(newHand[:], []byte(s))
		typ := findType(newHand)
		if typ > bestType {
			bestType = typ
		}
	}
	return bestType
}

type byHand2Strength []hand2

func (a byHand2Strength) Len() int      { return len(a) }
func (a byHand2Strength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byHand2Strength) Less(i, j int) bool {

	if a[i].bestType == a[j].bestType {
		for c := 0; c < handCount; c++ {
			cw0 := cardWeightMap2[a[i].originalCards[c]]
			cw1 := cardWeightMap2[a[j].originalCards[c]]
			if cw0 == cw1 {
				continue
			}

			return cw0 < cw1
		}
	}

	return a[i].bestType < a[j].bestType
}

func part2() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var hands []hand2
	for scanner.Scan() {
		var s string
		var h hand2
		fmt.Sscanf(scanner.Text(), "%5s %d", &s, &h.bid)
		copy(h.originalCards[:], []byte(s))
		hands = append(hands, h)
	}
	for i := 0; i < len(hands); i++ {
		hands[i].bestType = findBestType(hands[i].originalCards)
	}
	sort.Sort(byHand2Strength(hands))
	var sum int
	for id, h := range hands {
		sum += (id + 1) * h.bid
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
