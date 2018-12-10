/**
 * Day 9: Marble Mania
 * https://adventofcode.com/2018/day/9
 */

package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

// These are the set-in-stone rules of the game
const moveBackwardsSpaces = 7
const scoringMarbleFactor = 23

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go numPlayers lastMarble")
		os.Exit(1)
	}
	numPlayers, _ := strconv.Atoi(os.Args[1])
	lastMarble, _ := strconv.Atoi(os.Args[2])

	// container/ring implements a circular linked list
	current := ring.New(1)
	current.Value = 0
	scores := make(map[int]int)
	currentPlayer := 0
	for v := 1; v <= lastMarble; v++ {
		if v%scoringMarbleFactor == 0 {
			// Score a scoring marble and remove the other one
			scores[currentPlayer] += v
			removed := current.Move(-moveBackwardsSpaces - 1).Unlink(1)
			scores[currentPlayer] += removed.Value.(int)
			current = current.Move(-moveBackwardsSpaces + 1)
		} else {
			// Add a normal marble
			new := ring.New(1)
			new.Value = v
			next := current.Next()
			next.Link(new)
			current = new
		}
		currentPlayer = (currentPlayer + 1) % numPlayers
	}
	maxScore := 0
	for _, i := range scores {
		if i > maxScore {
			maxScore = i
		}
	}
	fmt.Println("Top score", maxScore)
}
