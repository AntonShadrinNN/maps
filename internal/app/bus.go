package app

import (
	"context"
	"maps/internal/entities/bus"
)

type BusRepo interface {
	CreateBus(_ context.Context, bus *bus.Bus) (id int64, err error)
	UpdateBus(_ context.Context, ID int64, num int32, depot string) (b *bus.Bus, err error)
}
