package handlers

import (
	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// uppercase means that it can be imported
func Handler(r *chi.Mux) {
	// global middleware

	// ignore all slashes at the end
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
