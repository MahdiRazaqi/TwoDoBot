package bot

const (
	start = "Welcome to To-do list bot."
)

func controller(m Message) {
	switch m.Text {
	case "/start":
		SendMessage(m.From.ID, start)
	case "/insertTask":
		SendMessage(m.From.ID, start)
	}

}
