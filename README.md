# Use of flags in "noun action --args" CLI applications

This little codebase shows an approach of using the Go `flag` package to
implement command line applications that have "sub-commands" like `kubectl`.

We have a "dispatcher" in the `cmd/foo` folder that gets built as the binary
entrypoint. It uses the first application argument to discover what sub-command
is being called.

Each sub-command is a package within `internal` and has its own entrypoint. That
entrypoint is hard-coded into the main `cmd/foo/main.go` file and associated
with the command name. Each sub-command uses a `flag.FlagSet` of its own to
parse the arguments from `os.Args[2:]`, which are passed to the entrypoint
rather than having the sub-command use `os.Args` directly.

There are two packages with sub-commands. One thing they show is that if you
want the same argument in both packages, you define it twice, in each `FlagSet`.

Then you run like this after `make build`ing:

```
./foo bar --name mike --quantity 12

./foo baz --quantity 12
```

So the dispatcher looks like this:

```go
package main

import (
	"fmt"
	"os"

	"github.ibm.com/mike-rhodes/goflags/internal/bar"
	"github.ibm.com/mike-rhodes/goflags/internal/baz"
)

var commands = map[string]func([]string){
	"bar": bar.Bar,
	"baz": baz.Baz,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage())
		os.Exit(1)
	}
	cmd, ok := commands[os.Args[1]]
	if !ok {
		fmt.Println(usage())
		os.Exit(1)
	}
	cmd(os.Args[2:])
}
```

And the sub-commands look like this:

```go
package baz

import (
	"flag"
	"log"
)

func Baz(args []string) {
	log.Printf("baz got args: %v", args)

	fs := flag.NewFlagSet("baz", flag.ExitOnError)
	name := fs.String("name", "", "who to baz")
	quantity := fs.Int("quantity", 1, "number of baz")
	fs.Parse(args)

	for i := 0; i <= *quantity; i++ {
		log.Printf("hey %s! you're bazzed!", *name)
	}
}
```
