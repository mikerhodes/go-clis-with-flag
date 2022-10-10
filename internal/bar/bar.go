package bar

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func Bar(args []string) {
	log.Printf("bar got args: %v", args)

	fs := flag.NewFlagSet("baz", flag.ExitOnError)

	// Here we want to output information about the positional arguments, so we
	// alter the output of the Usage function. Thankfully, PrintDefaults() gives
	// us the standard output.
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: %s [OPTIONS] FIRST_NAME LAST_NAME\n", os.Args[0])
		fs.PrintDefaults()
	}

	quantity := fs.Int("quantity", 0, "number of baz")
	fs.Parse(args)

	// we expect two positional arguments
	positionalArgs := fs.Args()
	if len(positionalArgs) != 2 {
		fs.Usage()
		os.Exit(1)
	}

	log.Printf("here are %d bazzes, %s %s", *quantity, positionalArgs[0], positionalArgs[1])
}
