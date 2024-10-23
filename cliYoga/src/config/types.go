package config

var Config ConfigJSON

type ConfigJSON struct {
	YogaSoul string  `json:"yogaSoul"` //根据自己喜欢来设置一个字符串
	CliInfo  CliInfo `json:"cliInfo"`  //这是cli的一些信息
	MyInfo   MyInfo  `json:"myInfo"`   //这是cli的一些信息
}

// 以下是json中的嵌套结构体
type CliInfo struct {
	CliName string `json:"cliName"`
	Url     string `json:"url"`
}
type MyInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
