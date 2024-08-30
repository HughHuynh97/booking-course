package main

import (
	"github.com/HughHuynh97/booking-course/pkg/config"
	"github.com/HughHuynh97/booking-course/pkg/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func routes(config *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
