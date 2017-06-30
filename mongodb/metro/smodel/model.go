package smodel

import "time"

const (
	// DimensionStore is
	DimensionStore = "store"
	// DimensionChannel is
	DimensionChannel = "channel"
	// DimensionCardholder is
	DimensionCardholder = "cardholder"
	// DimensionCampaign  is
	DimensionCampaign = "campaign"
)

func NewCustomerRecord() *CustomerRecord {
	return &CustomerRecord{
		id:  "",
		Key: Key{},
		Dimension: Dimension{
			"default", "default", "default", "default",
		},
		State: State{
			0,
		},
	}
}

type Dimension struct {
	Store      string `json:"store"`
	Channel    string `json:"channel"`
	Cardholder string `json:"cardholder"`
	Campaign   string `json:"campaign"`
}

type State struct {
	Count int `json:"count"`
}

type Key struct {
	Storekey      int
	Custkey       int
	Cardholderkey int
}

type CustomerRecord struct {
	id string
	Key
	Dimension
	State
	Time time.Time
}
