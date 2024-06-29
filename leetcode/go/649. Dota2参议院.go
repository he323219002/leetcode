package main

import "fmt"

func main() {
	// leetcode 答案
	senate := "RRDDD"
	res := predictPartyVictory(senate)
	fmt.Println(res)
}

func predictPartyVictory(senate string) string {
	radiant := make([]int, 0)
	dire := make([]int, 0)

	for index, str := range senate {
		if str == 'R' {
			radiant = append(radiant, index)
		} else {
			dire = append(dire, index)
		}
	}

	for len(radiant) != 0 && len(dire) != 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) != 0 {
		return "Radiant"
	}
	return "Dire"
}
