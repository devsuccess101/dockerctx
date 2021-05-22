package command

import (
	"fmt"
	"swarmd/cmd/dockerctx/helpers"
)

type SwapContextOp struct{}

func (_op SwapContextOp) Run() error {
	nextContext, err := helpers.ReadContext()

	if nextContext == "" || err != nil {
		return fmt.Errorf("no previous context found")
	}

	setCmd := SetContextOp{NextContext: nextContext}
	err = setCmd.Run()

	return err
}
