package database

import (
	"encoding/json"
	"webapp/config"
	"webapp/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"time"
	// "fmt"
	// "reflect"
	"strconv"
	"strings"
	"webapp/model/common/request"
)
// DB 数据库
var DB *gorm.DB
// InitMysql 初始化 mysql 数据库
func InitMysql() {
	dsn := config.DB_USER+":"+config.DB_PWD+"@tcp("+config.DB_ADD+":"+config.DB_PORT+")/"+config.DB_NAME+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	if config.DEBUG==1 {
		db = db.Debug()
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to configure database connection pool")
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将数据库连接赋值给全局变量
	DB = db
}
// 配合amis的crud组件封装的查询
func AmisCrudSelect(tableName string,where map[string]interface{},c *gin.Context)(int64,[]map[string]interface{}){
	dbTemp:=DB.Table(tableName)
	perPage,_:= strconv.Atoi(c.DefaultQuery("perPage", "20"))//默认每页的条数
	page,_:= strconv.Atoi(c.DefaultQuery("page", "1"))//默认页数
	order:="id DESC"
	if perPage > 200 {
		perPage = 200
	}
	orderBy := c.DefaultQuery("orderBy", "")//排序字段
	orderDir := c.DefaultQuery("orderDir", "")//排序方式
	if orderBy != "" && orderDir!="" {
		order=orderBy+" "+orderDir
	}

	/*
	whereAnd
	有哪些字段需要等于查询，如status=1 and role=2，需要传入whereAnd=["status","role"]
	程序获取get参数是否传入进行组装查询
	*/
	if whereAnd, ok := where["whereAnd"].([]string); ok {
    	for _, column := range whereAnd {
    		value:=c.DefaultQuery(column, "")
    		if value != ""{
    			dbTemp.Where(map[string]string{
    				column:value,
    			})
    		}
    	}
    }
	
    /*
	whereBetween
	有哪些字段需要范围查询，如 addtime BETWEEN 1704873256 AND 1704873256 AND edittime BETWEEN 1704873256 AND 1704873256，需要传入 whereBetween=["addtime","edittime"]
	程序获取get参数是否传入进行组装查询
	*/
    if whereBetween, ok := where["whereBetween"].([]string); ok {
    	for _, column := range whereBetween {
    		value:=c.DefaultQuery(column, "")
    		if value != "" {
				result := strings.Split(value, ",")
				dbTemp.Where(column+" BETWEEN ? AND ?", result[0], result[1])
			}
    	}
    }
	/*
	whereJsonAnd
	查询json格式字段中包含的值，多个值查询&&的方式，只能查询数组格式的json
	如 role为json格式["aa","bb","ccc"] 查询有aa和bb的值，类似where应该为 role like "%aa%" and role like "%bb%"
	程序获取get参数是否传入进行组装查询
	*/
    if whereJsonAnd, ok := where["whereJsonAnd"].([]string); ok {
    	for _, column := range whereJsonAnd {
    		value:=c.DefaultQuery(column, "")
    		if value != ""{
				searchValues := strings.Split(value, ",")
				result1, intSlice1 := areAllStringsNumbers(searchValues)
				if result1 {//是int类型的json格式,需要转成int的切片进行查询
					dbTemp.Where("JSON_CONTAINS("+column+", JSON_ARRAY(?), '$')", intSlice1)
				}else{
					dbTemp.Where("JSON_CONTAINS("+column+", JSON_ARRAY(?), '$')", searchValues)
				}
			}
    	}
    }
    /*
	whereJsonOr
	查询json格式字段中包含的值，多个值查询||的方式，只能查询数组格式的json
	如 role为json格式["aa","bb","ccc"] 查询有aa和bb的值，类似where应该为 role like "%aa%" or role like "%bb%"
	程序获取get参数是否传入进行组装查询
	*/
    if whereJsonOr, ok := where["whereJsonOr"].([]string); ok {
    	for _, column := range whereJsonOr {
    		value:=c.DefaultQuery(column, "")
    		if value != ""{
				searchValues := strings.Split(value, ",")
				tempDbNow:=DB.Table(tableName)
				result1, intSlice1 := areAllStringsNumbers(searchValues)
				if result1 {//是int类型的json格式,需要转成int的切片进行查询
					for _, value1 := range intSlice1 {
						tempDbNow.Or("JSON_CONTAINS("+column+", CAST(? AS JSON), '$')", value1)
					}
				}else{
					for _, value1 := range searchValues {
						tempDbNow.Or("JSON_CONTAINS("+column+", JSON_QUOTE(?), '$')", value1)
					}
				}
				
				dbTemp.Where(tempDbNow)
			}
    	}
    }
    /*
	whereAndLike
	有哪些字段需要模糊查询，如 name like "%超级%" and admin like "%admin%"，需要传入whereAndLike=["name","admin"]
	程序获取get参数是否传入进行组装查询
	*/
    if whereAndLike, ok := where["whereAndLike"].([]string); ok {
    	for _, column := range whereAndLike {
    		value:=c.DefaultQuery(column, "")
    		if value != ""{
				dbTemp.Where(column+" LIKE ?", "%"+value+"%")
    			
    		}
    	}
    }
    /*
	keyword
	全局模糊查询，如搜索关键字“大哥”： name like "%大哥%" OR admin like "%大哥%" OR info like "%大哥%"，需要传入 keyword=["name","admin","info"]
	程序获取get参数是否传入进行组装查询
	*/
	keyword := c.DefaultQuery("keyword", "")
	if keyword != "" {
	    if strSlice, ok := where["keyword"].([]string); ok {
			result := strings.Join(strSlice, " LIKE @keyword OR ")+" LIKE @keyword"
			dbTemp.Where(result, map[string]interface{}{"keyword": "%"+keyword+"%"})
	    }
	}

    
    var rows []map[string]interface{}
    var total int64
    dbTemp.Count(&total)
    /*
	select
	查询的字段，如 select id,name,info.....，不传入查询*
	*/
    if select_, ok := where["select"].(string); ok {
    	if select_ != "" {
    		dbTemp.Select(select_)
    	}
    }
    dbTemp.Count(&total)
    dbTemp.Limit(perPage).Offset((page-1)*perPage).Order(order).Find(&rows)
    /*
    jsonColumn
    哪些字段是json格式，进行json字符串转interface
    */
    if jsonColumn, ok := where["jsonColumn"].([]string); ok {
    	for _, column := range jsonColumn {
    		for key, value := range rows {
    			var jsons []interface{}
    			// 解码 JSON 数据
    			err := json.Unmarshal([]byte(value[column].(string)), &jsons)
    			if err != nil {
    				rows[key][column]=make([]interface{},0) 
    				break
    			}
    			rows[key][column]=jsons
    		}
    	}
    }
    return total,rows
}
/*
配合amis的crud组件封装的编辑
tableName 表名
data.strColumns 非json字段
data.jsonColumns json字段
data.defaultColumns 默认字段，如updatetime
*/
// func AmisCrudUpdate(tableName string,strColumns []string,jsonColumns []string,c *gin.Context)(bool){
func AmisCrudUpdate(tableName string,data map[string]interface{},c *gin.Context)(bool){
	id:=c.Param("id")
	jsonData:=map[string]interface{}{}
    if jsonColumns, ok := data["jsonColumns"].([]string); ok {//有json格式的字段
		jsonData,_=request.GetJsonData(append(data["strColumns"].([]string), data["jsonColumns"].([]string)...),c)//获取json参数
		for _, column := range jsonColumns {
			temp:=make([]interface{}, 0)
			temp1 := utils.GetMapDefule(jsonData,column,temp)
			temp2 := temp1.([]interface{})
			temp3,_ := json.Marshal(temp2)
			jsonData[column]=string(temp3)
		}
    }else{
		jsonData,_=request.GetJsonData(data["strColumns"].([]string),c)//获取json参数
    }
    if defaultColumns, ok := data["defaultColumns"].(map[string]string); ok {
		for k,v:= range defaultColumns {
			jsonData[k]=v
		}
	}
	if len(jsonData) != 0 {
		result :=DB.Table(tableName).Where(map[string]string{"id":id,}).Updates(jsonData)
		if result.RowsAffected >0 {
			return true
		}
	}
	return false
}
/*
配合amis的crud组件封装的添加
tableName 表名
data.strColumns 非json字段
data.jsonColumns json字段
data.defaultColumns 默认字段，如addtime
*/
func AmisCrudInsert(tableName string,data map[string]interface{},c *gin.Context)(bool){
	jsonData:=map[string]interface{}{}
    if jsonColumns, ok := data["jsonColumns"].([]string); ok {//有json格式的字段
		jsonData,_=request.GetJsonData(append(data["strColumns"].([]string), data["jsonColumns"].([]string)...),c)//获取json参数
		for _, column := range jsonColumns {
			temp:=make([]interface{}, 0)
			temp1 := utils.GetMapDefule(jsonData,column,temp)
			temp2 := temp1.([]interface{})
			temp3,_ := json.Marshal(temp2)
			jsonData[column]=string(temp3)
		}
    }else{
		jsonData,_=request.GetJsonData(data["strColumns"].([]string),c)//获取json参数
    }
    if defaultColumns, ok := data["defaultColumns"].(map[string]string); ok {
		for k,v:= range defaultColumns {
			jsonData[k]=v
		}
	}
	result := DB.Table(tableName).Create(&jsonData)
	if result.Error != nil {
		return false
	}
	return true
}
//判断这一个[]string切片中，是否全部都是数字类型，如果是就转成[]int切片
func areAllStringsNumbers(strSlice []string) (bool, []int) {
	var intSlice []int
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			// 如果有任何一个字符串不能成功转换为整数，则返回 false
			return false, nil
		}
		intSlice = append(intSlice, num)
	}
	// 如果所有字符串都成功转换为整数，则返回 true 和转换后的整数切片
	return true, intSlice
}