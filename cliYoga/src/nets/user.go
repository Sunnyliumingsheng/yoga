package nets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cli/config"
)

// 这里都是处理用户相关的

// 试着登录,检查这个目前的账号和密码是否正确
func Login() {
	targetUrl := config.Config.CliInfo.Url + "/api/root/login"
	dataMap := make(map[string]interface{})
	dataMap["sudoAuthentication"] = AuthenticationInfo
	payload, _ := json.Marshal(dataMap)
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(payload))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resqData, _ := io.ReadAll(response.Body)
	fmt.Println(string(resqData))
	if response.StatusCode != 200 {
		fmt.Println("验证成功,你可以自由的使用,并且除非服务器端出现问题,你将不再需要使用register和login")
	}
}

type User struct {
	UserID    int    `gorm:"primaryKey"`
	Openid    string `gorm:"not null;size:255"`
	Nickname  string `gorm:"not null;size:255"`
	Name      string `gorm:"size:255"`
	Gender    bool
	Signature string `gorm:"size:64"`
	IsStudent bool
	IsTeacher bool
	IsAdmin   bool
	AvaURL    string `gorm:"size:255"`
}

func SelectUserInfoByName(name string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/select/user/by/name"
	dataMap := make(map[string]interface{})
	dataMap["name"] = name
	dataMap["sudoAuthentication"] = AuthenticationInfo

	payload, _ := json.Marshal(dataMap)
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(payload))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resqData, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		fmt.Println("Error:", string(resqData))
	} else {
		fmt.Println(string(resqData))
	}

}
func InsertUserByName(name string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/insert/user/by/name"
	dataMap := make(map[string]interface{})
	dataMap["name"] = name
	dataMap["sudoAuthentication"] = AuthenticationInfo

	payload, _ := json.Marshal(dataMap)
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(payload))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resqData, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		fmt.Println("Error:", string(resqData))
	} else {
		fmt.Println(string(resqData))
	}
}
func DropUserByUserId(userId string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/drop/user/by/userId"
	dataMap := make(map[string]interface{})
	dataMap["userId"] = userId
	dataMap["sudoAuthentication"] = AuthenticationInfo

	payload, _ := json.Marshal(dataMap)
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(payload))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resqData, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		fmt.Println("Error:", string(resqData))
	} else {
		fmt.Println(string(resqData))
	}
}
func UpdateUserLevel(name string, wantStudent, wantTeacher, wantAdmin bool) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/update/user/level/by/name"
	dataMap := make(map[string]interface{})
	dataMap["name"] = name
	dataMap["sudoAuthentication"] = AuthenticationInfo
	dataMap["isStudent"] = wantStudent
	dataMap["isTeacher"] = wantTeacher
	dataMap["isAdmin"] = wantAdmin
	payload, _ := json.Marshal(dataMap)
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(payload))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resqData, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		fmt.Println("Error:", string(resqData))
	} else {
		fmt.Println(string(resqData))
	}
}
