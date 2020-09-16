package main

import (
	"learning/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8989");
}
