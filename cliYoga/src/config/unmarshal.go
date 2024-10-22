package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
)

func UnmarshalConfig() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.ReadFile(u.HomeDir + "/.yogaconfig.json")
	if err != nil {
		log.Fatal("!!! Could not read config file", "error:", err)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("!!! Could not marshal config", "error:", err)
	}

}
func ReUnmarshalConfig() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.ReadFile(u.HomeDir + "/.yogaconfig.json")
	if err != nil {
		log.Fatal("!!! Could not read config file", "error:", err)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("!!! Could not marshal config", "error:", err)
	}
	fmt.Println("配置信息是:")
	fmt.Println(string(file))
}
