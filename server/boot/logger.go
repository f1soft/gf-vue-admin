package boot

import (
	"gf-server/library/global"

	"github.com/gogf/gf/frame/g"
)

func InitializeLogger() {
	logger := g.Log()
	global.GFVA_LOG = logger
}
