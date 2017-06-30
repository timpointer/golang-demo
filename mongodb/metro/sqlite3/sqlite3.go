package sqlite3

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/timpointer/golang-demo/mongodb/metro/smodel"
)

const dataconntemplate = "./data/%s.db?cache=shared&mode=rwc"

func getDatabaseName(name string) string {
	return fmt.Sprintf(dataconntemplate, name)
}

func GetWriteDB(sqliteConnStr string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteConnStr)
	if err != nil {
		return nil, err
	}
	//init database tables
	sqlStmts := []string{
		"create table if not exists user_registration_count (date  TEXT,store  TEXT,channel  TEXT,cardholder  TEXT,campaign  TEXT,count  INTEGER);",
		"create table if not exists customer_record (userid  INTEGER,store  TEXT,channel  TEXT,cardholder  TEXT,campaign  TEXT,date  INTEGER);",
	}
	for _, sqlStmt := range sqlStmts {
		_, err := db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", sqlStmt, err)
		}
	}
	return db, nil
}

func GetReadDB(connStr string) (*sql.DB, error) {
	return sql.Open("sqlite3", connStr)
}
func SaveRecord(db *sql.DB, r *smodel.CustomerRecord) (sql.Result, error) {
	log.Println(r.Time)
	log.Println(r.Time.Unix())
	return db.Exec("INSERT INTO customer_record  (userid ,store  ,channel ,cardholder ,campaign  ,date  )VALUES (?,?,?,?,?,?);",
		r.Custkey,
		r.Store,
		r.Channel,
		r.Cardholder,
		r.Campaign,
		r.Time.Unix())
}
