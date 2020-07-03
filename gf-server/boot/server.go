package boot

import (
	"gf-server/global"
	"gf-server/router"
	"github.com/gogf/gf/frame/g"
	"time"
)
func RunServer() {
	global.GFVA_SERVER = g.Server()
	router.InitializeRouters()
	global.GFVA_SERVER.SetReadTimeout(10 * time.Second)
	global.GFVA_SERVER.SetWriteTimeout(10 * time.Second)
	global.GFVA_SERVER.SetMaxHeaderBytes(1 << 20)
	global.GFVA_SERVER.Run()
}
