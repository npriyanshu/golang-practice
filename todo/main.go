package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	id          int
	title       string
	description string
	complete    bool
}

var tasks []Task

var id = 1

func main() {
	fmt.Println("hello welcome to task list")
	reader := bufio.NewReader(os.Stdin)
loop:
	for {
		fmt.Println("1. Add Task 2. Update Task 3. complete Task 4. Show Tasks 5. Exit")

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)
		inputInt, _ := strconv.Atoi(input)

		switch inputInt {
		case 1:
			addTask(reader)
			showTasks()
		case 2:
			updateTask(reader)
			showTasks()
		case 3:
			completeTask(reader)
			showTasks()
		case 4:
			showTasks()
		default:
			fmt.Println("exit")
			break loop
		}
	}

}

func addTask(reader *bufio.Reader) {
	fmt.Println("enter task title :")
	title, _ := reader.ReadString('\n')
	fmt.Println("Enter description :")
	description, _ := reader.ReadString('\n')
	task := Task{id: id, title: title, description: description, complete: false}
	tasks = append(tasks, task)
	id++
	fmt.Println("Task successfully added")
}

func updateTask(reader *bufio.Reader) {
	var found = false
	var index int = 0

	fmt.Println("Enter the task id that you want to update")
	input, _ := reader.ReadString('\n')

	taskId, _ := strconv.Atoi(strings.TrimSpace(input))
	for i, task := range tasks {
		if task.id == taskId {
			found = true
			index = i
			break
		}
	}

	if !found {
		fmt.Println("Task not found")

	} else {

		fmt.Println("Enter new title! if you don't want to just press Enter")
		title, _ := reader.ReadString('\n')

		fmt.Println("Enter new description! if you don't want to just press Enter")
		description, _ := reader.ReadString('\n')

		title = strings.TrimSpace(title)
		description = strings.TrimSpace(description)

		if title != "" && title != "\n" {
			tasks[index].title = title
		}
		if description != "" && description != "\n" {
			tasks[index].description = description
		}
		fmt.Println("Task updated successfully!")
	}

}

func completeTask(reader *bufio.Reader) {
	fmt.Println("Enter the task id that you want to complete")
	input, _ := reader.ReadString('\n')

	taskId, _ := strconv.Atoi(strings.TrimSpace(input))

	for i, task := range tasks {
		if task.id == taskId {
			tasks[i].complete = true

			fmt.Println("Task updated to complete successfully!")
		}
	}
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("You don't have any tasks Please add Task to list.")
		return
	}
	fmt.Println("Here are your tasks")

	for _, task := range tasks {

		if task.complete {
			continue
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("Task ID: %d Title: %s Description: %s \n", task.id, task.title, task.description)
		// fmt.Println("----------------------------------------")
	}
}
