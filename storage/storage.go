package storage

type Event struct {
	Owner    string
	Category string
}

type EventStorer interface {
	Single(event Event)
	Start(event Event) string
	Stop(uuid string)
}
