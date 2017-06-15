package main

import (
	"fmt"
	"math/rand"
	"time"

	ttime "github.com/timpointer/golang-demo/time"
)

var days []string

func init() {
	tstart := time.Date(2014, 4, 10, 23, 0, 0, 0, time.UTC)
	tend := time.Date(2017, 5, 10, 23, 0, 0, 0, time.UTC)
	days = ttime.GetListDay(tstart, tend)
}
func main() {
	t := time.Now()
	rows := createDate()
	fmt.Println("start:", time.Since(t))
	//stores := reduce(rows, "store")
	//for key := range stores {
	//	fmt.Println(key)
	//}
	fmt.Println("done:", time.Since(t))
	nrows := countReduce(rows)
	//for _, value := range nrows {
	//fmt.Println(value)
	//}
	fmt.Println("len:", len(nrows))
	fmt.Println("end:", time.Since(t))
}

func countReduce(rows []*dataRigistrationCount) []*dataRigistrationCWCount {
	set := map[string]*dataRigistrationCWCount{}
	for _, row := range rows {
		hash, err := hashCW(row)
		if err != nil {
			return nil
		}
		_, ok := set[hash]
		if ok == true {
			set[hash].Count += row.Count
		} else {
			set[hash] = &dataRigistrationCWCount{
				dataOption{
					row.Store, row.Channel, row.Cardholder, row.Campaign, row.Count,
				},
				hash[:6],
			}
		}
	}
	values := []*dataRigistrationCWCount{}
	for _, value := range set {
		values = append(values, value)
	}
	return values
}

func hashCW(r *dataRigistrationCount) (string, error) {
	date, err := hashDateCW(r.Date)
	if err != nil {
		return "", err
	}
	return date + r.Campaign + r.Cardholder + r.Channel + r.Store, nil
}

func hashDateCW(d string) (string, error) {
	t, err := parseTime(d)
	if err != nil {
		return "", err
	}
	y, w := t.ISOWeek()
	return fmt.Sprintf("%04d%02d", y, w), nil
}

const shortForm = "20060102"

func parseTime(timestr string) (time.Time, error) {
	return time.Parse(shortForm, timestr)
}

func reduce(rows []*dataRigistrationCount, field string) map[string]bool {
	set := map[string]bool{}
	for _, row := range rows {
		switch field {
		case "store":
			set[row.Store] = true
		case "channel":
			set[row.Channel] = true
		case "cardholder":
			set[row.Cardholder] = true
		case "campaign":
			set[row.Campaign] = true
		}
	}
	return set
}

func createDate() []*dataRigistrationCount {
	rows := []*dataRigistrationCount{}
	for i := 0; i < 1000000; i++ {
		row := &dataRigistrationCount{
			dataOption{
				chooseStore(), chooseChannel(), chooseCardholder(), chooseActivity(), 1}, chooseRandomTime(),
		}
		rows = append(rows, row)
	}
	return rows
}

var stores []string
var channels []string
var cardholders []string
var activites []string
var t time.Time

func init() {
	stores = []string{"shanghai", "beijing", "hangzhou", "default"}
	channels = []string{"local", "web", "wechat", "ali", "default"}
	cardholders = []string{"MP", "NMP", "default"}
	activites = []string{"activity1", "activity2", "activity3", "activity4", "default"}
	t = time.Now()
	rand.Seed(42) // Try changing this number!
}

func chooseStore() string {
	return stores[rand.Intn(len(stores))]
}

func chooseChannel() string {
	return channels[rand.Intn(len(channels))]
}

func chooseCardholder() string {
	return cardholders[rand.Intn(len(cardholders))]
}

func chooseActivity() string {
	return activites[rand.Intn(len(activites))]
}

func chooseRandomTime() string {
	return days[rand.Intn(len(days))]
}

func chooseTime() int64 {
	nt := t.AddDate(0, 0, rand.Intn(1000)-1000)
	return nt.Unix()
}

type dataOption struct {
	Store      string `json:"store"`
	Channel    string `json:"channel"`
	Cardholder string `json:"cardholder"`
	Campaign   string `json:"campaign"`
	Count      int    `json:"count"`
}

type dataRigistrationCount struct {
	dataOption
	Date string `json:"date"`
}

type dataRigistrationCWCount struct {
	dataOption
	DateCW string `json:"datecw"`
}
