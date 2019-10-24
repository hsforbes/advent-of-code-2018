package main

import (
	"fmt"
	"io/ioutil"
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

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		// fmt.Printf("Line: %s\n", line)

		if len(line) == 0 {
			continue
		}

		fmt.Printf("%s ", line)

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

		fmt.Printf("\t# current total: %d\n", total)
	}

}
