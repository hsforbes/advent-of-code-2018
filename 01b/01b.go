package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	buffer, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(buffer), "\n")

	total := 0

	// for i := 0; i < len(lines); i++ {
	// 	fmt.Printf("Line\t%d:\t%s\n", i, lines[i])
	// }

	// for _, line := range strings.Split(string(buffer), "\n") {

	foundValues := []int{}

	for finished := false; finished == false; {
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			// fmt.Printf("%s ", line)

			if len(line) == 0 {
				if finished == false {
					i = -1
				} else {
					break
				}
			} else {

				lineString := string(line)
				firstCharRune := []rune(lineString)[0]
				// numberRune := []rune(lineString)[1:len(lineString)]
				// numberInt := int(numberRune - '0')
				numberString := lineString[1:len(lineString)]
				numberInt, err := strconv.Atoi(numberString)
				if err != nil {
					panic(err)
				}

				// fmt.Printf("%d ", numberInt)

				if "+" == string(firstCharRune) {
					// fmt.Printf("Found plus   %c ", line[0])
					total += numberInt
				} else {
					total -= numberInt
				}

				suggestedIndex := sort.SearchInts(foundValues, total)

				// Check for puzzle completion
				if suggestedIndex != len(foundValues) {
					if foundValues[suggestedIndex] == total {
						finished = true
						fmt.Printf("\n* * *\nFinished! First duplicate total:\t%d\n* * *\n", total)
					}
				}
				// Otherwise, move on

				// Put the new total in the right place in the sorted slice
				foundValues = append(foundValues, total)
				sort.Ints(foundValues)

				fmt.Printf("\t# Current total: %d", total)

				// fmt.Printf("\t# Found values len: %d", len(foundValues))
				// fmt.Printf("\t# Found Values: %d", foundValues)
				// for j := 0; j < len(foundValues); j++ {
				// 	fmt.Printf(" %d,", foundValues)
				// }

				// fmt.Println()
			}
		}
	}

}
