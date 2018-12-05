package commands

import (
	"fmt"
	"time"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a ready state is called
func ready(s *discordgo.Session, r *discordgo.Ready) {
	guilds := sh.SetGuilds(len(r.Guilds))
	time.Sleep(1 * time.Second)
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
