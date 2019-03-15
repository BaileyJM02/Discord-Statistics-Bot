package commands

import (
	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/BaileyJM02/Hue/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

func inviteCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	em := embed.NewEmbed().
		SetAuthor(m.Author.Username+" | Invite", m.Author.AvatarURL("")).
		SetDescription("Invite: [hue.observer/invite](https://hue.observer/invite)\nSupport: [hue.observer/join](https://hue.observer/join)").
		SetColor(0xffffff)

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
	if err != nil {
		logger.Error(err)
	}
	return
}

func init() {
	invite := ch.Command{
		"invite",
		"invite",
		"Get guild and bot invite links",
		"General",
		false,
		map[string]bool{},
		false,
		true,
		inviteCommand,
	}

	ch.Register(invite)
}
