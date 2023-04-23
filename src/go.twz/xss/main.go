package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gEngin := gin.New()
	// gEngin.Static("/index", "./")
	gEngin.LoadHTMLGlob("view/*")

	gEngin.GET("/index", Hello)
	gEngin.POST("/index", Hello2)

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
