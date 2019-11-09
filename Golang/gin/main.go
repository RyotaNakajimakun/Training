package main

import (
	. "github.com/RyotaNakajimakun/Training/Golang/gin/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const addr = ":23450"
func main() {
	r := gin.Default()
	{
		r.Use(static.Serve("/", static.LocalFile("./dist", true)))
		r.Static("/dist", "./dist")
		r.LoadHTMLGlob("dist/*.html")
	}

	r.GET("/", ReactHome)

	api := r.Group("/api")
	{
		api.POST("/post", GenrateAnsibleModule)
		api.GET("/post", GenrateAnsibleModule)
	}

	r.Run(addr)
}

func Test(c *gin.Context) {
	c.JSON(200, struct {
		f string
	}{f: "test"})
}
