package middleware

import (
	"fmt"
	"strings"
	"webapp/model/common/response"
	"webapp/config"
	"webapp/utils"
	"github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

// jwt权限验证 中间件
func UseJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			response.FailWithMessage("请登录后使用",c)
			return
		}
		token,err:=utils.JetDecode(tokenString,config.JWT_SK)
		if err != nil {
			fmt.Println(err)
			response.FailLogin(c)
			return
		}
		// 验证 Token
		if token.Valid {
			// fmt.Println("Token验证成功")
			claims := token.Claims.(jwt.MapClaims)
			c.Set("uid", int(claims["uid"].(float64))) // 定义全局
			c.Set("superadmin", int(claims["superadmin"].(float64))) // 定义全局
			c.Set("name", claims["name"].(string)) // 定义全局
			c.Set("role", claims["role"].(string)) // 定义全局
			if claims["superadmin"].(float64)==1 {//超级管理员
				c.Next()
				return
			}
			handlerName := c.HandlerName()
			handlerName = strings.Replace(handlerName, "webapp/controller.", "", -1)
			if strings.Contains(claims["role"].(string), handlerName+";") == false {//没有权限
				fmt.Println(claims["role"].(string),handlerName)
				response.FailWithMessage("无权限",c)
				return
			}
			c.Next()
			
		} else {
			response.FailLogin(c)
			return
		}
		
	}
}
