package main

import (
	. "github.com/RyotaNakajimakun/Training/Golang/gin/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("dist/*")

	r.GET("/", showHome)
	r.POST("/post", GenrateAnsibleModule)
	r.Run(":23450")
}
func showHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
