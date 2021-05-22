package cmdutil

import (
	"os"

	"github.com/fatih/color"
)

func ExitError(err error) {
	color.Red("error: %s", err)
	os.Exit(1)
}

func ActiveItem(item string) {
	color.Yellow(item)
}
