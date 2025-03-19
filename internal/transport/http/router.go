package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DmitriyGoryntsev/marketplace/internal/config"
	"github.com/labstack/echo/v4"
)

type RouterConfig struct {
	Host string
	Port string
}

type Router struct {
	config RouterConfig
	router *echo.Echo

	Handler *Handler
}

func NewRouterConfig(cfg *config.Config) RouterConfig {
	return RouterConfig{
		Host: cfg.HTTPServer.Host,
		Port: cfg.HTTPServer.Port,
	}
}

func NewRouter(RConfig RouterConfig, handler *Handler) *Router {
	r := echo.New()

	return &Router{
		config:  RConfig,
		router:  r,
		Handler: handler,
	}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", r.config.Host, r.config.Port), r.router))
}

func (r *Router) ShuttingDown() {
	r.router.Shutdown(context.Background())
}
