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
	v.Check(email != "", "email", "email must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "email must be a valid email address")
}

func (m *EmailModel) Insert(mail string) error {
	query := `
	INSERT INTO maillist (email)
	VALUES ($1)
	RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var email Email

	err := m.DB.QueryRowContext(ctx, query, mail).Scan(&email.Id)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "maillist_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
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

func (m *EmailModel) Exists(id int64) (bool, error) {
	var exists bool

	query := `
	SELECT EXISTS(SELECT true FROM maillist WHERE id = $1)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&exists)
	return exists, err
}
