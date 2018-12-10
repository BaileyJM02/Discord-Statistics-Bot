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

func messagesRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
	var stats *api.Member
	if len(content[1:]) == 0 {
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/member/%v", m.GuildID, m.Author.ID))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			json.NewDecoder(response.Body).Decode(&stats)
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Messages", m.Author.AvatarURL("")).
			SetDescription("View message count.").
			SetColor(0xffffff).
			AddField("Total Messages:", strconv.Itoa(stats.MessagesSent), false).
			AddField("Messages containing attachments:", strconv.Itoa(stats.AttachmentsSent), false).
			AddField("Messages containing links:", strconv.Itoa(stats.LinksSent), false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if len(content[1:]) > 0 {
		if len(m.Mentions) == 0 {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Error:** Please mention a user or use no arguments."))
			return
		}
		response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/member/%v", m.GuildID, m.Mentions[0].ID))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			json.NewDecoder(response.Body).Decode(&stats)
		}
		em := embed.NewEmbed().
			SetAuthor(m.Author.Username+" | Messages", m.Author.AvatarURL("")).
			SetDescription("View message count.").
			SetColor(0xffffff).
			AddField("Total Messages:", strconv.Itoa(stats.MessagesSent), false).
			AddField("Messages containing attachments:", strconv.Itoa(stats.AttachmentsSent), false).
			AddField("Messages containing links:", strconv.Itoa(stats.LinksSent), false)

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Error:** Command \"%v\" not found", content[1]))
	}
}

func init() {
	messages := ch.Command{
		"messages",
		"Messages [user]",
		"Find out message count",
		"Stats",
		false,
		map[string]bool{
			"user": false,
		},
		false,
		true,
		messagesRun,
	}

	ch.Register(messages)
}
