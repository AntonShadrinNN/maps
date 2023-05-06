package city

type City struct {
	name string
}

func New(name string) *City {
	return &City{name: name}
}
