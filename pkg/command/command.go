package command

import (
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
		Commands map[string]Command)
}
