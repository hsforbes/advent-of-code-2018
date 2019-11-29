package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	fabric := makeEmptyFabric()

	// fmt.Println("Let's try to print the empty fabric:")
	// printFabric(fabric)

	for {
		lineBytes, _, err := reader.ReadLine()
		line := string(lineBytes)

		if err == io.EOF {
			break
		}

		claim := makeClaimFromLine(line)

		fabric = addClaimToFabric(fabric, claim)

		// fmt.Printf("%s \n", line)
	}

	// fmt.Printf("\nFinal fabric:\n")
	// printFabric(fabric)

	countOfMultipleClaims := countMultipleClaims(fabric)

	fmt.Printf("\nCount of coordinates with multiple claims: %d", countOfMultipleClaims)

}

type Claim struct {
	xOffset, yOffset, xLength, yLength int
}

func makeClaimFromLine(line string) Claim {
	segments := strings.Split(line, string(' '))

	offsetsString := strings.ReplaceAll(segments[2], ":", "")
	offsets := strings.Split(offsetsString, ",")

	dimensions := strings.Split(segments[3], "x")

	xOffset, err := strconv.Atoi(offsets[0])
	yOffset, err := strconv.Atoi(offsets[1])
	xDimension, err := strconv.Atoi(dimensions[0])
	yDimension, err := strconv.Atoi(dimensions[1])
	if err != nil {
		panic(err)
	}

	return Claim{xOffset, yOffset, xDimension, yDimension}
}

func addClaimToFabric(fabric [][]int, claim Claim) [][]int {

	for xIndex := claim.xOffset; xIndex < claim.xOffset+claim.xLength; xIndex++ {
		for yIndex := claim.yOffset; yIndex < claim.yOffset+claim.yLength; yIndex++ {
			if fabric[xIndex][yIndex] == 0 {
				fabric[xIndex][yIndex] = 1
			} else if fabric[xIndex][yIndex] == 1 {
				fabric[xIndex][yIndex] = 2
			}
		}
	}

	return fabric
}

func makeEmptyFabric() [][]int {
	fabric := [][]int{}
	for i := 0; i < 1000; i++ {
		column := []int{}
		for j := 0; j < 1000; j++ {
			column = append(column, 0)
		}
		fabric = append(fabric, column)
	}
	return fabric
}

func printFabric(fabric [][]int) {
	for y := 0; y < len(fabric[0]); y++ {
		// fmt.Printf("\nLine: %d", y)
		fmt.Println()
		for x := 0; x < len(fabric); x++ {
			fmt.Printf("%d", fabric[x][y])
		}
	}
}

func countMultipleClaims(fabric [][]int) int {
	countOfMultipleClaims := 0
	for y := 0; y < len(fabric[0]); y++ {
		for x := 0; x < len(fabric); x++ {
			if fabric[x][y] == 2 {
				countOfMultipleClaims++
			}
		}
	}

	return countOfMultipleClaims
}
