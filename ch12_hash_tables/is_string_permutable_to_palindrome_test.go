package ch12_hash_tables

import "testing"

func TestCanFormPalindrome(t *testing.T) {

	exampleOne := "baab"
	result := CanFormPalindrome(exampleOne)
	if result != true {
		t.Error("expected true")
	}

	exampleTwo := "baacb"

	result = CanFormPalindrome(exampleTwo)
	if result != true {
		t.Error("expected true")
	}

	exampleThree := "baacdb"
	result = CanFormPalindrome(exampleThree)
	if result != false {
		t.Error("expected false")
	}
}
