package views

import (
	"net/http"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/go-chi/chi"
	"github.com/lpar/gzipped"
	"github.com/spf13/viper"
)

func RegisterEndPoints(mux *chi.Mux) {
	mux.Get("/today", TaskyHandler)
	mux.Get("/upcoming", TaskyHandler)
	mux.Get("/long_term", TaskyHandler)
	mux.Get("/incomplete", TaskyHandler)
	mux.Get("/completed", TaskyHandler)
	mux.Get("/thought_cloud", TaskyHandler)
}

func TaskyHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/tasky.html"
	webAssetsDirectory := http.Dir(viper.GetString(config.WebAssetsPathKey))
	gzipped.FileServer(webAssetsDirectory).ServeHTTP(w, r)
}
