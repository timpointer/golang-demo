package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var data = `
{"invoiceSN":"027270","custKey":788675,"storeKey":14,"szCODInvoiceType":"0","cardholderKey":4,"szCODDocumentNumber":"","amountToPay":624.15,"time":"153434","date":"20170510","rules":[{"qty":624.15,"total":14,"rule":"lifecycle"}],"amountCollected":624.15,"homeStoreKey":14}
`

func main() {
	tryError("senddata", func() error {
		return SendData("http://admin:admin@localhost:1512/record", data)
	})
}

func SendData(url string, data string) error {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		log.Printf("http.post %s:%v", url, err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("http.post readAll %s: %v", url, err)
		return err
	}
	if resp.StatusCode != 200 {
		log.Printf("http.post statuscode url %s, status %d, body %s", url, resp.StatusCode, string(body))
		return errors.New("status is not 200")
	}
	return nil
}

func tryError(name string, f func() error) error {
	const timeout = 60 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		err := f()
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding(%s); retrying...", err)
		time.Sleep(time.Second << uint(tries) * 2) //exponential back-off
	}
	return fmt.Errorf("server failed to respond after %s", name, timeout)
}
