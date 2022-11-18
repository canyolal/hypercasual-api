package data

import (
	"context"
	"database/sql"
	"time"
)

type Email struct {
	Id        int64     `json:"id"`
	Email     string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"-"`
}

type EmailModel struct {
	DB *sql.DB
}

func (m *EmailModel) Insert(mail string) error {
	query := `
	INSERT INTO maillist (mail)
	VALUES ($1)
	RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var email Email

	return m.DB.QueryRowContext(ctx, query, mail).Scan(&email.Id)
}
