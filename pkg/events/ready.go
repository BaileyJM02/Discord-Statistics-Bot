package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	api "github.com/BaileyJM02/Hue/pkg/apiHandler"
	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/BaileyJM02/Hue/pkg/logger"
	"github.com/DiscordBotList/dblgo"
	"github.com/bwmarrin/discordgo"
)

var ugh = 0

func getStatus(n int) discordgo.UpdateStatusData {
	return discordgo.UpdateStatusData{
		&ugh,
		&discordgo.Game{
			fmt.Sprintf("%v servers | -help", n),
			discordgo.GameTypeWatching,
			"",
			"",
			"",
			discordgo.TimeStamps{
				time.Now().UnixNano() / int64(time.Millisecond),
				time.Now().AddDate(0, 0, 2).UnixNano() / int64(time.Millisecond),
			},
			discordgo.Assets{
				"",
				"",
				"",
				"",
			},
			"",
			0,
		},
		false,
		"",
	}
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func updateDBL(t time.Time) {
	stats := make(map[string]int)
	response, err := http.Get("https://hue.observer/api/stats")
	if err != nil {
		logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	} else {
		json.NewDecoder(response.Body).Decode(&stats)
	}
	dblAPI := &http.Client{}
	data := make(map[string]int)
	data["guilds"] = stats["guilds"]
	data["users"] = stats["members"]
	dataJSON, _ := json.Marshal(data)
	postData := []byte(dataJSON)
	req, err := http.NewRequest("POST", "https://discordbotlist.com/api/bots/476106629625151508/stats", bytes.NewReader(postData))
	if err != nil {
		logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	}
	req.Header.Add("Authorization", "Bot eebd0d679ca2da68bf781b562968d7e38c53070933e51660c75645aa057c6ac6")
	resp, err := dblAPI.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
	respo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}
	_ = respo

	dblorg := dblgo.NewDBL("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQ3NjEwNjYyOTYyNTE1MTUwOCIsImJvdCI6dHJ1ZSwiaWF0IjoxNTQ2NzgyMTY1fQ.g7I54DYni7glDBxPf-PwddjY5_hYgF7Ra2YRi9SdpDI", "476106629625151508")
	err = dblorg.PostStats(stats["guilds"])
	if err != nil {
		logger.Error(err)
	}
}

// This function will be called (due to AddHandler above) every time a ready state is called
func runReady(s *discordgo.Session, r *discordgo.Ready) {
	time.Sleep(4 * time.Second)
	jsonValue, _ := json.Marshal(r)
	response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/up-to-date/"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		logger.Info(string(data))
	}

	time.Sleep(4 * time.Second)

	for _, guild := range r.Guilds {
		logger.Info(fmt.Sprintf("Checking %v (%v)", guild.Name, guild.ID))
		g, err := s.Guild(guild.ID)
		if err != nil {
			logger.Error(err)
		}

		var members []*discordgo.Member
		var m []*discordgo.Member
		lastm := ""
		count := 1
		hitzero := 0

		for count < 250 && hitzero < 2 {
			m, err = s.GuildMembers(guild.ID, lastm, 1000)
			if err != nil {
				logger.Error(fmt.Sprintf("GuildMembers: %s\n", err))
			}
			if len(m) > 0 {
				lastm = m[len(m)-1].User.ID
			} else {
				hitzero++
			}
			members = append(members, m...)
			count++
		}
		needsCheck := api.CheckHelper{
			g,
			members,
		}

		jsonValue, _ := json.Marshal(needsCheck)
		_, err = http.Post(fmt.Sprintf("http://localhost:8000/db/check/"), "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			logger.Error(fmt.Sprintf("The HTTP request failed with error %s\n", err))
		}
		logger.Info(fmt.Sprintf("Checking finished %v (%v)", guild.Name, guild.ID))
	}

	guilds := sh.SetGuilds(len(r.Guilds))
	time.Sleep(3 * time.Second)
	sh.ReadyUp()
	currentTime := time.Now()
	s.ChannelMessageSend("519965512751382549", fmt.Sprintf("[%v] Hue is ready, serving %v guilds, using version: %v.", currentTime.Format("15:04:05"), guilds, r.Version))

	s.UpdateStatusComplex(getStatus(len(r.Guilds)))

	go doEvery(30*time.Minute, updateDBL)
}

func init() {
	Ready := eh.Event{
		"ready",
		true,
		runReady,
	}

	eh.Register(Ready)
}
