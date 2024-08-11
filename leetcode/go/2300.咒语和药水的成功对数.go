package main

import (
	"sort"
)

func main() {
	p := []int{1, 3, 5, 2, 4}
	s := []int{5, 1, 3}
	successfulPairs(s, p, 7)
}

// func successfulPairs(spells []int, potions []int, success int64) []int {
// 	// 排序后双指针
// 	sort.Ints(spells)
// 	sort.Ints(potions)

// 	spellsPointer := 0
// 	potionsPointer := len(potions) - 1

// 	for spellsPointer < len(spells) {

// 	}

// 	return nil
// }

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	idx := make([]int, len(spells))
	for i, _ := range idx {
		idx[i] = i
	}
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] > potions[j]
	})
	sort.Slice(idx, func(i, j int) bool {
		return spells[idx[i]] < spells[idx[j]]
	})
	j := 0
	for _, p := range idx {
		v := spells[p]
		for j < len(potions) && int64(potions[j])*int64(v) >= success {
			j++
		}
		res[p] = j
	}
	return res
}
