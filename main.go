package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/BaileyJM02/Hue/pkg/command"
	"github.com/BaileyJM02/Hue/pkg/embed"

	// "text/template"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters

var ()

const (
	prefix = "="
)

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

func sayRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]command.Command) {
	if len(content[1:]) == 0 {
		s.ChannelMessageSend(m.ChannelID, "**Error:** No arguments given (expected text)")
		return
	}
	s.ChannelMessageSend(m.ChannelID,
		fmt.Sprintf("%v", strings.Join(content[1:], " ")))
}

func helpRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]command.Command) {
	if len(content[1:]) == 0 {
		commands := ""
		for key, value := range Commands {
			commands += fmt.Sprintf("\n**%v%v**: %v", prefix, key, value.Description)
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Commands:** %v.", commands))
		return
	}
	if cmd, ok := Commands[content[1]]; ok {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("\n**%v%v** [%v]\n\nDescription: %v\nUsage: %v\nArgs required?: %v", prefix, cmd.Name, cmd.Category, cmd.Description, cmd.Usage, cmd.NeedArgs))
	}
}

func pingRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]command.Command) {
	msg, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":PONGING;%d", time.Now().UnixNano()))

	split := strings.Split(msg.Content, ";")
	if split[0] != ":PONGING" || len(split) < 2 {
		return
	}

	parsed, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		fmt.Println("err,", err)
		return
	}
	taken := time.Duration(time.Now().UnixNano() - parsed)

	started := time.Now()
	s.ChannelMessageEdit(m.ChannelID, msg.ID, "Gateway (http send -> gateway receive time): "+taken.String())
	httpPing := time.Since(started)

	s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("HTTP API: `%vms` \nGateway: `%vms`", int64(httpPing/time.Millisecond), int64(taken/time.Millisecond)))
}


func main() {
	// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	// if err != nil { panic(err) }
	// err = tmpl.Execute(os.Stdout, sweaters)
	// if err != nil { panic(err) }

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + "NDc2MTA2NjI5NjI1MTUxNTA4.Dtswww.d9U5myuKU07X6VZUD3V_pHIUXoA")
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
	Commands := make(map[string]command.Command)

	// Declare commands 
	say := command.Command{
		"say",
		"say <message>",
		"Say something",
		"General",
		true,
		sayRun,
	}

	Commands["say"] = say

	ping := command.Command{
		"ping",
		"ping",
		"ping, pong",
		"General",
		false,
		pingRun,
	}

	Commands["ping"] = ping


	help := command.Command{
		"help",
		"help <command>",
		"Help you",
		"General",
		false,
		helpRun,
	}

	Commands["help"] = help

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Fields(m.Content)

	if !(strings.Contains(content[0], prefix)) {
		return
	}

	content[0] = strings.Replace(content[0], prefix, "", -1)

	if cmd, ok := Commands[content[0]]; ok {
		fmt.Println(cmd.Name)
		cmd.Run(s, m, content, Commands)
	}

	// If the message is "ping"
	if content[0] == prefix+"ping" {
		msg, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":PONGING;%d", time.Now().UnixNano()))

		split := strings.Split(msg.Content, ";")
		if split[0] != ":PONGING" || len(split) < 2 {
			return
		}

		parsed, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			fmt.Println("err,", err)
			return
		}
		taken := time.Duration(time.Now().UnixNano() - parsed)

		started := time.Now()
		s.ChannelMessageEdit(m.ChannelID, msg.ID, "Gateway (http send -> gateway receive time): "+taken.String())
		httpPing := time.Since(started)

		s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("HTTP API: `%vms` \nGateway: `%vms`", int64(httpPing/time.Millisecond), int64(taken/time.Millisecond)))
	}

	if content[0] == prefix+"hecclp" {
		if !(len(content[1:]) == 0) {
			switch content[1] {
			case "":
				text := "ping, say"
				embed := embed.NewEmbed().
					SetTitle("Commands").
					SetDescription(text + "\n\nUse `" + prefix + "help <command>` for more help.").
					SetColor(0x00000).MessageEmbed
				s.ChannelMessageSendEmbed(m.ChannelID, embed)
			case "ping":
				text := "**\\" + prefix + "Ping**, see how long the bot takes to respond."
				embed := embed.NewEmbed().
					SetTitle("Commands - Ping").
					SetDescription(text + "\n\nUse `" + prefix + "help <command>` for more help.").
					SetColor(0x00000).MessageEmbed
				s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}
		} else {
			// text := ""
			// for k, v := range commands {
			// 	command := v
			// 	text += `{dd}`
			// }
			// embed := embed.NewEmbed().
			// SetTitle("Commands").
			// SetDescription(text + "\n\nUse `"+prefix+"help <command>` for more help.").
			// SetColor(0x00000).MessageEmbed
			// s.ChannelMessageSendEmbed(m.ChannelID, embed)
		}
	}
}
