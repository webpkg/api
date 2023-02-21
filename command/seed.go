// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package command

import (
	"errors"
	"fmt"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/repository"
	"github.com/webpkg/cmd"
)

var (
	cmdSeed = &cmd.Command{
		Run:       runSeed,
		UsageLine: "seed",
		Short:     "seed the database with records",
		Long:      `seed the database with records.`,
	}
)

func runSeed(cmd *cmd.Command, args []string) error {

	if len(args) != 0 {
		return errors.New("too many arguments given")
	}

	if err := config.Init(); err != nil {
		return fmt.Errorf("config.Init: %v", err)
	}

	if err := repository.Init(); err != nil {
		return fmt.Errorf("repository.Init: %v", err)
	}

	defer repository.Close()

	// -- code here

	return nil
}
