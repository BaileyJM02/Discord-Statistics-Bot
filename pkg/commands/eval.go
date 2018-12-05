package commands

import (
	"fmt"
	"strings"
	"time"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/bwmarrin/discordgo"
	"github.com/novalagung/golpal"
)

func evalRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	if len(content[1:]) == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("*BOOOOM*"))
	} else {
		parsed := time.Now().UnixNano()
		cmdString := `
			import (sh "github.com/BaileyJM02/Hue/pkg/clientHandler";
			ch "github.com/BaileyJM02/Hue/pkg/commandHandler"; "strings"); func main() {fmt.Print();_ =sh.GetPrefix; _=ch.Commands; _=strings.Join([]string{"big","boy"}, ".");` + strings.Join(content[1:], " ") + ";}"
		output, err := golpal.New().Execute(cmdString)
		if err != nil {
			takenErr := time.Duration(time.Now().UnixNano() - parsed)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error => %.2fs```\n%v```", takenErr.Seconds(), err.Error()))
		}
		taken := time.Duration(time.Now().UnixNano() - parsed)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("result => %.2fs```\n%v```", taken.Seconds(), output))
	}
}

func init() {
	eval := ch.Command{
		"eval",
		"eval <expression>",
		"Shhhh",
		"Admin",
		true,
		map[string]bool{
			"expression": true,
		},
		true,
		true,
		evalRun,
	}

	ch.Register(eval)
}
