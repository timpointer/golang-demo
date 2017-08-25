package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_apiGetAuthList(t *testing.T) {
	authList, err := ApiGetAuthList("http://localhost:1511")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(authList)
}
func Test_apiHasAuth(t *testing.T) {
	result, err := ApiHasAuth("http://localhost:1511", "erwrw")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*result)
}

type AuthList []string

type OK string

func ApiHasAuth(host string, gcn string) (*OK, error) {
	url := host + "/hasAuth?gcn=" + gcn
	body, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	var ok OK
	err = json.Unmarshal(body, &ok)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		return nil, err
	}
	return &ok, nil
}

func ApiGetAuthList(host string) (*AuthList, error) {
	url := host + "/getAuthList"
	body, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	var authList AuthList
	err = json.Unmarshal(body, &authList)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		return nil, err
	}
	return &authList, nil
}

func httpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http get url %s status code %d", url, res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
