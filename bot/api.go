package bot

import (
	"encoding/json"
	"net/http"
)

func getUpdates(params Params) TelegramResponse {
	return request("getUpdates", http.MethodGet, params)
}

func sendMessage(userID int, text string, replyMarkup km) {
	keyboard, _ := json.Marshal(replyMarkup)

	request("sendMessage", http.MethodGet, Params{
		"chat_id":      userID,
		"text":         text,
		"reply_markup": string(keyboard),
	})
}
