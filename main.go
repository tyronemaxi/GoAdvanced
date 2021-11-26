package main

import (
	"GoAdvanced/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.InitRouter(r)


	r.Run(":8080")
}