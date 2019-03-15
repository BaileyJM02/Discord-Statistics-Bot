package commands

import (
	"fmt"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func StatsCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	s.ChannelMessageSend(m.ChannelID,
		fmt.Sprintf("Visit <https://hue.observer/stats> to see real-time statistics."))
		// sh.GetUptime()
}

func init() {
	stats := ch.Command{
		"stats",
		"stats",
		"Get info",
		"General",
		false,
		map[string]bool{},
		false,
		true,
		StatsCommand,
	}

	ch.Register(stats)
}
