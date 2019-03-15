package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	api "github.com/BaileyJM02/Hue/pkg/apiHandler"
	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/BaileyJM02/Hue/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

func guildinfoRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	var msg *discordgo.Message
	msg, _ = s.ChannelMessageSend(m.ChannelID, "<a:load:522124372140490753> Loading data.")
	if len(content[1:]) < 1 {
		var stats *api.GuildStats
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/stats", m.GuildID))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			json.NewDecoder(response.Body).Decode(&stats)
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | General", m.Author.AvatarURL("")).
			SetDescription("View the guild's stats.").
			SetColor(0xffffff).
			AddField("Total Members:", strconv.Itoa(stats.TotalMembers), false).
			AddField("Total Messages:", strconv.Itoa(stats.TotalMessages), false).
			AddField("Messages containing attachments:", strconv.Itoa(stats.AttachmentsSent), false).
			AddField("Messages containing links:", strconv.Itoa(stats.LinksSent), false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			logger.Error(err)
		}
		return
	}
	if content[1] == "messages" {
		var stats *api.GuildStats
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/stats", m.GuildID))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			json.NewDecoder(response.Body).Decode(&stats)
		}

		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Messages", m.Author.AvatarURL("")).
			SetDescription("View message stats.").
			SetColor(0xffffff).
			AddField("Total Messages:", strconv.Itoa(stats.TotalMessages), false).
			AddField("Containing attachments:", strconv.Itoa(stats.AttachmentsSent), true).
			AddField("Containing links:", strconv.Itoa(stats.LinksSent), true).
			AddField("Messages this month:", strconv.Itoa(api.MsgTM(m.GuildID)), true).
			AddField("Messages this day:", strconv.Itoa(api.MsgTD(m.GuildID)), true).
			AddField("Messages this hour:", strconv.Itoa(api.MsgTH(m.GuildID)), true).
			AddField("Messages this minute:", strconv.Itoa(api.MsgTm(m.GuildID)), true)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			logger.Error(err)
		}
		return
	}
	if content[1] == "members" {
		var stats *api.GuildStats
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/stats", m.GuildID))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		} else {
			json.NewDecoder(response.Body).Decode(&stats)
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Members", m.Author.AvatarURL("")).
			SetDescription("View member stats.").
			SetColor(0xffffff).
			AddField("Total Members:", strconv.Itoa(stats.TotalMembers), false).
			AddField("Members this month:", strconv.Itoa(api.MemTM(m.GuildID)), true).
			AddField("Members this day:", strconv.Itoa(api.MemTD(m.GuildID)), true).
			AddField("Members this hour:", strconv.Itoa(api.MemTH(m.GuildID)), true).
			AddField("Members this minute:", strconv.Itoa(api.MemTm(m.GuildID)), true)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			logger.Error(err)
		}
		return
	} else {
		s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("**Error:** Member not found", content[1]))
	}
}

func init() {
	guildinfo := ch.Command{
		"guildinfo",
		"guildinfo",
		"Find out the guild stats",
		"Stats",
		false,
		map[string]bool{
			"messages": false,
			"members":  false,
		},
		false,
		true,
		guildinfoRun,
	}

	ch.Register(guildinfo)
}
