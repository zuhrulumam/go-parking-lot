package entity

type VehicleType string

const (
	Bicycle    VehicleType = "B"
	Motorcycle VehicleType = "M"
	Automobile VehicleType = "A"
)

type Spot struct {
	SpotType  VehicleType
	IsActive  bool
	Occupied  bool
	VehicleNo string
}
