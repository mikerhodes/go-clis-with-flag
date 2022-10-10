package greet

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func Greet(args []string) {
	log.Printf("Greet got args: %v", args)

	fs := flag.NewFlagSet("Greet", flag.ExitOnError)

	// Alter the FlagSet's Usage function to print details of the positional
	// arguments.
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: %s [OPTIONS] FIRST_NAME LAST_NAME\n", os.Args[0])
		fs.PrintDefaults()
	}

	quantity := fs.Int("quantity", 0, "number of greetings")
	fs.Parse(args)

	// we expect two positional arguments
	positionalArgs := fs.Args()
	if len(positionalArgs) != 2 {
		fs.Usage()
		os.Exit(1)
	}

	fmt.Printf("I greet you %d times, %s %s!\n", *quantity, positionalArgs[0], positionalArgs[1])
}
