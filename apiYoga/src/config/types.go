package config

var Config ConfigJSON

type ConfigJSON struct {
	YogaSoul       string         `json:"yogaSoul"`       //根据自己喜欢来设置一个字符串
	RedisConfig    RedisConfig    `json:"redisConfig"`    //redis的配置位置
	PostgresConfig PostgresConfig `json:"postgresConfig"` //postgresql的配置位置
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
