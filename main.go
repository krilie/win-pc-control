// win-pc-control
// 电脑媒体控制程序
// 2019年1月25日20:48:52
package main

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lico603/win-pc-control/win_control"
	"net/http"
)

func main() {
	//gin的release模式
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	//开启gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	//前端静态页面 使用前后端分离
	router.Static("/home", "./home")
	router.GET("/", func(c *gin.Context) { //重定向到主页
		c.Redirect(http.StatusPermanentRedirect, "/home/index.html")
	})

	//登录鉴权
	authGroup := router.Group("/auth")
	authGroup.POST("/login", win_control.AuthLogin)
	authGroup.GET("/verity", win_control.AuthVerity)

	//后台api
	apiGroup := router.Group("/api")
	apiGroup.Use(win_control.CheckSecretKey())
	apiGroup.POST("/media/status", win_control.MediaStatusPost)
	apiGroup.POST("/volume/value", win_control.VolumeValuePost)
	apiGroup.POST("/sysctl/action", win_control.SysctlActionPost)
	// http 非https manners可以下确处理信号
	if err := manners.ListenAndServe(":80", router); err != nil {
		fmt.Println(err)
		return
	}
}
