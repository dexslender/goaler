package goal

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type GoalType int
const (
	TypeProject GoalType = iota+1 
	TypeMilestone
	TypeTODO
	TypeHabit
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
				{Name: "project",    Value: int(TypeProject)},
				{Name: "milestone",  Value: int(TypeMilestone)},
				{Name: "to-do list", Value: int(TypeTODO)},
				{Name: "habit",      Value: int(TypeHabit)},
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
	// data := e.SlashCommandInteractionData()
	return nil
}
