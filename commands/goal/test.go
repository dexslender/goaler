package goal

import (
	"strconv"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var subTest = discord.ApplicationCommandOptionSubCommand{
	Name:        "test",
	Description: "test propourses",
}

func runTest(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetIsComponentsV2(true).
		SetEphemeral(true).
		AddComponents(ui).
		Build())
}

var ui = discord.NewContainer(
	discord.NewTextDisplay("# TO-DO"),
	discord.NewActionRow(
		discord.NewSuccessButton("read", "/goal/test/hello1"),
		discord.NewSuccessButton("write", "/goal/test/hello2"),
		discord.NewSuccessButton("speak", "/goal/test/hello3"),
		discord.NewSuccessButton("listen", "/goal/test/hello4"),
	))

func handleButtons(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	subID := e.Vars["button"]
	index, _ := strconv.Atoi(string(subID[len(subID)-1]))
	ui.Components[1].(discord.ActionRowComponent).Components[index-1].(discord.ButtonComponent).AsDisabled()
	return e.UpdateMessage(discord.NewMessageUpdateBuilder().
		SetComponent(0, ui).
		Build())
}
