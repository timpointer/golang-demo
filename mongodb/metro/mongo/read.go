package mongo

import (
	"evolve/evolution/utils"
	"metro/evsync/model"

	"gopkg.in/mgo.v2/bson"
)

func GetCustomer(m *utils.MongoSessionManager, storekey, custkey int) (*model.Customer, error) {
	session, err := m.Get()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	result := &model.Customer{}
	c := session.DB("wxngiapi").C("customer")
	err = c.Find(bson.M{"storekey": storekey, "custkey": custkey}).One(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
