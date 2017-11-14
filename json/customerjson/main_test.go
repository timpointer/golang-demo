package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Time struct {
	time.Time
}

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "\"\"" {
		*t = Time{time.Time{}}
		return
	}
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time{now}
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	if t.IsZero() == false {
		b = t.AppendFormat(b, timeFormart)
	}
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return t.Format(timeFormart)
}

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time   `json:"birthday"`
}

func TestTimeJson(t *testing.T) {
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	//src := `{"id":5,"name":"xiaoming","birthday":""}`
	p := new(Person)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p)
	t.Log(p.Birthday)
	js, _ := json.Marshal(p)
	fmt.Println(string(js))
}
