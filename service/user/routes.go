package user

import (
	"fmt"
	"net/http"

	"github.com/Gustavicho/gocommerce/service/auth"
	"github.com/Gustavicho/gocommerce/types"
	"github.com/Gustavicho/gocommerce/utils"
)

type Handler struct {
	store  types.UserStore
	prefix string
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
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
	// get json payload
	var payload types.UserRegisterPayload
	utils.ParseJSON(r, payload)

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"user with email '%s' already exists", payload.Email,
		))
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, nil)
	}

	// if doesnt, create user
	err = h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
