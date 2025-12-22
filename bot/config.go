package bot

import "github.com/disgoorg/snowflake/v2"

type Config struct {
	Token string `validate:"required"`
	GuildID snowflake.ID `fig:"guildID"`
}
