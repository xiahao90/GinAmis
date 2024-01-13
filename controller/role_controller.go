package controller

import (
	"encoding/json"
	"webapp/database"
	"webapp/utils"
	"webapp/model/common/response"
	"webapp/model/common/request"
	"webapp/model/common/amis"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"strings"
)
//登录后可以访问的默认权限
var DefuleAuto = []string{
	"HomePage",
	"HomePwdSchema",
	"HomePwd",
	"HomeRepwd",
}

func RoleData(c *gin.Context) {
	where:=map[string]interface{}{
		"keyword":[]string{"name","info",
		},
		"jsonColumn":[]string{"data_role",},
		// "whereJsonAnd":[]string{"data_role",},
		"whereJsonOr":[]string{"data_role",},
	}
	total,rows:=database.AmisCrudSelect("role",where,c)
	response.OkWithData(map[string]interface{}{
		"count":total,
		"rows":rows,
	},c)
	return
}

func RoleMin(c *gin.Context) {
	var rows []map[string]interface{}
	database.DB.Table("role").Select("id as value,name as label").Find(&rows)
	response.OkWithData(map[string]interface{}{
		"options":rows,
	},c)
	return
}
func RoleDelete(c *gin.Context) {
	id:=c.Param("id")
	database.DB.Table("role").Where(map[string]string{"id":id,}).Delete(nil)
	response.Ok(c)
}
func RoleEdit(c *gin.Context) {
	id:=c.Param("id")
	jsonData,err:=request.GetJsonData([]string{"name","info", "data_role"},c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	temp_data_role := utils.GetMapDefule(jsonData,"data_role","[]")
	data_role := getDateList(temp_data_role.([]interface{}))
	data_role = removeRoleDuplicates(data_role)
	data_roles,_ := json.Marshal(data_role)
	save := map[string]interface{}{
		"name": utils.GetMapDefule(jsonData,"name",""),
		"info": utils.GetMapDefule(jsonData,"info",""),
		"data_role": data_roles,
	}
	result :=database.DB.Table("role").Where(map[string]string{"id":id,}).Updates(save)
	if result.RowsAffected >0 {
		response.Ok(c)
	}else{
		response.Fail(c)
	}
	return
}
func RoleCopy(c *gin.Context) {
	id:=c.Param("id")
	var role map[string]interface{}
	database.DB.Table("role").Where(map[string]string{"id":id,}).Find(&role)
	delete(role, "id")
	result := database.DB.Table("role").Create(&role)
	if result.Error != nil {
		response.FailWithMessage(result.Error.Error(), c)
		return
	}
	response.Ok(c)
	return
}
func RoleAdd(c *gin.Context) {
	jsonData,err:=request.GetJsonData([]string{"name", "data_role"},c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    // 获取当前时间的秒级时间戳
    addtime := time.Now().Unix()
	addtimes := fmt.Sprintf("%d", addtime)

	temp_data_role := utils.GetMapDefule(jsonData,"data_role","[]")
	data_role := getDateList(temp_data_role.([]interface{}))
	data_role = removeRoleDuplicates(data_role)
	data_roles,_ := json.Marshal(data_role)
	// 插入数据，使用 map
	data := map[string]interface{}{
		"name": utils.GetMapDefule(jsonData,"name","角色"+addtimes),
		"info": utils.GetMapDefule(jsonData,"info",""),
		"data_role": data_roles,
	}
	result := database.DB.Table("role").Create(&data)
	if result.Error != nil {
		response.FailWithMessage(result.Error.Error(), c)
		return
	}
	response.Ok(c)
	return
}
func hasKey(data map[string]interface{},key string) bool {
    _, exists := data[key]
    return exists
}
//传入前端的提交的权限，计算出来默认的权限
func getDateList(data []interface{})([]string){
    // tempData := make([]string,0)
    tempData := DefuleAuto
	tempstr:=""
	for _, x := range data {
		tempData = append(tempData,x.(string))
		tempstr = tempstr+x.(string)+";"
	}
    menu:=config_menu()
    for _, topLevelItem := range menu {
    	topLevelLabel :=make([]interface{},0)
    	if hasKey(topLevelItem.(map[string]interface{}),"auto") {
    		topLevelLabel = topLevelItem.(map[string]interface{})["auto"].([]interface{})
    	}
    	children := topLevelItem.(map[string]interface{})["children"].([]interface{})
    	for _, childItem := range children {
    		childValue := childItem.(map[string]interface{})["value"].(string)
    		if strings.Contains(tempstr, childValue){//在里面
    			for _, x := range topLevelLabel {//获取父级的默认权限加进去
					tempData = append(tempData,x.(string))
    			}
    			childAuth :=make([]interface{},0)
    			if hasKey(childItem.(map[string]interface{}),"auto") {
    				childAuth = childItem.(map[string]interface{})["auto"].([]interface{})//获取本级的默认权限加进去
	    			for _, x1 := range childAuth {
						tempData = append(tempData,x1.(string))
	    			}
    			}
    			fmt.Println(tempData)
    		}
    	}
    }
    return tempData
}
func RoleSchema(c *gin.Context) {
	edit:=amis.CKBS("RoleEdit",c)
	fmt.Println("RoleSchema:",edit)
	role_map:=getKV()
	response.OkWithData(gin.H{
	    "type": "crud",
	    "syncLocation": false,
	    "api": "/role/data",
	    "quickSaveItemApi": "put:/role/data/${id}",
	    "saveImmediately": true,
	    "name": "dataList",
	    "id":"dataList",
	    "filterTogglable": true,
	    // "perPageAvailable":[]int{5, 10, 20, 50, 100,500,1000,2000,3000,5000},
	    "headerToolbar":[]interface{}{
	        map[string]interface{}{
	            "title": "",
	            "type": "form",
	            "mode": "inline",
	            "className":"pull-left w-xl",
	            "target": "dataList",
	            "wrapWithPanel": false,
	            "controls": []interface{}{
	              	map[string]interface{}{
		                "type": "input-text",
		                "name": "keyword",
		                "size":"lg",
		                "placeholder": "通过关键字搜索",
		                "clearable": true,
		                "addOn": map[string]string{
		                  	"type": "submit",
		                  	"icon": "fa fa-search",
		                  	"level": "primary",
		                },
	              	},
	            },
	        },
	        map[string]interface{}{
	            "label":"角色添加",
	            "type":"button",
	            "actionType":"dialog",
	            "className":amis.CK("RoleAdd",c),
	            "level":"success",
	            "dialog":map[string]interface{}{
	                "title":"添加角色",
	                "size":"lg",
	                "body":map[string]interface{}{
	                    "type":"form",
	                    "api":"post:/role/data",
	                    "reload":"Listdata",
	                    "mode":"normal",
	                    "controls":[]interface{}{
	                        map[string]interface{}{
	                            "type":"input-text",
	                            "name":"name",
	                            "label":"名称",
	                            "required":1,
	                            "validations":map[string]interface{}{
	                                "minLength":1,
	                                "maxLength":255,
	                            },
	                        },
	                        map[string]interface{}{
	                            "type":"textarea",
	                            "name":"info",
	                            "label":"备注说明",
	                            "validations":map[string]interface{}{
	                                "minLength":1,
	                                "maxLength":6000,
	                            },
	                        },
	                        map[string]interface{}{
	                            "type":"checkboxes",
	                            "required":1,
	                            "size":"full",
	                            "name":"data_role",
	                            "label":"选择权限",
	                            "checkAll": true,
	                            "multiple":true,
	                            "joinValues":false,
	                            "extractValue":true,
	                            "cascade":true,
	                            "options":config_menu(),
	                        },
	                    },
	                },
	            },
	        },
	        map[string]interface{}{
	            "type":"reload",
	            "align":"right",
	            "icon":"fa fa-refresh",
	            "label":"刷新",
	            "tooltip":"",
	            "level":"success",
	        },
	        
	    },
	    "defaultParams":map[string]interface{}{
	        "perPage":20,
	    },
	    "footerToolbar":[]interface{}{"statistics", "switch-per-page", "pagination"},
	    "columns":[]interface{}{
	        map[string]interface{}{
	            "name":"id",
	            "label":"ID",
	        },
	        map[string]interface{}{
	            "name":"name",
	            "quickEditEnabledOn":edit,
	            "quickEdit":map[string]interface{}{
	                "type":"text",
	                "validations":map[string]interface{}{
	                    "maxLength":255,
	                },
	            },
	            "label":"名称",
	        },
	        map[string]interface{}{
	            "name":"info",
	            "quickEditEnabledOn":edit,
	            "quickEdit":map[string]interface{}{
	                "type":"text",
	                "validations":map[string]interface{}{
	                    "maxLength":6000,
	                },
	            },
	            "label":"备注说明",
	        },
	        map[string]interface{}{
	            "name":"data_role",
	            "type":"mapping",
	            "placeholder":"",
	            "map":role_map,
	            "width":"60%",
	            "quickEditEnabledOn":edit,
	            "quickEdit":map[string]interface{}{
	                "type":"checkboxes",
	                "required":1,
	                "size":"full",
	                "name":"data_role",
	                "label":"选择权限",
	                "checkAll": true,
	                "multiple":true,
	                "joinValues":false,
	                "extractValue":true,
	                "cascade":true,
	                "options":config_menu(),
	            },
	            "searchable": map[string]interface{}{
	                "type":"checkboxes",
	                "required":1,
	                "size":"full",
	                "name":"data_role",
	                "label":"选择权限",
	                "checkAll": true,
	                "multiple":true,
	                // "joinValues":false,
	                "extractValue":true,
	                "cascade":true,
	                "options":config_menu(),
	            },
	            "label":"权限设置",
	        },
	        map[string]interface{}{
	            "type":"operation",
	            "label":"操作",
	            "width":100,
	            "fixed":"right",
	            "buttons":[]interface{}{
	                map[string]interface{}{
	                    "label":"删除",
	                    "type":"action",
	                    "level":"link",
	                    "className": "text-danger "+amis.CK("RoleDelete",c),
	                    "actionType":"ajax",
	                    "confirmText":"确认要删除？",
	                    "api":"DELETE:/role/data/${id}",
	                },
	                map[string]interface{}{
	                    "label":"复制",
	                    "type":"action",
	                    "level":"link",
	                    "className":"text-info "+amis.CK("RoleAdd",c),
	                    "actionType":"ajax",
	                    "confirmText":"确认要复制？",
	                    "api":"post:/role/copy/${id}",
	                },
	            },
	        },
	    },
	},c)
}
func getKV()(map[string]interface{}){
	menu:=config_menu()
	kv := make(map[string]interface{})
	for _, topLevelItem := range menu {
		topLevelLabel := topLevelItem.(map[string]interface{})["label"].(string)
		children := topLevelItem.(map[string]interface{})["children"].([]interface{})
		for _, childItem := range children {
			childLabel := childItem.(map[string]interface{})["label"].(string)
			childValue := childItem.(map[string]interface{})["value"].(string)
			key := fmt.Sprintf("%s-%s", topLevelLabel, childLabel)
			kv[childValue] = key
		}
	}
	return kv
}
//权限配置菜单
func config_menu()([]interface{}) {
    return []interface{}{
        map[string]interface{}{
            "label":"账号管理",
            "auto":[]interface{}{
                "AdminSchema",
                "AdminMin",
                "AdminData",
                "RoleMin",
            },
            "children":[]interface{}{
                map[string]interface{}{
                    "label":"查看账号",
                    "value":"AdminData",
                },
                map[string]interface{}{
                    "label":"添加账号",
                    "value":"AdminAdd",
                },
                map[string]interface{}{
                    "label":"编辑账号",
                    "value":"AdminEdit",
                },
                map[string]interface{}{
                    "label":"删除账号",
                    "value":"AdminDelete",
                },
                map[string]interface{}{
                    "label":"重置密码",
                    "value":"AdminRepwd",
                },
            },
        },
        map[string]interface{}{
            "label":"角色管理",
            "auto":[]interface{}{
                "RoleSchema",
                "RoleMin",
                "RoleData",
            },
            "children":[]interface{}{
                map[string]interface{}{
                    "label":"查看角色",
                    "value":"RoleData",
                },
                map[string]interface{}{
                    "label":"添加角色",
                    "value":"RoleAdd",
                    "auto":[]interface{}{
                        "RoleCopy",
                    },
                },
                map[string]interface{}{
                    "label":"编辑角色",
                    "value":"RoleEdit",
                },
                map[string]interface{}{
                    "label":"删除角色",
                    "value":"RoleDelete",
                },
            },
        },
        map[string]interface{}{
            "label":"测试数据",
            "auto":[]interface{}{
                "TestSchema",
                "TestData",
            },
            "children":[]interface{}{
                map[string]interface{}{
                    "label":"查看数据",
                    "value":"TestData",
                },
                map[string]interface{}{
                    "label":"添加数据",
                    "value":"TestAdd",
                },
                map[string]interface{}{
                    "label":"编辑数据",
                    "value":"TestEdit",
                },
                map[string]interface{}{
                    "label":"删除数据",
                    "value":"TestDelete",
                },
            },
        },
    }
}