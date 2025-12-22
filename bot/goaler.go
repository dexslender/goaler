// Package bot: main bot package
package bot

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
)

type Goaler struct {
	*bot.Client
	*Config
	*log.Logger
}

func New(c *Config, l *log.Logger) Goaler {
	return Goaler{ Config: c, Logger: l }
}

func (g *Goaler) Setup() {
	var err error
	g.Client, err = disgo.New(g.Config.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsNonPrivileged|
				gateway.IntentGuilds,
			),
		),
		bot.WithEventListenerFunc(OnReady(g)),
		bot.WithLogger(slog.New(g)),
	)
	if err != nil { g.Fatal(err) }
}

func (g *Goaler) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second * 10)
	defer c()
	defer func(){
		g.Gateway.Close(ctx)
		g.Info("gateway closed, program finished.")
	}()


	if err := g.OpenGateway(ctx); err != nil {
		g.Fatal(err)
	}

	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}

func OnReady(g *Goaler) func(*events.Ready) {
	return func(r *events.Ready) {
		g.Infof("logged in as %s", r.User.Tag())
	}
}
