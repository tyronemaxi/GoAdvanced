package route

import (
	"GoAdvanced/api/handle"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {

	v1 := g.Group("v1")
	v1.POST("/login",Login)
	v1.GET("/user",handle.GetUserList)

	// grouping routes




}


