package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	total := 0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Printf("%s ", line)

		lineString := string(line)
		firstCharRune := []rune(lineString)[0]
		// numberRune := []rune(lineString)[1:len(lineString)]
		// numberInt := int(numberRune - '0')
		numberString := lineString[1:len(lineString)]
		numberInt, err := strconv.Atoi(numberString)

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
