package boot

import (
	"gf-server/library/global"
	"gf-server/router"
	"time"

	"github.com/gogf/gf/frame/g"
)

// init 隐式初始化
//func init() {
//	global.GFVA_SERVER = g.Server()
//	router.InitializeRouters()
//	global.GFVA_SERVER.SetReadTimeout(10 * time.Second)
//	global.GFVA_SERVER.SetWriteTimeout(10 * time.Second)
//	global.GFVA_SERVER.SetMaxHeaderBytes(1 << 20)
//	global.GFVA_SERVER.Run()
//}

func InitializeRunServer() {
	global.GFVA_SERVER = g.Server()
	router.InitializeRouters()
	global.GFVA_SERVER.SetReadTimeout(10 * time.Second)
	global.GFVA_SERVER.SetWriteTimeout(10 * time.Second)
	global.GFVA_SERVER.SetMaxHeaderBytes(1 << 20)
	global.GFVA_SERVER.Run()
}
