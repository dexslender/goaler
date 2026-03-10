package dev

import (
	"codeberg.org/dou/goaler/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var test = discord.ApplicationCommandOptionSubCommand{
	Name:        "test",
	Description: "run a random test",
}

var addTask = discord.NewModalCreateBuilder().
	SetTitle("add new task...").
	SetCustomID("/dev/add-task").
	AddComponents(
		discord.NewLabel("name",
			discord.NewShortTextInput("/dev/add-task/name").
				WithPlaceholder("do something").
				WithRequired(true)).WithDescription("provide a short name for your task"),
		discord.NewLabel("description",
			discord.NewParagraphTextInput("/dev/add-task/description").
				WithPlaceholder("this task consists of..."))).
	Build()

func _runTest(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetIsComponentsV2(true).
		AddComponents(discord.NewActionRow(discord.NewSuccessButton("hello", "/dev/hello"))).
		Build())
}

func _handleButton(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	e.Client().Logger.Info(data.CustomID())
	targetComp := e.Message.Components[0].(discord.ActionRowComponent)
	targetButt := targetComp.Components[0].(discord.ButtonComponent)
	targetButt.Disabled = true
	targetComp.UpdateComponent(targetButt.ID, targetButt)

	return e.UpdateMessage(discord.NewMessageUpdateBuilder().
		SetComponents(targetComp).
		Build())
}

func runTest(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetIsComponentsV2(true).
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
		)).
		Build())
}

func handleButtons(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	action := e.Vars["action"]
	switch action {
	case "add":
		return e.Modal(addTask)
	default:
		container, ok := e.Message.Components[0].(discord.ContainerComponent)
		if !ok {
			return e.UpdateMessage(discord.NewMessageUpdateBuilder().
				SetComponents(discord.NewTextDisplay("Unknown error")).
				Build())
		}
		for k, v := range container.Components {
			if s, ok := v.(discord.SectionComponent); ok {
				if b, ok := s.Accessory.(discord.ButtonComponent); 
				ok && b.CustomID == data.CustomID() {
					if b.Style == discord.ButtonStyleSecondary {
						b.Emoji = &util.Check
						b.Style = discord.ButtonStyleSuccess
					} else if b.Style == discord.ButtonStyleSuccess {
						b.Emoji = &util.Uncheck
						b.Style = discord.ButtonStyleSecondary
					}
					// b.Disabled = true
					s.Accessory = b
					container.Components[k] = s
					break
				}
			}
		}

		return e.UpdateMessage(discord.NewMessageUpdateBuilder().
			SetComponents(container).
			Build())
	}
}
