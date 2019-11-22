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

		// fmt.Println();
		// fmt.Printf("\n%s", line)

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

	charCounts := findCharCounts(line)

	lineAnswer := LineAnswer{false, false}
	/* 
	look through chars for two or three counts
	set lineAnswer
	*/
	for i:=0; i < len(charCounts); i++ {
		if charCounts[i].count == 2 {
			lineAnswer.hasExactlyTwoOfALetter = true
		} 
		if charCounts[i].count == 3 {
			lineAnswer.hasExactlyThreeOfALetter = true
		}
	}

	return lineAnswer
}


func findCharCounts(line string) []CharCount {

	var charCounts []CharCount

	for i:=0; i < len(line); i++ {
		// fmt.Printf("\t # Char: %s ", string(line[i]))
		charCounts = addKeyCharToCharCounts(charCounts, line[i])
	}

	return charCounts
}


func addKeyCharToCharCounts(charCounts []CharCount, keyChar byte) []CharCount {

	foundChar := false
	returnCharCounts := charCounts

	// TODO Improve by keeping charCounts sorted by letter
	for i:=0; i < len(charCounts); i++ {
		if charCounts[i].char == keyChar {
			returnCharCounts[i].count++
			foundChar = true
			break
		}
	}

	// If a new letter was found
	if !foundChar {
		returnCharCounts = append(returnCharCounts, CharCount{keyChar, 1})
	}

	return returnCharCounts
}


type CharCount struct {
	char byte
	count int
}
