package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	"flag"

	randomdata "github.com/Pallinder/go-randomdata"
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteConnStr = "./data/metroreport.db?cache=shared&mode=rwc"
)

func main() {
	command := flag.String("command", "default", "choose which command do you want to execute")
	flag.Parse()

	db, err := sql.Open("sqlite3", sqliteConnStr)
	if err != nil {
		log.Fatal(err)
	}
	//init database tables
	sqlStmts := []string{
		"create table if not exists user_registration (userid  INTEGER,name  TEXT,storepanel  TEXT,channel  TEXT,cardholder  TEXT,campaign  TEXT,date  INTEGER);",
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

		for i := 0; i < 10; i++ {
			_, err := insert(db, 1, randomdata.SillyName(), chooseStore(), chooseChannel(), chooseCardholder(), chooseActivity(), 1232342)
			if err != nil {
				log.Printf("insert: %s\n", err)
				return
			}
		}

	}
}

func insert(db *sql.DB, userid int, name, storepanel, channel, cardholder, campaign string, data int) (sql.Result, error) {
	return db.Exec("INSERT INTO user_registration  (userid ,name  ,storepanel  ,channel ,cardholder ,campaign  ,date  )VALUES (?,?,?,?,?,?,?);",
		userid, name, storepanel, channel, cardholder, campaign, data)
}

var stores []string
var channels []string
var cardholders []string
var activites []string
var t time.Time

func init() {
	stores = []string{"storeshanghai", "storebeijing", "storehangzhou"}
	channels = []string{"local", "web", "wechat", "ali"}
	cardholders = []string{"MP", "NMP"}
	activites = []string{"activity1", "activity2", "activity3", "activity4"}
	t = time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC)
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
