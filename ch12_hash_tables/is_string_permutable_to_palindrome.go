package ch12_hash_tables

func CanFormPalindrome(contents string) bool {
	charCounter := make(map[int32]int)
	for _, char := range contents {
		_, ok := charCounter[char]
		if !ok {
			charCounter[char] = 1
		} else {
			charCounter[char]++
		}
	}
	oddCount := 0
	for _, value := range charCounter {
		if value%2 != 0 {
			oddCount += 1
		}
	}
	return oddCount <= 1
}
