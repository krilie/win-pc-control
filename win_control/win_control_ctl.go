package win_control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// /media/status?action=stop POST
func MediaStatusPost(c *gin.Context) {
	action := c.Query("action")
	if action == "" {
		c.String(http.StatusBadRequest, "请求参数不能为空")
		return
	}
	if err := Media(action); err != nil {
		c.String(http.StatusBadRequest, "请求参数错误")
		return
	}
	c.String(http.StatusOK, "")
}

// /volume/value?value=? POST
func VolumeValuePost(c *gin.Context) {
	value := c.Query("value")
	if valueInt, err := strconv.Atoi(value); err != nil {
		c.String(http.StatusBadRequest, "请求参数错误,不能转为int")
		return
	} else {
		if valueInt >= 0 && valueInt <= 100 {
			if err := WinVolume(valueInt); err != nil {
				c.String(http.StatusInternalServerError, "调用volume命令失败："+err.Error())
				return
			} else {
				c.Status(http.StatusOK)
				return
			}
		} else {
			c.String(http.StatusBadRequest, "超出参数允许范围")
			return
		}
	}
}

// /sysctl/action? action=? post shutdown monitor_on monitor_off
func SysctlActionPost(c *gin.Context) {
	action := c.Query("action")
	if action == "" {
		c.String(http.StatusBadRequest, "请求参数不能为空")
		return
	}
	if err := SysCtl(action); err != nil {
		if err == errParam {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			c.String(http.StatusInternalServerError, "调用SysCtl命令失败："+err.Error())
		}
		return
	} else {
		c.Status(http.StatusOK)
	}
}
