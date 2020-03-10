package bot

import (
	"strconv"

	"github.com/MahdiRazaqi/TwoDoBot/jsondb"
)

var lastText string

// keyboard markup model
type km struct {
	Keyboard [][]kb `json:"keyboard"`
}

// keyboard button model
type kb struct {
	Text string `json:"text"`
}

func controller(m Message) {
	switch m.Text {
	case "/start":
		sendMessage(m.From.ID, "Welcome to To-do list bot.", km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
	case "insert-new-task":
		sendMessage(m.From.ID, "Send description task.", km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
	case "show-task-list":
		sendMessage(m.From.ID, "Send description task.", km{Keyboard: [][]kb{{{Text: "todo"}, {Text: "doing"}, {Text: "done"}}}})
	case "todo":
		user := jsondb.User{
			ID:        m.From.ID,
			FirstName: m.From.FirstName,
			LastName:  m.From.LastName,
			Username:  m.From.Username,
		}
		tasks := jsondb.SelectTask(user, "todo")
		if len(tasks) == 0 {
			sendMessage(m.From.ID, "No task", km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
		for i, task := range tasks {
			sendMessage(m.From.ID, "Task "+strconv.Itoa(i+1)+": "+task.Text, km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
	case "doing":
		user := jsondb.User{
			ID:        m.From.ID,
			FirstName: m.From.FirstName,
			LastName:  m.From.LastName,
			Username:  m.From.Username,
		}
		tasks := jsondb.SelectTask(user, "doing")
		if len(tasks) == 0 {
			sendMessage(m.From.ID, "No task", km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
		for i, task := range tasks {
			sendMessage(m.From.ID, "Task "+strconv.Itoa(i+1)+": "+task.Text, km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
	case "done":
		user := jsondb.User{
			ID:        m.From.ID,
			FirstName: m.From.FirstName,
			LastName:  m.From.LastName,
			Username:  m.From.Username,
		}
		tasks := jsondb.SelectTask(user, "done")
		if len(tasks) == 0 {
			sendMessage(m.From.ID, "No task", km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
		for i, task := range tasks {
			sendMessage(m.From.ID, "Task "+strconv.Itoa(i+1)+": "+task.Text, km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
	default:
		if lastText == "insert-new-task" {
			task := jsondb.Task{
				ID: m.MessageID,
				User: jsondb.User{
					ID:        m.From.ID,
					FirstName: m.From.FirstName,
					LastName:  m.From.LastName,
					Username:  m.From.Username,
				},
				Text:   m.Text,
				Status: "todo",
				Date:   m.Date,
			}

			jsondb.AddTask(task)
			sendMessage(m.From.ID, "Task added to list.", km{Keyboard: [][]kb{{{Text: "insert-new-task"}, {Text: "show-task-list"}}}})
		}
	}

	lastText = m.Text
}
