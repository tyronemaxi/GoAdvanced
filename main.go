package main

import (
	"GoAdvanced/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	route.InitRouter(r)

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_
	r.Use(gin.Logger())

	//
	r.Use(gin.Recovery())


	r.Run()
}

