package command

import (
	"fmt"
)

type VersionOp struct {
	Version string
}

func (op VersionOp) Run() error {
	fmt.Println(op.Version)

	return nil
}
