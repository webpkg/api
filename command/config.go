// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package command

import (
	"errors"

	"github.com/webpkg/api/config"
	"github.com/webpkg/cmd"
)

var (
	cmdConfig = &cmd.Command{
		Run:       runConfig,
		UsageLine: "config",
		Short:     "create config file",
		Long:      `create config.json file at current directory.`,
	}
)

func runConfig(cmd *cmd.Command, args []string) error {

	if len(args) != 0 {
		return errors.New("too many arguments given")
	}

	return config.WriteConfig(Force())
}
