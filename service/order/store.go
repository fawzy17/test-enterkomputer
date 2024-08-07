package order

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fawzy17/test-enterkomputer/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order []types.Order, meja string) error {
	if len(order) == 0 {
		return fmt.Errorf("no order products provided")
	}

	query := "INSERT INTO orders (orderId, productId, quantity, totalPrice, meja) VALUES "
	values := make([]interface{}, 0, len(order)*4)
	placeholders := make([]string, 0, len(order))

	for _, op := range order {
		placeholders = append(placeholders, "(?, ?, ?, ?, ?)")
		values = append(values, op.OrderId, op.ProductId, op.Quantity, op.TotalPrice, meja)
	}

	query += strings.Join(placeholders, ", ")

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (s *Store) GetBill(orderId string) ([]types.BillResponse, error) {
	rows, err := s.db.Query(`
		SELECT  
			orders.quantity, 
			orders.totalPrice, 
			products.name, 
			products.variant, 
			products.price 
		FROM 
			orders 
		JOIN 
			products 
		ON 
			orders.productId = products.id
		WHERE
			orders.orderId = ?
	`, orderId)

	if err != nil {
		return nil, err
	}

	bill := make([]types.BillResponse, 0)
	for rows.Next() {
		b, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		bill = append(bill, *b)
	}

	return bill, nil

}

func (s *Store) GetMeja(orderId string) (*string, error) {
	row := s.db.QueryRow("SELECT meja FROM orders WHERE orderId = ? LIMIT 1", orderId)
	
	var meja string
	
	err := row.Scan(&meja)
	if err != nil {
		return nil, err
	}
	
	return &meja, nil
}



func scanRowIntoUser(rows *sql.Rows) (*types.BillResponse, error) {
	bill := new(types.BillResponse)
	var totalPrice int
	var variant sql.NullString

	err := rows.Scan(
		&bill.Quantity,
		&totalPrice,
		&bill.Name,
		&variant,
		&bill.Price,
	)

	if err != nil {
		return nil, err
	}

	bill.TotalPrice = bill.Quantity * bill.Price
	if variant.Valid {
		bill.Variant = variant.String
	} else {
		bill.Variant = ""
	}

	return bill, nil
}
