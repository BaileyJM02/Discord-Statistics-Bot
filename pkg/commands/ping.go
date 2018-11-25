package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/bwmarrin/discordgo"
)

func pingRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	msg, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":PONGING;%d", time.Now().UnixNano()))

	split := strings.Split(msg.Content, ";")
	if split[0] != ":PONGING" || len(split) < 2 {
		return
	}

	parsed, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		fmt.Println("err,", err)
		return
	}
	taken := time.Duration(time.Now().UnixNano() - parsed)

	started := time.Now()
	s.ChannelMessageEdit(m.ChannelID, msg.ID, "Gateway (http send -> gateway receive time): "+taken.String())
	httpPing := time.Since(started)

	s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("HTTP API: `%vms` \nGateway: `%vms`", int64(httpPing/time.Millisecond), int64(taken/time.Millisecond)))
}

func init() {
	ping := ch.Command{
		"ping",
		"ping",
		"see how long the bot takes to respond.",
		"General",
		false,
		false,
		pingRun,
	}

	ch.Register(ping)
}
