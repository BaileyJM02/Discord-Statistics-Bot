package clientHandler

type Client struct {
	Prefix      string
	Token       string
	Description string
}

var (
	Bot Client = Client{
		"=",
		"Bot " + "NDc2MTA2NjI5NjI1MTUxNTA4.Dtswww.d9U5myuKU07X6VZUD3V_pHIUXoA",
		"A bot called Hue.",
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
