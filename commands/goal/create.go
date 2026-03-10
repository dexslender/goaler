package goal

import (
	"github.com/dexslender/goaler/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var subCreate = discord.ApplicationCommandOptionSubCommand{
	Name:        "create",
	Description: "create a new goal",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionInt{
			Name:        "type",
			Description: "type of goal do you need",
			Required:    true,
			Choices: []discord.ApplicationCommandOptionChoiceInt{
				{Name: "project", Value: int(util.TypeProject)},
				{Name: "milestone", Value: int(util.TypeMilestone)},
				{Name: "to-do list", Value: int(util.TypeTODO)},
				{Name: "single quest", Value: int(util.TypeQuest)},
				{Name: "habit", Value: int(util.TypeHabit)},
			},
		},
		discord.ApplicationCommandOptionString{
			Name:        "name",
			Description: "provide an unique name for your goal",
			Required:    true,
		},
	},
}

func runCreate(e *handler.CommandEvent) error {
	data := e.SlashCommandInteractionData()

	goalType := util.GoalType(data.Int("type"))
	goalName := data.String("name")

	switch goalType {
	case util.TypeProject:
		return e.CreateMessage(discord.NewMessageCreateV2(
				discord.NewTextDisplay("Project created successfully with name: " + goalName),
			))
	case util.TypeMilestone:
		return e.CreateMessage(discord.NewMessageCreateV2(
				discord.NewTextDisplay("Milestone created successfully with name: " + goalName),
			))
	case util.TypeTODO:
		return e.CreateMessage(discord.NewMessageCreateV2(
				discord.NewTextDisplay("To-do list created successfully with name: " + goalName),
			))
	case util.TypeQuest:
		return e.CreateMessage(discord.NewMessageCreateV2(
				discord.NewTextDisplay("Quest created successfully with name: " + goalName),
			))
	case util.TypeHabit:
		return e.CreateMessage(discord.NewMessageCreateV2(
				discord.NewTextDisplay("Habit created successfully with name: " + goalName),
			))
	default:
		return e.CreateMessage(discord.NewMessageCreateV2(
				discord.NewTextDisplay("Invalid goal type"),
			))
	}
}
