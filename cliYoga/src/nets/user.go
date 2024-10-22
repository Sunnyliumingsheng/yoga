package nets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 这里都是处理用户相关的

// 试着登录,检查这个目前的账号和密码是否正确
func Login() {
	targetUrl := "http://localhost:8080/api/root/login"
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
