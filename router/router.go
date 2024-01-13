package router

import (
	"webapp/controller"
	"webapp/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 创建一个 gin 实例
	r := gin.Default()
	// 
	/*--------- 静态资源加载 --------*/
	// 静态资源
	r.Static("/static", "./static")
	/*--------- 加载前端模版 --------*/
	// 全部加载 view 下模版资源
	r.LoadHTMLGlob("view/*")

	// html页面 /
	r.GET("/", controller.HomeGet)
	r.GET("/imgcode", controller.HomeImgcode)
	v1 := r.Group("/v1")//版本
	{
		v1.POST("/signin", controller.HomeSignin)//登录
		v1.Use(middleware.UseJwt())//验证权限
		{
			v1.GET("/page", controller.HomePage)//首页的schema
			v1.GET("/pwdschema", controller.HomePwdSchema)//修改密码的schema
			v1.POST("/repwd", controller.HomeRepwd)//提交修改密码
			role := v1.Group("/role")//角色
			{
				role.GET("/schema", controller.RoleSchema)//角色的schema
				role.GET("/min", controller.RoleMin)
				role.POST("/data", controller.RoleAdd)
				role.GET("/data", controller.RoleData)
				role.PUT("/data/:id", controller.RoleEdit)
				role.DELETE("/data/:id", controller.RoleDelete)
				role.POST("/copy/:id", controller.RoleCopy)
			}
			admin := v1.Group("/admin")//账号
			{
				admin.GET("/schema", controller.AdminSchema)
				admin.POST("/data", controller.AdminAdd)
				admin.GET("/data", controller.AdminData)
				admin.PUT("/data/:id", controller.AdminEdit)
				admin.DELETE("/data/:id", controller.AdminDelete)
				admin.POST("/repwd/:id", controller.AdminRepwd)
			}
			test := v1.Group("/test")//测试数据
			{
				test.GET("/schema", controller.TestSchema)
				test.POST("/data", controller.TestAdd)
				test.GET("/data", controller.TestData)
				test.PUT("/data/:id", controller.TestEdit)
				test.DELETE("/data/:id", controller.TestDelete)
			}
		}
	}
	
	return r
}
