package main

import (
	"fmt"

	"github.com/webpkg/cmd"
)

var (
	version    = "v0.0.1"
	osarch     string // set by ldflags
	gitVersion string // set by ldflags
	buildDate  string // set by ldflags

	cmdVersion = &cmd.Command{
		Run:       runVersion,
		UsageLine: "version",
		Short:     "display version",
		Long:      "display version and build info.\n",
	}
)

func runVersion(cmd *cmd.Command, args []string) {
	fmt.Println(version, osarch, gitVersion, buildDate)
}
