package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/articles", func(c *gin.Context) {
		c.String(200, "articles")
	})

	router.Run()
}
