package amis

import (
	"strings"
	"github.com/gin-gonic/gin"
)
// 计算是否有权限，没有权限返回amis的classname的none，进行隐藏
func CK(path string,c *gin.Context)(string){
	superadmin,_:=c.Get("superadmin")
	if superadmin == 1 {
		return ""
	}
	role,_:=c.Get("role")
	if strings.Contains(role.(string), path) {
		return ""
	}
	return "none"
}
// 计算是否有权限，有权限返回字符串格式的true
func CKBS(path string,c *gin.Context)(string){
	if CK(path,c)=="" {
		return "true"
	}
	return "false"
}
