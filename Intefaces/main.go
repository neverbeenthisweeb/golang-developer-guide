package main

import "fmt"

type bot interface {
	getLetterCount() int
}

type englishBot struct {
	word     int
	sentence int
}

type spanishBot struct {
	word     int
	sentence int
}

func main() {

	englishy := englishBot{
		word:     10,
		sentence: 25,
	}

	spanishy := spanishBot{
		word:     5,
		sentence: 15,
	}

	printLetterCount(englishy)
	printLetterCount(spanishy)

}

func (e englishBot) getLetterCount() int {
	return e.word + e.sentence
}

func (s spanishBot) getLetterCount() int {
	return s.word + s.sentence
}

func printLetterCount(b bot) {
	fmt.Printf("Sum of letters: %v\n", b.getLetterCount())
}
