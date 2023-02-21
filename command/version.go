// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package command

import (
	"fmt"

	"github.com/webpkg/cmd"
)

var (
	version = "v0.0.1"

	cmdVersion = &cmd.Command{
		Run:       runVersion,
		UsageLine: "version",
		Short:     "display version",
		Long:      "display version and build info.\n",
	}
)

func runVersion(cmd *cmd.Command, args []string) error {
	fmt.Println(version)
	return nil
}
