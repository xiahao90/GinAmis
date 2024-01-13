package config
//调试模式 关闭为0
const DEBUG = 0

//系统名称
const SYSTEM_NAME = "GinAmis"

/*----------- mysql数据库 ------------*/
// DB_ADD 数据库地址
const DB_ADD = "127.0.0.1"
// DB_PORT 数据库端口
const DB_PORT = "3306"
// DB_NAME 数据库名字
const DB_NAME = "ginamis"
// DB_USER 数据库用户
const DB_USER = "root"
// DB_PWD 数据库密码
const DB_PWD = "php123456"

// 密钥（数据库加密盐）
const SALT = "3f4b8f3eaa728d3af669e749af7eae67"

// JWT密钥，各位老6记得改一下这个哦，不然会被利用
const JWT_SK = "n0zuwj849msb2rdet3oivfxlc1qp5ag6"
const JWT_EXTIME = 168//jwt过期时间，小时
