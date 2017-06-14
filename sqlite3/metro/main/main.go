package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"flag"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteConnStr = "./data/metroreport.db?cache=shared&mode=rwc"
)

func main() {
	command := flag.String("command", "default", "choose which command do you want to execute")
	store := flag.String("store", "shanghai", "store params")
	starttime := flag.String("starttime", "20140113", "starttime params")
	endtime := flag.String("endtime", "20170223", "endtime params")
	flag.Parse()

	const shortForm = "20060102"
	start, err := time.Parse(shortForm, *starttime)
	log.Printf("starttime %d\n", start.Unix())
	if err != nil {
		log.Fatal(err)
	}

	end, err := time.Parse(shortForm, *endtime)
	log.Printf("endtime %d\n", end.Unix())
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", sqliteConnStr)
	if err != nil {
		log.Fatal(err)
	}
	//init database tables
	sqlStmts := []string{
		"create table if not exists user_registration (userid  INTEGER,name  TEXT,store  TEXT,channel  TEXT,cardholder  TEXT,campaign  TEXT,date  INTEGER);",
	}
	for _, sqlStmt := range sqlStmts {
		_, err := db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}

	log.Println("command", *command)
	switch *command {
	case "insert":
		log.Println("insert db")
		// insert dump data to sqlite
		insertDumpData(db)
	case "collectiondata":
		// collection data from sqlite
		err := collectionReportData(db, start, end)
		if err != nil {
			log.Fatal(err)
			return
		}

	case "report":
		report(db, start, end, *store)
	case "reportURC":
		// get report for user new registration
		log.Println("reportURC")
		rows, err := reportURC(db, *starttime, *endtime, *store)
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
		reportYTD(db, *store)
	case "reportcw":
		reportCW(db, 1, *store)
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
