package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	eh "github.com/BaileyJM02/Hue/pkg/eventHandler"

	// populate Events map[]
	"github.com/BaileyJM02/Hue/pkg/embed"
	_ "github.com/BaileyJM02/Hue/pkg/events"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters

func init() {
	// flag.StringVar(&Token, "t", "", "Bot Token")
	// flag.Parse()
}

func error(session *discordgo.Session, channelid string, er string, reason string) {
	embed := embed.NewEmbed().
		SetTitle("There was an error while performing this action!").
		SetDescription("Something went wrong! You can try running the command again but if the error persists, please report it in the [official server](https://invite.gg/hue).").
		AddField("Error:", fmt.Sprintf("```bash\n%v\n```", er), false).
		AddField("You may be able to fix this by following these instructions:", fmt.Sprintf("%v.", reason), false).
		SetColor(0xF44336).MessageEmbed

	session.ChannelMessageSendEmbed(channelid, embed)
}

func main() {
	// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	// if err != nil { panic(err) }
	// err = tmpl.Execute(os.Stdout, sweaters)
	// if err != nil { panic(err) }

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(sh.GetToken())
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Register events.
	for event := range eh.Events {
		dg.AddHandler(eh.Events[event].Run)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Close()
}
