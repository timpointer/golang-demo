package main

import (
	"evolve/evolution/utils"
	"log"
	"time"

	"github.com/timpointer/golang-demo/mongodb/metro/mongo"
	"github.com/timpointer/golang-demo/mongodb/metro/parse"
	"github.com/timpointer/golang-demo/mongodb/metro/smodel"
	"github.com/timpointer/golang-demo/mongodb/metro/sqlite3"
)

func main() {
	mgo := utils.NewMongoSessionManager("report", "localhost", 100)
	db, err := sqlite3.GetWriteDB("./data/warehouse.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	customer, err := mongo.GetCustomer(mgo, 59, 66087)
	if err != nil {
		log.Fatal(err)
	}

	record, err := parse.ParseCustomer(customer, 1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = sqlite3.SaveRecord(db, record)
	if err != nil {
		log.Fatal(err)
	}
	//start := time.Date(2017, 4, 23, 0, 0, 0, 0, time.UTC)
	//end := time.Date(2017, 5, 24, 0, 0, 0, 0, time.UTC)
	/*rows, err := getCustomerRow(Mgosession, start, end)
	reportRows := parseCustomerRow(rows)
	countRows := parseReportRow(reportRows)
	fmt.Println("**********************")
	for _, val := range countRows {
		fmt.Println(val)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", len(rows))
	fmt.Println("reportcount:", len(countRows))
	*/
}

func parseReportRow(rows []*CustomerReportRow) []*dataRigistrationCount {
	return countReduce(rows)
}

/*
func parseCustomerRow(rows []CustomerRow) []*CustomerReportRow {
	newrows := []*CustomerReportRow{}
	for _, row := range rows {
		store, channel := "default", "default"
		cardholder := "default"
		campaign := "default"

		if row.channel != "" {
			campaign = row.channel
			if s, c, err := parseChannel(row.channel); err == nil {
				store, channel = s, c
			}
		}

		data := &CustomerReportRow{
			Date: row.LPRegisterDate,
			dataOption: dataOption{
				Store:      store,
				Channel:    channel,
				Campaign:   campaign,
				Cardholder: cardholder,
				Count:      1,
			},
		}
		newrows = append(newrows, data)
		fmt.Println(row)
	}
	return newrows
}
*/

type CustomerReportRow struct {
	smodel.Dimension
	smodel.State
	Date time.Time
}

type Customer struct {
	CardHolders []CardHolder
}

type CardHolder struct {
	Channel        string    `bson:",omitempty"`
	LPRegisterDate time.Time `bson:",omitempty"`
}

type dataRigistrationCount struct {
	smodel.Dimension
	smodel.State
	Date string `json:"date"`
}

func countReduce(rows []*CustomerReportRow) []*dataRigistrationCount {
	set := map[string]*dataRigistrationCount{}
	for _, row := range rows {
		hash, err := hashDay(row)
		if err != nil {
			return nil
		}
		_, ok := set[hash]
		if ok == true {
			set[hash].Count += row.Count
		} else {
			set[hash] = &dataRigistrationCount{
				smodel.Dimension{
					row.Store, row.Channel, row.Cardholder, row.Campaign,
				},
				smodel.State{
					row.Count,
				},
				hash[:8],
			}
		}
	}
	values := []*dataRigistrationCount{}
	for _, value := range set {
		values = append(values, value)
	}
	return values
}

func hashDay(r *CustomerReportRow) (string, error) {
	return r.Date.Format("20060102") + r.Campaign + r.Cardholder + r.Channel + r.Store, nil
}
