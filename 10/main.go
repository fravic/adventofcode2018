/**
 * Day 10: The Stars Align
 * https://adventofcode.com/2018/day/10
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const imgW = 200
const imgH = 50

type star struct {
	x, y, vx, vy int
}

func main() {
	stars := readInitialPositions()
	b := make([]byte, 1) // Throwaway buffer
	seconds := 0

	for {
		img := make([]bool, imgW*imgH)
		anythingVisible := false

		camX, camY := getCameraCenter(stars)
		for _, s := range stars {
			if s.x >= camX && s.x < camX+imgW && s.y >= camY && s.y < camY+imgH {
				img[(s.y-camY)*imgW+(s.x-camX)] = true
				anythingVisible = true
			}
			s.x += s.vx
			s.y += s.vy
		}

		if anythingVisible {
			outputImage(img)
			fmt.Println(seconds)

			// When any key is pressed, display the next image
			os.Stdin.Read(b)
		}
		seconds++
	}
}

func readInitialPositions() (stars []*star) {
	stars = make([]*star, 0)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		r, _ := regexp.Compile(`^position=\<([-| ]*\d+), ([-| ]*\d+)\> velocity=\<([-| ]*\d+), ([-| ]*\d+)\>$`)
		matches := r.FindStringSubmatch(t)
		stars = append(stars, &star{
			atoi(matches[1]),
			atoi(matches[2]),
			atoi(matches[3]),
			atoi(matches[4]),
		})
	}
	return
}

func getCameraCenter(stars []*star) (x, y int) {
	// Center the camera around the stars' average position
	for _, s := range stars {
		x += s.x
		y += s.y
	}
	x /= len(stars)
	y /= len(stars)
	x -= imgW / 2
	y -= imgH / 2
	return
}

func outputImage(img []bool) {
	fmt.Printf("\033[0;0H") // Reset print position to top left
	printHorizontalLine()
	for y := 0; y < imgH; y++ {
		fmt.Print("|")
		for x := 0; x < imgW; x++ {
			if img[y*imgW+x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|\n")
	}
	printHorizontalLine()
}

func printHorizontalLine() {
	for x := 0; x < imgW+2; x++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func atoi(str string) int {
	res, _ := strconv.Atoi(strings.Trim(str, " "))
	return res
}
