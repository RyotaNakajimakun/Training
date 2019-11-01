package main

import (
	"github.com/RyotaNakajimakun/Training/Golang/gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//r.StaticFS("/assets", assetFS())
	r.POST("/post", controllers.GenrateAnsibleModule)
	r.Run(":23450")
}