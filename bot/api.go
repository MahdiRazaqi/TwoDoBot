package bot

import (
	"net/http"
)

func getUpdates(params Params) TelegramResponse {
	return request("getUpdates", http.MethodGet, params)
}

func sendMessage(userID int, text string) {
	request("sendMessage", http.MethodGet, Params{
		"chat_id": userID,
		"text":    text,
	})
}
