package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", index)
	router.GET("/ping", ping)
	router.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{ "title": "Main website",	})
}