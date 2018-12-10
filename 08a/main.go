/**
 * Day 8: Memory Maneuver
 * https://adventofcode.com/2018/day/8
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	fmt.Println("Metadata sum:", getMetadataSum(scanner))
}

func getMetadataSum(rootNodeScanner *bufio.Scanner) (sum int) {
	// Read the quantity of child nodes
	rootNodeScanner.Scan()
	childCount, _ := strconv.Atoi(rootNodeScanner.Text())

	// Read the quantity of metadata entries
	rootNodeScanner.Scan()
	metadataCount, _ := strconv.Atoi(rootNodeScanner.Text())

	// Recursively scan the child nodes if necessary
	for i := 0; i < childCount; i++ {
		sum += getMetadataSum(rootNodeScanner)
	}

	// Sum the metadata entries
	for i := 0; i < metadataCount; i++ {
		rootNodeScanner.Scan()
		metadata, _ := strconv.Atoi(rootNodeScanner.Text())
		sum += metadata
	}

	return
}
