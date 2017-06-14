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
	Date       string
	Store      string
	Channel    string
	Cardholder string
	Campaign   string
	Count      int
}
