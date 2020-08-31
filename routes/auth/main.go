package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/auth", func(c *gin.Context) {
		c.String(200, "auth")
	})

	router.Run()
}
