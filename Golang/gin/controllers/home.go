package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HomeGet handles GET / route
func HomeGet(c *gin.Context) {
	c.HTML(http.StatusOK, "home/show", gin.H{})
}
