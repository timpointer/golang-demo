package main

import (
	"errors"
	"evolve/evolution/utils"
	"log"
	"time"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	Mgosession := utils.NewMongoSessionManager("report", "localhost", 100)
	start := time.Date(2017, 5, 23, 0, 0, 0, 0, time.UTC)
	end := time.Date(2017, 5, 24, 0, 0, 0, 0, time.UTC)
	rows, err := getCustomerRow(Mgosession, start, end)
	parseCustomerRow(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", len(rows))
}

func parseCustomerRow(rows []CustomerRow) []dataRigistrationCount {
	newrows := []dataRigistrationCount{}
	for _, row := range rows {
		fmt.Println(row)
	}
	return newrows
}

func parseChannel(c string) (store string, channel string, err error) {
	if len(c) != 11 {
		err = errors.New("format is not correct")
		return
	}
	store = c[4:7]
	channel = c[9:11]
	return
}

func getCustomerRow(m *utils.MongoSessionManager, start, end time.Time) ([]CustomerRow, error) {
	session, err := m.Get()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	coll := session.DB("wxngiapi").C("customer")
	iter := coll.Find(bson.M{"cardholders": bson.M{"$elemMatch": bson.M{"lpregisterdate": bson.M{"$gte": start, "$lt": end}}}}).Iter()
	defer iter.Close()

	var rows []CustomerRow
	customer := Customer{}
	for iter.Next(&customer) {
		for _, ch := range customer.CardHolders {
			if ch.LPRegisterDate.Before(end) && ch.LPRegisterDate.After(start) {
				rows = append(rows, CustomerRow{ch.LPRegisterDate, ch.Channel})
			}
		}
	}
	return rows, nil
}

type CustomerRow struct {
	LPRegisterDate time.Time
	channel        string
}

type Customer struct {
	CardHolders []CardHolder
}

type CardHolder struct {
	Channel        string    `bson:",omitempty"`
	LPRegisterDate time.Time `bson:",omitempty"`
}

type dataOption struct {
	Store      string `json:"store"`
	Channel    string `json:"channel"`
	Cardholder string `json:"cardholder"`
	Campaign   string `json:"campaign"`
	Count      int    `json:"count"`
}

type dataRigistrationCount struct {
	dataOption
	Date string `json:"date"`
}
