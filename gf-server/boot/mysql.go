package boot

import (
	"gf-server/library/global"

	"github.com/gogf/gf/frame/g"
)

// init 隐式初始化Mysql
//func init() {
//	db := g.DB("default")
//	global.GFVA_DB = db
//}

func InitializeMysql() {
	db := g.DB("default")
	global.GFVA_DB = db
}
