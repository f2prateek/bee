package main_test

import (
	"fmt"

	"github.com/f2prateek/bee"
)

func ExampleCharacters() {
	main.Spell("aBcdefghIjkl")
	fmt.Println()
	main.Spell("MnopqRstuvw")
	fmt.Println()
	main.Spell("xyz")

	// Output:
	// Alfa Bravo Charlie Delta Echo Foxtrot Golf Hotel India Juliett Kilo Lima
	// Mike November Oscar Papa Quebec Romeo Sierra Tango Uniform Victor Whiskey
	// X-ray Yankee Zulu
}

func ExampleDigits() {
	main.Spell("1234567890")

	// Output:
	// 1 2 3 4 5 6 7 8 9 0
}

func ExampleUnknown() {
	main.Spell(", ./\\")

	// Output:
	// ,   . / \
}
