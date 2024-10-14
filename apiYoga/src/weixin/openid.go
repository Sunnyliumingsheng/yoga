package weixin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"api/config"
	"api/loger"
)

//这个文件夹是用来获取微信的openid的

type Code2SessionResponse struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid,omitempty"` // 只有在用户将该公众号绑定到微信开放平台帐号后，才会出现该字段
	ErrCode    int    `json:"errcode,omitempty"`
	ErrMsg     string `json:"errmsg,omitempty"`
}

func GetOpenId(jsCode string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.Config.Weixin.AppId, config.Config.Weixin.AppSecret, jsCode)
	loger.Loger.Println(url)

	// 发送 HTTP POST 请求
	resp, err := http.Get(url)
	if err != nil {
		// 处理错误
		loger.Loger.Printf("请求失败: %v", err)
		return "", err
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		// 处理非200状态码
		loger.Loger.Printf("服务器返回错误状态码: %d", resp.StatusCode)
		return "", fmt.Errorf("HTTP 状态码 %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		loger.Loger.Printf("读取响应体失败: %v", err)
		return "", err
	}
	var response Code2SessionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		loger.Loger.Printf("解码 JSON 响应体失败: %v", err)
		return "", err
	}
	if response.ErrCode != 0 {
		// 处理微信返回的错误
		loger.Loger.Printf("微信返回错误: %d - %s", response.ErrCode, response.ErrMsg)
		return "", fmt.Errorf("微信返回错误: %d - %s", response.ErrCode, response.ErrMsg)
	}
	return response.OpenID, nil

}
