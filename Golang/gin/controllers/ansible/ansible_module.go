package ansible

import (
	"fmt"
	"strings"
)

const (
	moduleName = "- name: %s"
)

type Module struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type ModuleGenerate func(*Module)

func GetModuleGenerate(command string) ModuleGenerate {
	fields := strings.Fields(command)
	for _, field := range fields {
		switch field {
		case "systemctl":
			return generateSystemd
		}
	}
	return generateSystemd
}

func generateSystemd(m *Module) {
	moduleValues := []string{"  systemd:"}
	name := "    name: %s"
	stat := "    stat: %s"
	enabled := "    enabled: %s"
	fields := strings.Fields(m.Name)

	for _, field := range fields {
		switch field {
		case "systemctl":
		case "restart":
			moduleValues = append(moduleValues, fmt.Sprintf(stat, "restarted"))
		case "start":
			moduleValues = append(moduleValues, fmt.Sprintf(stat, "started"))
		case "stop":
			moduleValues = append(moduleValues, fmt.Sprintf(stat, "stoped"))
		case "reload":
			moduleValues = append(moduleValues, fmt.Sprintf(stat, "reloaded"))

		case "enable":
			moduleValues = append(moduleValues, fmt.Sprintf(enabled, "yes"))
		case "disable":
			moduleValues = append(moduleValues, fmt.Sprintf(enabled, "no"))

		default:
			moduleValues = append(moduleValues, fmt.Sprintf(name, field))
		}
	}
	m.Name = fmt.Sprintf(moduleName, m.Name)
	m.Values = moduleValues
}
