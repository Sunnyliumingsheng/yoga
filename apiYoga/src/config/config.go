package config

import (
	"encoding/json"
	"fmt"
	"os"

	"api/loger"
)

func UnmarshalConfig() {
	file, err := os.ReadFile("../config/config.json")
	if err != nil {
		loger.Loger.Fatal("!!! Could not read config file", "error:", err)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		loger.Loger.Fatal("!!! Could not marshal config", "error:", err)
	}
	fmt.Println("config:")
	fmt.Println(Config)
}
