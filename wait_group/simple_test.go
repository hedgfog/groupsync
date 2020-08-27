package wait_group

import (
	"sync"
	"testing"
)

func TestSimple(t *testing.T) {
	t.Log("Hello, world")
	go func() {
		t.Log("Goroutine, world!")
	}()
	t.Log("Goodbye, world")
}

func TestSimpleFix(t *testing.T) {
	t.Log("Hello, world")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Log("Goroutine, world!")
	}()
	wg.Wait()
	t.Log("Goodbye, world")
}
