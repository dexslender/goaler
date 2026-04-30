package dev

import (
	"github.com/dexslender/goaler/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

// Test command
var test = discord.ApplicationCommandOptionSubCommand{
	Name:        "test",
	Description: "run a random test",
}

// Add Task Modal
var addTask = discord.NewModalCreate(
	"/dev/add-task",
	"add new task...",
	discord.NewLabel("name",
		discord.NewShortTextInput("/dev/add-task/name").
			WithPlaceholder("do something").
			WithRequired(true)),
	discord.NewLabel("description",
		discord.NewParagraphTextInput("/dev/add-task/description").
			WithPlaceholder("this task consists of...")),
)

func _runTest(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateV2().
		AddComponents(discord.NewActionRow(discord.NewSuccessButton("hello", "/dev/hello"))))
}

func _handleButton(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	e.Client().Logger.Info(data.CustomID())
	targetComp := e.Message.Components[0].(discord.ActionRowComponent)
	targetButt := targetComp.Components[0].(discord.ButtonComponent)
	targetButt.Disabled = true
	targetComp.UpdateComponent(targetButt.ID, targetButt)

	return e.UpdateMessage(discord.NewMessageUpdateV2(targetComp))
}

func runTest(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateV2().
		AddComponents(discord.NewContainer(
			discord.NewSection(
				discord.NewTextDisplay("### task 1"),
				discord.NewTextDisplay("doing something...")).
				WithAccessory(discord.NewSuccessButton("", "/dev/check/1").
					WithEmoji(util.Check)),
			discord.NewSection(
				discord.NewTextDisplay("### task 2"),
				discord.NewTextDisplay("do this after work...")).
				WithAccessory(discord.NewSecondaryButton("", "/dev/check/2").
					WithEmoji(util.Uncheck)),
			discord.NewActionRow(
				discord.NewSecondaryButton("", "/dev/check/add").
					WithEmoji(util.Plus)),
		)))
}

func handleButtons(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	action := e.Vars["action"]
	switch action {
	case "add":
		return e.Modal(addTask)
	default:
		container, ok := e.Message.Components[0].(discord.ContainerComponent)
		if !ok {
			return e.UpdateMessage(
				discord.NewMessageUpdateV2(discord.NewTextDisplay("Unknown error")),
			)
		}
		for i, component := range container.Components {
			section, ok := component.(discord.SectionComponent)
			if !ok {
				continue
			}
			button, ok := section.Accessory.(discord.ButtonComponent)
			if !ok || button.CustomID != data.CustomID() {
				continue
			}

			switch button.Style {
			case discord.ButtonStyleSecondary:
				button.Emoji = &util.Check
				button.Style = discord.ButtonStyleSuccess
			case discord.ButtonStyleSuccess:
				button.Emoji = &util.Uncheck
				button.Style = discord.ButtonStyleSecondary
			}
			// b.Disabled = true
			section.Accessory = button
			container.Components[i] = section
			break
		}
		return e.UpdateMessage(discord.NewMessageUpdateV2(container))
	}
}
