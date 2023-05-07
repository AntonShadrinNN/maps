package bus

import "time"

type Bus struct {
	ID     int64
	Number int32
	Depot  string
	CDate  time.Time
	UDate  time.Time
}

func New(n int32, dep string) *Bus {
	return &Bus{
		Number: n,
		Depot:  dep,
		CDate:  time.Now().UTC(),
		UDate:  time.Now().UTC(),
	}
}
