package ch12_hash_tables

func IsLetterConstructibleFromMagazine(letterTxt string, magazineTxt string) bool {
	usableChars := make(map[int32]int)
	for _, char := range magazineTxt {
		_, ok := usableChars[char]
		if !ok {
			usableChars[char] = 1
		} else {
			usableChars[char]++
		}
	}

	for _, char := range letterTxt {
		val, ok := usableChars[char]
		if !ok {
			return false
		}
		if val-1 < 0 {
			return false
		}
		usableChars[char]--
	}
	return true
}
