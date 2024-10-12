package db

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"

	"api/config"
	"api/loger"
)

var sessionDurartionHour time.Duration

func init() {
	sessionDurartionHour = time.Duration(config.Config.Authentication.SessionDurationHour) * time.Hour
}

func AddSession(userId string, level int) {
	err := rdb.Set(userId, level, sessionDurartionHour).Err()
	if err != nil {
		loger.Loger.Println("error!!!: redis : set", userId, level, ":panic :", err)
	}
}
func AuthSession(userId string) (isActive bool, level int) {
	result, err := rdb.Get(userId).Result()
	loger.Loger.Println(result)
	if err == redis.Nil {
		loger.Loger.Println("AuthSession : there is no userId", userId, "record")
		return false, 0
	}
	if err != nil {
		loger.Loger.Println("AuthSession : error :", err)
		return false, 0
	}
	level, err = strconv.Atoi(result)
	if err != nil {
		loger.Loger.Println("AuthSession : error: atoi error:userId:", userId, "resule of level:", result)
	}

	return true, level
}
