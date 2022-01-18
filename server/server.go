package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Http struct {
	echo   *echo.Echo
	server http.Server
}

func New(echo *echo.Echo, config *Config) Http {
	return Http{
		echo: echo,
		server: http.Server{
			Addr:         config.Address,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
		},
	}
}

func (h *Http) Server() *http.Server {
	return &http.Server{
		Addr:         h.server.Addr,
		ReadTimeout:  h.server.ReadTimeout,
		WriteTimeout: h.server.WriteTimeout,
	}
}
