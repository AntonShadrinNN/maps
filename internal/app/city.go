package app

import (
	"context"
	"maps/internal/entities/city"
)

type CityRepo interface {
	CreateCity(_ context.Context, c *city.City) (id int64, err error)
}
