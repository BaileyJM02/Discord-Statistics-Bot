package eventHandler

type Event struct {
	Name    string
	Enabled bool
	Run     interface{}
}

var (
	Events map[string]Event
)

func Register(evt Event) {
	if Events == nil {
		Events = make(map[string]Event)
	}
	Events[evt.Name] = evt
}
