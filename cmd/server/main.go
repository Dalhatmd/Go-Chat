package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", func (c *gin.Context) {
		c.HTML(200, "Index.html", gin.H{
			"title": "Homepage",
		})
	})
	router.Run(":8080")
}
