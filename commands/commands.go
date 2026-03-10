// Package commands
package commands

import (
	"codeberg.org/dou/goaler/bot"
	"codeberg.org/dou/goaler/commands/dev"
	"codeberg.org/dou/goaler/commands/goal"
	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Commands = []discord.ApplicationCommandCreate{
	hello,
	ping,
	goal.Goal,
	// development and debug propurses commands
	// TODO - disable on production
	dev.Dev,
}

func Setup(o *bot.Goaler, r handler.Router) {
	r.Command("/hello", runHello)
	r.Command("/ping", runPing)

	r.Route("/goal", goal.Handle)
	r.Route("/dev", dev.Handle) // disable on production

	r.ButtonComponent("/ping/refresh", runPingRefresh)
}

func NotFound(l *log.Logger) func(*handler.InteractionEvent) error {
	return func(ie *handler.InteractionEvent) error {
		l.Error("command not found")
		return ie.CreateMessage(discord.NewMessageCreateV2(discord.NewContainer(
			discord.NewTextDisplay("# Not found\n-# maybe is still in development"),
		)).WithEphemeral(true))
	}
}
