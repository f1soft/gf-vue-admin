package boot

import (
	"gf-server/global"

	"github.com/gogf/gf/frame/g"
)

// init
//func init() {
//	db := g.DB("default")
//	global.GFVA_DB = db
//}

func InitializeMysql() {
	db := g.DB("default")
	global.GFVA_DB = db
}
