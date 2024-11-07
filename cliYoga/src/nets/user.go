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
	UserID    int
	Openid    string
	Nickname  string
	Name      string
	Gender    bool
	Signature string
	IsStudent bool
	IsTeacher bool
	IsAdmin   bool
	AvaURL    string
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
func SelectUserTail(tail int) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/select/user/tail"
	dataMap := make(map[string]interface{})
	dataMap["tail"] = tail
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
func InsertTeacherAccountAndPassword(teacherId int, account, password string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/insert/teacher"
	dataMap := make(map[string]interface{})
	dataMap["account"] = account
	dataMap["password"] = password
	dataMap["teacherId"] = teacherId
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
func InsertAdminAccountAndPassword(adminId int, account, password string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/insert/admin"
	dataMap := make(map[string]interface{})
	dataMap["account"] = account
	dataMap["password"] = password
	dataMap["adminId"] = adminId
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
func SelectTeacherInfoByName(name string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/select/teacehr/by/name"
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

func SelectAdminInfoByName(name string) {
	targetUrl := config.Config.CliInfo.Url + "/api/root/select/admin/by/name"
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
