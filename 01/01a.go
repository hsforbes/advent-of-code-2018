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

	// file, err := os.Open("./input.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// reader := bufio.NewReader(file)
	buffer, _ := ioutil.ReadFile("./input.txt")

	total := 0

	for _, line := range strings.Split(string(buffer), "\n") {
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
