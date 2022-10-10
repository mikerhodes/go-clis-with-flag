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
