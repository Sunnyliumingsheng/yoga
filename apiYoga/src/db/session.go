package db

import (
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func AddSession(userId string, level int) {

	err := rdb.Set(userId, level, 30*time.Second).Err()
	if err != nil {
		log.Println("error!!!: redis : set", userId, level, ":panic :", err)
	}

}
func AuthSession(userId string) (isActive bool, level int) {
	result, err := rdb.Get(userId).Result()
	log.Println(result)
	if err == redis.Nil {
		log.Println("AuthSession : there is no userId", userId, "record")
		return false, 0
	}
	if err != nil {
		log.Println("AuthSession : error :", err)
		return false, 0
	}
	level, err = strconv.Atoi(result)
	if err != nil {
		log.Println("AuthSession : error: atoi error:userId:", userId, "resule of level:", result)
	}

	return true, level

}
