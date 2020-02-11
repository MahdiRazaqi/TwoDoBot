package jsondb

// DBPath have database location
const DBPath = "./database/"

// TodoList model
type TodoList struct {
	Todo  []Task `json:"todo"`
	Doing []Task `json:"doing"`
	Done  []Task `json:"done"`
}

// Task model
type Task struct {
	ID     int    `json:"id"`
	User   User   `json:"user"`
	Text   string `json:"text"`
	Status string `json:"status"`
	Date   int    `json:"date"`
}

// User model
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
