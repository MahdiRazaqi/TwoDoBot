package jsondb

import (
	"fmt"
	"strconv"
)

// AddTask to To-do
func AddTask(task Task) {
	id := strconv.Itoa(task.User.ID)
	path := DBPath + id + ".json"

	data := TodoList{}
	if HasExist(path) {
		data = ReadDB(path)
	}
	data.Todo = append(data.Todo, task)

	WriteDB(path, data)
}

// DeleteTask from jsonDB
func DeleteTask(task Task) {
	id := strconv.Itoa(task.User.ID)
	path := DBPath + id + ".json"

	data := TodoList{}
	if HasExist(path) {
		data = ReadDB(path)
	}

	var cardList *[]Task
	switch task.Status {
	case "todo":
		cardList = &data.Todo
	case "doing":
		cardList = &data.Doing
	case "done":
		cardList = &data.Done
	}

	for i, t := range *cardList {
		if t.ID == task.ID {
			(*cardList)[i] = (*cardList)[len(*cardList)-1]
			*cardList = (*cardList)[:len(*cardList)-1]
		}
	}

	WriteDB(path, data)
}

// SelectTask from DB
func SelectTask(user User, list string) []Task {
	id := strconv.Itoa(user.ID)
	path := DBPath + id + ".json"

	data := TodoList{}
	if HasExist(path) {
		data = ReadDB(path)
	}

	var cardList *[]Task
	switch list {
	case "todo":
		cardList = &data.Todo
	case "doing":
		cardList = &data.Doing
	case "done":
		cardList = &data.Done
	}

	fmt.Println(cardList)
	return *cardList
}
