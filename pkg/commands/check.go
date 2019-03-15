package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/BaileyJM02/Hue/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

func checkRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	var msg *discordgo.Message
	var edited bool
	msg, _ = s.ChannelMessageSend(m.ChannelID, "<a:load:522124372140490753> Checking.")
	if len(content[1:]) == 0 {
		guild, err := s.Guild(m.GuildID)
		if err != nil {
			logger.Error(err)
		}
		jsonValue, _ := json.Marshal(guild)
		response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/check/"), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			s := string(data[:])
			s = strings.Replace(s, "\n", "", -1)
			edited, err = strconv.ParseBool(s)
			if err != nil {
				logger.Error(err)
			}
		}
		if edited == true {
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Check", m.Author.AvatarURL("")).
				SetDescription("Thank you for running that command, some edits needed to be made. Hue is now fully up-to-date.").
				SetColor(0x8BC34A)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			if err != nil {
				logger.Error(err)
			}
			return
		}
		if edited == false {
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Check", m.Author.AvatarURL("")).
				SetDescription("Thank you for running that command. No edits needed to be made.").
				SetColor(0xFFC107)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			if err != nil {
				logger.Error(err)
			}
			return
		}
	}
	if len(content[1:]) > 0 {
		guild, err := s.Guild(content[1])
		if err != nil {
			s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("**Error:** Guild not found."))
			return
		}
		jsonValue, _ := json.Marshal(guild)
		response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/check/"), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			s := string(data[:])
			s = strings.Replace(s, "\n", "", -1)
			edited, _ = strconv.ParseBool(s)
		}
		if edited == true {
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Check", m.Author.AvatarURL("")).
				SetDescription("Thank you for running that command, some edits needed to be made. Hue is now fully up-to-date.").
				SetColor(0x8BC34A)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			if err != nil {
				logger.Error(err)
			}
			return
		}
		if edited == false {
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Check", m.Author.AvatarURL("")).
				SetDescription("Thank you for running that command. No edits needed to be made.").
				SetColor(0xFFC107)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			if err != nil {
				logger.Error(err)
			}
			return
		}
	}
}

func init() {
	check := ch.Command{
		"check",
		"check [guildID]",
		"Check that guild data is correct",
		"Stats",
		false,
		map[string]bool{
			"guildID": false,
		},
		true,
		false,
		checkRun,
	}

	ch.Register(check)
}
