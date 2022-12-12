package main

import (
	"fmt"
)

type monkey struct {
	items        []int
	op           func(int) int
	testDiv      int
	throwOnTrue  int
	throwOnFalse int
}

var monkeysTest = []monkey{
	{
		items:        []int{79, 98},
		op:           func(i int) int { return i * 19 },
		testDiv:      23,
		throwOnTrue:  2,
		throwOnFalse: 3,
	},
	{
		items:        []int{54, 65, 75, 74},
		op:           func(i int) int { return i + 6 },
		testDiv:      19,
		throwOnTrue:  2,
		throwOnFalse: 0,
	},
	{
		items:        []int{79, 60, 97},
		op:           func(i int) int { return i * i },
		testDiv:      13,
		throwOnTrue:  1,
		throwOnFalse: 3,
	},
	{
		items:        []int{74},
		op:           func(i int) int { return i + 3 },
		testDiv:      17,
		throwOnTrue:  0,
		throwOnFalse: 1,
	},
}

var monkeys = []monkey{
	{
		items:        []int{66, 79},
		op:           func(i int) int { return i * 11 },
		testDiv:      7,
		throwOnTrue:  6,
		throwOnFalse: 7,
	},
	{
		items:        []int{84, 94, 94, 81, 98, 75},
		op:           func(i int) int { return i * 17 },
		testDiv:      13,
		throwOnTrue:  5,
		throwOnFalse: 2,
	},
	{
		items:        []int{85, 79, 59, 64, 79, 95, 67},
		op:           func(i int) int { return i + 8 },
		testDiv:      5,
		throwOnTrue:  4,
		throwOnFalse: 5,
	},
	{
		items:        []int{70},
		op:           func(i int) int { return i + 3 },
		testDiv:      19,
		throwOnTrue:  6,
		throwOnFalse: 0,
	},
	{
		items:        []int{57, 69, 78, 78},
		op:           func(i int) int { return i + 4 },
		testDiv:      2,
		throwOnTrue:  0,
		throwOnFalse: 3,
	},
	{
		items:        []int{65, 92, 60, 74, 72},
		op:           func(i int) int { return i + 7 },
		testDiv:      11,
		throwOnTrue:  3,
		throwOnFalse: 4,
	},
	{
		items:        []int{77, 91, 91},
		op:           func(i int) int { return i * i },
		testDiv:      17,
		throwOnTrue:  1,
		throwOnFalse: 7,
	},
	{
		items:        []int{76, 58, 57, 55, 67, 77, 54, 99},
		op:           func(i int) int { return i + 6 },
		testDiv:      3,
		throwOnTrue:  2,
		throwOnFalse: 1,
	},
}

func main() {
	monkeysList := monkeys
	rounds := 20
	countInspectedItems := make([]int, len(monkeysList))
	for r := 0; r < rounds; r++ {
		for i := 0; i < len(monkeysList); i++ {
			items := monkeysList[i].items
			monkeysList[i].items = []int{}
			countInspectedItems[i] += len(items)
			for _, item := range items {
				newItem := monkeysList[i].op(item) / 3
				throw := monkeysList[i].throwOnFalse
				if newItem%monkeysList[i].testDiv == 0 {
					throw = monkeysList[i].throwOnTrue
				}

				monkeysList[throw].items = append(monkeysList[throw].items, newItem)
			}
		}
	}
	fmt.Println(countInspectedItems)
	fmt.Println(227 * 225)
}
