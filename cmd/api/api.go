package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Gustavicho/gocommerce/service/product"
	"github.com/Gustavicho/gocommerce/service/user"
)

type APIService struct {
	Addr string
	DB   *sql.DB
}

func NewAPIService(addr string, db *sql.DB) *APIService {
	return &APIService{
		Addr: addr,
		DB:   db,
	}
}

func (as *APIService) Run() error {
	fmt.Printf("Starting server. Listening to port %s\n", as.Addr)

	router := http.NewServeMux()

	userStore := user.NewStore(as.DB)
	userHandler := user.NewHandler(userStore)
	userHandler.AddPrefix("/api/v1")
	userHandler.RegisterRoutes(router)

	productStore := product.NewStore(as.DB)
	productHandler := product.NewHandler(productStore)
	productHandler.AddPrefix("/api/v1")
	productHandler.RegisterRoutes(router)

	return http.ListenAndServe(as.Addr, router)
}
