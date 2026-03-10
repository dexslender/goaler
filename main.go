package main

import (
	"os"

	"codeberg.org/dou/goaler/bot"
	"codeberg.org/dou/goaler/commands"
	"codeberg.org/dou/goaler/util"
	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
	"github.com/disgoorg/snowflake/v2"
	"github.com/kkyr/fig"
)

func main() {
	logger := log.NewWithOptions(
		os.Stderr,
		log.Options{ReportTimestamp: true},
	)

	var config bot.Config
	if err := fig.Load(&config,
		fig.UseEnv("GOALER"),
		fig.File("goaler.yaml"),
	); err != nil {
		logger.Fatal(err)
	}

	util.SetupDB(logger)

	goaler := bot.New(&config, logger)
	router := handler.New()
	router.NotFound(commands.NotFound(logger))
	router.With(middleware.Logger)
	commands.Setup(&goaler, router)
	goaler.Setup()
	if goaler.GuildID != 0 {
		handler.SyncCommands(
			goaler.Client,
			commands.Commands,
			[]snowflake.ID{goaler.GuildID},
		)
	}

	goaler.AddEventListeners(router)

	goaler.StartNLock()
}
