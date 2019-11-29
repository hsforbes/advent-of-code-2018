package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	linesRead := []string{}

	for {
		lineBytes, _, err := reader.ReadLine()
		line := string(lineBytes)

		if err == io.EOF {
			break
		}

		linesRead = append(linesRead, line)

		// fmt.Println();
		// fmt.Printf("\n%s", line)

	}

	sort.Strings(linesRead)

	// fmt.Println("Made sorted slice of strings:")
	// fmt.Println(linesRead)

	for lineIndex := 0; lineIndex < len(linesRead)-1; lineIndex++ {
		// commonChars := []rune{}

		for charIndex := 0; charIndex < 26; charIndex++ {
			if linesRead[lineIndex][charIndex] == linesRead[lineIndex+1][charIndex] {
				if charIndex > 19 {
					fmt.Printf("\nMatching index\t%d", charIndex)
					fmt.Printf("\nLine 1:\t%s\tLine 2:\t%s\n", linesRead[lineIndex], linesRead[lineIndex+1])
				}
				if charIndex == 25 {
					fmt.Printf("\nFound the match:\t%s", linesRead[lineIndex])
				}
			} else {
				break
			}
		}
	}

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
	for i := 0; i < len(charCounts); i++ {
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

	for i := 0; i < len(line); i++ {
		// fmt.Printf("\t # Char: %s ", string(line[i]))
		charCounts = addKeyCharToCharCounts(charCounts, line[i])
	}

	return charCounts
}

func addKeyCharToCharCounts(charCounts []CharCount, keyChar byte) []CharCount {

	foundChar := false
	returnCharCounts := charCounts

	// TODO Improve by keeping charCounts sorted by letter
	for i := 0; i < len(charCounts); i++ {
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
	char  byte
	count int
}
