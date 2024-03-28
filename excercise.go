package main

// https://go.dev/tour/moretypes/23

import (
	"strings"
)

func WordCount(s string) map[string]int {

	answer := map[string]int{}
	substrings := strings.Split(s, " ")

	for i := 0; i < len(substrings); i++ {
		answer[substrings[i]] += 1
	}
	return answer
}
