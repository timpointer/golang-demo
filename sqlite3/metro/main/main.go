package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"flag"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteConnStr = "./data/metroreport.db?cache=shared&mode=rwc"
)

type Config struct {
	starttime string
	endtime   string
	store     string
	db        *sql.DB
	start     *time.Time
	end       *time.Time
	command   string
}

var config *Config

func init() {
	config = &Config{}
	config.start = &time.Time{}
	config.end = &time.Time{}
}

func main() {
	flag.StringVar(&config.command, "command", "default", "choose which command do you want to execute")
	flag.StringVar(&config.store, "store", "shanghai", "store params")
	flag.StringVar(&config.starttime, "starttime", "20170113", "starttime params")
	flag.StringVar(&config.endtime, "endtime", "20170223", "endtime params")
	flag.Parse()

	const shortForm = "20060102"
	start, err := time.Parse(shortForm, config.starttime)
	log.Printf("starttime %d\n", start.Unix())
	if err != nil {
		log.Fatal(err)
	}

	*config.end, err = time.Parse(shortForm, config.endtime)
	log.Printf("endtime %d\n", config.end.Unix())
	if err != nil {
		log.Fatal(err)
	}

	config.db, err = sql.Open("sqlite3", sqliteConnStr)
	if err != nil {
		log.Fatal(err)
	}
	//init database tables
	sqlStmts := []string{
		"create table if not exists user_registration (userid  INTEGER,name  TEXT,store  TEXT,channel  TEXT,cardholder  TEXT,campaign  TEXT,date  INTEGER);",
	}
	for _, sqlStmt := range sqlStmts {
		_, err := config.db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			route(config)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	route(config)
}

func route(config *Config) {
	log.Println("command", config.command)
	switch config.command {
	case "insert":
		// insert dump data to sqlite
		insertDumpData(config.db)
	case "collectiondata":
		// collection data from sqlite
		err := collectionReportData(config.db, *config.start, *config.end)
		if err != nil {
			log.Fatal(err)
			return
		}

	case "report":
		report(config.db, *config.start, *config.end, config.store)
	case "reporttemplate":
		tmpl := template.Must(template.ParseFiles("./template/metro.html"))
		buf := &bytes.Buffer{}
		err := tmpl.Execute(buf, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(buf)
	case "reportURC":
		// get report for user new registration
		log.Println("reportURC")
		rows, err := reportURC(config.db, config.starttime, config.endtime, config.store)
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
	case "reportytd":
		reportYTD(config.db, config.store)
	case "reportcw":
		reportCW(config.db, 1, config.store)
	}

}

var stores []string
var channels []string
var cardholders []string
var activites []string
var t time.Time

func init() {
	stores = []string{"shanghai", "beijing", "hangzhou"}
	channels = []string{"local", "web", "wechat", "ali"}
	cardholders = []string{"MP", "NMP"}
	activites = []string{"activity1", "activity2", "activity3", "activity4"}
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
	log.Println(nt)
	return nt.Unix()
}
