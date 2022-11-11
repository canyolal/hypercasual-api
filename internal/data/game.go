package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Game struct {
	Id           int64  `json:"id"`
	Name         string `json:"name,omitempty"`
	Genre        string `json:"genre,omitempty"`
	PubisherName string `json:"publisher_name,omitempty"`
	Version      int32  `json:"version"`
}

type GameModel struct {
	DB *sql.DB
}

// Insert a game into db
func (m *GameModel) Insert(publisherName, name, genre string) error {
	query := `
	INSERT INTO games (name, genre, publisher_name)
	VALUES ($1, $2 ,$3)
	RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var game Game

	return m.DB.QueryRowContext(ctx, query, name, genre, publisherName).Scan(&game.Id)
}

// Get all games from db
func (m *GameModel) GetAllWithFilters(name, genre string, filters Filters) ([]*Game, map[string]string, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, name, genre, publisher_name, version
		FROM games
		WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (to_tsvector('simple', genre) @@ plainto_tsquery('simple', $2) OR $2 = '')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`,
		filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, name, genre, filters.limit(), filters.offset())
	if err != nil {
		return nil, nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	games := []*Game{}
	gameList := make(map[string]string)

	for rows.Next() {
		var game Game

		err = rows.Scan(
			&totalRecords,
			&game.Id,
			&game.Name,
			&game.Genre,
			&game.PubisherName,
			&game.Version,
		)
		if err != nil {
			return nil, nil, Metadata{}, err
		}
		gameList[game.Name] = game.Genre
		games = append(games, &game)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return games, gameList, metadata, nil
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
			&game.PubisherName,
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
