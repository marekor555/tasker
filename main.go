package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
	tasks = []string{}
)

// generic remove from slice function
func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

// suming slice of strings (for left arguments)
func sumArgs(slice []string, start int) string {
	var output string
	slice = slice[start:]
	for _, element := range slice {
		output += element + " "
	}
	return output
}

func loadTasks() error {
	// open, read and parse tasks.json file
	tasksRaw, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		return err
	} else if err != nil {
		panic(err)
	}
	err = json.Unmarshal(tasksRaw, &tasks)
	if err != nil {
		panic(err)
	}
	return nil
}

func saveTasks(tasks []string) {
	// parse tasks back to json
	tasksRaw, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	// write parsed json to file
	err = os.WriteFile("tasks.json", tasksRaw, 0664)
	if err != nil {
		panic(err)
	}
}

func help() {
	color.Blue("init [none] - initialize tasker in this directory")
	color.Blue("add [text]... - add task")
	color.Blue("remove [index] - remove task by index")
	color.Blue("clear [none] - clear tasks")
	color.Blue("list [none] - list tasks")
}

func main() {
	// there has to be some options
	if len(os.Args) == 1 {
		color.Red("no arguments specified")
		help()
		return
	}

	// parsing commands
	command := strings.ToLower(os.Args[1])
	switch command {
	case "init":
		saveTasks([]string{})
	case "help":
		help()
	case "add":
		loadTasks()
		// there has to be at least three argumets
		if len(os.Args) <= 2 {
			color.RedString("provide task name")
			return
		}
		// sum extra arguments, they are task name
		name := sumArgs(os.Args, 2)
		name = strings.ReplaceAll(name, "\n", "")
		tasks = append(tasks, name) // append new task
		saveTasks(tasks)
	case "list":
		loadTasks()
		// for loop, print tasks
		// ([index]) [name]
		for i, task := range tasks {
			fmt.Printf(color.GreenString("(%d)")+color.YellowString(" %v\n"), i, task)
		}
	case "remove":
		loadTasks()
		// there has to be at least three arguments
		if len(os.Args) <= 2 {
			color.Red("provide task index")
			return
		}
		// get index name
		indexStr := os.Args[2]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			panic(err)
		}
		tasks = remove(tasks, index)
		saveTasks(tasks)
	case "clear":
		// clear tasks
		saveTasks([]string{})
	default:
		color.Red("command not found")
		help()
	}
}
