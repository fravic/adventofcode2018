/**
 * Day 3: No Matter How You Slice It
 * https://adventofcode.com/2018/day/3
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

const fabricSize = 1000

type claim struct {
	x, y, w, h int
	id         string
}

type fabric map[int]map[int]int

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be opened")
	}
	claims := make([]claim, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var id string
		var x, y, w, h int
		fmt.Sscanf(line, "#%s @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		claims = append(claims, claim{
			x, y, w, h, id,
		})
	}
	fabric := constructFabric(claims)
	claimID := findClaimWithNoOverlap(claims, fabric)
	fmt.Println(*claimID)
}

func constructFabric(claims []claim) (fab fabric) {
	fab = make(fabric)
	for i := 0; i < fabricSize; i++ {
		fab[i] = make(map[int]int)
	}
	for _, claim := range claims {
		for i := 0; i < claim.w; i++ {
			for j := 0; j < claim.h; j++ {
				fab[claim.x+i][claim.y+j]++
			}
		}
	}
	return
}

func findClaimWithNoOverlap(claims []claim, fab fabric) *string {
	var res *string
	for _, claim := range claims {
		overlap := false
		for i := 0; i < claim.w; i++ {
			for j := 0; j < claim.h; j++ {
				if fab[claim.x+i][claim.y+j] > 1 {
					overlap = true
				}
			}
		}
		if !overlap {
			res = &claim.id
			break
		}
	}
	return res
}
