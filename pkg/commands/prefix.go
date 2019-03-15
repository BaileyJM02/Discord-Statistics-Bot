package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/BaileyJM02/Hue/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

func prefixRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	var msg *discordgo.Message
	msg, _ = s.ChannelMessageSend(m.ChannelID, "<a:load:522124372140490753> Loading data.")
	if len(content[1:]) < 1 {
		var prefix string
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/prefix", m.GuildID))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			json.NewDecoder(response.Body).Decode(&prefix)
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Prefix", m.Author.AvatarURL("")).
			SetDescription("View the guild's prefix information.").
			SetColor(0xffffff).
			AddField("Prefix:", "`"+prefix+"`", false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			logger.Error(err)
		}
		return
	}
	if content[1] == "reset" {
		perms, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			logger.Error(err)
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Prefix (Denied)", m.Author.AvatarURL("")).
				SetDescription("We couldn't calculate your perms so I'm guessing you don't have permission to do that?").
				SetColor(0xE74C3C)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			return
		}
		if perms&discordgo.PermissionManageServer != discordgo.PermissionManageServer {
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Prefix (Denied)", m.Author.AvatarURL("")).
				SetDescription("You don't have permission to do that.").
				SetColor(0xE74C3C)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			return
		}
		var prefix string
		prefix = "-"
		jsonValue, _ := json.Marshal(prefix)
		response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/guild/%v/prefix/%v", m.GuildID, m.ID), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %v", err))
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			_ = data
		}

		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Prefix", m.Author.AvatarURL("")).
			SetDescription("View the guild's prefix information.").
			SetColor(0x8BC34A).
			AddField("Reset Prefix:", "`"+prefix+"`", false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			logger.Error(err)
		}
		return
	}
	if content[1] == "set" {
		perms, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			logger.Error(err)
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Prefix (Denied)", m.Author.AvatarURL("")).
				SetDescription("We couldn't calculate your perms so I'm guessing you don't have permission to do that?").
				SetColor(0xE74C3C)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			return
		}
		if perms&discordgo.PermissionManageServer != discordgo.PermissionManageServer {
			em := embed.NewEmbed().
				SetAuthor(m.Author.Username+" | Prefix (Denied)", m.Author.AvatarURL("")).
				SetDescription("You don't have permission to do that.").
				SetColor(0xE74C3C)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
			s.ChannelMessageDelete(m.ChannelID, msg.ID)
			return
		}
		var prefix string
		prefix = content[2]
		jsonValue, _ := json.Marshal(prefix)
		response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/guild/%v/prefix/%v", m.GuildID, m.ID), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %v", err))
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			_ = data
		}

		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Prefix", m.Author.AvatarURL("")).
			SetDescription("View the guild's prefix information.").
			SetColor(0x8BC34A).
			AddField("Set Prefix:", "`"+prefix+"`", false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			logger.Error(err)
		}
		return
	}
}

func init() {
	prefix := ch.Command{
		"prefix",
		"prefix [set <prefix>] [reset]",
		"Set the bot's prefix",
		"Settings",
		false,
		map[string]bool{
			"reset": false,
			"set":   false,
		},
		false,
		true,
		prefixRun,
	}

	ch.Register(prefix)
}
