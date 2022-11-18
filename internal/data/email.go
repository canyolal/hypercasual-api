package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/canyolal/hypercasual-inventories/internal/validator"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type Email struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"-"`
}

type EmailModel struct {
	DB *sql.DB
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func (m *EmailModel) Insert(mail string) error {
	query := `
	INSERT INTO maillist (email)
	VALUES ($1)
	RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var email Email

	return m.DB.QueryRowContext(ctx, query, mail).Scan(&email.Id)
}

func (m *EmailModel) Delete(mail string) error {
	query := `
	DELETE FROM maillist
	WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, mail)
	return err
}
