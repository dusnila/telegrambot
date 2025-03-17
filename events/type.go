package events

type Fether interface {
	Feth(limit int) ([]Event, error)
}

type Processor interface {
	Processor(e Event) error
}

type Type int

const (
	Unknow Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
