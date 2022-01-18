package main

import (
	"github.com/labstack/echo/v4"
	"notebook/notebook"
	"notebook/server"
	"notebook/templates"
)

func main() {
	if err := notebook.Migrate(); err != nil {
		panic(err.Error())
	}

	cfg, err := server.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.Renderer = templates.NewRenderer()
	s := server.New(e, cfg)

	notebook.RegisterRoutes(e)
	err = e.StartServer(s.Server())
	if err != nil {
		panic(err.Error())
	}
}
