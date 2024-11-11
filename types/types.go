package types

import "time"

// /////[USER]///////
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRegisterPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"  validate:"required"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=4"`
}

type UserLoginPayload struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// /////[PRODUCT]///////
type ProductStore interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(id int) (*Product, error)
	CreateProduct(Product) error
}

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductCreatePayload struct {
	Name        string  `json:"name"        validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image"       validate:"required"`
	Price       float64 `json:"price"       validate:"required"`
	Quantity    int     `json:"quantity"    validate:"required"`
}
