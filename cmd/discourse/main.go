// foo demonstrates how to create a CLI application with a relatively advanced
// invocation interface, with global options followed by subcommands and their
// options.
//
// > foo -debug            bar          -quantity 12   Mike Rhodes
//
//	[global options] [subcommand] [options]     [positional args]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mikerhodes/goclisubcommands/internal/farewell"
	"github.com/mikerhodes/goclisubcommands/internal/greet"
)

var commands = map[string]func([]string){
	"greet":    greet.Greet,
	"farewell": farewell.Farewell,
}

func main() {
	// Define and parse the global options
	debug := flag.Bool("debug", false, "print debug information")
	flag.Parse()

	// Pull the rest of the original arguments into a "subcommand line"
	subcommand := flag.Args()

	// Check we have a subcommand to run
	if len(subcommand) == 0 {
		printUsage()
		os.Exit(1)
	}
	cmd, ok := commands[subcommand[0]]
	if !ok {
		printUsage()
		os.Exit(1)
	}

	// Use our global option
	if *debug {
		log.Println("Debug info on")
	}

	// Call the subcommand with the remaining arguments
	cmd(subcommand[1:])
}

func printUsage() {
	fmt.Print("Usage: sl [global options] [command] [options]\n\n")
	fmt.Println("Available commands:")
	s := ""
	for k := range commands {
		s += " - " + k + "\n"
	}
	fmt.Println(s)
	fmt.Println("Global options:")
	flag.PrintDefaults()
}
