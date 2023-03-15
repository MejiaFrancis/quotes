// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Let's model the users table
type Quote struct {
	Quote_ID     int64
	Quote        string
	Quote_Author string
	CreatedAt    time.Time
}

// Setup dependency injection
type QuoteModel struct {
	DB *sql.DB
}

// Write SQL code to access the database
// TODO
// Creating a Get Method for Users table
func (m *QuoteModel) Get() (*Quote, error) {
	var q Quote

	statement := `
	            SELECT quote_id, quote, quote_author
				FROM quotes
				ORDER BY RANDOM()
				LIMIT 1
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.Quote_ID, &q.Quote, &q.Quote_Author)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *QuoteModel) Insert(quote string, quote_author string) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO quotes(quote, quote_author)
				VALUES($1, $2)
				RETURNING quote_id				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, quote, quote_author).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
