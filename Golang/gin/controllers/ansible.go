package controllers

import (
	"github.com/RyotaNakajimakun/Training/Golang/gin/controllers/ansible"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"net/http"
)

func GenrateAnsibleModule(c *gin.Context) {
	command := struct{ Value string `json:"value"` }{}
	c.Bind(&command)
	pp.Print(command)

	module := ansible.Module{}
	module.Values = append(module.Values, command.Value)

	generator := ansible.GetModuleGenerate(command.Value)
	generator(&module)

	pp.Print(module)
	c.JSON(http.StatusOK, module)
}
