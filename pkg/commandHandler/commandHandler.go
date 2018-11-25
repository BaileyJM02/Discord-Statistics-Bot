package commandHandler

import (
	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Category    string
	NeedArgs    bool
	Run         func(s *discordgo.Session,
		m *discordgo.MessageCreate,
		content []string,
		Commands map[string]Command,
		client sh.Client)
}

var (
	Commands map[string]Command
)

func Register(cmd Command) {
	if Commands == nil {
		Commands = make(map[string]Command)
	}
	Commands[cmd.Name] = cmd
}
