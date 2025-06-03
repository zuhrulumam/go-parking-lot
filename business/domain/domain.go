package domain

import (
	"github.com/zuhrulumam/go-parking-lot/business/domain/parking"
	"github.com/zuhrulumam/go-parking-lot/business/domain/transaction"
	"gorm.io/gorm"
)

type Domain struct {
	Parking     parking.DomainItf
	Transaction transaction.DomainItf
}

type Option struct {
	DB *gorm.DB
}

func Init(opt Option) *Domain {
	d := &Domain{
		Parking: parking.InitParkingDomain(parking.Option{
			DB: opt.DB,
		}),
		Transaction: transaction.Init(transaction.Option{
			DB: opt.DB,
		}),
	}

	return d
}
