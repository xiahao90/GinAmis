<h1 align="center">GinAmis</h1>

<div align="center">基于Gin+Amis+GORM+Mysql8 实现的权限管理与数据库CRUD快速操作的脚手架，目的是提供一套基础开发框架，方便、快速的完成业务需求的开发。
<br/>

</div>

## 特点

- 遵循 `RESTful API` 设计规范&基于接口的编程规范
- 基于 GORM与 `Amis`前端提交的特性封装了快速查询，添加，编辑的接口代码（快捷CRUD）
- 基于 `Amis`的app多页应用组件设计，全部用schemaApi方式获取json渲染页面，不用再写前端代码，对后端开发者友好
- 基于 `JWT` 的用户认证 -- 基于 JWT 的权限验证机制
- 基于 `go mod` 的依赖管理(国内源可使用：<https://goproxy.cn/>)

## 开发赖版本

- [go](https://go-zh.org/) 1.21.5 
- [Gin](https://gin-gonic.com/) v1.9.1
- [Mysql](https://www.mysql.com/) 8.x
- [Amis](https://github.com/) 6.0.0

## 快速开始

- 1，导入mysql8.sql数据库
- 2，修改配置文件GinAmis/config/conf.go相关配置
- 3，启动项目`go run main.go`
- 4，访问你的地址：[http://127.0.0.1:8888](http://127.0.0.1:8888)
- 5，账号密码 `admin/123456`

## 注意事项

- 1，替换jwt的密钥
- 2，release启动项目的时候，不会填充验证码

## 开发说明（增加新功能）

- 1，需要在`home_controller.go`底部，编写代码添加amis的菜单，参考`测试数据`配置
- 2，需要在`role_controller.go`底部，编写权限配置项，参考`测试数据`配置
- 3，参考`test_controller.go`代码编写相关功能
- 4，查询封装说明，深度结合Amis的查询提交，代码位置test_controller.go，`AmisCrudSelect`
```
	where:=map[string]interface{}{
		"whereAnd":[]string{"number","radios","switch","checkbox","select"},
		"whereJsonAnd":[]string{"checkboxes","select_1"},
		"whereJsonOr":[]string{"tag"},
		"whereAndLike":[]string{"input","textarea"},
		"whereBetween":[]string{"addtime","updatetime","datetime"},
		"keyword":[]string{"input","password","textarea"},
		"select":"id,admin,name,superadmin,role,addtime,status",
		"jsonColumn":[]string{"tag","checkboxes","select_1"},
	}
	//whereAnd，数据库字段等于的查询，如状态字段staus
	//whereJsonAnd，数据库json数组格式字段，多个选项and的方式查询，如标签tag字段在数据库中是个数组格式的json，前端提交了两个值，需要同时包含“zhugeliang“，”caocao”
	//whereJsonOr，与whereJsonAnd类似，同时包含变成了包含其中一项就行
	//whereAndLike，数据库字段like查询，如查询身份证时候，输入前6位就行
	//whereBetween，数据库字段范围查询，如查询添加时间
	//keyword，数据哪些字段需要模糊查询，如前端提交了一个“大哥”，数据库需要在name与info字段中模糊匹配，用于全局搜索框
	//select，查询哪些字段
	//jsonColumn，哪些字段是json格式，需要填写
```

- 5，编辑封装说明，深度结合Amis的快捷编辑功能，代码位置test_controller.go，`AmisCrudUpdate`
```
	data:=map[string]interface{}{
		"strColumns":[]string{"name","info"},
		"jsonColumns":[]string{"tag"},
		"defaultColumns":map[string]string{
		"updatetime":addtime,
		},
	}
	//strColumns，数据库非json格式的字段
	//jsonColumns，数据库json格式的字段
	//defaultColumns，非前端提交的默认值
```

- 6，添加封装说明，深度结合Amis的form组件的提交功能，代码位置test_controller.go，`AmisCrudInsert`
```
	与上一条编辑封装说明类似
```

## 相关截图

<img width="1088" alt="image" src="https://github.com/xiahao90/GinAmis/assets/11575908/95a15a2f-cc02-457b-add1-6345886a1331">
<img width="1656" alt="image" src="https://github.com/xiahao90/GinAmis/assets/11575908/8f5b0b36-df98-47f1-948d-8500d458ad10">
<img width="1643" alt="image" src="https://github.com/xiahao90/GinAmis/assets/11575908/bd25c813-a4e1-47c8-8a49-72827f44bc51">
<img width="1141" alt="image" src="https://github.com/xiahao90/GinAmis/assets/11575908/5a864fd8-510e-4348-967a-b260cb45cd5e">
<img width="902" alt="image" src="https://github.com/xiahao90/GinAmis/assets/11575908/27402940-ca23-46b2-83e0-540b87b87eb2">



