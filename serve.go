package main

import (
	"log"
	"net/http"
	"time"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/rbac"
	"github.com/webpkg/api/repository"
	"github.com/webpkg/api/route"
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

func runServe(cmd *cmd.Command, args []string) {

	cfg, err := config.ReadConfig()

	if err != nil {
		log.Fatalf("config: %v", err)
	}

	repository.Init(cfg)

	app := web.Create()

	rbac.Init(cfg.Rbac)
	route.Init(app)

	if cfg.App.PublicDir != "" {
		app.NotFound = http.FileServer(http.Dir(cfg.App.PublicDir))
	}

	err = app.ListenAndServe(cfg.Server.Addr, func(srv *http.Server) {
		srv.ReadTimeout = cfg.Server.ReadTimeout * time.Second
		srv.ReadHeaderTimeout = cfg.Server.ReadHeaderTimeout * time.Second
		srv.WriteTimeout = cfg.Server.WriteTimeout * time.Second
		srv.IdleTimeout = cfg.Server.IdleTimeout * time.Second
	})

	repository.Close()

	if err != nil {
		log.Printf("serve: %v", err)
	}
}
