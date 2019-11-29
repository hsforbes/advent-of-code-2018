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

		fmt.Printf("%s \n", line)
	}

	// fmt.Printf("\nFinal fabric:\n")
	// printFabric(fabric)

}

type Claim struct {
	xOffset, yOffset, xLength, yLength int
}

func makeClaimFromLine(line string) Claim {
	// segments := strings.Split(line, string(' '))

	return Claim{3, 2, 5, 4}
}

func addClaimToFabric(fabric [][]int, claim Claim) [][]int {

	for xIndex := claim.xOffset; xIndex < claim.xOffset+claim.xLength-1; xIndex++ {
		for yIndex := claim.yOffset; yIndex < claim.yOffset+claim.yLength-1; yIndex++ {
			if fabric[xIndex][yIndex] == 0 {
				fabric[xIndex][yIndex] = 1
			} else {
				fabric[xIndex][yIndex] = 2
			}
		}
	}

	return fabric
}

func makeEmptyFabric() [][]int {
	fabric := [][]int{}
	row := []int{}
	for i := 0; i < 1000; i++ {
		row = append(row, 0)
	}
	for i := 0; i < 1000; i++ {
		fabric = append(fabric, row)
	}
	return fabric
}

func printFabric(fabric [][]int) {
	for x := 0; x < len(fabric); x++ {
		fmt.Printf("Line: %d\n", x)
		fmt.Println(fabric[x])
	}
}
