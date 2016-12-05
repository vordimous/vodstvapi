package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var c = make(map[string]string)

//Init starts all of the routes
func Init() {
	//go get a config file and parse values
	file, _ := os.Open("/etc/esvods/config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("config:", c)
	}
}

//Get attr from stored config
func Get(attr string) string {
	return c[attr]
}
