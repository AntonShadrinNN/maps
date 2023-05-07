package repo

import (
	"context"
	"maps/internal/entities/city"
	"maps/internal/errs"
	"sync"
)

type CityRepo struct {
	storage map[int64]*city.City
	mx      *sync.Mutex
	lastID  int64
}

func (cr *CityRepo) CreateCity(_ context.Context, c *city.City) (int64, error) {
	cr.mx.Lock()
	defer cr.mx.Unlock()

	if _, ok := cr.storage[cr.lastID]; ok {
		return -1, errs.CityStorageOverflowError
	}

	c.ID = cr.lastID
	cr.storage[cr.lastID] = c
	cr.lastID++

	return cr.lastID - 1, nil
}
