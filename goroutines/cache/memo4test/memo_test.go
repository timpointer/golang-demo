package memo4test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"gopl.io/ch9/memo4"
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
	fmt.Println("int", int)
	if int < 20 {
		return nil, errors.New("randow error")
	}
	return body, nil
}

func Test(t *testing.T) {
	m := memo.New(httpGetBodytest)
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBodytest)
	Concurrent(t, m)
	for index := 0; index < 20; index++ {
		time.Sleep(time.Second)
		Concurrent(t, m)
	}
}
