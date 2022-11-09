package data

import (
	"context"
	"database/sql"
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
