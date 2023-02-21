// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package command

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gostartkit/api/config"
	"github.com/gostartkit/api/helper"
	"github.com/gostartkit/api/rbac"
	"github.com/gostartkit/api/repository"
	"github.com/gostartkit/api/route"
	"github.com/webpkg/cmd"
	"github.com/webpkg/web"
)

var (
	cmdServe = &cmd.Command{
		Run:       runServe,
		UsageLine: "serve",
		Short:     "start web service",
		Long:      "start web service.\n",
	}
)

func runServe(cmd *cmd.Command, args []string) error {

	if err := config.Init(); err != nil {
		return fmt.Errorf("config.Init: %v", err)
	}

	if err := repository.Init(); err != nil {
		return fmt.Errorf("repository.Init: %v", err)
	}

	defer repository.Close()

	app := web.CreateApplication()

	rbac.Init()
	route.Init(app)

	network := config.Server().Network

	if network == "" {
		network = "tcp"
	}

	addr := config.Server().Addr

	if addr == "" {
		addr = "127.0.0.1:8443"
	}

	appName := config.App().AppName

	if appName == "" {
		appName = "auth.go"
	}

	// clean sock file if exists
	if strings.ToLower(network) == "unix" && helper.FileExist(addr) {
		if err := os.Remove(addr); err != nil {
			return fmt.Errorf("sock: %v", err)
		}
	}

	log.Printf("%s(%d) %s:%s\n", appName, os.Getpid(), network, addr)

	err := app.ListenAndServe(network, addr, func(srv *http.Server) {
		srv.ReadTimeout = config.Server().ReadTimeout * time.Second
		srv.ReadHeaderTimeout = config.Server().ReadHeaderTimeout * time.Second
		srv.WriteTimeout = config.Server().WriteTimeout * time.Second
		srv.IdleTimeout = config.Server().IdleTimeout * time.Second
	})

	if err != nil {
		return fmt.Errorf("serve: %v", err)
	}

	return nil
}
