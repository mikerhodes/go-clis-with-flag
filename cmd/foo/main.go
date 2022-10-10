package main

import (
	"fmt"
	"os"

	"github.ibm.com/mike-rhodes/goclisubcommands/internal/bar"
	"github.ibm.com/mike-rhodes/goclisubcommands/internal/baz"
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

func usage() string {
	s := "Usage: sl [cmd] [options]\nCommands:\n"
	for k := range commands {
		s += " - " + k + "\n"
	}
	return s
}
