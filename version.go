package main

import (
	"fmt"

	"github.com/webpkg/cmd"
)

var (
	_version = "v0.0.1"
	_osarch  string // set by ldflags

	cmdVersion = &cmd.Command{
		Run:       runVersion,
		UsageLine: "version",
		Short:     "display version",
		Long:      "display version and build info.\n",
	}
)

func runVersion(cmd *cmd.Command, args []string) {
	fmt.Println(_version, _osarch)
}
