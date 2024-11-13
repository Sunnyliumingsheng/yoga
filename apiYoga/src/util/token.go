package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"api/config"
	"api/loger"
)

func GenerateToken(id string) (Base64token string) {
	mapClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().AddDate(1, 0, 0).UnixMilli(),
	}
	// 默认为一年时效
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	loger.Loger.Println(config.Config.YogaSoul, "生成新的token// 这里可以以后删除,在token.go中")
	Base64token, _ = token.SignedString([]byte(config.Config.YogaSoul))
	return Base64token
} //一般来说这个函数无论如何也不会出错

func ParseToken(Base64token string) (id string, err error) {
	token, err := jwt.Parse(Base64token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.YogaSoul), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("token无效")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	//语法知识，x,ok:=y.(T)为类型断言，y为待转化的值，T为待转化类型，x为转化后的值
	//这比x:=(T)y好在能够确定是否转化成功
	if !ok {
		return "", errors.New("无法提取token中的claims")
	}
	id, ok = claims["id"].(string)
	if !ok {
		return "", errors.New("无法获取id")
	}

	// 检查token是否过期
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", errors.New("无法获取token过期时间")
	}
	expTime := time.UnixMilli(int64(exp))
	if expTime.Before(time.Now()) {
		return "", errors.New("token过期")
	}
	return id, nil
}
func AsyncGenerateToken(id string, base64TokenChan chan string) {
	mapClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().AddDate(1, 0, 0).UnixMilli(),
	}
	// 默认为一年时效
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	loger.Loger.Println(config.Config.YogaSoul)
	Base64token, _ := token.SignedString([]byte(config.Config.YogaSoul))
	base64TokenChan <- Base64token
} //一般来说这个函数无论如何也不会出错
func AsyncParseToken(Base64token string, isValidChan chan bool, userIdChan chan int) {
	token, err := jwt.Parse(Base64token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.YogaSoul), nil
	})
	if err != nil {
		isValidChan <- false
		return
	}
	if !token.Valid {
		isValidChan <- false
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	//语法知识，x,ok:=y.(T)为类型断言，y为待转化的值，T为待转化类型，x为转化后的值
	//这比x:=(T)y好在能够确定是否转化成功
	if !ok {
		isValidChan <- false
		return
	}
	id, ok := claims["id"].(int)
	if !ok {
		isValidChan <- false
		return
	}

	// 检查token是否过期
	exp, ok := claims["exp"].(float64)
	if !ok {
		isValidChan <- false
		return
	}
	expTime := time.UnixMilli(int64(exp))
	if expTime.Before(time.Now()) {
		isValidChan <- false
		return
	}
	userIdChan <- id
	isValidChan <- true
}
func GenerrateTokenWithLevelAndAccount(account string, level int) (Base64token string) {
	loger.Loger.Println(level, "GenerrateTokenWithLevelAndAccount")
	mapClaims := jwt.MapClaims{
		"account": account,
		"exp":     time.Now().AddDate(1, 0, 0).UnixMilli(),
		"level":   level,
	}
	// 默认为一年时效
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	loger.Loger.Println(config.Config.YogaSoul, "生成新的token// 这里可以以后删除,在token.go中")
	Base64token, _ = token.SignedString([]byte(config.Config.YogaSoul))
	return Base64token
}
func ParseTokenWithLevelAndAccount(Base64token string) (account string, level int, err error) {
	token, err := jwt.Parse(Base64token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.YogaSoul), nil
	})
	if err != nil {
		return "", -1, err
	}
	if !token.Valid {
		return "", -1, errors.New("token无效")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", -1, errors.New("无法提取token中的claims")
	}
	account, ok = claims["account"].(string)
	if !ok {
		return "", -1, errors.New("无法获取account")
	}

	// 处理 level 字段为 float64 转换为 int,这是jwt的规定，数字存储为float64
	levelfloat, ok := claims["level"].(float64)
	if !ok {
		return "", -1, errors.New("无法获取level")
	}
	level = int(levelfloat)
	// 将 level 转换为 int
	return account, level, nil
}
