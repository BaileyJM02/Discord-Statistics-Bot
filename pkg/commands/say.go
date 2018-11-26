package commands

import (
	"fmt"
	"strings"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func SayCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
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
		map[string]bool{
			"text": true,
		},
		false,
		SayCommand,
	}

	ch.Register(say)
}
