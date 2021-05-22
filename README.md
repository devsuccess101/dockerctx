# swarmd

This repository provides `dockerctx` tool.
- `dockerctx` helps you switch between Docker contexts back and forth. Inspired by [`kubectx`](https://github.com/ahmetb/kubectx).

## dockerctx

Please download the `dockerctx` binary in [the latest release](https://github.com/kimyvgy/swarmd/releases).

```bash
USAGE:
  dockerctx               : list the contexts
  dockerctx <NAME>        : switch to context <NAME>
  dockerctx -             : switch to previous context
  dockerctx -c, --current : show the current context name
  dockerctx -v, --version : show version information

  dockerctx -h, --help    : show this message
```
