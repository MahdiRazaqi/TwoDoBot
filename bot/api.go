package bot

import (
	"net/http"
)

func getUpdates(params Params) TelegramResponse {
	return request("getUpdates", http.MethodGet, params)
}

// SendMessage func
func SendMessage(userID int, text string) {
	request("sendMessage", http.MethodGet, Params{
		"chat_id": userID,
		"text":    text,
	})
}
