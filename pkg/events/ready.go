package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func ready(s *discordgo.Session, r *discordgo.Ready) {
	jsonValue, _ := json.Marshal(r)
	response, err := http.Post(fmt.Sprintf("http://localhost:8000/db/up-to-date/"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	guilds := sh.SetGuilds(len(r.Guilds))
	time.Sleep(3 * time.Second)
	ready := sh.ReadyUp()
	currentTime := time.Now()
	s.ChannelMessageSend("519965512751382549", fmt.Sprintf("[%v] Hue is ready (%v), serving %v guilds, using version: %v.", currentTime.Format("15:04:05"), ready, guilds, r.Version))

}

func init() {
	Ready := eh.Event{
		"ready",
		true,
		ready,
	}

	eh.Register(Ready)
}
