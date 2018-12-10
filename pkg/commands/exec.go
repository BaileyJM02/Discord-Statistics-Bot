package commands

// import (
// 	"fmt"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"

// 	sh "github.com/BaileyJM02/Hue/pkg/clientHandler"
// 	ch "github.com/BaileyJM02/Hue/pkg/commandHandler"
// 	"github.com/bwmarrin/discordgo"
// )

// func execRun(s *discordgo.Session, m *discordgo.MessageCreate, content []string, Commands map[string]ch.Command, client sh.Client) {
// 	parsed := time.Now().UnixNano()
// 	// here we perform the pwd command.
// 	// we can store the output of this in our out variable
// 	// and catch any errors in err
// 	out, err := exec.Command("sh", "-c", strings.Join(content[1:], " ")).Output()

// 	// if there is an error with our execution
// 	// handle it here
// 	if err != nil {
// 		takenErr := time.Duration(time.Now().UnixNano() - parsed)
// 		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error => `%.2fs`, Runtime: `%v` ```bash\n%v```", takenErr.Seconds(), strings.Title(runtime.GOOS), err.Error()))
// 	}

// 	// as the out variable defined above is of type []byte we need to convert
// 	// this to a string or else we will see garbage printed out in our console
// 	// this is how we convert it to a string
// 	output := string(out[:])

// 	// once we have converted it to a string we can then output it.
// 	taken := time.Duration(time.Now().UnixNano() - parsed)
// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Result => `%.2fs`, Runtime: `%v` ```bash\n%v```", taken.Seconds(), strings.Title(runtime.GOOS), output))
// }

// func init() {
// 	exec := ch.Command{
// 		"exec",
// 		"exec <expression>",
// 		"Ugh idk tbh",
// 		"Admin",
// 		true,
// 		map[string]bool{
// 			"expression": true,
// 		},
// 		true,
// 		false,
// 		execRun,
// 	}

// 	ch.Register(exec)
// }
