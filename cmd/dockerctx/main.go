/*
Copyright © 2021 Nguyen Huu Kim

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"fmt"
	"swarmd/cmd/dockerctx/command"
	"swarmd/internal/cmdutil"

	"github.com/spf13/cobra"
)

type DockerctxOptions struct {
	WantShowCurrentCtx *bool
	WantShowVersion    *bool
}

type Op interface {
	Run() error
}

type UnsupportedOp struct{ Err error }

func (op UnsupportedOp) Run() error {
	return op.Err
}

var (
	Version   = "v0.0.0+unkown"
	options   DockerctxOptions
	dockerctx = &cobra.Command{
		Use:   "dockerctx",
		Short: "Faster way to switch between Docker contexts.",
		Long:  "This tool helps you switch between Docker contexts back and forth.",
		Run:   runCommand,
	}
)

func main() {
	options.WantShowCurrentCtx = dockerctx.Flags().BoolP("current", "c", false, "show the current context name")
	options.WantShowVersion = dockerctx.Flags().BoolP("version", "v", false, "print version information and quit")

	dockerctx.SetHelpTemplate(`This tool helps you switch between Docker contexts back and forth.

USAGE:
  dockerctx               : list the contexts
  dockerctx <NAME>        : switch to context <NAME>
  dockerctx -             : switch to previous context
  dockerctx -c, --current : show the current context name
  dockerctx -v, --version : show version information

  dockerctx -h, --help    : show this message
`)

	cobra.CheckErr(dockerctx.Execute())
}

func runCommand(cmd *cobra.Command, args []string) {
	op := parseArgs(args, options)

	if err := op.Run(); err != nil {
		cmdutil.ExitError(err)
	}
}

func parseArgs(args []string, options DockerctxOptions) Op {
	switch {
	case *options.WantShowVersion:
		return command.VersionOp{Version: Version}
	case *options.WantShowCurrentCtx:
		return command.CurrentOp{}
	default:
		switch length := len(args); {
		case length == 0:
			return command.ListContextsOp{}
		case length == 1 && args[0] == "-":
			return command.SwapContextOp{}
		case length == 1 && args[0] != "-":
			return command.SetContextOp{NextContext: args[0]}
		}
	}

	return UnsupportedOp{Err: fmt.Errorf("too many arguments")}
}
