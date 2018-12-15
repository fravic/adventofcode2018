/**
 * Day 14: Chocolate Charts
 * https://adventofcode.com/2018/day/14
 */

package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

const startingRecipe01 = 3
const startingRecipe02 = 7

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go numRecipes")
		os.Exit(1)
	}
	searchString := os.Args[1]
	elf01 := ring.New(2)
	elf01.Value = startingRecipe01
	elf02 := elf01.Next()
	elf02.Value = startingRecipe02
	lastRecipe := elf02
	inputChecker := elf01
	recipeCount := 2
	for true {
		lastRecipe, inputChecker, recipeCount = addNewRecipes(elf01, elf02, lastRecipe, inputChecker, recipeCount, searchString)
		elf01 = moveElf(elf01)
		elf02 = moveElf(elf02)
	}
}

func addNewRecipes(
	elf01 *ring.Ring,
	elf02 *ring.Ring,
	lastRecipe *ring.Ring,
	inputChecker *ring.Ring,
	recipeCount int,
	searchString string,
) (*ring.Ring, *ring.Ring, int) {
	sum := elf01.Value.(int) + elf02.Value.(int)
	sumStr := strconv.Itoa(sum)

	for _, c := range sumStr {
		newRecipe := ring.New(1)
		newRecipe.Value = int(c - '0')
		lastRecipe.Link(newRecipe)
		lastRecipe = newRecipe
		recipeCount++

		// When we add each new recipe, now need to check if the last characters match
		if recipeCount > len(searchString) {
			inputChecker = inputChecker.Next()
			if checkIfCharactersMatch(searchString, inputChecker) {
				fmt.Printf("Found match after %d recipes\n", recipeCount-len(searchString))
				os.Exit(0)
			}
		}
	}

	return lastRecipe, inputChecker, recipeCount
}

func checkIfCharactersMatch(matchString string, checkFrom *ring.Ring) (allMatch bool) {
	allMatch = true
	checker := checkFrom
	for i := 0; i < len(matchString); i++ {
		if checker.Value != int(matchString[i]-'0') {
			allMatch = false
			break
		}
		checker = checker.Next()
	}
	return
}

func moveElf(elf *ring.Ring) (newElf *ring.Ring) {
	newElf = elf
	for i := 0; i < elf.Value.(int)+1; i++ {
		newElf = newElf.Next()
	}
	return
}
