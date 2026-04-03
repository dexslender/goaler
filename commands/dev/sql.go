package dev

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var sql = discord.ApplicationCommandOptionSubCommand{
	Name:        "sql",
	Description: "do a sql query to the database",
}

func runSQL(e *handler.CommandEvent) error {
	return e.Modal(discord.NewModalCreate(
		"/dev/sqlquery",
		"do a sql query...",
		discord.NewLabel("query",
			discord.NewParagraphTextInput("query").
				WithPlaceholder("No WHERE, everywhere!"),
		),
	))
}

func handleModal(e *handler.ModalEvent) error {
	query := e.Data.Text("query")
	e.Client().Logger.Info(query)
	return e.CreateMessage(discord.NewMessageCreateV2(
		discord.NewTextDisplay("ok"),
	).WithEphemeral(true))
}
