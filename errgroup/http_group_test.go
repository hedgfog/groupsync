package wait_group

import (
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
)

func TestHttpGroup(t *testing.T) {
	var eg errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		url := url
		eg.Go(func() error {
			_, err := http.Get(url)
			if err != nil {
				return err
			}
			//bs, err := ioutil.ReadAll((resp.Body))
			//if err != nil{
			//	return err
			//}
			//t.Log(string(bs))
			t.Log(url)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
