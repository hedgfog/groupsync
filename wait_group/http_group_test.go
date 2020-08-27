package wait_group

import (
	"net/http"
	"sync"
	"testing"
)

func TestHttpGroup(t *testing.T) {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				panic(err) // this is very bad
			}
			//bs, err := ioutil.ReadAll((resp.Body))
			//if err != nil{
			//	panic(err) // this is very bad
			//}
			//t.Log(string(bs))
			t.Log(url)
		}(url)
	}
	// Ожидаем пока все HTTP запросы завершатся.
	wg.Wait()
}
