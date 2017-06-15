package main

import (
	"database/sql"
	"log"
	"time"

	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
	ttime "github.com/timpointer/golang-demo/time"
)

func insertDumpData() error {
	db, err := getWriteDB(sqliteConnStr)
	if err != nil {
		return fmt.Errorf("open db fialed:%v", err)
	}
	defer db.Close()
	times := Int64Slice(make([]int64, 100))
	leng := len(times)
	for i := 0; i < leng; i++ {
		times = append(times, chooseTime())
	}

	for i := 0; i < leng; i++ {
		_, err := insert(db, 1, randomdata.SillyName(), chooseStore(), chooseChannel(), chooseCardholder(), chooseActivity(), chooseTime())
		if err != nil {
			return fmt.Errorf("insert: %v", err)
		}
	}
	return nil
}

func collectionReportData(db *sql.DB, start, end time.Time) error {
	maps := map[string][]string{
		"store":      nil,
		"channel":    nil,
		"cardholder": nil,
		"campaign":   nil,
	}
	for key := range maps {
		list, err := selectOptions(db, key)
		if err != nil {
			return fmt.Errorf("selectOptions store:%v", err)
		}
		maps[key] = list
	}
	dbw, err := getWriteDB(sqliteConnStr)
	if err != nil {
		return fmt.Errorf("open write database failed; %v", err)
	}
	defer dbw.Close()
	log.Println(maps)
	for _, store := range maps["store"] {
		for _, channel := range maps["channel"] {
			for _, cardholder := range maps["cardholder"] {
				for _, campaign := range maps["campaign"] {
					count, err := querycount(db, start, end, store, channel, campaign, cardholder)
					if err != nil {
						return fmt.Errorf("select count:%v", err)
					}
					tstart := time.Date(2015, 5, 10, 23, 0, 0, 0, time.UTC)
					tend := time.Date(2017, 5, 10, 23, 0, 0, 0, time.UTC)
					days := ttime.GetListDay(tstart, tend)
					for _, day := range days {

						data := &dataRigistrationCount{
							dataOption{
								store,
								channel,
								cardholder,
								campaign,
								count,
							},
							day,
						}
						_, err = insertCount(dbw, data)
						if err != nil {
							return fmt.Errorf("insert count:%v", err)
						}
					}
					fmt.Println(count)
				}
			}
		}
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

func selectCount(db *sql.DB) int {
	// todo
	return 0
}

// Int64Slice is sortable int64 slice
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
