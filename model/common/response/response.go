package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int       `json:"status"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	LOGINOUT	= 211
	ERROR   	= 1
	SUCCESS 	= 0
)

func Result(status int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		status,
		data,
		msg,
	})
	c.Abort()
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data map[string]interface{}, c *gin.Context) {
	// 先检查是否包含 "rows" 键
	if _, ok := data["rows"]; ok {
		// 如果包含，再检查对应的切片是否为空
		if len(data["rows"].([]map[string]interface{})) == 0 {
			// 如果为空，给 rows 赋值一个空数组
			data["rows"] = make([]map[string]interface{}, 0)
		}
	}
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)

}

func FailLogin(c *gin.Context) {
	Result(LOGINOUT, map[string]interface{}{}, "登录过期，请重新登录", c)

}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
