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

type dataRigistrationCount struct {
	Date       string `json:"date"`
	Store      string `json:"store"`
	Channel    string `json:"channel"`
	Cardholder string `json:"cardholder"`
	Campaign   string `json:"campaign"`
	Count      int    `json:"count"`
}
