package main

import (
	"database/sql"
	"fmt"
)

func getWriteDB(sqliteConnStr string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteConnStr)
	if err != nil {
		return nil, err
	}
	//init database tables
	sqlStmts := []string{
		"create table if not exists user_registration_count (date  TEXT,store  TEXT,channel  TEXT,cardholder  TEXT,campaign  TEXT,count  INTEGER);",
	}
	for _, sqlStmt := range sqlStmts {
		_, err := db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", sqlStmt, err)
		}
	}
	return db, nil
}

func getReadDB(connStr string) (*sql.DB, error) {
	return sql.Open("sqlite3", connStr)
}
