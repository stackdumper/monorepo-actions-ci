package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		c.String(200, "users")
	})

	router.Run()
}
