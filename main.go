package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
	_ "github.com/BaileyJM02/Hue/pkg/commands"
	"github.com/BaileyJM02/Hue/pkg/embed"
	// "text/template"

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
		AddField("Error:", fmt.Sprintf("```bash\n%v\n```", er)).
		AddField("You may be able to fix this by following these instructions:", fmt.Sprintf("%v.", reason)).
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

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	prefix := sh.GetPrefix()

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Fields(m.Content)

	if !(strings.Contains(content[0], prefix)) {
		return
	}

	content[0] = strings.Replace(content[0], prefix, "", -1)
	
	if cmd, ok := ch.Commands[content[0]]; ok {
		if (cmd.OwnerOnly && (m.Author.ID != "398197113495748626")) {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Error:** You're not the owner!"))
			return
		}
		fmt.Println("Command Run: ",cmd.Name)
		cmd.Run(s, m, content, ch.Commands, sh.Bot)
		return
	}

				// text := "**\\" + prefix + "Ping**, see how long the bot takes to respond."
				// embed := embed.NewEmbed().
				// 	SetTitle("Commands - Ping").
				// 	SetDescription(text + "\n\nUse `" + prefix + "help <command>` for more help.").
				// 	SetColor(0x00000).MessageEmbed
				// s.ChannelMessageSendEmbed(m.ChannelID, embed)
}