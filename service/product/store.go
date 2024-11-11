package product

import (
	"database/sql"
	"fmt"

	"github.com/Gustavicho/gocommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAllProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	var allP []types.Product
	for rows.Next() {
		err = scanRowsIntoProduct(rows, p)
		if err != nil {
			return nil, err
		}

		allP = append(allP, *p)
	}

	if len(allP) == 0 {
		return nil, fmt.Errorf("no products found")
	}

	return allP, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		err = scanRowsIntoProduct(rows, p)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return p, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		p.Name,
		p.Description,
		p.Image,
		p.Price,
		p.Quantity,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows, p *types.Product) error {
	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&p.Quantity,
		&p.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
