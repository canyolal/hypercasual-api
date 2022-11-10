package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Publisher struct {
	Id      int64  `json:"id"`
	Name    string `json:"name,omitempty"`
	Link    string `json:"link,omitempty"`
	Version int32  `json:"version"`
}

type PublisherModel struct {
	DB *sql.DB
}

func (m *PublisherModel) Insert(publisher *Publisher) error {
	query := `
	INSERT INTO publishers (name, link)
	VALUES ($1, $2)
	RETURNING id, version`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, publisher.Name, publisher.Link).Scan(&publisher.Id, &publisher.Version)
}

// Get a single publisher from DB
func (m *PublisherModel) Get(id int64) (*Publisher, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
	SELECT id, name, link, version
	FROM publishers
	WHERE id = $1`

	var publisher Publisher

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&publisher.Id,
		&publisher.Name,
		&publisher.Link,
		&publisher.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &publisher, nil
}
