package main

import (
	"npos-demo/pkg/util"
	"npos-demo/routers"
)

func main() {
	util.Viper()
	r := routers.InitRouter()
	r.Run(":8080")
}
