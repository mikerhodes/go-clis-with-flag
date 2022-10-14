package farewell

import (
	"flag"
	"fmt"
	"log"
)

func Farewell(args []string) {
	log.Printf("Farewell got args: %v", args)

	fs := flag.NewFlagSet("Farewell", flag.ExitOnError)
	name := fs.String("name", "", "who to baz")
	fs.StringVar(name, "n", "", "who to baz")
	quantity := fs.Int("quantity", 1, "number of baz")
	fs.IntVar(quantity, "q", 1, "number of baz")
	fs.Parse(args)

	for i := 0; i <= *quantity; i++ {
		fmt.Printf("Until we meet again, %s, farewell!\n", *name)
	}
}
