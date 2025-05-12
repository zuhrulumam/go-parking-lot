package usecase

import (
	"github.com/zuhrulumam/doit-test/business/domain"
	"github.com/zuhrulumam/doit-test/business/usecase/parking"
)

type Usecase struct {
	Parking parking.UsecaseItf
}

type Option struct {
}

func Init(dom *domain.Domain, opt Option) *Usecase {
	u := &Usecase{
		Parking: parking.InitParkingUsecase(parking.Option{
			ParkingDom:     dom.Parking,
			TransactionDom: dom.Transaction,
		}),
	}

	return u
}
