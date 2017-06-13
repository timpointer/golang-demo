package main

import (
	"database/sql"
	"time"

	"fmt"

	"sort"

	randomdata "github.com/Pallinder/go-randomdata"
	ttime "github.com/timpointer/golang-demo/time"
)

func insertDumpData(db *sql.DB) error {
	times := Int64Slice(make([]int64, 1000))
	for i := 0; i < len(times); i++ {
		times = append(times, chooseTime())
	}
	sort.Sort(times)

	for i := 0; i < 1000; i++ {
		_, err := insert(db, 1, randomdata.SillyName(), chooseStore(), chooseChannel(), chooseCardholder(), chooseActivity(), chooseTime())
		if err != nil {
			return fmt.Errorf("inser: %v", err)
		}
	}
	return nil
}

func collectionReportData(db *sql.DB) error {
	maps := map[string][]string{
		"storepanel": nil,
		"channel":    nil,
		"cardholder": nil,
		"campaint":   nil,
	}
	fmt.Println(maps)
	for key := range maps {
		list, err := selectOptions(db, key)
		if err != nil {
			return fmt.Errorf("selectOptions storepanel:%v", err)
		}
		maps[key] = list
	}

	for store := range maps["storepanel"] {
		fmt.Println(store)
	}
	// todo
	return nil
}

func splitByMonth(times []int64) map[string][]int64 {
	m := map[string][]int64{}
	for _, t := range times {
		monthStr := ttime.GetYearMonth(time.Unix(t, 0))
		m[monthStr] = append(m[monthStr], t)
	}
	return m
}

func selectOptions(db *sql.DB, colmun string) ([]string, error) {
	rows, err := db.Query("select DISTINCT ? from user_registration", colmun)
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

func insert(db *sql.DB, userid int, name, storepanel, channel, cardholder, campaign string, data int64) (sql.Result, error) {
	return db.Exec("INSERT INTO user_registration  (userid ,name  ,storepanel  ,channel ,cardholder ,campaign  ,date  )VALUES (?,?,?,?,?,?,?);",
		userid, name, storepanel, channel, cardholder, campaign, data)
}

// Int64Slice is sortable int64 slice
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
