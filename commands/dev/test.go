package dev

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var test = discord.ApplicationCommandOptionSubCommand{
	Name:        "test",
	Description: "run a random test",
}

func runTest(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.MessageCreate{
		Flags: discord.MessageFlagIsComponentsV2,
		Components: []discord.LayoutComponent{discord.ContainerComponent{
			Components: []discord.ContainerSubComponent{
				discord.SectionComponent{
					Components: []discord.SectionSubComponent{
						discord.NewTextDisplay("### some title"),
						discord.NewTextDisplay("some description"),
					},
					Accessory: discord.NewSuccessButton("/dev/check", discord.ComponentEmoji{ Name: "✅" }),
				},
			},
		}},
	})
	// return e.CreateMessage(discord.NewMessageCreateBuilder().
	// 	SetIsComponentsV2(true).
	// 	AddComponents(discord.NewContainer(
	// 		discord.NewSection(
	// 			discord.NewTextDisplay("Some title"),
	// 			discord.NewTextDisplay("some description"),
	// 			).WithAccessory(discord.NewSuccessButton("", "check").
	// 				WithEmoji(discord.NewComponentEmoji("✅"))),
	// 		)).
	// 	Build())
}
