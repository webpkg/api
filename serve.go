package main

import (
	"log"
	"net/http"

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

	cfg, err := readConfig()

	if err != nil {
		log.Fatalf("config: %v", err)
	}

	app := web.Create()

	repository.Init(cfg.Database)
	route.Init(app)

	log.Fatal(app.ListenAndServe(cfg.Server.Addr, func(srv *http.Server) {
		srv.ReadTimeout = cfg.Server.ReadTimeout
		srv.ReadHeaderTimeout = cfg.Server.ReadHeaderTimeout
		srv.WriteTimeout = cfg.Server.WriteTimeout
		srv.IdleTimeout = cfg.Server.IdleTimeout
	}))
}
