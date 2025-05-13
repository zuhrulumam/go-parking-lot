package parking

import (
	"context"

	"github.com/zuhrulumam/doit-test/business/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=business/domain/parking/parking.go -destination=mocks/mock_parking.go -package=mocks
type DomainItf interface {
	GetAvailableParkingSpot(ctx context.Context, data entity.GetAvailableParkingSpot) ([]entity.ParkingSpot, error)
	InsertVehicle(ctx context.Context, data entity.InsertVehicle) error
	UpdateParkingSpot(ctx context.Context, data entity.UpdateParkingSpot) error
	UpdateVehicle(ctx context.Context, data entity.UpdateVehicle) error
	GetVehicle(ctx context.Context, data entity.SearchVehicle) (entity.Vehicle, error)
}

type parking struct {
	db *gorm.DB
}

type Option struct {
	DB *gorm.DB
}

func InitParkingDomain(opt Option) DomainItf {
	p := &parking{
		db: opt.DB,
	}

	return p
}
