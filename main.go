package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	if err := router.Run(":80"); err != nil {
		fmt.Println(err)
		return
	}
}
