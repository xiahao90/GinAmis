package controller

import (
	// "encoding/json"
	"webapp/database"
	// "webapp/config"
	// "webapp/utils"
	"webapp/model/common/response"
	// "webapp/model/common/request"
	"webapp/model/common/amis"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)
var options123 =[]interface{}{
    map[string]interface{}{
        "label": "选项1",
        "value": 1,
    },
    map[string]interface{}{
        "label": "选项2",
        "value": 2,
    },
    map[string]interface{}{
        "label": "选项3",
        "value": 3,
    },
    map[string]interface{}{
        "label": "选项4",
        "value": 4,
    },
    map[string]interface{}{
        "label": "选项5",
        "value": 5,
    },
    map[string]interface{}{
        "label": "选项6",
        "value": 6,
    },
}
var optionsStr =[]interface{}{
    map[string]interface{}{
        "label": "诸葛亮",
        "value": "zhugeliang",
    },
    map[string]interface{}{
        "label": "曹操",
        "value": "caocao",
    },
    map[string]interface{}{
        "label": "钟无艳",
        "value": "zhongwuyan",
    },
}
func TestData(c *gin.Context) {
	where:=map[string]interface{}{
		"whereAnd":[]string{"number","radios","switch","checkbox","select"},
		"whereJsonAnd":[]string{"checkboxes","select_1"},
		"whereJsonOr":[]string{"tag"},
		"whereAndLike":[]string{"input","textarea"},
		"whereBetween":[]string{"addtime","updatetime","datetime"},
		"keyword":[]string{"input","password","textarea"},
		// "select":"id,admin,name,superadmin,role,addtime,status",
		"jsonColumn":[]string{"tag","checkboxes","select_1"},
	}
	total,rows:=database.AmisCrudSelect("test",where,c)
	response.OkWithData(map[string]interface{}{
		"count":total,
		"rows":rows,
	},c)
	return
}
func TestEdit(c *gin.Context) {
    // 获取当前时间的秒级时间戳
    timenow := time.Now().Unix()
    timenows := fmt.Sprintf("%d", timenow)
    data:=map[string]interface{}{
        //post提交的正常字段
        "strColumns":[]string{"checkbox","switch","number","input","password","radios","select","datetime","textarea"},
        //post提交的json字段
        "jsonColumns":[]string{"tag","checkboxes","select_1"},
        //默认值字段
        "defaultColumns":map[string]string{
            "updatetime":timenows,
        },
    }
    i:=database.AmisCrudUpdate("test",data,c)
	if i {
		response.Ok(c)
	}else{
		response.Fail(c)
	}
	return
}
func TestDelete(c *gin.Context) {
	id:=c.Param("id")
	database.DB.Table("test").Where(map[string]string{"id":id,}).Delete(nil)
	response.Ok(c)
}
func TestAdd(c *gin.Context) {
    // 获取当前时间的秒级时间戳
    addtime := time.Now().Unix()
    addtimes := fmt.Sprintf("%d", addtime)
    data:=map[string]interface{}{
        //post提交的正常字段
        "strColumns":[]string{"checkbox","switch","number","input","password","radios","select","datetime","textarea"},
        //post提交的json字段
        "jsonColumns":[]string{"tag","checkboxes","select_1"},
        //默认值字段
        "defaultColumns":map[string]string{
            "addtime":addtimes,
        },
    }
    i:=database.AmisCrudInsert("test",data,c)
    if i {
        response.Ok(c)
    }else{
        response.Fail(c)
    }
    return
}
//添加数据的json
func __add_schema(c *gin.Context)(map[string]interface{}){
    return map[string]interface{}{
        "label": "添加测试数据",
        "type": "button",
        "actionType": "dialog",
        "level": "success",
        "className":amis.CK("TestAdd",c),
        "dialog": map[string]interface{}{
            "title": "添加测试数据",
            "size": "lg",
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
                "api": "post:/test/data",
                "reload": "DataList",
                "data":map[string]interface{}{
                    "switch":0,
                    "checkbox":0,
                },
                "controls": []interface{}{
                    map[string]interface{}{
                        "type": "input-text",
                        "name": "input",
                        "label": "文本",
                        "description":"这是提示",
                        "required":1,
                    },
                    map[string]interface{}{
                        "type": "input-password",
                        "name": "password",
                        "required":1,
                        "label": "密码",
                        "validations":map[string]interface{}{
                            "minLength":8,
                            "maxLength":40,
                        },
                    },
                    map[string]interface{}{
                        "type": "input-number",
                        "required":1,
                        "name": "number",
                        "label": "数字",
                        "value": 5,
                        "min": 1,
                        "max": 10,
                    },
                    map[string]interface{}{
                        "type": "input-tag",
                        "required":1,
                        "name": "tag",
                        "label": "标签",
                        "placeholder": "",
                        "joinValues":false,
                        "extractValue":true,
                        "options": optionsStr,
                    },
                    map[string]interface{}{
                        "type": "checkboxes",
                        "required":1,
                        "name": "checkboxes",
                        "label": "多选框",
                        "joinValues":false,
                        "extractValue":true,
                        "clearable": true,
                        "options": options123,
                    },
                    map[string]interface{}{
                        "type": "radios",
                        "required":1,
                        "name": "radios",
                        "label": "单选框",
                        "options": options123,
                    },
                    map[string]interface{}{
                        "type": "switch",
                        "required":1,
                        "name": "switch",
                        "label": "开关",
                        "trueValue": 1,
                        "falseValue": 0,
                        "onText": "我开启了哦",
                        "offText": "关",
                    },
                    map[string]interface{}{
                        "name": "checkbox",
                        "type": "checkbox",
                        "required":1,
                        "option": "勾选框",
                        "trueValue": 1,
                        "falseValue": 0,
                    },
                    map[string]interface{}{
                        "name":"select",
                        "label":"下拉单选",
                        "type":"select",
                        "required":1,
                        "options": options123,
                    },
                    map[string]interface{}{
                        "name":"select_1",
                        "label":"下拉多选",
                        "type":"select",
                        "required":1,
                        "searchable": true,
                        "maxTagCount": 3,
                        "checkAll":true,
                        "clearable":true,
                        "multiple":true,
                        "joinValues":false,
                        "extractValue":true,
                        "options": options123,
                    },
                    map[string]interface{}{
                        "type": "input-datetime",
                        "required":1,
                        "name": "datetime",
                        "label": "日期+时间",
                    },
                    map[string]interface{}{
                        "type": "textarea",
                        "required":1,
                        "name": "textarea",
                        "label": "多行文本",
                    },
                },
            },
        },
    }
}
func __show_schema(c *gin.Context)(map[string]interface{}){
    add:=__add_schema(c)
    add["dialog"].(map[string]interface{})["body"].(map[string]interface{})["columnCount"]=2
    add["dialog"].(map[string]interface{})["body"].(map[string]interface{})["api"]=""
    show:=map[string]interface{}{
        "label":"详情",
        "type":"action",
        "level":"link",
        "actionType":"dialog",
        "dialog": map[string]interface{}{
            "title":"${input}",
            "size":"md",
            "closeOnEsc":true,
            "closeOnOutside":true,
            "body":add["dialog"].(map[string]interface{})["body"],
        },
    }
    show["dialog"].(map[string]interface{})["body"].(map[string]interface{})["static"]=true;
    show["dialog"].(map[string]interface{})["body"].(map[string]interface{})["controls"]=append(show["dialog"].(map[string]interface{})["body"].(map[string]interface{})["controls"].([]interface{}),map[string]interface{}{
        "type":"input-date",
        "name":"addtime",
        "label":"添加时间",
    })
    show["dialog"].(map[string]interface{})["body"].(map[string]interface{})["controls"]=append(show["dialog"].(map[string]interface{})["body"].(map[string]interface{})["controls"].([]interface{}),map[string]interface{}{
        "type":"input-date",
        "name":"updatetime",
        "label":"编辑时间",
    })
    return show
}
func __edit_schema(c *gin.Context)(map[string]interface{}){
    edit:=__add_schema(c);
    edit["label"]="编辑";
    edit["level"]="link";
    edit["className"]=amis.CK("TestEdit",c)
    edit["dialog"].(map[string]interface{})["title"]="${input}";
    edit["dialog"].(map[string]interface{})["body"].(map[string]interface{})["api"]="put:/test/data/${id}";
    return edit;
}


