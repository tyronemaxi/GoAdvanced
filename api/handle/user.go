package handle

import (
	"GoAdvanced/database"
	"github.com/gin-gonic/gin"
	"net/http"
)



func GetUserList(c *gin.Context)  {

	//requestBody := c.Request.Body

	user := database.User{}
	response := user.SelectAll()

	c.JSON(http.StatusOK,gin.H{
		"message": "登陆成功",
		"data":response,
	})
	return
}
