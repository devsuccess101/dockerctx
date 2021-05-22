package command

import (
	"fmt"
	"swarmd/internal/dockercontext"
)

type CurrentOp struct{}

func (op CurrentOp) Run() (err error) {
	cmd := dockercontext.DockerContext{}
	contextName, err := cmd.GetCurrentName()

	fmt.Println(contextName)

	return
}
