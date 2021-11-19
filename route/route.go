package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter(g *gin.Engine) {
	// group 1
	g1 := g.Group("g1")

	g1 .GET("/hello",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello g1",
		})
	})

	// parameters in path
	g.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fmt.Println(message)
		c.String(http.StatusOK, message)
	})

	// querystring parameters
	g.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	// Multipart/urlencoded Form
	g.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	// Map as querystring or postform parameters
	g.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names:%v",ids,names)
	})

	// upload file
	g.MaxMultipartMemory = 8 << 20
	g.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := fmt.Sprintf("/tmp")
		// upload the file to specific dst
		c.SaveUploadedFile(file,dst)
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!",len(file)))
	})

	// grouping routes
	v1 := g.Group("v1")


	v1.POST("/login", Login)


}


