package jsondb

import (
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
