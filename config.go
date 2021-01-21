package main

import (
	"log"

	"github.com/webpkg/api/config"
	"github.com/webpkg/cmd"
)

var (
	cmdConfig = &cmd.Command{
		Run:       runConfig,
		UsageLine: "config",
		Short:     "create config file",
		Long:      "create config.json file at current directory.\n",
	}
)

func runConfig(cmd *cmd.Command, args []string) {

	if len(args) != 0 {
		log.Fatal("Too many arguments given.\n")
	}

	config.WriteConfig()
}
