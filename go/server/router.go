package server

import (
	"github.com/gorilla/mux"
)

// Router returns a Mux configuration based on an array of customRoute
func Router() (r *mux.Router) {
	r = mux.NewRouter()

	r.Handle("/login", loginHandler())

	return
}
