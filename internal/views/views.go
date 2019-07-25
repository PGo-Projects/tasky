package views

import (
	"github.com/go-chi/chi"
)

func RegisterEndPoints(mux *chi.Mux) {
	RegisterTaskyEndpoints(mux)
}
