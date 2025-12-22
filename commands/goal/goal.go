// Package goal
package goal

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Goal = discord.SlashCommandCreate{
	Name: "goal",
	Description: "manage your goals",
	Options: []discord.ApplicationCommandOption{subCreate, subList, subTest},
}

func Handle(r handler.Router) {
	r.Command("/create", runCreate)
	r.Command("/test", runTest)

	r.ButtonComponent("/test/{button}", handleButtons)
}
