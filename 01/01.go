package main

import (
	"fmt"
	"sort"
)

func main() {

	foundValues := []int{3, 5, 7}
	
	newValue := 6
	
	suggInd := sort.SearchInts(foundValues, newValue)
	
	fmt.Printf("Suggested index:\t%d\n", suggInd)
}
