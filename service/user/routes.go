package user

import "net/http"

type Handler struct {
	prefix string
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AddPrefix(pf string) {
	h.prefix = pf
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST "+h.prefix+"/login", h.handleLogin)
	r.HandleFunc("POST "+h.prefix+"/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// do stuff
}