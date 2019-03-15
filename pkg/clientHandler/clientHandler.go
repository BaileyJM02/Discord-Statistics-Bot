package clientHandler

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"

	cf "github.com/BaileyJM02/Hue/pkg/configHandler"
	"github.com/BaileyJM02/Hue/pkg/logger"
)

type Client struct {
	Prefix      string
	Token       string
	Description string
	Ready       bool
	Guilds      int
	Uptime      time.Time
}

var (
	Bot Client = Client{
		cf.Config.Prefix,
		"Bot " + cf.Config.Token,
		"A bot called Hue.",
		false,
		0,
		time.Now(),
	}
)

func GetPrefix(id string) string {
	var prefix string
	response, err := http.Get(fmt.Sprintf("http://localhost:8000/db/guild/%v/prefix", id))
	if err != nil {
		logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		return "-"
	}
	json.NewDecoder(response.Body).Decode(&prefix)
	return prefix
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
	Bot.Uptime = time.Now()
	return Bot.Ready
}

func GetGuilds() int {
	return Bot.Guilds
}

func GetUptime() time.Duration {
	uptime := time.Now().Sub(Bot.Uptime)
	return uptime
}

func SetGuilds(i int) int {
	Bot.Guilds = i
	return Bot.Guilds
}
