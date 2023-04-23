package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var Logger *log.Logger
var MysqlClient *gorm.DB

func main() {
	gEngin := gin.New()
	// gEngin.Static("/index", "./")
	gEngin.LoadHTMLGlob("view/*")

	MysqlClient = MysqlConn()
	F := OpenLogFile()

	gEngin.GET("/index", Hello)
	gEngin.POST("/index", Hello2)
	gEngin.GET("/insert", Insert)
	gEngin.POST("/insert", Insert)
	gEngin.GET("/show", Show)

	Logger = log.New(F, "", log.LstdFlags)

	gEngin.Run(":8011")
}

func Hello(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func Hello2(ctx *gin.Context) {
	info := ctx.PostForm("name")
	ctx.Data(200, "text/html", []byte(info))
	// ctx.String(200, info)
}

func Insert(ctx *gin.Context) {
	if ctx.PostForm("user") != "" {
		fmt.Printf("获取到的信息是：%s", ctx.PostForm("user"))
		info := XssInfo{
			User: ctx.PostForm("user"),
			Name: ctx.PostForm("name"),
		}
		_ = mysqlInsert(MysqlClient, info)
	}
	ctx.HTML(200, "index2.html", nil)
}

func Show(ctx *gin.Context) {
	hh := gin.H{}
	info := mysqlFindDetail(MysqlClient)
	hh["user"] = info.User
	hh["name"] = info.Name
	// ctx.HTML(200, "show.html", hh)
	ctx.Data(200, "text/html", []byte(info.Name))
}

func MysqlConn() *gorm.DB {
	protocol := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		"root",
		"tongwz",
		"192.168.33.10",
		"xss",
		"utf8mb4",
	)

	gDb, _ := gorm.Open("mysql", protocol)
	return gDb
}

type XssInfo struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Name string `json:"name"`
}

func mysqlInsert(db *gorm.DB, insert XssInfo) error {
	err := db.Table("xss_info").Create(&insert).Error
	if err != nil {
		Logger.Printf("数据表插入数据错误：%s", err.Error())
		return err
	}
	return nil
}

func mysqlFindDetail(db *gorm.DB) XssInfo {
	info := XssInfo{}
	err := db.Table("xss_info").Where("id >= 1").First(&info).Error
	if err != nil {
		Logger.Printf("数据表查询数据错误：%s", err.Error())
	}

	return info
}

func OpenLogFile() *os.File {
	filePath := "logger.log"
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		fmt.Printf("日志文件不存在：%s \n", filePath)
		_ = os.MkdirAll("/logger.log", os.ModePerm)
		// filePath = getLogFileFullPath("")
	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}
