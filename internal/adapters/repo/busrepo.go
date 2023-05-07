package repo

import (
	"context"
	"maps/internal/entities/bus"
	"maps/internal/errs"
	"sync"
	"time"
)

type BusRepo struct {
	storage map[int64]*bus.Bus
	mx      *sync.Mutex
	lastID  int64
}

func (br *BusRepo) CreateBus(_ context.Context, bus *bus.Bus) (int64, error) {
	br.mx.Lock()
	defer br.mx.Unlock()

	if _, ok := br.storage[br.lastID]; ok {
		return -1, errs.BusStorageOverflowError
	}

	bus.ID = br.lastID
	br.lastID++

	return br.lastID - 1, nil
}

func (br *BusRepo) UpdateBus(_ context.Context, ID int64, n int32, dep string) (*bus.Bus, error) {
	br.mx.Lock()
	defer br.mx.Unlock()

	b, ok := br.storage[ID]
	if !ok {
		return nil, errs.BusNotExistsError
	}
	b.Number = n
	b.Depot = dep
	b.UDate = time.Now().UTC()

	return b, nil
}

func New() *BusRepo {
	return &BusRepo{
		storage: make(map[int64]*bus.Bus),
		mx:      &sync.Mutex{},
		lastID:  0,
	}
}
