package command

import (
	"fmt"
	"swarmd/internal/cmdutil"
	"swarmd/internal/dockercontext"
)

type ListContextsOp struct{}

func (_op ListContextsOp) Run() error {
	docker := dockercontext.DockerContext{}
	currentCtx, currentCtxError := docker.GetCurrentName()
	contextList, listCtxError := docker.GetContextList()

	if currentCtxError != nil {
		return currentCtxError
	}

	if listCtxError != nil {
		return listCtxError
	}

	for _, contextName := range contextList {
		if contextName == currentCtx {
			cmdutil.ActiveItem(contextName + " *")
		} else {
			fmt.Println(contextName)
		}
	}

	return nil
}
