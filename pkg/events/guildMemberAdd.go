package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/BaileyJM02/Hue/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func guildMemberAdd(s *discordgo.Session, ma *discordgo.GuildMemberAdd) {
	var m *discordgo.Member
	m = toMember(ma)
	if sh.GetReady() {
		jsonValue, _ := json.Marshal(m)
		response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/guild/%v/member/%v", m.GuildID, m.User.ID), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			_ = data
		}
	}
}

func toMember(newmember *discordgo.GuildMemberAdd) *discordgo.Member {
	var member *discordgo.Member
	member = &discordgo.Member{
		newmember.GuildID,
		newmember.JoinedAt,
		newmember.Nick,
		newmember.Deaf,
		newmember.Mute,
		newmember.User,
		newmember.Roles,
	}
	return member
}

func init() {
	GuildMemberAdd := eh.Event{
		"guildMemberAdd",
		true,
		guildMemberAdd,
	}

	eh.Register(GuildMemberAdd)
}
