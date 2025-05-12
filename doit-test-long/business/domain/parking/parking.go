package parking

type DomainItf interface {
	Park()
	// Unpark()
	// AvailableSpot()
	// SearchVehicle()
}

type parking struct {
}

type Option struct {
}

func InitParkingDomain() DomainItf {
	p := &parking{}

	return p
}
