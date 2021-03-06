package memo4test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var httpGetBodytest = func(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	int := rand.Intn(100)
	if int < 50 {
		return nil, fmt.Errorf("url:%s ,%v", url, "internal error")
	}
	return body, nil
}

func Test(t *testing.T) {
	m := New(httpGetBodytest)
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBodytest)
	Concurrent(t, m)
	for index := 0; index < 20; index++ {
		time.Sleep(time.Second)
		Concurrent(t, m)
	}
}
