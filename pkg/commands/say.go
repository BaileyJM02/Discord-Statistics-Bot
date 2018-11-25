package commands

import (
	"fmt"
	"strings"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func SayCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	if len(content[1:]) == 0 {
		s.ChannelMessageSend(m.ChannelID, "**Error:** No arguments given (expected text)")
		return
	}
	s.ChannelMessageSend(m.ChannelID,
		fmt.Sprintf("%v", strings.Join(content[1:], " ")))
}

func init() {
	say := ch.Command{
		"say",
		"say <message>",
		"Say something",
		"General",
		true,
		SayCommand,
	}

	ch.Register(say)
}
