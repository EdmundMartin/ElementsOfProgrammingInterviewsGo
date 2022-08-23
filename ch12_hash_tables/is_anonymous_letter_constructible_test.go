package ch12_hash_tables

import "testing"

func TestIsLetterConstructibleFromMagazine(t *testing.T) {

	letterText := "Hello, I am a killer"
	magazineTxt := "Hello sexy, where to buy killer heels. I am in shock, and"

	result := IsLetterConstructibleFromMagazine(letterText, magazineTxt)

	if result != true {
		t.Error("expected true")
	}
}
