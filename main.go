package main

import (
	myConfig "parterService/config"
	"parterService/routers"
)

func main() {
	myConfig.InitConfig();
	routers.InitRouter()
}
