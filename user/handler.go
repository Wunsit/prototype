package user

import (
	"log"
	"net/http"
)

type Handler struct {
	s Storage
}

func NewHandler(s Storage) *Handler {
	return &Handler{s}
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("new request")
	switch r.Method {
	case http.MethodGet:
		u, err := h.s.GetByEmail(r.Context(), "mohamed@live.com")
		log.Println(u, err)
	case http.MethodPut:
	case http.MethodPost:
	default:
	}
}
