package product

import (
	"database/sql"
	"fmt"

	"github.com/fawzy17/test-enterkomputer/types"
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

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductById(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	product := new(types.Product)
	for rows.Next() {
		product, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if product.ID == 0 {
		return nil, fmt.Errorf("product id %d not found", id)
	}

	return product, nil
}

func (s *Store) GetPromo() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE (name = 'Jeruk' AND variant = 'Dingin') OR (name = 'Nasi' AND variant = 'Goreng')")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}



func scanRowIntoUser(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	var variant sql.NullString

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Category,
		&variant,
		&product.Price,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	if variant.Valid {
		product.Variant = variant.String
	} else {
		product.Variant = ""
	}

	return product, nil
}
