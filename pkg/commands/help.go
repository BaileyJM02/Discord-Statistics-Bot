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
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("\n**%v%v** [%v]\n\nDescription: %v\nUsage: %v%v\nArgs required?: %v", sh.GetPrefix(), cmd.Name, cmd.Category, cmd.Description, sh.GetPrefix(), cmd.Usage, cmd.NeedArgs))
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
		false,
		helpRun,
	}

	ch.Register(help)
}
