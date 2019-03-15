package logger

import (
	"fmt"
	"time"

	"github.com/BaileyJM02/Hue/pkg/embed"
	"github.com/bwmarrin/discordgo"
)

const (
	token   = "NTI0Mjc0NTA2NDE2NjUyMjk2.DwK4vg.NK9Eh3MEje3Ntff_tDW4tC91NCM"
	channel = "524275582079336448"
)

var DG *discordgo.Session
var err error

func init() {
	DG, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session. No logs can be made.", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = DG.Open()
	if err != nil {
		fmt.Println("Error opening connection. No logs can be made.", err)
		return
	}
	Info("Log Bot is now running.")
}

func Info(text string) string {
	embed := embed.NewEmbed().
		SetDescription(fmt.Sprintf("**%v** %v", string(time.Now().Format("15:04:05")), text)).
		SetColor(0x3FA1E3).MessageEmbed

	DG.ChannelMessageSendEmbed(channel, embed)
	fmt.Println(string(time.Now().Format("15:04:05")), text)
	return text
}

func Error(er interface{}) string {
	embed := embed.NewEmbed().
		SetDescription(fmt.Sprintf("**%v** %v", string(time.Now().Format("15:04:05")), er)).
		SetColor(0xE74C3C).MessageEmbed

	DG.ChannelMessageSendEmbed(channel, embed)
	fmt.Println(string(time.Now().Format("15:04:05")), er)
	return fmt.Sprintf("%v", er)
}
