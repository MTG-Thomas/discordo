# Contributing

Thanks for helping improve Discordo.

This repository is a fork of `ayn2op/discordo`. For fixes that are generally useful to Discordo users, prefer opening a pull request upstream first:

https://github.com/ayn2op/discordo

Fork-specific changes should stay narrow and be documented in their pull request description so they can be separated from upstreamable work later.

## Development

Run the standard Go checks before opening a pull request:

```sh
go test ./...
go build -trimpath -ldflags=-s .
```

On Linux, install `libx11-dev` before running the full test/build workflow if clipboard support is enabled by the build.

## Licensing

Discordo is licensed under GPL-3.0. Contributions to this repository are made under the same license.
