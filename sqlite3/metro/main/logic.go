package main

import (
	"database/sql"
	"log"
	"time"

	ttime "github.com/timpointer/golang-demo/time"
)

type newCardHolderRigistration struct {
	*dataRow
	cw   int
	year int
}

type dataRow struct {
	currentYear  int
	delta        int
	percentDelta int
}

func reportYTD(db *sql.DB, store string) (*dataRow, error) {
	t := time.Now()
	t.Year()
	firstDayOfYear := time.Date(t.Year(), 0, 0, 0, 0, 0, 0, time.UTC)
	return report(db, firstDayOfYear, t, store)
}

func reportCW(db *sql.DB, cw int, store string) (*newCardHolderRigistration, error) {
	t := time.Now()
	tCW := ttime.FirstDayOfISOWeek(t.Year(), cw, time.UTC)
	tCWEnd := tCW.AddDate(0, 0, 7)
	row, err := report(db, tCW, tCWEnd, store)
	if err != nil {
		return nil, err
	}
	ncr := &newCardHolderRigistration{}
	ncr.dataRow = row
	ncr.year = t.Year()
	ncr.cw = cw
	return ncr, nil
}

func report(db *sql.DB, start time.Time, end time.Time, store string) (*dataRow, error) {
	row := &dataRow{}
	count, err := querycount(db, start, end, store)
	if err != nil {
		return nil, err
	}
	row.currentYear = count
	log.Println("count", count)

	lastyearstart := start.AddDate(-1, 0, 0)
	lastyearend := end.AddDate(-1, 0, 0)

	lastyearcount, err := querycount(db, lastyearstart, lastyearend, store)
	if err != nil {
		return nil, err
	}
	delta := calDelta(count, lastyearcount)
	row.delta = delta
	log.Println("delta", delta)
	percent := calPercent(count, lastyearcount)
	row.percentDelta = percent
	log.Println("percent", percent)
	return row, nil
}

func calDelta(thisyear, lastyear int) int {
	return thisyear - lastyear
}

func calPercent(thisyear, lastyear int) int {
	return (thisyear - lastyear) * 100 / thisyear
}

func querycount(db *sql.DB, start time.Time, end time.Time, store string) (int, error) {
	var count int
	err := db.QueryRow("select count(userid) from user_registration where date BETWEEN ?  and ? and channel = 'ali'", start.Unix(), end.Unix()).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
