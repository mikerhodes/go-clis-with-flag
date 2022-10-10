package bar

import (
	"flag"
	"log"
)

func Bar(args []string) {
	log.Printf("bar got args: %v", args)

	fs := flag.NewFlagSet("baz", flag.ContinueOnError)
	quantity := fs.Int("quantity", 0, "number of baz")
	fs.Parse(args)

	log.Printf("here are %d bazzes", *quantity)
}
