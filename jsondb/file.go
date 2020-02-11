package jsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadDB read file
func ReadDB(path string) TodoList {
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	data := TodoList{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		panic(err)
	}
	return data
}

// WriteDB for write into file
func WriteDB(path string, todoList TodoList) {
	file, err := json.Marshal(todoList)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(path, file, 0644); err != nil {
		panic(err)
	}
}

// HasExist check exist file or folder
func HasExist(path string) bool {
	if _, err := os.Stat(DBPath); os.IsNotExist(err) {
		err := os.Mkdir(DBPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
