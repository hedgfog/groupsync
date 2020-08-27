package wait_group

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncTest(t *testing.T) {
	count := 100
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go work(i, &wg)
	}
	wg.Wait()
	t.Log("Successfully!")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(id)
}
