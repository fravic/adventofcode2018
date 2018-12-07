/**
 * Day 7: The Sum of Its Parts
 * https://adventofcode.com/2018/day/7
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type task struct {
	name            rune
	numDependencies int
	dependents      []*task
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	tasksByName := make(map[rune]*task)
	for scanner.Scan() {
		r, err := regexp.Compile(`Step (.) must be finished before step (.) can begin\.`)
		if err != nil {
			fmt.Println("Regex compilation error", err)
		}
		matches := r.FindStringSubmatch(scanner.Text())
		name0, name1 := matches[1][0], matches[2][0]
		task0 := getOrCreateTask(tasksByName, rune(name0))
		task1 := getOrCreateTask(tasksByName, rune(name1))
		task1.numDependencies++
		task0.dependents = append(task0.dependents, task1)
	}
	roots := getRoots(tasksByName)
	fmt.Println(resolveTasks(roots))
}

func getOrCreateTask(tasksByName map[rune]*task, taskName rune) (t *task) {
	t, exists := tasksByName[taskName]
	if !exists {
		t = &task{
			taskName,
			0,
			make([]*task, 0),
		}
		tasksByName[taskName] = t
	}
	return
}

func getRoots(tasksByName map[rune]*task) (roots []*task) {
	roots = make([]*task, 0)
	for _, t := range tasksByName {
		if t.numDependencies == 0 {
			roots = append(roots, t)
		}
	}
	return
}

func resolveTasks(roots []*task) string {
	nextTaskIdx := minIdxByName(roots)
	if nextTaskIdx < 0 {
		return ""
	}
	nextTask := roots[nextTaskIdx]
	roots = append(roots[0:nextTaskIdx], roots[nextTaskIdx+1:]...)
	fmt.Println("Resolving", string(nextTask.name))
	for _, t := range nextTask.dependents {
		t.numDependencies--
		if t.numDependencies == 0 {
			roots = append(roots, t)
		}
	}
	return string(nextTask.name) + resolveTasks(roots)
}

func minIdxByName(tasks []*task) (idx int) {
	idx = -1
	for i, t := range tasks {
		if idx < 0 || t.name < tasks[idx].name {
			idx = i
		}
	}
	return
}
