package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func guildCreate(s *discordgo.Session, gc *discordgo.GuildCreate) {
	var g *discordgo.Guild
	g = gc.Guild

	if sh.GetReady() {
		guilds := sh.SetGuilds(sh.GetGuilds() + 1)
		s.ChannelMessageSend("519970380266340352", fmt.Sprintf("**%v** invited Hue. Now serving %v guilds.", g.Name, guilds))
		jsonValue, _ := json.Marshal(g)
		response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/guild/%v", g.ID), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
}

func init() {
	GuildCreate := eh.Event{
		"guildCreate",
		true,
		guildCreate,
	}

	eh.Register(GuildCreate)
}
