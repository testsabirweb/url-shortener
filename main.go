package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello ",
		})
	})

	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
