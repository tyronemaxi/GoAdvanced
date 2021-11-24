package route

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {

	g.POST("/login",Login)


	// grouping routes




}


