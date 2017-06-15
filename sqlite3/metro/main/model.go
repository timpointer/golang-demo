package main

type newCardHolderRigistration struct {
	*dataRow
	Cw   int
	Year int
}

type dataRow struct {
	CurrentYear  int
	Delta        int
	PercentDelta int
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

type dataRigistrationCWCount struct {
	dataOption
	DateCW string `json:"datecw"`
}
