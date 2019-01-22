package main

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//静态网页的处理
	router.Static("/home", "./home")
	router.Static("/static", "./static")

	router.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "api hello")
	})
	if err := manners.ListenAndServe(":80", router); err != nil {
		fmt.Println(err)
		return
	}
}
