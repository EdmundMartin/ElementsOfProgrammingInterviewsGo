package ch12_hash_tables

import (
	"sort"
	"strings"
)

func FindAnagrams(dictionary []string) [][]string {
	sortedStringAnagrams := make(map[string][]string)

	for _, s := range dictionary {
		original := s
		slice := strings.Split(s, "")
		sort.Strings(slice)

		s = strings.Join(slice, "")
		_, ok := sortedStringAnagrams[s]
		if !ok {
			sortedStringAnagrams[s] = []string{}
		}
		val, _ := sortedStringAnagrams[s]
		val = append(val, original)
		sortedStringAnagrams[s] = val
	}

	results := [][]string{}

	for _, value := range sortedStringAnagrams {
		if len(value) >= 2 {
			results = append(results, value)
		}
	}
	return results
}
