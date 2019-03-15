package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	api "github.com/BaileyJM02/Hue/pkg/apiHandler"
	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/BaileyJM02/Hue/pkg/logger"

	// needed to populate Commands map[]
	_ "github.com/BaileyJM02/Hue/pkg/commands"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself

	prefix := sh.GetPrefix(m.GuildID)

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Fields(m.Content)

	if m.Content == "" {
		sendData(s, m)
		return
	}

	if m.Author.Bot == true {
		sendData(s, m)
		return
	}

	if !(strings.Contains(content[0], prefix) ||
		strings.Contains(content[0], fmt.Sprintf("<@%v>", s.State.User.ID)) ||
		strings.Contains(content[0], fmt.Sprintf("<@!%v>", s.State.User.ID))) {
		sendData(s, m)
		return
	}

	content[0] = strings.Replace(content[0], prefix, "", -1)

	if strings.Contains(content[0], fmt.Sprintf("<@%v>", s.State.User.ID)) ||
		strings.Contains(content[0], fmt.Sprintf("<@!%v>", s.State.User.ID)) {
		content = append(content[:0], content[1:]...)
	}

	if cmd, ok := ch.Commands[content[0]]; ok {
		if cmd.OwnerOnly && (m.Author.ID != "398197113495748626") {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Error:** You're not the owner!"))
			sendData(s, m)
			return
		}

		if cmd.NeedArgs == true && len(content[1:]) == 0 {
			s.ChannelMessageSendEmbed(m.ChannelID, ch.HelpEmbed(m, cmd))
			sendData(s, m)
			return
		}
		guild, _ := s.Guild(m.GuildID)
		logger.Info(fmt.Sprintf("Command Run: `%v` in %v", cmd.Name, guild.Name))
		cmd.Run(s, m, content, ch.Commands, sh.Bot)
		sendData(s, m)
		return
	}
}

// Fallback checks to prevent major errors.
func sendData(s *discordgo.Session, m *discordgo.MessageCreate) {
	// user, err := s.GuildMember(m.GuildID, m.ID)
	// if err != nil {
	// 	fmt.Printf("The GuildMember request failed with error %s\n", err)
	// }
	// jsonValue, _ := json.Marshal(user)
	// response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/guild/%v/member/%v", m.GuildID, m.ID), "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// }
	mmm, _ := s.GuildMember(m.GuildID, m.Author.ID)
	mm := api.MessageHelper{
		m,
		mmm,
	}
	jsonValue, _ := json.Marshal(mm)
	response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/guild/%v/message/%v", m.GuildID, m.ID), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Error(fmt.Sprintf("The HTTP request failed with error %v", err))
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		_ = data
	}
}

func init() {
	MessageCreate := eh.Event{
		"messageCreate",
		true,
		messageCreate,
	}

	eh.Register(MessageCreate)
}
