package bus

import "time"

type Bus struct {
	Number int32
	Depot  string
	CDate  time.Time
	UDate  time.Time
}

func New(n int32, dep string) *Bus {
	return &Bus{
		Number: n,
		Depot:  dep,
		CDate:  time.Now(),
		UDate:  time.Now(),
	}
}
