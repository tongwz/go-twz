package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	protocol := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1:3306",
		"transfer",
		"utf8mb4",
	)

	gDb, err := gorm.Open("mysql", protocol)
	if err != nil {
		_ = fmt.Errorf("mysql连接错误：%s", err.Error())
		return
	}

	// 加入sql写入日志
	gDb.LogMode(true)

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456", // 没有password也得有个空值
		DB:       0,
	})

	RedisClient.Ping()

	ch := make(chan bool)

	<-ch

}
