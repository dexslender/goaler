package commands

import (
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate{
	Name:        "ping",
	Description: "get latency of the bot",
}

func runPing(e *handler.CommandEvent) error {
	a := time.Now()
	err := e.DeferCreateMessage(false)
	if err != nil { return err }
	rest := time.Since(a).Round(time.Millisecond)
	gate := e.Client().Gateway.Latency().Round(time.Millisecond)

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateV2(
		buildContainer(rest, gate),
	))
	return err
}

func runPingRefresh(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	a := time.Now()
	err := e.DeferUpdateMessage()
	if err != nil { return err }
	rest := time.Since(a).Round(time.Millisecond)
	gate := e.Client().Gateway.Latency().Round(time.Millisecond)
	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateV2(

		buildContainer(rest, gate),
	))
	return err
}

func buildContainer(rest, gate time.Duration) discord.ContainerComponent {
	return discord.NewContainer(
		discord.NewTextDisplayf("```yaml\nrest: %s\ngateway: %s```", rest, gate),
		discord.NewActionRow(
			discord.NewSecondaryButton("pong", "/ping/refresh").
				WithEmoji(discord.NewComponentEmoji("🏓")),
		),
	)
}
