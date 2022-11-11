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
func (m *GameModel) Insert(publisherID int64, name, genre string) error {
	query := `
	INSERT INTO games (name, genre, publisher_id)
	VALUES ($1, $2 ,$3)
	RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var game Game

	return m.DB.QueryRowContext(ctx, query, name, genre, publisherID).Scan(&game.Id)
}

// Get all games from db
func (m *GameModel) GetAll() ([]*Game, map[string]string, error) {
	query := `SELECT * FROM games`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	games := []*Game{}
	gameList := make(map[string]string)

	for rows.Next() {
		var game Game

		err = rows.Scan(
			&game.Id,
			&game.Name,
			&game.Genre,
			&game.PublisherId,
			&game.Version,
		)
		if err != nil {
			return nil, nil, err
		}
		gameList[game.Name] = game.Genre
		games = append(games, &game)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}
	return games, gameList, nil
}
