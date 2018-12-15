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
	numRecipes, _ := strconv.Atoi(os.Args[1])
	elf01 := ring.New(2)
	elf01.Value = startingRecipe01
	elf02 := elf01.Next()
	elf02.Value = startingRecipe02
	lastRecipe := elf02
	recipesSoFar := 2
	for recipesSoFar < numRecipes+10 {
		lastRecipe, recipesSoFar = addNewRecipes(elf01, elf02, lastRecipe, recipesSoFar, numRecipes)
		elf01 = moveElf(elf01)
		elf02 = moveElf(elf02)
	}
	fmt.Print("\n")
}

func addNewRecipes(
	elf01 *ring.Ring,
	elf02 *ring.Ring,
	lastRecipe *ring.Ring,
	recipesSoFar int,
	numRecipes int,
) (*ring.Ring, int) {
	// Add the new recipes
	sum := elf01.Value.(int) + elf02.Value.(int)
	sumStr := strconv.Itoa(sum)

	for _, c := range sumStr {
		newRecipe := ring.New(1)
		newRecipe.Value = int(c - '0')
		lastRecipe.Link(newRecipe)
		lastRecipe = newRecipe
		recipesSoFar++

		if recipesSoFar > numRecipes && recipesSoFar <= numRecipes+10 {
			fmt.Print(string(c))
		}
	}
	return lastRecipe, recipesSoFar
}

func moveElf(elf *ring.Ring) (newElf *ring.Ring) {
	newElf = elf
	for i := 0; i < elf.Value.(int)+1; i++ {
		newElf = newElf.Next()
	}
	return
}
