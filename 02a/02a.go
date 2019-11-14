package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	linesWithExactlyTwoOfALetter, linesWithExactlyThreeOfALetter := 0, 0

	for {
		lineBytes, _, err := reader.ReadLine()
		line := string(lineBytes)

		if err == io.EOF {
			break
		}

		fmt.Println();
		fmt.Printf("\n%s", line)

		lineAnswer := countTwoOrThreeRepeatedLetters(line)
		
		if(lineAnswer.hasExactlyTwoOfALetter) {
			linesWithExactlyTwoOfALetter++
		}
		if(lineAnswer.hasExactlyThreeOfALetter) {
			linesWithExactlyThreeOfALetter++
		}

	}

	fmt.Printf("\nlinesWithExactlyTwoOfALetter:\t%d\t#\tlinesWithExactlyThreeOfALetter:\t%d", linesWithExactlyTwoOfALetter, linesWithExactlyThreeOfALetter)
	fmt.Printf("\nChecksum:\t%d", linesWithExactlyTwoOfALetter * linesWithExactlyThreeOfALetter)

}

type LineAnswer struct {
	hasExactlyTwoOfALetter, hasExactlyThreeOfALetter bool
}

func countTwoOrThreeRepeatedLetters(line string) LineAnswer {

	findCharCounts(line)

	/* 
	look through chars for two or three counts
	set lineAnswer
	*/

	lineAnswer := LineAnswer{true, true}
	return lineAnswer
}

func findCharCounts(line string) []CharCount {
	var charCounts []CharCount

	for i:=0; i < len(line); i++ {
		fmt.Printf("\t # Char: %s ", string(line[i]))
	}

	a := CharCount{13, 1}
	// b := CharCount(13, 1)

	charCounts = append(charCounts, a)
	charCounts = append(charCounts, a)
	// charCounts = append(charCounts, b)

	return charCounts
}

type CharCount struct {
	char byte
	count int
}
