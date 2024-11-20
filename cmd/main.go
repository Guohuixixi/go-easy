package main

import (
	"fmt"
	"github.com/Guohuixixi/go-easy/bootstrap/wire"
)

func main() {
	app, err := wire.InitApp()
	if err != nil {
		fmt.Println("初始化模块失败")
		panic(err)
	}
	app.Server.Run(fmt.Sprintf("%s:%d", app.Config.ServerConfig.Host, app.Config.ServerConfig.Port))
}
