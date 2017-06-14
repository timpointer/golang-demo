package main

import (
	"database/sql"
	"log"
	"time"
)

func querycount(db *sql.DB, start time.Time, end time.Time, store, channel, campaign, cardholder string) (int, error) {
	var count int
	where := ""
	if store != "" {
		where = " and store='" + store + "'"
	}
	if channel != "" {
		where += " and channel='" + channel + "'"
	}
	if campaign != "" {
		where += " and campaign='" + campaign + "'"
	}
	if cardholder != "" {
		where += " and cardholder='" + cardholder + "'"
	}
	log.Println(where)

	err := db.QueryRow("select count(userid) from user_registration where date BETWEEN ?  and ? "+where, start.Unix(), end.Unix()).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func queryURC(db *sql.DB, start, end string, store string) ([]*dataRigistrationCount, error) {
	datarows := []*dataRigistrationCount{}
	rows, err := db.Query("select date, channel,campaign,cardholder,store,count from user_registration_count where date BETWEEN ? and ?", start, end)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := &dataRigistrationCount{}
		if err := rows.Scan(&row.Date, &row.Channel, &row.Campaign, &row.Cardholder, &row.Store, &row.Count); err != nil {
			return nil, err
		}
		datarows = append(datarows, row)
	}
	rows.Close()
	return datarows, nil
}

func selectOptions(db *sql.DB, colmun string) ([]string, error) {
	rows, err := db.Query("select DISTINCT " + colmun + " from user_registration")
	if err != nil {
		return nil, err
	}
	list := []string{}
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, err
}

func insert(db *sql.DB, userid int, name, store, channel, cardholder, campaign string, data int64) (sql.Result, error) {
	return db.Exec("INSERT INTO user_registration  (userid ,name  ,store  ,channel ,cardholder ,campaign  ,date  )VALUES (?,?,?,?,?,?,?);",
		userid, name, store, channel, cardholder, campaign, data)
}

func insertCount(db *sql.DB, data *dataRigistrationCount) (sql.Result, error) {
	return db.Exec("INSERT INTO user_registration_count  (date,store  ,channel ,cardholder ,campaign  ,count  )VALUES (?,?,?,?,?,?);",
		data.Date, data.Store, data.Channel, data.Cardholder, data.Campaign, data.Count)
}
