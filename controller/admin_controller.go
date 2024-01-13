package controller

import (
	"encoding/json"
	"webapp/database"
	"webapp/config"
	"webapp/utils"
	"webapp/model/common/response"
	"webapp/model/common/request"
	"webapp/model/common/amis"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)
func AdminData(c *gin.Context) {
	where:=map[string]interface{}{
		"whereAnd":[]string{"status","superadmin",},
		// "whereJsonAnd":[]string{"role",},
		"whereJsonOr":[]string{"role",},
		"whereAndLike":[]string{"name","admin",},
		"whereBetween":[]string{"addtime",},
		"keyword":[]string{"name","admin",},
		"select":"id,admin,name,superadmin,role,addtime,status",
		"jsonColumn":[]string{"role",},
	}
	total,rows:=database.AmisCrudSelect("admin",where,c)
	response.OkWithData(map[string]interface{}{
		"count":total,
		"rows":rows,
	},c)
	return
}
func AdminEdit(c *gin.Context) {
    data:=map[string]interface{}{
        //post提交的正常字段
        "strColumns":[]string{"name","superadmin"},
        //post提交的json字段
        "jsonColumns":[]string{"role"},
        //默认值字段
    }
    i:=database.AmisCrudUpdate("admin",data,c)
	if i {
		response.Ok(c)
	}else{
		response.Fail(c)
	}
	return
}
func AdminDelete(c *gin.Context) {
	id:=c.Param("id")
	database.DB.Table("admin").Where(map[string]string{"id":id,}).Delete(nil)
	response.Ok(c)
}
func AdminRepwd(c *gin.Context) {
	id:=c.Param("id")
	jsonData,err:=request.GetJsonData([]string{"password"},c)
    fmt.Println(jsonData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    addtime := time.Now().Unix()
	addtimes := fmt.Sprintf("%d", addtime)
	salt:=utils.Sha256(addtimes)
	password:=utils.GetMapDefule(jsonData,"password","")
	pwdjs:=utils.Sha256(config.SALT+salt+password.(string))
	save:=map[string]interface{}{
		"password":pwdjs,
		"salt":salt,
	}
	result :=database.DB.Table("admin").Where(map[string]string{"id":id,}).Updates(save)
	if result.RowsAffected >0 {
		response.Ok(c)
	}else{
		response.Fail(c)
	}
	return
}
func AdminAdd(c *gin.Context) {
	jsonData,err:=request.GetJsonData([]string{"admin", "name","password","superadmin"},c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var admin map[string]interface{}
	database.DB.Table("admin").Where(map[string]string{"admin":jsonData["admin"].(string)}).Find(&admin)
	if len(admin)>0 {
		response.FailWithMessage("账号已经存在.", c)
	    return
	}
    // 获取当前时间的秒级时间戳
    addtime := time.Now().Unix()
	addtimes := fmt.Sprintf("%d", addtime)
	tempo:=make([]interface{}, 0)
	temp_data_role := utils.GetMapDefule(jsonData,"role",tempo)
	data_role := temp_data_role.([]interface{})
	data_roles,_ := json.Marshal(data_role)

	salt:=utils.Sha256(addtimes)
	password:=utils.GetMapDefule(jsonData,"password","")
	pwdjs:=utils.Sha256(config.SALT+salt+password.(string))
	fmt.Println("-----",utils.GetMapDefule(jsonData,"superadmin","0"))
	// 插入数据，使用 map
	data := map[string]interface{}{
		"name": utils.GetMapDefule(jsonData,"name",""),
		"admin": utils.GetMapDefule(jsonData,"admin",""),
		"password":pwdjs,
		"salt":salt,
		"addtime":addtime,
		"status":1,
		"superadmin":utils.GetMapDefule(jsonData,"superadmin",0),
		"role": data_roles,
	}
	result := database.DB.Table("admin").Create(&data)
	if result.Error != nil {
		response.FailWithMessage(result.Error.Error(), c)
		return
	}
	response.Ok(c)
	return
}
func AdminSchema(c *gin.Context) {
	edit:=amis.CKBS("AdminEdit",c)
	response.OkWithData(gin.H{
        "type": "crud",
        "syncLocation": false,
        "api": "/admin/data",
        "quickSaveItemApi": "put:/admin/data/${id}",
        "saveImmediately": true,
        "name": "dataList",
        "filterTogglable": true,
        "headerToolbar": []interface{}{
            
            map[string]interface{}{
                "type": "reload",
                "align": "right",
                "icon": "fa fa-refresh",
                "label": "刷新",
                "tooltip": "",
                "level": "success",
            },
            map[string]interface{}{
                "title": "",
                "type": "form",
                "mode": "inline",
                "className":"pull-left w-xl",
                "target": "dataList",
                "wrapWithPanel": false,
                "controls": []interface{}{
                  map[string]interface{}{
                    "type": "text",
                    "name": "keyword",
                    "size":"lg",
                    "placeholder": "通过关键字搜索",
                    "clearable": true,
                    "addOn": map[string]interface{}{
                      "type": "submit",
                      "icon": "fa fa-search",
                      "level": "primary",
                    },
                  },
                },
            },
            map[string]interface{}{
                "label": "添加用户",
                "type": "button",
                "actionType": "dialog",
                "level": "success",
                "className":amis.CK("AdminAdd",c),
                "dialog": map[string]interface{}{
                    "title": "添加用户",
                    "size": "xs",
                    "actions": []interface{}{
                        map[string]interface{}{
                            "type": "action",
                            "actionType": "submit",
                            "level": "primary",
                            "label": "保存",
                        },
                    },
                    "body": map[string]interface{}{
                        "type": "form",
                        "api": "post:/admin/data",
                        "reload": "userList",
                        "data":map[string]interface{}{"superadmin":0},
                        "controls": []interface{}{
                            map[string]interface{}{
                                "type": "text",
                                "name": "name",
                                "value": "${UUID(2)}",
                                "label": "姓名/昵称",
                                "required":1,
                                "validations":map[string]interface{}{
                                    "maxLength":20,
                                },
                            },
                            map[string]interface{}{
                                "type": "text",
                                "name": "admin",
                                "label": "账号",
                                "value": "${UUID(8)}",
                                "required":1,
                                "suffix": "",
                                "validations":map[string]interface{}{
                                    "minLength":5,
                                    "maxLength":40,
                                },
                            },
                            map[string]interface{}{
                                "type": "input-text",
                                "name": "password",
                                "label": "密码",
                                "value": "${UUID(12)|upperCase}${UUID(4)}",
                                "required":1,
                                "suffix": "",
                                "validations":map[string]interface{}{
                                    "minLength":8,
                                    "maxLength":40,
                                },
                            },
                            map[string]interface{}{
                                "name": "superadmin",
                                "type": "checkbox",
                                "label": "",
                                "option": "超级管理员",
                                "trueValue": 1,
                                "falseValue": 0,
                                // "value":1,
                          	},
                            map[string]interface{}{
                                "name":"role",
                                "label":"角色",
                                "visibleOn": "${superadmin == 0}",
                                "required":1,
                                "type":"select",
                                "searchable": true,
                                "maxTagCount": 3,
                                "checkAll":true,
                                "clearable":true,
                                "multiple":true,
                                "joinValues":false,
                                "extractValue":true,
                                "source":map[string]interface{}{"method":"get","url":"/role/min","cache":4000},
                            },
                            
                        },
                    },
                },
            },
            "pagination",
        },
        "defaultParams": map[string]interface{}{
            "perPage": 20,
        },
        "footerToolbar": []interface{}{"statistics", "switch-per-page", "pagination"},
        "columns":[]interface{}{
            map[string]interface{}{
                "name": "id",
                "label": "ID",
            },
            map[string]interface{}{
                "name": "admin",
                "label": "账号",
                "searchable":true,
                "tpl":"${admin}",
            },
            map[string]interface{}{
                "name": "name",
                "searchable": true,
                "type":"tpl",
                "tpl":"<span style='vertical-align: middle'>${name}</span>",
                "quickEditEnabledOn":edit,
                "quickEdit": map[string]interface{}{
                    "type": "text",
                    "validations":map[string]interface{}{
                        "maxLength":20,
                    },
                },
                "label": "姓名/昵称",
            },
            map[string]interface{}{
                "name": "superadmin",
                "label": "超级管理员",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":map[int]interface{}{1:"<span class='text-success'>是</span>",0:"否"},
                "quickEdit": map[string]interface{}{
                    "type":"select",
                	"options":map[int]interface{}{1:"是",0:"否"},
                },
                "searchable": map[string]interface{}{
                	"type":"select",
                	"options":map[int]interface{}{1:"是",0:"否"},
                },
            },
            map[string]interface{}{
                "name": "role",
                "label": "角色",
                "type":"mapping",
                "maxTagCount": 3,
                "quickEditEnabledOn":edit,
                "source":map[string]interface{}{"method":"get","url":"/role/min","cache":4000},
                "quickEdit": map[string]interface{}{
                    "name":"role",
                    "required":1,
                    "type":"select",
                    "maxTagCount": 3,
                    "checkAll":true,
                    "searchable": true,
                    "clearable":true,
                    "multiple":true,
                    "joinValues":false,
                    "extractValue":true,
                    "source":map[string]interface{}{"method":"get","url":"/role/min","cache":4000},
                },
                "searchable": map[string]interface{}{
                    "required":1,
                    "type":"select",
                    "maxTagCount": 3,
                    "checkAll":true,
                    "searchable": true,
                    "clearable":true,
                    "multiple":true,
                    "extractValue":true,
                    "source":map[string]interface{}{"method":"get","url":"/role/min","cache":4000},
                },
            },
            map[string]interface{}{
                "name": "addtime",
                "label": "添加时间",
                "type":"datetime",
                "sortable":true,
                "searchable": map[string]interface{}{
                  "type": "input-datetime-range",
                  "timeFormat": "HH:mm:ss",
                },
            },
            map[string]interface{}{
                "name": "status",
                "label": "状态",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":map[int]interface{}{1:"<span class='text-success'>正常</span>",0:"<span class='text-danger'>锁定</span>"},
                "quickEdit": map[string]interface{}{
                    "type":"select",
                	"options":map[int]interface{}{1:"正常",0:"锁定"},
                },
                "searchable": map[string]interface{}{
                	"type":"select",
                	"options":map[int]interface{}{1:"正常",0:"锁定"},
                },
            },
            map[string]interface{}{
                "type": "operation",
                "label": "操作",
                "width": 120,
                "fixed": "right",
                "buttons": []interface{}{
                    map[string]interface{}{
                        "level": "link",
                        "label": "删除",
                        "type": "button",
                        "className": "text-danger "+amis.CK("AdminDelete",c),
                        "actionType": "ajax",
                        "confirmText": "确认要删除？",
                        "api": "delete:/admin/data/${id}",
                    },

                    map[string]interface{}{
                        "label": "重置密码",
                        "level": "link",
                        "type": "button",
                        "className": "text-info "+amis.CK("AdminEdit",c),
                        "actionType": "dialog",
                        "dialog": map[string]interface{}{
                            "title": "修改密码",
                            "size": "xs",
                            "body": map[string]interface{}{
                                "type": "form",
                                "api": "post:/admin/repwd/${id}",
                                "reload": "userList",
                                "controls": []interface{}{
                                    map[string]interface{}{
                                        "type": "input-text",
                                        "name": "password",
                                        "label": "新密码",
                                		"value": "${UUID(12)|upperCase}${UUID(4)}",
                                        "required":1,
                                        "suffix": "",
                                        "validations":map[string]interface{}{
                                            "minLength":8,
                                            "maxLength":40,
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