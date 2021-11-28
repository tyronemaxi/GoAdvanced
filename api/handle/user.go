package handle

import (
	"GoAdvanced/database"
	"GoAdvanced/route"
	"github.com/gin-gonic/gin"
	"net/http"
)



func GetUserList(c *gin.Context)  {
	user, err := route.Welcome(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token 认证失败",
	})
		return
	}

	users := database.User{}
	response := users.SelectAll()

	c.JSON(http.StatusOK,gin.H{
		"message": "登陆成功",
		"user": user,
		"data":response,
	})
	return
}
