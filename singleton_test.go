package singleton_test

import (
	"sync"
	"testing"

	. "github.com/weiwenchen2022/singleton"
)

type one struct {
	mu sync.Mutex
	n  int
}

func (o *one) Increment() int {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.n++
	return o.n
}

func run(t *testing.T, o *Singleton[one], c chan<- struct{}) {
	o.Instance().Increment()
	c <- struct{}{}
}

func TestSingleton(t *testing.T) {
	t.Parallel()

	var o Singleton[one]
	c := make(chan struct{})
	const N = 10

	for i := 0; i < N; i++ {
		go run(t, &o, c)
	}
	for i := 0; i < N; i++ {
		<-c
	}
	if o.Instance().n != N {
		t.Errorf("once failed outside run: %d is not 1", *o.Instance())
	}
}

func BenchmarkSingleton(b *testing.B) {
	var o Singleton[one]

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			o.Instance()
		}
	})
}
