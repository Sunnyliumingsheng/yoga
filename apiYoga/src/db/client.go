package db

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/config"
)

// 这个文件主要用来连接数据库
var rdb *redis.Client
var postdb *gorm.DB

// 用来测试的结构体和函数
type tests struct {
	id uint `gorm:"primaryKey;column:test_id"`
}

func (tests) TableName() string {
	return "tests"
}
func addTest(a uint) {
	postdb.Create(&tests{
		id: a,
	})
}
func getTest() {
	var a []tests
	postdb.Select(&a)
	fmt.Println(a)
}

func StartClient() {
	fmt.Println("!!!!!!!!!!")
	redisInit()
	postgresInit()
	addTest(1)
	getTest()
	AddSession("123", 2)
	time.Sleep(time.Second)
	fmt.Println("start get authentication")
	isActive, level := AuthSession("123")
	fmt.Println("isActive :", isActive, "level :", level)
}
func redisInit() {
	log.Println(config.Config.RedisConfig.IpAddress + ":" + fmt.Sprint(config.Config.RedisConfig.Port))
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisConfig.IpAddress + ":" + fmt.Sprint(config.Config.RedisConfig.Port),
		Password: config.Config.RedisConfig.Password,
		DB:       config.Config.RedisConfig.Db,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Panic("redis client error : ", err.Error())
	}
}
func postgresInit() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Config.PostgresConfig.IpAddress, config.Config.PostgresConfig.Username, config.Config.PostgresConfig.Password, config.Config.PostgresConfig.Dbname, config.Config.PostgresConfig.Port)
	var err error
	postdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("!!!!!! error : client to postgres SQL :error :", err)
	}

}
