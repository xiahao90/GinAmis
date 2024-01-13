package controller

import (
	"encoding/json"
	"webapp/config"
	"webapp/database"
	"webapp/utils/nbcache"
	"webapp/utils"
	"webapp/model/common/response"
	"webapp/model/common/request"
	"webapp/model/common/amis"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mojocn/base64Captcha"
   	"image/color"
	"time"
	"fmt"
)

var store = base64Captcha.DefaultMemStore

//首页
func HomeGet(c *gin.Context) {
	// 返回 html
	resp := gin.H{
		"system_name":  config.SYSTEM_NAME,
	}
	c.HTML(http.StatusOK, "index.html", resp)
}
//获取验证码
func HomeImgcode(c *gin.Context) {
	uuid,code,answer,_:=MakeCaptcha()
	// 设置 key-value 并设置默认过期时间
	nbcache.SetCahce("code-"+uuid, answer, 60*time.Minute)
	if gin.Mode() == "release" {
		answer = ""
	}
	response.OkWithDetailed(gin.H{
		"uuid": uuid,
		"img":code,
		"code":answer,
	}, "", c)
}
//提交登录
func HomeSignin(c *gin.Context) {
	jsonData,err:=request.GetJsonData([]string{"uuid", "code", "password", "username"},c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	key :="code-"+jsonData["uuid"].(string)
	code,_:=nbcache.GetCache(key)
	nbcache.DeleteCache(key)
	if code!= jsonData["code"].(string) {
		response.FailWithMessage("验证码错误", c)
		return 
	}
	password,err:=utils.Base64decode(jsonData["password"].(string))
	if err != nil {
		response.FailWithMessage("密码格式错误", c)
		return
	}
	var admin map[string]interface{}
	database.DB.Table("admin").Where(map[string]string{"admin":jsonData["username"].(string)}).Find(&admin)
	if len(admin)==0 {
		response.FailWithMessage("用户名与密码错误.", c)
	    return
	}
	if admin["status"].(int8)==0 {
		response.FailWithMessage("账号已经锁定", c)
	    return
	}
	pwdjs:=utils.Sha256(config.SALT+admin["salt"].(string)+password)
	if pwdjs!=admin["password"].(string) {
		response.FailWithMessage("用户名与密码错误", c)
	    return
	}
	var role_path []string
	rolename := ""
	superadmin:=0
	if admin["superadmin"].(int8)==1 {
		superadmin=1
		rolename="超级管理员"
	}else{
		var roleid []int
		// 解码 JSON 数据
		err = json.Unmarshal([]byte(admin["role"].(string)), &roleid)
		if err != nil {
			response.FailWithMessage("JSON 解码错误", c)
			return
		}
		var roles []map[string]interface{}
		database.DB.Table("role").Select("name,data_role").Where("id IN ?", roleid).Find(&roles)
		// 打印查询结果
		for _, role := range roles {
			if rolename == "" {
				rolename = role["name"].(string)
			}
			var path []string
			// 解码 JSON 数据
			err = json.Unmarshal([]byte(role["data_role"].(string)), &path)
			if err != nil {
				response.FailWithMessage("JSON 解码错误.", c)
				return
			}
			// 使用 range 遍历切片
			for _, value := range path {
				role_path = append(role_path,value)
			}
		}
		role_path=removeRoleDuplicates(role_path)
	}
	jwtCode,err:=utils.JetEncode(admin["id"].(int32),admin["name"].(string),superadmin,role_path,rolename,config.JWT_SK,config.JWT_EXTIME)
	response.OkWithDetailed(gin.H{
		"token": jwtCode,
	}, "登录成功", c)
}
//权限去重
func removeRoleDuplicates(input []string) []string {
	uniqueMap := make(map[string]bool)
	result := []string{}

	for _, val := range input {
		// 如果值在 map 中不存在，表示是新元素，则添加到结果切片中，并在 map 中标记为已存在
		if _, ok := uniqueMap[val]; !ok {
			uniqueMap[val] = true
			result = append(result, val)
		}
	}

	return result
}
//获取验证码
func MakeCaptcha() (string, string,string, error) {
    //定义一个driver
    var driver base64Captcha.Driver
    //创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
    driverString := base64Captcha.DriverString{
        Height:          30,                                     //高度
        Width:           100,                                    //宽度
        NoiseCount:      0,                                      //干扰数
        ShowLineOptions: 2 | 4,                                  //展示个数
        Length:          4,                                      //长度
        Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", //验证码随机字符串来源
        BgColor: &color.RGBA{ // 背景颜色
            R: 3,
            G: 102,
            B: 214,
            A: 125,
        },
        Fonts: []string{"wqy-microhei.ttc"}, // 字体
    }
    driver = driverString.ConvertFonts()
    //生成验证码
    c := base64Captcha.NewCaptcha(driver, store)
    id,b64s,answer,err := c.Generate()
    return id,b64s,answer,err
}
//生成环境开启Schema缓存，避免每次点击菜单都获取页面Schema
func SchemaApiCache(api string)(map[string]interface{}){
    cache:=60*60*1000 //缓存1个小时
	if gin.Mode() != "release" {
		cache=1//不缓存
	}
	return map[string]interface{}{
		"method": "get",
		"url": api,
		"cache":cache,
	}
}
//修改自己的密码
func HomeRepwd(c *gin.Context) {
	jsonData,err:=request.GetJsonData([]string{"code","password","uuid","old_pwd"},c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id, exists := c.Get("uid")
	if exists ==false {
		response.FailWithMessage("请重新登录.", c)
	    return
	}
	key :="code-"+jsonData["uuid"].(string)
	code,_:=nbcache.GetCache(key)
	nbcache.DeleteCache(key)
	if code!= jsonData["code"].(string) {
		response.FailWithMessage("验证码错误", c)
		return 
	}
	oldpassword:=jsonData["old_pwd"].(string)
	var admin map[string]interface{}
	database.DB.Table("admin").Where(map[string]string{"id":fmt.Sprintf("%v", id)}).Find(&admin)
	if len(admin)==0 {
		response.FailWithMessage("请重新登录.", c)
	    return
	}
	
	pwdjs:=utils.Sha256(config.SALT+admin["salt"].(string)+oldpassword)
	if pwdjs!=admin["password"].(string) {
		response.FailWithMessage("旧密码错误，如忘记，联系管理员重置.", c)
	    return
	}

    addtime := time.Now().Unix()
	addtimes := fmt.Sprintf("%d", addtime)
	salt:=utils.Sha256(addtimes)
	password:=utils.GetMapDefule(jsonData,"password","")
	pwdjs=utils.Sha256(config.SALT+salt+password.(string))
	save:=map[string]interface{}{
		"password":pwdjs,
		"salt":salt,
	}
	result :=database.DB.Table("admin").Where(map[string]string{"id":fmt.Sprintf("%v", id)}).Updates(save)
	if result.RowsAffected >0 {
		response.Ok(c)
	}else{
		response.Fail(c)
	}
	return
}
//修密码的页面
func HomePwdSchema(c *gin.Context) {
	response.OkWithData(gin.H{
        "type": "form",
        "wrapWithPanel":false,
        "api": "post:/repwd",
        "controls": []interface{}{
        	map[string]interface{}{
        	    "type": "input-password",
        	    "name": "old_pwd",
        	    "label": "旧密码",
        	    "required":1,
        	},
            map[string]interface{}{
                "type": "input-password",
                "name": "password",
                "label": "新密码",
                "required":1,
                "suffix": "",
                "validations":map[string]interface{}{
                  	"matchRegexp":`/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[~`+"`!@#$%^&*()_\\-+=])^.{8,32}$/",
                },
                "description":"推荐密码： ${UUID(12)|upperCase}${UUID(4)}",
                "validationErrors": map[string]interface{}{
                   	"matchRegexp": "密码必须包含大小写字母加数字加“~`!@#$%^&*()_-+=”里面的特殊字符,8-32位",
             	},
            },
            map[string]interface{}{
                "type": "input-password",
                "name": "password2",
                "label": "重复密码",
                "required":1,
                "suffix": "",
                "validations":map[string]interface{}{
                    "equalsField":"password",
                },
               	"validationErrors": map[string]interface{}{
                  	"equalsField": "两次输入密码不一致",
        		},
            },
            map[string]interface{}{
    			"type": "flex",
				"justify":"start",
				"items": []interface{}{
					map[string]interface{}{
					  	"type": "input-text",
					  	"name": "code",
					  	"label": "验证码",
					  	"required": true,
					  	"placeholder": "请输入验证码",
					  	"size": "sm",
					},
					map[string]interface{}{
					  	"type": "hidden",
					  	"name": "uuid",
					},
					map[string]interface{}{
					  	"type": "service",
					  	"api": "/../imgcode",
					  	"id": "service-reload",
					  	"body": map[string]interface{}{
						    "type":"tpl",
						    "className":"m-t-sm m-l",
						    "tpl":"<img src='${img}' class='pull-right' id='captcha' style='cursor: pointer;margin-top: 28px;'>",
						    "onEvent": map[string]interface{}{
						      	"click": map[string]interface{}{
						        	"actions": []interface{}{
						          		map[string]interface{}{
						            		"componentId": "service-reload",
						            		"actionType": "reload",
						          		},
						        	},
						      	},
						    },
					  	},
					},
				},
            },
        },
	},c)
}
//菜单
func HomePage(c *gin.Context) {
	response.OkWithData(gin.H{
		"pages": []interface{}{
			map[string]interface{}{
		        "label":"系统管理",
		        "url":"/",
		        "schema": map[string]interface{}{
		            "type": "page",
		            "body": "欢迎来到"+config.SYSTEM_NAME,
		        },
		        "children": []interface{}{
		        	map[string]interface{}{
		        	    "label":"用户管理",
		        	    "icon":"fa fa-grin-alt",
		        	    "url":"/admin",
		        	    "schemaApi": SchemaApiCache("/admin/schema"),
		        	    "className":amis.CK("AdminSchema",c),
		        	},
		        	map[string]interface{}{
		        	    "label":"角色管理",
		        	    "icon":"fa fa-baby",
		        	    "url":"/role",
		        	    "schemaApi": SchemaApiCache("/role/schema"),
		        	    "className":amis.CK("RoleSchema",c),
		        	},
		        },
		    },
		    map[string]interface{}{
		        "label":"数据管理",
		        "children": []interface{}{
		        	map[string]interface{}{
		        	    "label":"测试数据",
		        	    "icon":"fa fa-snowflake",
		        	    "url":"/test",
		        	    "schemaApi": SchemaApiCache("/test/schema"),
		        	    "className":amis.CK("TestSchema",c),
		        	},
		        },
		    },
		},
	},c)
}

