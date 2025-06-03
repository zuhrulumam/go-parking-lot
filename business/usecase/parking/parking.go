package parking

import (
	"context"

	parkingDom "github.com/zuhrulumam/go-parking-lot/business/domain/parking"
	transactionDom "github.com/zuhrulumam/go-parking-lot/business/domain/transaction"
	"github.com/zuhrulumam/go-parking-lot/business/entity"
)

type UsecaseItf interface {
	Park(ctx context.Context, data entity.Park) error
	Unpark(ctx context.Context, data entity.UnPark) error
	AvailableSpot(ctx context.Context, data entity.GetAvailablePark) ([]entity.ParkingSpot, error)
	SearchVehicle(ctx context.Context, data entity.SearchVehicle) (entity.Vehicle, error)
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
