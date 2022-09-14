package ch8_stacks_queues

import "strings"

func IsWellFormed(s string) bool {
	leftChars := []string{}
	lookUp := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for _, char := range strings.Split(s, "") {

		_, ok := lookUp[char]
		if ok {
			leftChars = append(leftChars, char)
		} else if len(leftChars) == 0 || lookUp[leftChars[len(leftChars)-1]] != char {
			return false
		} else {
			leftChars = leftChars[:len(leftChars)-1]
		}
	}
	return len(leftChars) == 0
}
