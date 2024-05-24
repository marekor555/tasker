package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

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
	if len(os.Args) == 1 {
		fmt.Println("no option provided...")
		return
	}

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

	command := strings.ToLower(os.Args[1])
	if command == "add" {
		if len(os.Args) < 2 {
			fmt.Println("provide task name")
			return
		}
		name := sumArgs(os.Args, 2)
		name = strings.ReplaceAll(name, "\n", "")
		tasks = append(tasks, name)
	} else if command == "list" {
		for i, task := range tasks {
			fmt.Printf("(%d) %v\n", i, task)
		}
	} else if command == "remove" {
		if len(os.Args) < 2 {
			fmt.Println("provide task index")
			return
		}
		indexStr := os.Args[2]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			panic(err)
		}
		tasks = remove(tasks, index)
	} else if command == "clear" {
		tasks = []string{}
	}

	tasksRaw, err = json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", tasksRaw, 0664)
	if err != nil {
		panic(err)
	}
}
