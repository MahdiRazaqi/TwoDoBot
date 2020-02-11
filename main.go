package main

import (
	"github.com/MahdiRazaqi/Todo/jsondb"
)

func main() {
	data := jsondb.Task{
		ID: 27,
		User: jsondb.User{
			ID:        2,
			FirstName: "Mahdi",
			LastName:  "Razaqi",
			Username:  "MahdiRazaqi",
		},
		Text:   "Hello",
		Status: "todo",
		Date:   0,
	}

	jsondb.AddTask(data)
}
