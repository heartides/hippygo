package main

import (
	_"learning/conf"
	"learning/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":3030")
}
