package handler

import (
	"fmt"
	"net/http"
)

type PublicHandler struct{}

func (h *PublicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		fmt.Fprintln(w, "Public GET / response")
		return
	}
	http.Error(w, "Forbidden", http.StatusForbidden)
}
