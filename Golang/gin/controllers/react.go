package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReactHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
