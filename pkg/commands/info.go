package commands

import (
	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/BaileyJM02/Hue/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

func infoCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	em := embed.NewEmbed().
		SetAuthor(m.Author.Username+" | Info", m.Author.AvatarURL("")).
		SetDescription(`We are a small group of people including one coder and multiple administrators and supporters who all have a love for the Discord platform. That is why we wanted to give back to the community with a bot that makes your server information easily accessible and manageable. We strive to make a support tool for your Discord server that creates an image and overview of a lot of different information including member and server activity.

		If you have any questions about the service or how to set it up then feel free to contact our support. You can find our support options [here](https://hue.observer/support) or [invite the bot](https://hue.observer/invite).`).
		SetColor(0xffffff)

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
	if err != nil {
		logger.Error(err)
	}
	return
}

func init() {
	info := ch.Command{
		"info",
		"info",
		"Get infomation about Hue.",
		"General",
		false,
		map[string]bool{},
		false,
		true,
		infoCommand,
	}

	ch.Register(info)
}
