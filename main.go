package main

import (
	"fmt"
	"log"

	"github.com/tj/docopt"
)

const (
	Usage = `Bee.
Bee prints a word with it's spelling alphabet.

Usage:
  bee <word>...
  bee -h | --help
  bee --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	Version = `1.0.0`
)

var Table = map[rune]string{
	'a': "Alfa",
	'b': "Bravo",
	'c': "Charlie",
	'd': "Delta",
	'e': "Echo",
	'f': "Foxtrot",
	'g': "Golf",
	'h': "Hotel",
	'i': "India",
	'j': "Juliett",
	'k': "Kilo",
	'l': "Lima",
	'm': "Mike",
	'n': "November",
	'o': "Oscar",
	'p': "Papa",
	'q': "Quebec",
	'r': "Romeo",
	's': "Sierra",
	't': "Tango",
	'u': "Uniform",
	'v': "Victor",
	'w': "Whiskey",
	'x': "X-ray",
	'y': "Yankee",
	'z': "Zulu",
	'1': "One",
	'2': "Two",
	'3': "Three",
	'4': "Four",
	'5': "Five",
	'6': "Six",
	'7': "Seven",
	'8': "Eight",
	'9': "Nine",
	'0': "Zero",
}

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, Version, false)
	check(err)

	words := arguments["<word>"].([]string)

	for _, w := range words {
		Spell(w)
		fmt.Println()
	}
}

func Spell(w string) {
	last := len(w) - 1
	for i, c := range w {
		if val, ok := Table[c]; ok {
			fmt.Print(val)
		} else {
			fmt.Print(string(c))
		}
		if i != last {
			fmt.Print(" ")
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
