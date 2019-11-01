package controllers

import (
	"github.com/RyotaNakajimakun/Training/Golang/gin/controllers/ansible"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenrateAnsibleModule(c *gin.Context) {
	command := struct{ Value string `json:"value"` }{}
	c.Bind(&command)

	module := ansible.Module{Name: command.Value}

	generator := ansible.GetModuleGenerate(command.Value)
	generator(&module)

	c.JSON(http.StatusOK, module)
}
