package commands

import (
	"fmt"
	"strings"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/bwmarrin/discordgo"
)

func helpRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	if len(content[1:]) == 0 {
		cats := make(map[string]string)
		for _, value := range ch.Commands {
			if _, ok := cats[ch.Commands[value.Name].Category]; !ok {
				cats[ch.Commands[value.Name].Category] = value.Name
			} else {
				cats[ch.Commands[value.Name].Category] += ", " + value.Name
			}
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Commands", m.Author.AvatarURL("")).
			SetDescription("Commands avaliable for use.").
			SetColor(0xffffff)
		for key, value := range cats {
			em.AddField(strings.Title(key), strings.Title(value), false)
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if cmd, ok := Commands[content[1]]; ok {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, ch.HelpEmbed(m, cmd))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Error:** Command \"%v\" not found", content[1]))
	}
}

func init() {
	help := ch.Command{
		"help",
		"help <command>",
		"Help you",
		"General",
		false,
		map[string]bool{
			"command": false,
		},
		false,
		true,
		helpRun,
	}

	ch.Register(help)
}
