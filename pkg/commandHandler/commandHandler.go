package commandHandler

import (
	"fmt"
	"strings"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Category    string
	NeedArgs    bool
	Args        map[string]bool
	OwnerOnly   bool
	Enabled     bool
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
	if Commands[cmd.Name].Enabled {
		Commands[cmd.Name] = cmd
	}
}

// HelpEmbed [MessageCreate] [Command]
func HelpEmbed(m *discordgo.MessageCreate, cmd Command) *discordgo.MessageEmbed {
	var args []string
	if len(cmd.Args) == 0 {
		args = append(args, "No args.")
	}
	for key, value := range cmd.Args {
		if value == true {
			args = append(args, fmt.Sprintf("<%v>", key))
		} else {
			args = append(args, fmt.Sprintf("[%v]", key))
		}
	}
	em := embed.NewEmbed().
		SetAuthor(m.Author.Username+" | "+strings.Title(cmd.Name[:1])+cmd.Name[1:], m.Author.AvatarURL("")).
		SetDescription(strings.Title(cmd.Description[:1])+cmd.Description[1:]).
		SetFooter("\n\nUse "+sh.GetPrefix()+"help for more commands.").
		AddField("Usage", sh.GetPrefix()+cmd.Usage, false).
		AddField("Args", strings.Join(args, ", "), false).
		SetColor(0xffffff).MessageEmbed

	return em
}
