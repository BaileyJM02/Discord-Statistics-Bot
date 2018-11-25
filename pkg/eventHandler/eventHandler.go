package eventHandler

import (
	"github.com/bwmarrin/discordgo"
)

type Event struct {
	Name    string
	Enabled bool
	Run     func(s *discordgo.Session,
		m *discordgo.MessageCreate)
}

var (
	Events map[string]Event
)

func Register(evt Event) {
	if Events == nil {
		Events = make(map[string]Event)
	}
	Events[evt.Name] = evt
}
