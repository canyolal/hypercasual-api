package main

import (
	"github.com/robfig/cron/v3"
)

// run a cron job and check and update if new game exists
func (app *application) runCronGameUpdater() {
	c := cron.New()
	c.AddFunc("@every 1m", app.CheckGames) // every 3h
	c.Start()
}

// get all games from db
func (app *application) fetchGamesList() {
	_, gameList, err := app.models.Game.GetAll()
	if err != nil {
		app.logger.PrintError(err, nil)
		return
	}
	if gameList != nil {
		app.gamesAndGenres = gameList
	}
}
