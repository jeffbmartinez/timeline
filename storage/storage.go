package storage

type Event struct {
	Owner    string `json:"owner"`
	Category string `json:"category"`
}

type EventStorer interface {
	Single(event Event)
	Start(event Event) string
	Stop(uuid string)
}
