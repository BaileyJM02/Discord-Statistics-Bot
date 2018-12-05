package clientHandler

import (
	cf "github.com/BaileyJM02/Hue/pkg/configHandler"
)

type Client struct {
	Prefix      string
	Token       string
	Description string
	Ready       bool
	Guilds      int
}

var (
	Bot Client = Client{
		cf.Config.Prefix,
		"Bot " + cf.Config.Token,
		"A bot called Hue.",
		false,
		0,
	}
)

func GetPrefix() string {
	return Bot.Prefix
}

func GetToken() string {
	return Bot.Token
}

func GetDescription() string {
	return Bot.Description
}

func GetReady() bool {
	return Bot.Ready
}

func ReadyUp() bool {
	Bot.Ready = true
	return Bot.Ready
}

func GetGuilds() int {
	return Bot.Guilds
}

func SetGuilds(i int) int {
	Bot.Guilds = i
	return Bot.Guilds
}
