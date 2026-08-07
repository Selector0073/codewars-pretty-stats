package router

import (
	"codewars-pretty-stats/internal/config"
	"codewars-pretty-stats/internal/service"
	"net/http"
)

func New(cfg *config.AppConfig) *http.ServeMux {
	var mux = http.NewServeMux()

	mux.HandleFunc("/", service.Svg(cfg.Axiom))

	return mux
}
