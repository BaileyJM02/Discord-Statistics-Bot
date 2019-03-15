package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/BaileyJM02/Hue/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func guildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	if sh.GetReady() {
		deleteMember(m)
	}
}

func deleteMember(m *discordgo.GuildMemberRemove) {
	// Request (DELETE http://www.example.com/bucket/sample)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/db/guild/%v/member/%v", m.GuildID, m.User.ID), nil)
	if err != nil {
		logger.Error(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	_ = respBody
	if err != nil {
		logger.Error(err)
		return
	}
}

func init() {
	GuildMemberRemove := eh.Event{
		"guildMemberRemove",
		true,
		guildMemberRemove,
	}

	eh.Register(GuildMemberRemove)
}
