package main

import (
	"gf-server/boot"
)

func main() {
	boot.InitializeLogger()    // 初始化日志管理
	boot.InitializeRedis()     // 初始化Redis连接, 如果use_multipoint为false,则不会初始化redis的服务
	boot.InitializeMysql()     // 初始化Mysql连接
	boot.InitializeRunServer() // 初始化gf服务器
}
