package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var Logger *log.Logger

func main() {
	Logger = log.New(OpenLogFile(), "", log.LstdFlags)

	for {
		time.Sleep(time.Second * 1)

		Logger.Printf("tongWz 测试日志写入 %s", time.Now().Format("2006-01-02 15:04:05"))

		go writeLog()
	}
}

func writeLog() {
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	Logger.Printf("某个 情况 %s", time.Now().Format("2006-01-02 15:04:05"))
	Logger.Printf("我是来捣乱的，我是来捣乱我是来捣乱的，我是来捣乱我是来捣乱的，我是来捣乱， wo shi lai dao luande %s", time.Now().Format("2006-01-02 15:04:05"))
}

func OpenLogFile() *os.File {
	filePath := "logger.log"
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("日志文件")
	}
	switch {
	case os.IsNotExist(err):
		fmt.Printf("日志文件不存在：%s \n", filePath)
		_ = os.MkdirAll("/logger.log", os.ModePerm)
	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}

func ForNew1() {
	fmt.Printf("forNew 1")
}

func ForNew2() {
	fmt.Printf("forNew 2")
}

func Slave1() {
	fmt.Printf("slave 1")
}

func Slave2() {
	fmt.Printf("slave 2")
}

func ForNew3() {
	fmt.Printf("forNew 3")
}

func ForNew4() {
	fmt.Printf("forNew 4")
}
