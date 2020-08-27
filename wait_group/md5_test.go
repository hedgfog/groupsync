package wait_group

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestMD5File(t *testing.T) {
	hashes, err := MD5RootParallel("/Users/hedg.fog/Downloads")
	if err != nil {
		t.Fatal(err)
	}
	for path, hash := range hashes {
		t.Logf("%v  %v\n", path, hash)
	}

}

type Result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func MD5RootParallel(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)
	//TODO context use
	resultChan, errChan := sumFiles(done, root)

	hashes := make(map[string][md5.Size]byte)
	for res := range resultChan {
		if res.err != nil {
			return nil, res.err
		}
		hashes[res.path] = res.sum
	}
	if err := <-errChan; err != nil {
		return nil, err
	}
	return hashes, nil

}

func sumFiles(done <-chan struct{}, root string) (<-chan Result, <-chan error) {
	resultChan := make(chan Result)
	errChan := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case resultChan <- Result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()
			// Abort the walk if done is closed.
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})
		go func() {
			wg.Wait()
			close(resultChan)
		}()
		// No select needed here, since errc is buffered.
		errChan <- err
	}()
	return resultChan, errChan
}

func MD5Root(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}
