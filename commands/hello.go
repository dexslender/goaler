package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var hello = discord.SlashCommandCreate{
	Name: "hello",
	Description: "just sends a hello message!",
}

func runHello(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContentf("Hello %s", e.User().Mention()).
		Build())
}
