package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type userInfo struct {
	username string
	passwd string

}

func Login(c *gin.Context) {
	fmt.Println("loging....")
}
