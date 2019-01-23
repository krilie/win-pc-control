package win_control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SecretKey = "12345678"
var CookieKey = "secret_key"

// 检验中间件
func CheckSecretKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从数据库里读出来，或定死
		if skey, err := c.Cookie(CookieKey); err != nil {
			c.String(http.StatusUnauthorized, "请登录。。。")
			c.Abort()
			return
		} else {
			if skey == SecretKey {
				//key 验证成功。。。
				c.Next()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}

//登录
// auth/login?secret_key=?  Post
func AuthLogin(c *gin.Context) {
	key := c.Query("secret_key")
	if key == "" {
		c.String(http.StatusBadRequest, "请输入secret key")
		return
	}
	if key != SecretKey {
		c.String(http.StatusNotFound, "请输入正确的key")
		return
	}
	//写cookie,并返回200 一个月有效
	c.SetCookie(CookieKey, SecretKey, 60*60*24*30, "/", "localhost", false, true)
	c.Status(http.StatusOK)
	return
}

//检验登录
// auth/verity Get 200 401
func AuthVerity(c *gin.Context) {
	val, err := c.Cookie(CookieKey)
	if err != nil && err == http.ErrNoCookie {
		c.Status(http.StatusUnauthorized)
		return
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if val != SecretKey {
		c.String(http.StatusUnauthorized, "secret错误")
		return
	}
	c.Status(http.StatusOK)
	return
}
