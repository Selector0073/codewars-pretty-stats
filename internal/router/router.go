package router

import (
	"codewars-pretty-stats/internal/service"
	"net/http"
)

func New() *http.ServeMux {
	var mux = http.NewServeMux()

	mux.HandleFunc("/", service.Svg)

	return mux
}
