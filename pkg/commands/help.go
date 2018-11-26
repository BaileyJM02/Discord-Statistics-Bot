package commands

import (
	"fmt"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func helpRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	if len(content[1:]) == 0 {
		commands := ""
		for key, value := range Commands {
			commands += fmt.Sprintf("\n**%v%v**: %v", sh.GetPrefix(), key, value.Description)
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Commands:** %v.", commands))
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
		helpRun,
	}

	ch.Register(help)
}
