package main

import (
	"database/sql"
	"log"
	"time"

	"fmt"

	ttime "github.com/timpointer/golang-demo/time"
)

func reportYTD(db *sql.DB) (*dataRow, error) {
	t := time.Now()
	t.Year()
	firstDayOfYear := time.Date(t.Year(), 0, 0, 0, 0, 0, 0, time.UTC)
	return report(db, firstDayOfYear, t)
}

func reportCW(db *sql.DB, cw int) (*newCardHolderRigistration, error) {
	t := time.Now()
	tCW := ttime.FirstDayOfISOWeek(t.Year(), cw, time.UTC)
	endTCW := tCW.AddDate(0, 0, 7)
	row, err := report(db, tCW, endTCW)
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

func reportURCCW(db *sql.DB, week int) ([]*dataRigistrationCount, error) {
	day := ttime.FirstDayOfISOWeek(t.Year(), week, time.UTC)
	start := ttime.GetYearMonthDay(day)
	end := ttime.GetYearMonthDay(day.AddDate(0, 0, 6))
	return queryURC(db, start, end)
}

func reportURC(db *sql.DB, start, end string) ([]*dataRigistrationCount, error) {
	return queryURC(db, start, end)
}

func report(db *sql.DB, start time.Time, end time.Time) (*dataRow, error) {
	row := &dataRow{}
	count, err := querycount(db, start, end, "", "", "", "")
	if err != nil {
		return nil, err
	}
	row.CurrentYear = count
	log.Println("count", count)

	lastyearstart := start.AddDate(-1, 0, 0)
	lastyearend := end.AddDate(-1, 0, 0)

	lastyearcount, err := querycount(db, lastyearstart, lastyearend, "", "", "", "")
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
