package server

import (
	"fmt"
	"net/http"
)

func loginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v", r.Context())
	})
}
