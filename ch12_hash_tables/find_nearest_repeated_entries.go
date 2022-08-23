package ch12_hash_tables

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NearestRepeatedEntries(paragraph []string) int {
	wordToLatestIdx := make(map[string]int)
	minDistance := math.MaxInt64

	for idx, word := range paragraph {
		val, ok := wordToLatestIdx[word]
		if ok {
			minDistance = min(minDistance, idx-val)
		}
		wordToLatestIdx[word] = idx
	}
	if minDistance == math.MaxInt64 {
		// No words repeat - return minus one
		return -1
	}
	return minDistance
}
