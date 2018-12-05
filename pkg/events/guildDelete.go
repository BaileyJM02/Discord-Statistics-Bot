package commands

import (
	"fmt"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func guildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	if sh.GetReady() {
		guilds := sh.SetGuilds(sh.GetGuilds() - 1)
		s.ChannelMessageSend("519970380266340352", fmt.Sprintf("**%v** removed Hue. Now serving %v guilds.", g.Name, guilds))
	}
}

func init() {
	GuildDelete := eh.Event{
		"guildDelete",
		true,
		guildDelete,
	}

	eh.Register(GuildDelete)
}
