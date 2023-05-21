// Package singleton provides Singleton type that implements the singleton design pattern.
package singleton

import "sync"

// Singleton implements the singleton design pattern.
// A Singleton must not be copied after first use.
type Singleton[T any] struct {
	// New optionally specifies a function to create
	// a instance when Instance being called for the first time.
	// It may not be changed concurrently with calls to Instance.
	New func() *T

	instance *T
	once     sync.Once
}

func (s *Singleton[T]) Instance() *T {
	s.once.Do(func() {
		if s.New != nil {
			s.instance = s.New()
		} else {
			s.instance = new(T)
		}
	})

	return s.instance
}
