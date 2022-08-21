package ch12_hash_tables

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {

	dictionary := []string{
		"debitcard",
		"elvis",
		"silent",
		"badcredit",
		"lives",
		"freedom",
		"listen",
		"levis",
		"money",
	}

	results := FindAnagrams(dictionary)
	fmt.Println(results)
}