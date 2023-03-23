package object

import (
	"sync"

	"github.com/opensvc/om3/core/path"
)

type (
	Dataer interface {
		Status
	}

	DataElement[T Dataer] struct {
		Path  path.T
		Value *T
	}

	// Data defines a shared holder for all objects Dataer
	Data[T Dataer] struct {
		sync.RWMutex
		data map[path.T]*T
	}
)

var (
	// StatusData is the package data holder for all objects statuses
	StatusData *Data[Status]
)

func NewData[T Dataer]() *Data[T] {
	return &Data[T]{
		data: make(map[path.T]*T),
	}
}

func (c *Data[T]) Set(p path.T, v *T) {
	c.Lock()
	defer c.Unlock()
	c.data[p] = v
}

func (c *Data[T]) Unset(p path.T) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, p)
}

func (c *Data[T]) Get(p path.T) *T {
	c.RLock()
	v := c.data[p]
	c.RUnlock()
	return v
}

func (c *Data[T]) GetAll() []DataElement[T] {
	c.RLock()
	result := make([]DataElement[T], 0)
	for p, v := range c.data {
		result = append(result, DataElement[T]{
			Path:  p,
			Value: v,
		})
	}
	c.RUnlock()
	return result
}

// InitData reset package objects data, it can be used for tests.
func InitData() {
	StatusData = NewData[Status]()
}

func init() {
	InitData()
}
