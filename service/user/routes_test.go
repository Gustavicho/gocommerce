package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gustavicho/gocommerce/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should faill if payload is invalid", func(t *testing.T) {
		payload := types.UserRegisterPayload{}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /api/v1/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("Should create user if payload is valid", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			FirstName: "Leonardo",
			LastName:  "Chaves",
			Email:     "gusta@example.com",
			Password:  "123456",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /api/v1/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("This user already exists")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
