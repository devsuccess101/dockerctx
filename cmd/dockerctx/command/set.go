package command

import (
	"fmt"
	"swarmd/cmd/dockerctx/helpers"
	"swarmd/internal/dockercontext"
)

type SetContextOp struct {
	NextContext string
}

func (op SetContextOp) Run() error {
	docker := dockercontext.DockerContext{}
	prevContext, err := docker.GetCurrentName()

	if err != nil {
		return nil
	}

	activedContext, err := docker.SwitchContext(op.NextContext)

	if err != nil {
		return fmt.Errorf("context \"%s\" does not exist", op.NextContext)
	}

	fmt.Println(activedContext)
	fmt.Printf("current context is now \"%s\"\n", activedContext)

	if prevContext != op.NextContext {
		helpers.SaveContext(prevContext)
	}

	return nil
}
