package data

import (
	"context"
	"database/sql"
	"time"
)

type Game struct {
	Id          int64  `json:"id"`
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	PublisherId int64  `json:"publisher_id,omitempty"`
	Version     int32  `json:"version"`
}

type GameModel struct {
	DB *sql.DB
}

// Insert a game into db
func (m *GameModel) Insert(publisher, name, genre string) error {
	query := `
	INSERT INTO games (name, genre, publisher)
	VALUES ($1, $2 ,$3)
	RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var game Game

	return m.DB.QueryRowContext(ctx, query, name, genre, publisher).Scan(&game.Id)
}
