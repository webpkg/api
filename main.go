package main

import (
	"flag"

	"github.com/webpkg/cmd"
)

var (
	_webForce = false
)

func main() {

	cmd.SetFlags(func(f *flag.FlagSet) {
		f.BoolVar(&_webForce, "force", false, "")
	})

	cmd.AddCommands(cmdConfig, cmdServe, cmdVersion)
	cmd.Execute()
}
