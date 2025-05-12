package domain

import "github.com/zuhrulumam/doit-test/business/domain/parking"

type Domain struct {
	Parking parking.DomainItf
}

func Init() *Domain {
	d := &Domain{
		Parking: parking.InitParkingDomain(),
	}

	return d
}
