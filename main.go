package main

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//静态网页的处理
	router.Static("/home", "./home")
	//重定向
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/home/index.html")
	})

	router.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "api hello")
	})
	if err := manners.ListenAndServe(":80", router); err != nil {
		fmt.Println(err)
		return
	}
}
