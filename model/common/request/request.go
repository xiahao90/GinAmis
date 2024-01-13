package request

import (
	// "fmt"
	"errors"
	"strings"
	"github.com/gin-gonic/gin"
)
// 验证map中参数是否存在
func CheckJson(data map[string]interface{}, params []string) (error) {
    var missingParams []string
    for _, param := range params {
        if _, exists := data[param]; !exists {
            // return errors.New(param+"必传")
            missingParams = append(missingParams, param)
        }
    }
    if len(missingParams) != 0 {
    	resultString := strings.Join(missingParams, ";")+"必传" 
        return errors.New(resultString)
    }
    return nil
    // return missingParams
}
//获取json格式的参数,并进行验证
func GetJsonData(keys[]string,c *gin.Context) (map[string]interface{},error) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		return jsonData,errors.New("json格式错误")
	}
	newJsonData:=map[string]interface{}{}

	// 查找不存在的参数
	err := CheckJson(jsonData, keys)
	if err != nil {
		// 过滤掉不要的
		for _, key := range keys {
		  	if value, ok := jsonData[key]; ok {
		  		newJsonData[key]=value
		  	}
		}
		return newJsonData,err
	}

	// 过滤掉不要的
	for _, key := range keys {
	  	if value, ok := jsonData[key]; ok {
	  		newJsonData[key]=value
	  	}
	}
	return newJsonData,nil
}
