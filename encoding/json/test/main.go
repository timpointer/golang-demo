package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
 {"channel":"170400008CP","mobile":"13681953126","openid":"oJ9_hjjewiNEQkRbq69W3vUOyXfM","uid":"self-reg","unionid":"oa1QTv6n6k2kcJSJ8GtjQ6K59wDA"}
	`
	type LPMessage struct {
		Channel, Mobile, Openid, UID, Unionid string
	}
	const unbindjsonStream = `
{"membertype":"LP","openid":"oJ9_hjjewiNEQkRbq69W3vUOyXfM","outerstr":"cardpack"}
	`
	type UnbindMessage struct {
		Membertype, LP, Openid, Outerstr string
	}

	var m LPMessage
	err := parse(jsonStream, &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", m)

	/*
		dec := json.NewDecoder(strings.NewReader(jsonStream))
		for {
			var m LPMessage
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", m)
		}
	*/
}

func parse(str string, i interface{}) error {
	dec := json.NewDecoder(strings.NewReader(str))
	for {
		if err := dec.Decode(&i); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}
