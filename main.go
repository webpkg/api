package main

import (
	"flag"

	"github.com/webpkg/cmd"
	_ "github.com/webpkg/mysql"
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
