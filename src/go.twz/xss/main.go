package main

import "github.com/gin-gonic/gin"

func main() {
	gEngin := gin.New()
	gEngin.Static("/index", "")
	gEngin.LoadHTMLGlob("")
}
