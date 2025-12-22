// Package dev
package dev

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Dev = discord.SlashCommandCreate{
	Name: "dev",
	Description: "just development utilities",
	Options: []discord.ApplicationCommandOption{sql, test},
}

func Handle(r handler.Router) {
	r.Command("/sql", runSQL)
	r.Command("/test", runTest)

	r.Modal("/sqlquery", handleModal)
}
