package bot

const (
	start = "Welcome to To-do list bot."
)

func controller(m Message) {
	switch m.Text {
	case "/start":
		sendMessage(m.From.ID, start)
	case "/insert":
		sendMessage(m.From.ID, "Send description task")
	case "/delete":
		sendMessage(m.From.ID, start)
	}
}
