package main

import (
	"github.com/webpkg/cmd"
	_ "github.com/webpkg/mysql"
)

func main() {
	cmd.AddCommands(cmdConfig, cmdServe, cmdVersion)
	cmd.Execute()
}
