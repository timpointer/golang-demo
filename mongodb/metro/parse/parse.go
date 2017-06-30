package parse

import (
	"errors"
	"metro/evsync/model"

	"github.com/timpointer/golang-demo/mongodb/metro/smodel"
)

func parseChannel(c string) (store string, channel string, err error) {
	if len(c) != 11 {
		err = errors.New("format is not correct")
		return
	}
	store = c[4:7]
	channel = c[9:11]
	return
}

func ParseCustomer(c *model.Customer, cardholder int) ([]*smodel.CustomerRecord, error) {
	var holder *model.CardHolder
	for _, ch := range c.CardHolders {
		if ch.CardHolderKey == cardholder {
			holder = &ch
			break
		}
	}
	records := []*smodel.CustomerRecord{}
	if holder == nil {

	} else {
		if len(holder.MemberTypeHistory) > 0 {
			for _, v := range holder.MemberTypeHistory {
				log := smodel.NewCustomerRecord()
				log.Storekey = c.StoreKey
				log.Custkey = c.CustKey
				log.Cardholderkey = holder.CardHolderKey
				parseCardholder(log, holder)
				log.Time = v.TimeStamp
				records = append(records, log)
			}
		} else {
			log := smodel.NewCustomerRecord()
			log.Storekey = c.StoreKey
			log.Custkey = c.CustKey
			log.Cardholderkey = holder.CardHolderKey
			log.Time = holder.LPRegisterDate
			parseCardholder(log, holder)
			records = append(records, log)
		}
	}
	return records, nil
}

func parseCardholder(log *smodel.CustomerRecord, c *model.CardHolder) {
	if c.Channel != "" {
		log.Campaign = c.Channel
		if s, c, err := parseChannel(c.Channel); err == nil {
			log.Store, log.Channel = s, c
		}
	}
}
