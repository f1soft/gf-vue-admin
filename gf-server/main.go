package main

import (
	_ "gf-server/boot"
	_ "gf-server/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
