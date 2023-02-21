// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package command

import (
	"flag"

	"github.com/webpkg/cmd"
)

var (
	_force bool
)

// Run exec commands
func Run() {
	cmd.SetFlags(func(f *flag.FlagSet) {
		f.BoolVar(&_force, "force", false, "force without warn")
		f.BoolVar(&_force, "f", false, "force without warn")
	})
	cmd.AddCommands(cmdConfig, cmdSeed, cmdServe, cmdVersion)
	cmd.Execute()
}

func Force() bool {
	return _force
}
