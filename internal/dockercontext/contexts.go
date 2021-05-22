package dockercontext

import (
	"os/exec"
	"strings"
)

type DockerContext struct{}

func (_cmd *DockerContext) GetCurrentName() (contextName string, err error) {
	args := []string{"context", "inspect", "-f", "{{.Name}}"}
	out, err := exec.Command("docker", args...).Output()
	contextName = strings.TrimRight(string(out), "\n")
	return
}

func (_cmd *DockerContext) GetContextList() (contextList []string, err error) {
	args := []string{"context", "ls", "-q"}
	out, err := exec.Command("docker", args...).Output()
	contextList = strings.Split(strings.Trim(string(out), "\n"), "\n")
	return
}

func (_cmd *DockerContext) SwitchContext(contextName string) (outputMsg string, err error) {
	args := []string{"context", "use", contextName}
	out, err := exec.Command("docker", args...).Output()
	outputMsg = strings.Trim(string(out), "\n")
	return
}
