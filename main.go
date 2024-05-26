package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
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

var (
	tasks = []string{}
)

func main() {
	// there has to be some options
	if len(os.Args) == 1 {
		color.Red("no option provided, displaying help")
		color.Blue("add [text]... - add task")
		color.Blue("remove [index] - remove task by index")
		color.Blue("clear [none] - clear tasks")
		color.Blue("list [none] - list tasks")
		return
	}

	// open, read and parse tasks.json file
	tasksRaw, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		os.Create("tasks.json")
	} else if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal(tasksRaw, &tasks)
		if err != nil {
			panic(err)
		}
	}

	// parsing commands
	command := strings.ToLower(os.Args[1])
	if command == "add" {
		// there has to be at least three argumets
		if len(os.Args) <= 2 {
			color.RedString("provide task name")
			return
		}
		// sum extra arguments, they are task name
		name := sumArgs(os.Args, 2)
		name = strings.ReplaceAll(name, "\n", "")
		tasks = append(tasks, name) // append new task
	} else if command == "list" {
		// for loop, print tasks
		// ([index]) [name]
		for i, task := range tasks {
			fmt.Printf(color.GreenString("(%d)")+color.YellowString(" %v\n"), i, task)
		}
	} else if command == "remove" {
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
	} else if command == "clear" {
		// clear tasks
		tasks = []string{}
	}

	// parse tasks back to json
	tasksRaw, err = json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	// write parsed json to file
	err = os.WriteFile("tasks.json", tasksRaw, 0664)
	if err != nil {
		panic(err)
	}
}
