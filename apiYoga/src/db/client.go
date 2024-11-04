package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"api/config"
	"api/loger"
)

// 这个文件主要用来连接数据库
var rdb *redis.Client
var postdb *gorm.DB

// 用来测试的结构体和函数
type tests struct {
	id int `gorm:"primaryKey;column:test_id"`
}

func (tests) TableName() string {
	return "tests"
}
func addTest(a int) {
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
	fmt.Println("start get authentication")
	isActive, level := AuthSession("123")
	fmt.Println("isActive :", isActive, "level :", level)
}
func redisInit() {
	loger.Loger.Println(config.Config.RedisConfig.IpAddress + ":" + fmt.Sprint(config.Config.RedisConfig.Port))
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisConfig.IpAddress + ":" + fmt.Sprint(config.Config.RedisConfig.Port),
		Password: config.Config.RedisConfig.Password,
		DB:       config.Config.RedisConfig.Db,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		loger.Loger.Panic("redis client error : ", err.Error())
	}
}
func postgresInit() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Config.PostgresConfig.IpAddress, config.Config.PostgresConfig.Username, config.Config.PostgresConfig.Password, config.Config.PostgresConfig.Dbname, config.Config.PostgresConfig.Port)
	var err error
	newLogger := logger.New(
		loger.Loger, // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	// 将日志信息打印到自定义的log文件中
	postdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		loger.Loger.Panic("!!!!!! error : client to postgres SQL :error :", err)
	}

}
