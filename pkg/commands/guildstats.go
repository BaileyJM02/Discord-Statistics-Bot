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
	"github.com/bwmarrin/discordgo"
)

func guildstatsRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	var guild *api.Guild
	var totalAttachments int
	var totalLinks int
	if len(content[1:]) < 1 {
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v", m.GuildID))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			json.NewDecoder(response.Body).Decode(&guild)
		}
		if guild == nil {

		}
		for _, v := range guild.Data.Members {
			totalAttachments += v.AttachmentsSent
		}
		for _, v := range guild.Data.Members {
			totalLinks += v.LinksSent
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Messages", m.Author.AvatarURL("")).
			SetDescription("View general stats.").
			SetColor(0xffffff).
			AddField("Total Members:", strconv.Itoa(guild.Stats.TotalMembers), false).
			AddField("Total Messages:", strconv.Itoa(api.MsgAT(guild)), false).
			AddField("Messages containing attachments:", strconv.Itoa(totalAttachments), false).
			AddField("Messages containing links:", strconv.Itoa(totalLinks), false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if content[1] == "messages" {
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v", m.GuildID))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			json.NewDecoder(response.Body).Decode(&guild)
		}
		if guild == nil {

		}
		for _, v := range guild.Data.Members {
			totalAttachments += v.AttachmentsSent
		}
		for _, v := range guild.Data.Members {
			totalLinks += v.LinksSent
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Messages", m.Author.AvatarURL("")).
			SetDescription("View message stats.").
			SetColor(0xffffff).
			AddField("Total Messages:", strconv.Itoa(api.MsgAT(guild)), false).
			AddField("Containing attachments:", strconv.Itoa(totalAttachments), true).
			AddField("Containing links:", strconv.Itoa(totalLinks), true).
			AddField("Messages this month:", strconv.Itoa(api.MsgTM(guild)), true).
			AddField("Messages this day:", strconv.Itoa(api.MsgTD(guild)), true).
			AddField("Messages this hour:", strconv.Itoa(api.MsgTH(guild)), true).
			AddField("Messages this minute:", strconv.Itoa(api.MsgTm(guild)), true)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if content[1] == "members" {
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v", m.GuildID))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			json.NewDecoder(response.Body).Decode(&guild)
		}
		if guild == nil {

		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Members", m.Author.AvatarURL("")).
			SetDescription("View member stats.").
			SetColor(0xffffff).
			AddField("Total Members:", strconv.Itoa(guild.Stats.TotalMembers), false).
			AddField("Members this month:", strconv.Itoa(api.MemTM(guild)), true).
			AddField("Members this day:", strconv.Itoa(api.MemTD(guild)), true).
			AddField("Members this hour:", strconv.Itoa(api.MemTH(guild)), true).
			AddField("Members this minute:", strconv.Itoa(api.MemTm(guild)), true)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

func init() {
	guildstats := ch.Command{
		"guildstats",
		"guildstats",
		"Find out the guild stats",
		"Stats",
		false,
		map[string]bool{
			"messages": false,
			"members":  false,
		},
		false,
		true,
		guildstatsRun,
	}

	ch.Register(guildstats)
}
