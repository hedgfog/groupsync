package wait_group

import (
	"errors"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestReturnErr(t *testing.T) {
	var eg errgroup.Group
	eg.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("first")
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 5)
		t.Log("You wait me?")
		return errors.New("second")
	})

	if err := eg.Wait(); err != nil {
		t.Logf("err = %v", err)
	}

}
