package commands

import (
	"fmt"
	"strings"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"

	// needed to populate Commands map[]
	_ "github.com/BaileyJM02/Hue/pkg/commands"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	prefix := sh.GetPrefix()

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Fields(m.Content)

	if !(strings.Contains(content[0], prefix)) {
		return
	}

	content[0] = strings.Replace(content[0], prefix, "", -1)

	if cmd, ok := ch.Commands[content[0]]; ok {
		if cmd.OwnerOnly && (m.Author.ID != "398197113495748626") {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Error:** You're not the owner!"))
			return
		}
		fmt.Println("Command Run: ", cmd.Name)
		cmd.Run(s, m, content, ch.Commands, sh.Bot)
		return
	}

	// text := "**\\" + prefix + "Ping**, see how long the bot takes to respond."
	// embed := embed.NewEmbed().
	// 	SetTitle("Commands - Ping").
	// 	SetDescription(text + "\n\nUse `" + prefix + "help <command>` for more help.").
	// 	SetColor(0x00000).MessageEmbed
	// s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func init() {
	MessageCreate := eh.Event{
		"messageCreate",
		true,
		messageCreate,
	}

	eh.Register(MessageCreate)
}
