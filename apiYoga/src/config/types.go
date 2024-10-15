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
// 作为一个小作坊产品,我默认一个人即可维护和开发，所以为了逻辑简单只给了一个开发者接口
// 如果需要更改，可以将sudo改为sudoers，并想办法将sudo验证函数作修改，我设置了一个sudo验证器，你只需要改那里就好了
// 另外如果觉得我提供的sudo验证方案太简陋，完全可以自己重写验证器
type sudo struct {
	SuperUsername string `json:"superUsername"`
	SuperPassword string `json:"superPassword"`
}
