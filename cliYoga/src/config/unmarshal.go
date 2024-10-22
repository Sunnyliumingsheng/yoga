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
func MarshalConfig() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	modifiedJsonData, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// 将修改后的 JSON 数据写回文件
	err = os.WriteFile(u.HomeDir+"/.yogaconfig.json", modifiedJsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

}
