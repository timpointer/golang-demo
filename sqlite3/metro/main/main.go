package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"os"
	"time"

	"fmt"

	"net/url"

	_ "github.com/mattn/go-sqlite3"
)

const (
	shortForm     = "20060102"
	sqliteConnStr = "./data/metroreport.db?cache=shared&mode=rwc"
)

type Config struct {
	db *sql.DB
}

var config *Config

func init() {
	config = &Config{}
	var err error
	config.db, err = getWriteDB(sqliteConnStr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	values := parseQuery()
	route(config, values)
}

func parseQuery() url.Values {
	values := url.Values{}
	if len(os.Args) > 0 {
		query := os.Args[1]
		var err error
		values, err = url.ParseQuery(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	return values
}

func parseTime(timestr string) (time.Time, error) {
	return time.Parse(shortForm, timestr)
}

func route(config *Config, values url.Values) {
	command := values.Get("command")
	log.Println("command", command)
	switch command {
	case "insert":
		// insert dump data to sqlite
		err := insertDumpData()
		if err != nil {
			log.Fatal(err)
			return
		}
	case "collectiondata":
		// collection data from sqlite
		starttime, endtime := t, t
		start := values.Get("start")
		if start != "" {
			var err error
			starttime, err = parseTime(start)
			if err != nil {
				log.Fatalf("parseTime start %v", err)
				return
			}
		}
		end := values.Get("end")
		if end != "" {
			var err error
			endtime, err = parseTime(end)
			if err != nil {
				log.Fatalf("parseTime end %v", err)
				return
			}
		}

		err := collectionReportData(config.db, starttime, endtime)
		if err != nil {
			log.Fatal(err)
			return
		}

	case "reporttemplate":
		tmpl := template.Must(template.ParseFiles("./template/metro.html"))
		buf := &bytes.Buffer{}
		err := tmpl.Execute(buf, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(buf)
	case "reportURCCW":
		rows, err := reportURCCW(config.db, 1)
		if err != nil {
			log.Fatal(err)
			return
		}

		data, err := json.Marshal(rows)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Print(string(data))

	case "reportURC":
		// get report for user new registration
		rows, err := reportURC(config.db, "", "")
		if err != nil {
			log.Fatal(err)
			return
		}

		data, err := json.Marshal(rows)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Print(string(data))
	// deprecated
	case "report":
		report(config.db, t, t)
	case "reportytd":
		reportYTD(config.db)
	case "reportcw":
		reportCW(config.db, 1)
	}

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

func chooseTime() int64 {
	nt := t.AddDate(0, 0, rand.Intn(1000)-1000)
	return nt.Unix()
}
