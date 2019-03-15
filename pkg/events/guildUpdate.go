package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/BaileyJM02/Hue/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func guildUpdate(s *discordgo.Session, gc *discordgo.GuildUpdate) {
	var g *discordgo.Guild
	g = gc.Guild

	jsonValue, _ := json.Marshal(g)
	response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/up-to-date/"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		_ = data
	}
}

func init() {
	GuildUpdate := eh.Event{
		"guildUpdate",
		true,
		guildUpdate,
	}

	eh.Register(GuildUpdate)
}
