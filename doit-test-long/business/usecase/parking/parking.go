package parking

import (
	"context"

	parkingDom "github.com/zuhrulumam/doit-test/business/domain/parking"
	transactionDom "github.com/zuhrulumam/doit-test/business/domain/transaction"
	"github.com/zuhrulumam/doit-test/business/entity"
)

type UsecaseItf interface {
	Park(ctx context.Context, data entity.Park) error
	// Unpark()
	// AvailableSpot()
	// SearchVehicle()
}

type Option struct {
	ParkingDom     parkingDom.DomainItf
	TransactionDom transactionDom.DomainItf
}

type parking struct {
	ParkingDom     parkingDom.DomainItf
	TransactionDom transactionDom.DomainItf
}

func InitParkingUsecase(opt Option) UsecaseItf {
	p := &parking{
		ParkingDom:     opt.ParkingDom,
		TransactionDom: opt.TransactionDom,
	}

	return p
}
