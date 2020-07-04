package boot

import (
	"fmt"
	"gf-server/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// init 隐式初始化Redis
//func init() {
//	if g.Cfg().GetBool("system.UseMultipoint"){
//		global.GFVA_REDIS = g.Redis()
//	}
//	ping()
//}

func InitializeRedis() {
	if g.Cfg("system").GetBool("system.UseMultipoint") {
		global.GFVA_REDIS = g.Redis()
	}
	ping()
}

func ping() {
	conn, err := global.GFVA_REDIS.Do("PING")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gconv.String(conn))
}
