package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteConnStr = "./data/prc3v.db?cache=shared&mode=rwc"
)

func main() {
	db, err := sql.Open("sqlite3", sqliteConnStr)
	if err != nil {
		log.Fatal(err)
	}
	//init database tables
	sqlStmts := []string{
		"create table if not exists whitelistscc(home_store_id INT,cust_no INT,auth_person_id INT);",
		"create table if not exists couponreport(openid text,cardno text,savetime datetime,lastname text,firstname text,id text,idtype text,phone text,PRIMARY KEY (openid) );",
		"create unique index if not exists couponreportUniqueIdx on couponreport(openid);",
		"create table if not exists crdetail(openid text,epc text,point int,status text,activetime datetime,redeemtime datetime,PRIMARY KEY (epc) );",
		"create index if not exists couponreportIdx on crdetail(epc,openid,status);",
	}
	for _, sqlStmt := range sqlStmts {
		_, err := db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}
}
