package global

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

var (
	GFVA_DB     gdb.DB
	GFVA_LOG    *glog.Logger
	GFVA_SERVER *ghttp.Server
	GFVA_REDIS  *gredis.Redis
)
