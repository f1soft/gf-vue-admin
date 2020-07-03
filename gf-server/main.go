package main

import (
	"gf-server/boot"
	_ "gf-server/router"
)

func main() {
	boot.RunServer()
}
