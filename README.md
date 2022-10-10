# Use of `flag` in "noun action --args" CLI Go applications

This little codebase shows an approach of using the Go `flag` package to
implement command line applications that have "sub-commands" like `kubectl`.
I find that it's simpler to use for many applications than pulling in a more
feature-rich library.

We have a "dispatcher" in the `cmd/discourse` folder that gets built as the binary
entrypoint. It uses the first application argument to discover what sub-command
is being called.

Each sub-command is a package within `internal` and has its own entrypoint. That
entrypoint is hard-coded into the main `cmd/discourse/main.go` file and associated
with the command name. Each sub-command uses a `flag.FlagSet` of its own to
parse the arguments from `os.Args[2:]`, which are passed to the entrypoint
rather than having the sub-command use `os.Args` directly.

There are two packages with sub-commands. One thing they show is that if you
want the same argument in both packages, you define it twice, in each `FlagSet`.

Then you run like this after `make build`ing:

```sh
# We have a subcommand, named flag and two positional arguments
./discourse greet --quantity 3 Mike Rhodes
```
