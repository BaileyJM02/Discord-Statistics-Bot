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
func guildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	if sh.GetReady() {
		guilds := sh.SetGuilds(sh.GetGuilds() - 1)
		s.ChannelMessageSend("519970380266340352", fmt.Sprintf("**%v** removed Hue. Now serving %v guilds.", g.Name, guilds))
		deleteGuild(g)
	}
}

func deleteGuild(g *discordgo.GuildDelete) {
	// Request (DELETE http://www.example.com/bucket/sample)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/db/guild/%v", g.ID), nil)
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
	logger.Info("A guild delete event has been triggered.")
}

func init() {
	GuildDelete := eh.Event{
		"guildDelete",
		true,
		guildDelete,
	}

	eh.Register(GuildDelete)
}
