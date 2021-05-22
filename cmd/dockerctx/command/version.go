package command

import (
	"fmt"
)

var (
	version = "v0.0.0+unkown"
)

type VersionOp struct{}

func (_op VersionOp) Run() error {
	fmt.Println(version)

	return nil
}
