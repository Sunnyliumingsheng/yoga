package config

var Config ConfigJSON

type ConfigJSON struct {
	YogaSoul           string             `json:"yogaSoul"`           //根据自己喜欢来设置一个字符串
	RedisConfig        RedisConfig        `json:"redisConfig"`        //redis的配置位置
	PostgresConfig     PostgresConfig     `json:"postgresConfig"`     //postgresql的配置位置
	Weixin             Weixin             `json:"weixin"`             //有关微信的一些信息
	Authentication     Authentication     `json:"authentication"`     //控制验证方式有效时间
	NewUserDefaultInfo NewUserDefaultInfo `json:"newUserDefaultInfo"` //新注册用户的默认信息
	Sudo               sudo               `json:"sudo"`               //超级管理员的一些信息
}

//以下是json中的嵌套结构体

type RedisConfig struct {
	IpAddress string `json:"ipAddress"`
	Port      string `json:"port"`
	Password  string `json:"password"`
	Db        int    `json:"db"`
}

type PostgresConfig struct {
	IpAddress string `json:"ipAddress"`
	Port      string `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Dbname    string `json:"dbname"`
}
type Weixin struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}
type Authentication struct {
	TokenDurationDay    int `json:"tokenDurationDay"`
	SessionDurationHour int `json:"sessionDurationHour"`
}
type NewUserDefaultInfo struct {
	Nickname  string `json:"nickname"`
	Gender    bool   `json:"gender"`
	Signature string `json:"signature"`
	AvaURL    string `json:"avaUrl"`
}
type sudo struct {
	SuperUsername string `json:"superUsername"`
	SuperPassword string `json:"superPassword"`
}
