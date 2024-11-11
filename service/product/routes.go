package product

import (
	"fmt"
	"net/http"

	"github.com/Gustavicho/gocommerce/types"
	"github.com/Gustavicho/gocommerce/utils"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	prefix string
	store  types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) AddPrefix(pf string) {
	h.prefix = pf
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET "+h.prefix+"/products", h.handleGetAll)
	r.HandleFunc("POST "+h.prefix+"/products", h.handleCreate)
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var payload types.ProductCreatePayload

	// parse json
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// create product
	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
