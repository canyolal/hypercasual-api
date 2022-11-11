package main

import (
	"github.com/robfig/cron/v3"
)

// update game DB from app store link for every 6h
func (app *application) runCronGameUpdater() {
	c := cron.New()
	c.AddFunc("@every 6h", app.CheckGames)
	c.Start()
}

// get all games from db
func (app *application) fetchGamesList() {
	_, gameList, err := app.models.Game.GetAll()
	if err != nil {
		app.logger.PrintError(err, nil)
		return
	}
	app.gamesAndGenres = gameList
}
