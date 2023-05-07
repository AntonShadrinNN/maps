package city

type City struct {
	ID   int64
	name string
}

func New(name string) *City {
	return &City{name: name}
}
