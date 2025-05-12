package parking

import (
	"context"

	"github.com/zuhrulumam/doit-test/business/entity"
	"gorm.io/gorm"
)

type DomainItf interface {
	GetAvailableParkingSpot(ctx context.Context, data entity.GetAvailableParkingSpot) ([]entity.ParkingSpot, error)
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
