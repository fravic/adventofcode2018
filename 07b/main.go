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

type worker struct {
	task   *task
	doneAt int
}

const numWorkers = 5

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
	rootTasks, workers := assignTasksToWorkers(getRoots(tasksByName), make([]*worker, 0), 0)
	fmt.Println(doWork(rootTasks, workers))
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

func getRoots(tasksByName map[rune]*task) (rootTasks []*task) {
	rootTasks = make([]*task, 0)
	for _, t := range tasksByName {
		if t.numDependencies == 0 {
			rootTasks = append(rootTasks, t)
		}
	}
	return
}

func assignTasksToWorkers(rootTasks []*task, workers []*worker, currentTime int) ([]*task, []*worker) {
	for len(workers) < numWorkers && len(rootTasks) > 0 {
		nextTaskIdx := minIdxByName(rootTasks)
		workers = append(workers, &worker{
			task:   rootTasks[nextTaskIdx],
			doneAt: currentTime + getTaskDuration(rootTasks[nextTaskIdx]),
		})
		rootTasks = append(rootTasks[0:nextTaskIdx], rootTasks[nextTaskIdx+1:]...)
	}
	return rootTasks, workers
}

func doWork(rootTasks []*task, workers []*worker) int {
	nextWorkerIdx := minIdxByDoneAt(workers)
	nextWorker := workers[nextWorkerIdx]
	workers = append(workers[0:nextWorkerIdx], workers[nextWorkerIdx+1:]...)

	// Finish the worker's task
	fmt.Println("Finished task", string(nextWorker.task.name))
	for _, t := range nextWorker.task.dependents {
		t.numDependencies--
		if t.numDependencies == 0 {
			rootTasks = append(rootTasks, t)
		}
	}

	if len(rootTasks) == 0 && len(workers) == 0 {
		return nextWorker.doneAt
	}

	// Not yet done, assign more tasks until we are!
	rootTasks, workers = assignTasksToWorkers(rootTasks, workers, nextWorker.doneAt)
	return doWork(rootTasks, workers)
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

func minIdxByDoneAt(workers []*worker) (idx int) {
	idx = -1
	for i, t := range workers {
		if idx < 0 || t.doneAt < workers[idx].doneAt {
			idx = i
		}
	}
	return
}

func getTaskDuration(t *task) int {
	return int(t.name) - int('A') + 61
}
