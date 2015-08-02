package main_test

import (
	"fmt"

	"github.com/f2prateek/bee"
)

func ExampleCharacters() {
	main.Spell("abcdefghijkl")
	fmt.Println()
	main.Spell("mnopqrstuvw")
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
	// One Two Three Four Five Six Seven Eight Nine Zero
}
