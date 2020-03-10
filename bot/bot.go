package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var endpoint string

// Start bot server
func Start() {
	config()

	var offset int
	for range time.Tick(1000 * time.Millisecond) {
		resp := getUpdates(Params{
			"offset": offset + 1,
		})

		for _, m := range resp.Result {
			offset = m.UpdateID
			controller(m.Message)
		}
	}
}

func config() {
	config, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	data := map[string]string{}
	if err := json.Unmarshal(config, &data); err != nil {
		panic(err)
	}

	endpoint = data["endpoint"]
}

func request(query string, method string, params Params) TelegramResponse {
	var url string
	url = endpoint + query

	if method == http.MethodGet {
		url += "?"
		for key, value := range params {
			url += fmt.Sprintf("%s=%v&", key, value)
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := []byte(body)
	var response TelegramResponse
	if err := json.Unmarshal(jsonData, &response); err != nil {
		log.Println(err)
	}
	return response
}
