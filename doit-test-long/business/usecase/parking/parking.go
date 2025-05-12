package parking

import (
	parkingDom "github.com/zuhrulumam/doit-test/business/domain/parking"
)

type UsecaseItf interface {
	Park()
	// Unpark()
	// AvailableSpot()
	// SearchVehicle()
}

type parking struct {
}

func InitParkingUsecase(parkingDom parkingDom.DomainItf) UsecaseItf {
	p := &parking{}

	return p
}
