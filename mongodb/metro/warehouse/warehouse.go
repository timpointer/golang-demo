package warehouse

import (
	"database/sql"
	"metro/evsync/model"

	"github.com/timpointer/golang-demo/mongodb/metro/smodel"
)

type Reader interface {
	Read(storekey, custkey, cardholderkey string) (*model.Customer, error)
}

type Writer interface {
	Write(r *smodel.CustomerRecord) (sql.Result, error)
}

type Closer interface {
	Close() error
}