func TestSchema(c *gin.Context) {
	edit:=amis.CKBS("TestEdit",c)
	response.OkWithData(gin.H{
        "type": "crud",
        "syncLocation": false,
        "api": "/test/data",
        "quickSaveItemApi": "put:/test/data/${id}",
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
            __add_schema(c),
            "pagination",
        },
        "defaultParams": map[string]interface{}{
            "perPage": 20,
        },
        "footerToolbar": []interface{}{"statistics", "switch-per-page", "pagination"},
        "columns":[]interface{}{
            map[string]interface{}{
                "name": "id",
                "fixed": "left",
                "label": "ID",
            },
            map[string]interface{}{
                "name": "input",
                "fixed": "left",
                "label": "文本",
                "searchable":true,
                "quickEditEnabledOn":edit,
                "quickEdit":true,
            },
            map[string]interface{}{
                "name": "password",
                "label": "密码",
                "quickEditEnabledOn":edit,
                "quickEdit":map[string]interface{}{
                    "type": "input-password",
                    "validations":map[string]interface{}{
                        "minLength":8,
                        "maxLength":40,
                    },
                },
            },
            map[string]interface{}{
                "name": "number",
                "label": "数字",
                "searchable":true,
                "quickEditEnabledOn":edit,
                "quickEdit":map[string]interface{}{
                    "type": "input-number",
                    "min": 1,
                    "max": 10,
                },
            },
            map[string]interface{}{
                "name": "tag",
                "label": "标签",
                "type":"mapping",
                "maxTagCount": 3,
                "quickEditEnabledOn":edit,
                "map":optionsStr,
                "quickEdit": map[string]interface{}{
                    "type": "input-tag",
                    "joinValues":false,
                    "extractValue":true,
                    "options": optionsStr,
                },
                "searchable": map[string]interface{}{
                    "type": "input-tag",
                    // "joinValues":false,
                    "extractValue":true,
                    "options": optionsStr,
                },
            },
            map[string]interface{}{
                "name": "checkboxes",
                "label": "多选框",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":options123,
                "quickEdit": map[string]interface{}{
                    "type": "checkboxes",
                    "joinValues":false,
                    "extractValue":true,
                    "options": options123,
                },
                "searchable": map[string]interface{}{
                    "type": "checkboxes",
                    // "joinValues":false,
                    "extractValue":true,
                    "options": options123,
                },
            },
            map[string]interface{}{
                "name": "radios",
                "label": "单选框",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":options123,
                "quickEdit": map[string]interface{}{
                    "type": "radios",
                    "options": options123,
                },
                "searchable": map[string]interface{}{
                    "type": "radios",
                    "options": options123,
                },
            },
            map[string]interface{}{
                "name": "switch",
                "label": "开关",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":[]interface{}{
                    map[string]interface{}{
                        "label": "开",
                        "value": 1,
                    },
                    map[string]interface{}{
                        "label": "关",
                        "value": 0,
                    },
                },
                "quickEdit": map[string]interface{}{
                    "type": "switch",
                    "label": "开关",
                    "trueValue": 1,
                    "falseValue": 0,
                    "onText": "我开启了哦",
                    "offText": "关",
                },
                "searchable": map[string]interface{}{
                    "type": "select",
                    "options": []interface{}{
                        map[string]interface{}{
                            "label": "开",
                            "value": 1,
                        },
                        map[string]interface{}{
                            "label": "关",
                            "value": 0,
                        },
                    },
                },
            },
            map[string]interface{}{
                "name": "checkbox",
                "label": "勾选框",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":[]interface{}{
                    map[string]interface{}{
                        "label": "是",
                        "value": 1,
                    },
                    map[string]interface{}{
                        "label": "否",
                        "value": 0,
                    },
                },
                "quickEdit": map[string]interface{}{
                    "type": "checkbox",
                    "option": "勾选框",
                    "trueValue": 1,
                    "falseValue": 0,
                },
                "searchable": map[string]interface{}{
                    "type": "select",
                    "options": []interface{}{
                        map[string]interface{}{
                            "label": "是",
                            "value": 1,
                        },
                        map[string]interface{}{
                            "label": "关",
                            "value": 0,
                        },
                    },
                },
            },
            map[string]interface{}{
                "name": "select",
                "label": "下拉单选",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":options123,
                "quickEdit": map[string]interface{}{
                    "type":"select",
                    "options": options123,
                },
                "searchable": map[string]interface{}{
                    "type": "select",
                    "options": options123,
                },
            },
            map[string]interface{}{
                "name": "select_1",
                "label": "下拉多选",
                "type":"mapping",
                "quickEditEnabledOn":edit,
                "map":options123,
                "quickEdit": map[string]interface{}{
                    "type":"select",
                    "searchable": true,
                    "maxTagCount": 3,
                    "checkAll":true,
                    "clearable":true,
                    "multiple":true,
                    "joinValues":false,
                    "extractValue":true,
                    "options": options123,
                },
                "searchable": map[string]interface{}{
                    "type":"select",
                    "searchable": true,
                    "maxTagCount": 3,
                    "checkAll":true,
                    "clearable":true,
                    "multiple":true,
                    // "joinValues":false,
                    "extractValue":true,
                    "options": options123,
                },
            },
            map[string]interface{}{
                "name": "datetime",
                "label": "日期+时间",
                "type":"datetime",
                "searchable": map[string]interface{}{
                    "type": "input-datetime-range",
                    "timeFormat": "HH:mm:ss",
                },
                "quickEditEnabledOn":edit,
                "quickEdit": map[string]interface{}{
                    "type": "input-datetime",
                },
            },
            map[string]interface{}{
                "name": "textarea",
                "label": "多行文本",
                "searchable":true,
                "quickEditEnabledOn":edit,
                "quickEdit":map[string]interface{}{
                    "type": "textarea",
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
                "name": "updatetime",
                "label": "编辑时间",
                "type":"datetime",
                "sortable":true,
                "searchable": map[string]interface{}{
                  "type": "input-datetime-range",
                  "timeFormat": "HH:mm:ss",
                },
            },
            map[string]interface{}{
                "type": "operation",
                "label": "操作",
                "width": 150,
                "fixed": "right",
                "buttons": []interface{}{
                    __show_schema(c),
                    __edit_schema(c),
                    map[string]interface{}{
                        "level": "link",
                        "label": "删除",
                        "type": "button",
                        "className": "text-danger "+amis.CK("TestDelete",c),
                        "actionType": "ajax",
                        "confirmText": "确认要删除？",
                        "api": "delete:/test/data/${id}",
                    },
                },
            },
        },
    },c)
}