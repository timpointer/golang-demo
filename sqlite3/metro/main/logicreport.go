package main

import (
	"database/sql"
	"log"
	"time"

	"fmt"

	ttime "github.com/timpointer/golang-demo/time"
)

func reportYTD(db *sql.DB, store string) (*dataRow, error) {
	t := time.Now()
	t.Year()
	firstDayOfYear := time.Date(t.Year(), 0, 0, 0, 0, 0, 0, time.UTC)
	return report(db, firstDayOfYear, t, store)
}

func reportCW(db *sql.DB, cw int, store string) (*newCardHolderRigistration, error) {
	t := time.Now()
	tCW := ttime.FirstDayOfISOWeek(t.Year(), cw, time.UTC)
	endTCW := tCW.AddDate(0, 0, 7)
	row, err := report(db, tCW, endTCW, store)
	if err != nil {
		return nil, err
	}

	ncr := &newCardHolderRigistration{
		row, t.Year(), cw,
	}
	return ncr, nil
}

func reportWapper(db *sql.DB, start time.Time, end time.Time, store string) (*dataRow, error) {
	dblist, err := getDatabase(start, end)
	if err != nil {
		return nil, err
	}
	for sqlitedb := range dblist {
		log.Println(sqlitedb)
	}
	return nil, nil
}

func getDatabase(start, end time.Time) ([]*sql.DB, error) {
	dbList := []*sql.DB{}
	connDataList := ttime.GetListMonth(start, end)
	for _, connData := range connDataList {
		connStr := fmt.Sprintf("./data/metroreport%s.db?cache=shared&mode=rwc", connData)
		db, err := getReadDB(connStr)
		if err != nil {
			return nil, fmt.Errorf("open %s:%v", connStr, err)
		}
		dbList = append(dbList, db)
	}
	return dbList, nil
}

func reportURC(db *sql.DB, start, end string, store string) ([]*dataRigistrationCount, error) {
	return queryURC(db, start, end, store)
}

func report(db *sql.DB, start time.Time, end time.Time, store string) (*dataRow, error) {
	row := &dataRow{}
	count, err := querycount(db, start, end, store, "", "", "")
	if err != nil {
		return nil, err
	}
	row.CurrentYear = count
	log.Println("count", count)

	lastyearstart := start.AddDate(-1, 0, 0)
	lastyearend := end.AddDate(-1, 0, 0)

	lastyearcount, err := querycount(db, lastyearstart, lastyearend, store, "", "", "")
	if err != nil {
		return nil, err
	}
	delta := calDelta(count, lastyearcount)
	row.Delta = delta
	log.Println("delta", delta)
	percent := calPercent(count, lastyearcount)
	row.PercentDelta = percent
	log.Println("percent", percent)
	return row, nil
}

func calDelta(thisyear, lastyear int) int {
	return thisyear - lastyear
}

func calPercent(thisyear, lastyear int) int {
	return (thisyear - lastyear) * 100 / thisyear
}
